package utils

import (
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/model/cfg"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func buildDominatorIndexes(cfg intsrv.IControlFlowGraph) []util.IBitSet {
	list := cfg.BasicBlocks()
	length := list.Size()
	arrayOfDominatorIndexes := make([]util.IBitSet, length)

	initial := util.NewBitSetWithSize(length)
	_ = initial.Set(0)
	arrayOfDominatorIndexes[0] = initial

	for i := 0; i < length; i++ {
		initial = util.NewBitSetWithSize(length)
		initial.FlipRange(0, length)
		arrayOfDominatorIndexes[i] = initial
	}

	initial = arrayOfDominatorIndexes[0]
	initial.ClearAll()
	_ = initial.Set(0)

	change := false

	for {
		change = false

		for _, basicBlock := range list.ToSlice() {
			index := basicBlock.Index()
			dominatorIndexes := arrayOfDominatorIndexes[index]
			initial = dominatorIndexes.Clone()

			for _, predecessorBB := range basicBlock.Predecessors().ToSlice() {
				dominatorIndexes.And(arrayOfDominatorIndexes[predecessorBB.Index()])
			}

			_ = dominatorIndexes.Set(index)
			change = change || (initial != dominatorIndexes)
		}

		if !change {
			break
		}
	}

	return arrayOfDominatorIndexes
}

func IdentifyNaturalLoops(cfg intsrv.IControlFlowGraph, arrayOfDominatorIndexes []util.IBitSet) util.IList[intsrv.ILoop] {
	list := cfg.BasicBlocks()
	length := list.Size()
	arrayOfMemberIndexes := make([]util.IBitSet, length)

	// Identify loop members
	for i := 0; i < length; i++ {
		current := list.Get(i)
		dominatorIndexes := arrayOfDominatorIndexes[i]

		switch current.Type() {
		case intsrv.TypeConditionalBranch:
			index := current.Branch().Index()
			if (index >= 0) && dominatorIndexes.Get(index) {
				// 'branch' is a dominator -> Back edge found
				arrayOfMemberIndexes[index] = searchLoopMemberIndexes(length, arrayOfMemberIndexes[index], current, current.Branch())
			}
		case intsrv.TypeStatements, intsrv.TypeGoto:
			index := current.Next().Index()
			if (index >= 0) && dominatorIndexes.Get(index) {
				// 'next' is a dominator -> Back edge found
				arrayOfMemberIndexes[index] = searchLoopMemberIndexes(length, arrayOfMemberIndexes[index], current, current.Next())
			}
			break
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range current.SwitchCases().ToSlice() {
				index := switchCase.BasicBlock().Index()

				if (index >= 0) && dominatorIndexes.Get(index) {
					// 'switchCase' is a dominator -> Back edge found
					arrayOfMemberIndexes[index] = searchLoopMemberIndexes(length, arrayOfMemberIndexes[index], current, switchCase.BasicBlock())
				}
			}
			break
		}
	}

	// Loops & 'try' statements
	for i := 0; i < length; i++ {
		if arrayOfMemberIndexes[i] != nil {
			memberIndexes := arrayOfMemberIndexes[i]
			maxOffset := -1

			for j := 0; j < length; j++ {
				if memberIndexes.Get(j) {
					offset := list.Get(j).FromOffset()
					if maxOffset < offset {
						maxOffset = offset
					}
				}
			}

			start := list.Get(i)
			startDominatorIndexes := arrayOfDominatorIndexes[i]

			if start.Type() == intsrv.TypeTryDeclaration &&
				maxOffset != start.FromOffset() &&
				maxOffset < start.ExceptionHandlers().First().BasicBlock().FromOffset() {
				// 'try' statement outside the loop
				newStart := start.Next()
				newStartPredecessors := newStart.Predecessors()

				// Loop in 'try' statement
				iterator := start.Predecessors().Iterator()

				for iterator.HasNext() {
					predecessor := iterator.Next()

					if !startDominatorIndexes.Get(predecessor.Index()) {
						_ = iterator.Remove()
						predecessor.Replace(start, newStart)
						newStartPredecessors.Add(predecessor)
					}
				}

				_ = memberIndexes.Clear(start.Index())
				arrayOfMemberIndexes[newStart.Index()] = memberIndexes
				arrayOfMemberIndexes[i] = nil
			}
		}
	}

	// Build loops
	loops := util.NewDefaultList[intsrv.ILoop]()

	for i := 0; i < length; i++ {
		if arrayOfMemberIndexes[i] != nil {
			memberIndexes := arrayOfMemberIndexes[i]

			// Unoptimize loop
			start := list.Get(i)
			startDominatorIndexes := arrayOfDominatorIndexes[i]
			searchZoneIndexes := util.NewBitSetWithSize(length)
			searchZoneIndexes.Or(startDominatorIndexes)
			searchZoneIndexes.FlipRange(0, length)
			_ = searchZoneIndexes.Set(start.Index())

			if start.Type() == intsrv.TypeConditionalBranch {
				if (start.Next() != start) &&
					(start.Branch() != start) &&
					memberIndexes.Get(start.Next().Index()) &&
					memberIndexes.Get(start.Branch().Index()) {
					// 'next' & 'branch' blocks are inside the loop -> Split loop ?
					nextIndexes := util.NewBitSetWithSize(length)
					branchIndexes := util.NewBitSetWithSize(length)

					recursiveForwardSearchLoopMemberIndexes2(nextIndexes, memberIndexes, start.Next(), start)
					recursiveForwardSearchLoopMemberIndexes2(branchIndexes, memberIndexes, start.Branch(), start)

					commonMemberIndexes := nextIndexes.Clone()
					commonMemberIndexes.And(branchIndexes)

					onlyLoopHeaderIndex := util.NewBitSetWithSize(length)
					_ = onlyLoopHeaderIndex.Set(i)

					if commonMemberIndexes == onlyLoopHeaderIndex {
						// Only 'start' is the common basic block -> Split loop
						loops.Add(makeCfgLoopReducerLoop(list, start, searchZoneIndexes, memberIndexes))

						branchIndexes.FlipRange(0, length)
						searchZoneIndexes.And(branchIndexes)
						_ = searchZoneIndexes.Set(start.Index())

						loops.Add(makeCfgLoopReducerLoop(list, start, searchZoneIndexes, nextIndexes))
					} else {
						loops.Add(makeCfgLoopReducerLoop(list, start, searchZoneIndexes, memberIndexes))
					}
				} else {
					loops.Add(makeCfgLoopReducerLoop(list, start, searchZoneIndexes, memberIndexes))
				}
			} else {
				loops.Add(makeCfgLoopReducerLoop(list, start, searchZoneIndexes, memberIndexes))
			}
		}
	}

	loops.Sort(func(i, j int) bool {
		return loops.Get(i).Members().Size() < loops.Get(i).Members().Size()
	})

	return loops
}

func searchLoopMemberIndexes(length int, memberIndexes util.IBitSet, current, start intsrv.IBasicBlock) util.IBitSet {
	visited := util.NewBitSetWithSize(length)

	recursiveBackwardSearchLoopMemberIndexes(visited, current, start)

	if memberIndexes == nil {
		return visited
	} else {
		memberIndexes.Or(visited)
		return memberIndexes
	}
}

func recursiveBackwardSearchLoopMemberIndexes(visited util.IBitSet, current, start intsrv.IBasicBlock) {
	if visited.Get(current.Index()) == false {
		_ = visited.Set(current.Index())

		if current != start {
			for _, predecessor := range current.Predecessors().ToSlice() {
				recursiveBackwardSearchLoopMemberIndexes(visited, predecessor, start)
			}
		}
	}
}

func makeCfgLoopReducerLoop(list util.IList[intsrv.IBasicBlock], start intsrv.IBasicBlock,
	searchZoneIndexes, memberIndexes util.IBitSet) intsrv.ILoop {
	length := list.Size()
	maxOffset := -1

	for i := 0; i < length; i++ {
		if memberIndexes.Get(i) {
			offset := checkMaxOffset(list.Get(i))
			if maxOffset < offset {
				maxOffset = offset
			}
		}
	}

	// Extend members
	memberIndexes.ClearAll()
	recursiveForwardSearchLoopMemberIndexes3(memberIndexes, searchZoneIndexes, start, maxOffset)

	members := util.NewSet[intsrv.IBasicBlock]()

	for i := 0; i < length; i++ {
		if memberIndexes.Get(i) {
			members.Add(list.Get(i))
		}
	}

	// Search 'end' block
	end := cfg.End

	if start.Type() == intsrv.TypeConditionalBranch {
		// First, check natural 'end' blocks
		index := start.Branch().Index()
		if memberIndexes.Get(index) == false {
			end = start.Branch()
		} else {
			index = start.Next().Index()
			if memberIndexes.Get(index) == false {
				end = start.Next()
			}
		}
	}

	if end == cfg.End {
		// Not found, check all member blocks
		end = searchEndBasicBlock(memberIndexes, maxOffset, members)

		if !end.MatchType(intsrv.TypeEnd|intsrv.TypeReturn|intsrv.TypeLoopStart|intsrv.TypeLoopContinue|intsrv.TypeLoopEnd) &&
			(end.Predecessors().Size() == 1) &&
			(end.Predecessors().Iterator().Next().LastLineNumber()+1 >= end.FirstLineNumber()) {
			set := util.NewSet[intsrv.IBasicBlock]()

			if recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, end, nil) {
				members.AddAll(set.ToSlice())

				for _, member := range set.ToSlice() {
					if member.Index() >= 0 {
						_ = memberIndexes.Set(member.Index())
					}
				}

				end = searchEndBasicBlock(memberIndexes, maxOffset, set)
			}
		}
	}

	// Extend last member
	if end != cfg.End {
		m := util.NewSet[intsrv.IBasicBlock]()
		set := util.NewSet[intsrv.IBasicBlock]()

		for _, member := range m.ToSlice() {
			if member.Type() == intsrv.TypeConditionalBranch && member != start {
				set.Clear()
				if recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, member.Next(), end) {
					members.AddAll(set.ToSlice())
				}
				set.Clear()
				if recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, member.Branch(), end) {
					members.AddAll(set.ToSlice())
				}
			}
		}
	}

	return cfg.NewLoop(start, members, end)
}

func searchEndBasicBlock(memberIndexes util.IBitSet, maxOffset int, members util.ISet[intsrv.IBasicBlock]) intsrv.IBasicBlock {
	end := cfg.End

	for _, member := range members.ToSlice() {
		switch member.Type() {
		case intsrv.TypeConditionalBranch:
			bb := member.Branch()
			if !memberIndexes.Get(bb.Index()) && (maxOffset < bb.FromOffset()) {
				end = bb
				maxOffset = bb.FromOffset()
				break
			}
			fallthrough
		case intsrv.TypeStatements, intsrv.TypeGoto:
			bb := member.Next()
			if !memberIndexes.Get(bb.Index()) && (maxOffset < bb.FromOffset()) {
				end = bb
				maxOffset = bb.FromOffset()
			}
			break
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range member.SwitchCases().ToSlice() {
				bb := switchCase.BasicBlock()
				if !memberIndexes.Get(bb.Index()) && (maxOffset < bb.FromOffset()) {
					end = bb
					maxOffset = bb.FromOffset()
				}
			}
		case intsrv.TypeTryDeclaration:
			bb := member.Next()
			if !memberIndexes.Get(bb.Index()) && (maxOffset < bb.FromOffset()) {
				end = bb
				maxOffset = bb.FromOffset()
			}
			for _, exceptionHandler := range member.ExceptionHandlers().ToSlice() {
				bb = exceptionHandler.BasicBlock()
				if !memberIndexes.Get(bb.Index()) && (maxOffset < bb.FromOffset()) {
					end = bb
					maxOffset = bb.FromOffset()
				}
			}
			break
		}
	}

	return end
}

func checkMaxOffset(basicBlock intsrv.IBasicBlock) int {
	maxOffset := basicBlock.FromOffset()
	var offset int

	if basicBlock.Type() == intsrv.TypeTryDeclaration {
		for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
			if exceptionHandler.InternalThrowableName() == "" {
				// Search throw block
				offset = checkThrowBlockOffset(exceptionHandler.BasicBlock())
			} else {
				offset = checkSynchronizedBlockOffset(exceptionHandler.BasicBlock())
			}
			if maxOffset < offset {
				maxOffset = offset
			}
		}
	} else if basicBlock.Type() == intsrv.TypeSwitchDeclaration {
		var lastBB intsrv.IBasicBlock
		var previousBB intsrv.IBasicBlock

		for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
			bb := switchCase.BasicBlock()
			if (lastBB == nil) || (lastBB.FromOffset() < bb.FromOffset()) {
				previousBB = lastBB
				lastBB = bb
			}
		}
		if previousBB != nil {
			offset = checkSynchronizedBlockOffset(previousBB)
			if maxOffset < offset {
				maxOffset = offset
			}
		}
	}

	return maxOffset
}

func checkSynchronizedBlockOffset(basicBlock intsrv.IBasicBlock) int {
	if basicBlock.Next().Type() == intsrv.TypeTryDeclaration && LastOpcode(basicBlock) == 194 { // MONITORENTER
		return checkThrowBlockOffset(basicBlock.Next().ExceptionHandlers().First().BasicBlock())
	}
	return basicBlock.FromOffset()
}

func checkThrowBlockOffset(basicBlock intsrv.IBasicBlock) int {
	offset := basicBlock.FromOffset()
	watchdog := util.NewBitSet()

	for !basicBlock.MatchType(intsrv.GroupEnd) && !watchdog.Get(basicBlock.Index()) {
		_ = watchdog.Set(basicBlock.Index())
		basicBlock = basicBlock.Next()
	}

	if basicBlock.Type() == intsrv.TypeThrow {
		return basicBlock.FromOffset()
	}

	return offset
}

func recursiveForwardSearchLoopMemberIndexes2(visited, searchZoneIndexes util.IBitSet, current, target intsrv.IBasicBlock) {
	if !current.MatchType(intsrv.GroupEnd) && (visited.Get(current.Index()) == false) && (searchZoneIndexes.Get(current.Index()) == true) {
		_ = visited.Set(current.Index())

		if current != target {
			recursiveForwardSearchLoopMemberIndexes2(visited, searchZoneIndexes, current.Next(), target)
			recursiveForwardSearchLoopMemberIndexes2(visited, searchZoneIndexes, current.Branch(), target)

			for _, switchCase := range current.SwitchCases().ToSlice() {
				recursiveForwardSearchLoopMemberIndexes2(visited, searchZoneIndexes, switchCase.BasicBlock(), target)
			}

			for _, exceptionHandler := range current.ExceptionHandlers().ToSlice() {
				recursiveForwardSearchLoopMemberIndexes2(visited, searchZoneIndexes, exceptionHandler.BasicBlock(), target)
			}

			if current.Type() == intsrv.TypeGotoInTernaryOperator {
				_ = visited.Set(current.Next().Index())
			}
		}
	}
}

func recursiveForwardSearchLoopMemberIndexes3(visited, searchZoneIndexes util.IBitSet, current intsrv.IBasicBlock, maxOffset int) {
	if !current.MatchType(intsrv.TypeEnd|intsrv.TypeLoopStart|intsrv.TypeLoopContinue|intsrv.TypeLoopEnd|intsrv.TypeSwitchBreak) &&
		(visited.Get(current.Index()) == false) &&
		(searchZoneIndexes.Get(current.Index()) == true) &&
		(current.FromOffset() <= maxOffset) {
		_ = visited.Set(current.Index())

		recursiveForwardSearchLoopMemberIndexes3(visited, searchZoneIndexes, current.Next(), maxOffset)
		recursiveForwardSearchLoopMemberIndexes3(visited, searchZoneIndexes, current.Branch(), maxOffset)

		for _, switchCase := range current.SwitchCases().ToSlice() {
			recursiveForwardSearchLoopMemberIndexes3(visited, searchZoneIndexes, switchCase.BasicBlock(), maxOffset)
		}

		for _, exceptionHandler := range current.ExceptionHandlers().ToSlice() {
			recursiveForwardSearchLoopMemberIndexes3(visited, searchZoneIndexes, exceptionHandler.BasicBlock(), maxOffset)
		}

		if current.Type() == intsrv.TypeGotoInTernaryOperator {
			_ = visited.Set(current.Next().Index())
		}
	}
}

func recursiveForwardSearchLastLoopMemberIndexes(members util.ISet[intsrv.IBasicBlock], searchZoneIndexes util.IBitSet,
	set util.ISet[intsrv.IBasicBlock], current, end intsrv.IBasicBlock) bool {
	if (current == end) || members.Contains(current) || set.Contains(current) {
		return true
	} else if current.MatchType(intsrv.GroupSingleSuccessor) {
		if !inSearchZone(current.Next(), searchZoneIndexes) || !predecessorsInSearchZone(current, searchZoneIndexes) {
			_ = searchZoneIndexes.Clear(current.Index())
			return true
		} else {
			set.Add(current)
			return recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, current.Next(), end)
		}
	} else if current.Type() == intsrv.TypeConditionalBranch {
		if !inSearchZone(current.Next(), searchZoneIndexes) ||
			!inSearchZone(current.Branch(), searchZoneIndexes) ||
			!predecessorsInSearchZone(current, searchZoneIndexes) {
			_ = searchZoneIndexes.Clear(current.Index())
			return true
		} else {
			set.Add(current)
			return recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, current.Next(), end) ||
				recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, current.Branch(), end)
		}
	} else if current.MatchType(intsrv.GroupEnd) {
		if !predecessorsInSearchZone(current, searchZoneIndexes) {
			if current.Index() >= 0 {
				_ = searchZoneIndexes.Clear(current.Index())
			}
		} else {
			set.Add(current)
		}
		return true
	}

	return false
}

func predecessorsInSearchZone(basicBlock intsrv.IBasicBlock, searchZoneIndexes util.IBitSet) bool {
	predecessors := basicBlock.Predecessors()

	for _, predecessor := range predecessors.ToSlice() {
		if !inSearchZone(predecessor, searchZoneIndexes) {
			return false
		}
	}

	return true
}

func inSearchZone(basicBlock intsrv.IBasicBlock, searchZoneIndexes util.IBitSet) bool {
	return basicBlock.MatchType(intsrv.TypeEnd|intsrv.TypeReturn|intsrv.TypeRet|
		intsrv.TypeLoopEnd|intsrv.TypeLoopStart|intsrv.TypeInfiniteGoto|intsrv.TypeJump) ||
		searchZoneIndexes.Get(basicBlock.Index())
}

func recheckEndBlock(members util.ISet[intsrv.IBasicBlock], end intsrv.IBasicBlock) intsrv.IBasicBlock {
	flag := false

	for _, predecessor := range end.Predecessors().ToSlice() {
		if !members.Contains(predecessor) {
			flag = true
			break
		}
	}

	if flag {
		return end
	}

	// Search new 'end' block
	var newEnd intsrv.IBasicBlock

	for _, member := range members.ToSlice() {
		if member.MatchType(intsrv.GroupSingleSuccessor) {
			bb := member.Next()
			if (bb != end) && !members.Contains(bb) {
				newEnd = bb
				break
			}
		} else if member.Type() == intsrv.TypeConditionalBranch {
			bb := member.Next()
			if (bb != end) && !members.Contains(bb) {
				newEnd = bb
				break
			}
			bb = member.Branch()
			if (bb != end) && !members.Contains(bb) {
				newEnd = bb
				break
			}
		}
	}

	if (newEnd == nil) || (end.FromOffset() >= newEnd.FromOffset()) {
		return end
	}

	// Replace 'end' block
	if end.MatchType(intsrv.TypeReturn | intsrv.TypeReturnValue | intsrv.TypeThrow) {
		members.Add(end)
		end = newEnd
	} else if end.MatchType(intsrv.GroupSingleSuccessor) && (end.Next() == newEnd) {
		members.Add(end)
		end = newEnd
	} else {
		return end
	}

	return end
}

func reduceLoop(loop intsrv.ILoop) intsrv.IBasicBlock {
	start := loop.Start()
	members := loop.Members()
	end := loop.End()
	toOffset := start.ToOffset()

	// Recheck 'end' block
	end = recheckEndBlock(members, end)

	// Build new basic block for loop
	loopBB := start.ControlFlowGraph().NewBasicBlock3(intsrv.TypeLoop, start.FromOffset(), start.ToOffset())

	// Update predecessors
	startPredecessorIterator := start.Predecessors().Iterator()

	for startPredecessorIterator.HasNext() {
		predecessor := startPredecessorIterator.Next()

		if !members.Contains(predecessor) {
			predecessor.Replace(start, loopBB)
			loopBB.Predecessors().Add(predecessor)
			_ = startPredecessorIterator.Remove()
		}
	}

	loopBB.SetSub1(start)

	// Set LOOP_START, LOOP_END and TYPE_JUMP
	for _, member := range members.ToSlice() {
		if member.MatchType(intsrv.GroupSingleSuccessor) {
			bb := member.Next()

			if bb == start {
				member.SetNext(cfg.LoopStart)
			} else if bb == end {
				member.SetNext(cfg.LoopEnd)
			} else if !members.Contains(bb) && (bb.Predecessors().Size() > 1) {
				member.SetNext(newJumpBasicBlockGraphLoopReducer(member, bb))
			}
		} else if member.Type() == intsrv.TypeConditionalBranch {
			bb := member.Next()

			if bb == start {
				member.SetNext(cfg.LoopStart)
			} else if bb == end {
				member.SetNext(cfg.LoopEnd)
			} else if !members.Contains(bb) && (bb.Predecessors().Size() > 1) {
				member.SetNext(newJumpBasicBlockGraphLoopReducer(member, bb))
			}

			bb = member.Branch()

			if bb == start {
				member.SetBranch(cfg.LoopStart)
			} else if bb == end {
				member.SetBranch(cfg.LoopEnd)
			} else if !members.Contains(bb) && (bb.Predecessors().Size() > 1) {
				member.SetBranch(newJumpBasicBlockGraphLoopReducer(member, bb))
			}
		} else if member.Type() == intsrv.TypeSwitchDeclaration {
			for _, switchCase := range member.SwitchCases().ToSlice() {
				bb := switchCase.BasicBlock()

				if bb == start {
					switchCase.SetBasicBlock(cfg.LoopStart)
				} else if bb == end {
					switchCase.SetBasicBlock(cfg.LoopEnd)
				} else if !members.Contains(bb) && (bb.Predecessors().Size() > 1) {
					switchCase.SetBasicBlock(newJumpBasicBlockGraphLoopReducer(member, bb))
				}
			}
		}
		if toOffset < member.ToOffset() {
			toOffset = member.ToOffset()
		}
	}

	if end != nil {
		loopBB.SetNext(end)
		end.ReplaceWithOlds(members, loopBB)
	}

	start.Predecessors().Clear()

	loopBB.SetToOffset(toOffset)

	return loopBB
}

func newJumpBasicBlockGraphLoopReducer(bb, target intsrv.IBasicBlock) intsrv.IBasicBlock {
	predecessors := util.NewSet[intsrv.IBasicBlock]()

	predecessors.Add(bb)
	target.Predecessors().Remove(bb)

	return bb.ControlFlowGraph().NewBasicBlock5(intsrv.TypeJump, bb.FromOffset(), target.FromOffset(), predecessors)
}

func ReduceControlFlowGraphLoopReducer(cfg intsrv.IControlFlowGraph) {
	arrayOfDominatorIndexes := buildDominatorIndexes(cfg)
	loops := IdentifyNaturalLoops(cfg, arrayOfDominatorIndexes)
	loopsLength := loops.Size()

	for i := 0; i < loopsLength; i++ {
		loop := loops.Get(i)
		startBB := loop.Start()
		loopBB := reduceLoop(loop)

		// Update other loops
		for j := loopsLength - 1; j > i; j-- {
			otherLoop := loops.Get(j)

			if otherLoop.Start() == startBB {
				otherLoop.SetStart(loopBB)
			}

			if otherLoop.Members().Contains(startBB) {
				otherLoop.Members().RemoveAll(loop.Members().ToSlice())
				otherLoop.Members().Add(loopBB)
			}

			if otherLoop.End() == startBB {
				otherLoop.SetEnd(loopBB)
			}
		}
	}
}
