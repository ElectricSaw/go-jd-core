package utils

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/model/cfg"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"math"
)

func ReduceGraphGotoReducer(cfg intsrv.IControlFlowGraph) bool {
	start := cfg.Start()
	jsrTargets := util.NewBitSet()
	visited := util.NewBitSetWithSize(cfg.BasicBlocks().Size())

	return ReduceGraphGotoReducer2(visited, start, jsrTargets)
}

func ReduceGraphGotoReducer2(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	if !basicBlock.MatchType(intsrv.GroupEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeStart:
		case intsrv.TypeStatements:
		case intsrv.TypeIf:
		case intsrv.TypeIfElse:
		case intsrv.TypeSwitch:
		case intsrv.TypeTry:
		case intsrv.TypeTryJsr:
		case intsrv.TypeTryEclipse:
		case intsrv.TypeGotoInTernaryOperator:
			return ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets)
		case intsrv.TypeConditionalBranch:
		case intsrv.TypeCondition:
		case intsrv.TypeConditionOr:
		case intsrv.TypeConditionAnd:
		case intsrv.TypeConditionTernaryOperator:
			return reduceConditionalBranch(visited, basicBlock, jsrTargets)
		case intsrv.TypeSwitchDeclaration:
			return reduceSwitchDeclaration(visited, basicBlock, jsrTargets)
		case intsrv.TypeTryDeclaration:
			return reduceTryDeclaration(visited, basicBlock, jsrTargets)
		case intsrv.TypeJsr:
			return reduceJsr(visited, basicBlock, jsrTargets)
		case intsrv.TypeLoop:
			return reduceLoop2(visited, basicBlock, jsrTargets)
		}
	}

	return true
}

func reduceConditionalBranch(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	for aggregateConditionalBranches(basicBlock) {
	}

	// assert basicBlock.MatchType(intsrv.GroupCondition);

	if ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets) && ReduceGraphGotoReducer2(visited, basicBlock.Branch(), jsrTargets) {
		return reduceConditionalBranch2(basicBlock)
	}

	return false
}

func reduceConditionalBranch2(basicBlock intsrv.IBasicBlock) bool {
	next := basicBlock.Next()
	branch := basicBlock.Branch()
	watchdog := NewWatchDog()

	if next == branch {
		// Empty 'if'
		createIf(basicBlock, cfg.End, cfg.End, branch)
		return true
	}

	if next.MatchType(intsrv.GroupEnd) && (next.Predecessors().Size() <= 1) {
		// Create 'if'
		createIf(basicBlock, next, next, branch)
		return true
	}

	if next.MatchType(intsrv.GroupSingleSuccessor|intsrv.TypeReturn|
		intsrv.TypeReturnValue|intsrv.TypeThrow) && (next.Predecessors().Size() == 1) {
		nextLast := next
		nextNext := next.Next()
		cfg1 := next.ControlFlowGraph()
		lineNumber := cfg1.LineNumber(basicBlock.FromOffset())
		maxOffset := branch.FromOffset()

		if (maxOffset == 0) || (next.FromOffset() > branch.FromOffset()) {
			maxOffset = math.MaxInt
		}

		for (nextLast != nextNext) && nextNext.MatchType(intsrv.GroupSingleSuccessor) &&
			(nextNext.Predecessors().Size() == 1) && (cfg1.LineNumber(nextNext.FromOffset()) >= lineNumber) &&
			(nextNext.FromOffset() < maxOffset) {
			watchdog.Check(nextNext, nextNext.Next())
			nextLast = nextNext
			nextNext = nextNext.Next()
		}

		if nextNext == branch {
			createIf(basicBlock, next, nextLast, branch)
			return true
		}

		if nextNext.MatchType(intsrv.GroupEnd) && (nextNext.FromOffset() < maxOffset) {
			createIf(basicBlock, next, nextNext, branch)
			return true
		}

		if branch.MatchType(intsrv.GroupEnd) {
			if (nextNext.FromOffset() < maxOffset) && (nextNext.Predecessors().Size() == 1) {
				createIf(basicBlock, next, nextNext, branch)
			} else {
				createIfElse(intsrv.TypeIfElse, basicBlock, next, nextLast, branch, branch, nextNext)
			}
			return true
		}

		if branch.MatchType(intsrv.GroupSingleSuccessor) && (branch.Predecessors().Size() == 1) {
			branchLast := branch
			branchNext := branch.Next()

			watchdog.Clear()

			for (branchLast != branchNext) &&
				branchNext.MatchType(intsrv.GroupSingleSuccessor) &&
				(branchNext.Predecessors().Size() == 1) &&
				(cfg1.LineNumber(branchNext.FromOffset()) >= lineNumber) {
				watchdog.Check(branchNext, branchNext.Next())
				branchLast = branchNext
				branchNext = branchNext.Next()
			}

			if nextNext == branchNext {
				if nextLast.MatchType(intsrv.TypeGotoInTernaryOperator | intsrv.TypeTernaryOperator) {
					createIfElse(intsrv.TypeTernaryOperator, basicBlock, next, nextLast, branch, branchLast, nextNext)
					return true
				} else {
					createIfElse(intsrv.TypeIfElse, basicBlock, next, nextLast, branch, branchLast, nextNext)
					return true
				}
			} else {
				if (nextNext.FromOffset() < branch.FromOffset()) && (nextNext.Predecessors().Size() == 1) {
					createIf(basicBlock, next, nextNext, branch)
					return true
				} else if (nextNext.FromOffset() > branch.FromOffset()) && branchNext.MatchType(intsrv.GroupEnd) {
					createIfElse(intsrv.TypeIfElse, basicBlock, next, nextLast, branch, branchNext, nextNext)
					return true
				}
			}
		}
	}

	if branch.MatchType(intsrv.GroupSingleSuccessor|intsrv.TypeReturn|intsrv.TypeReturnValue|intsrv.TypeThrow) && (branch.Predecessors().Size() == 1) {
		branchLast := branch
		branchNext := branch.Next()

		watchdog.Clear()

		for (branchLast != branchNext) &&
			branchNext.MatchType(intsrv.GroupSingleSuccessor) &&
			(branchNext.Predecessors().Size() == 1) {
			watchdog.Check(branchNext, branchNext.Next())
			branchLast = branchNext
			branchNext = branchNext.Next()
		}

		if branchNext == next {
			basicBlock.InverseCondition()
			createIf(basicBlock, branch, branchLast, next)
			return true
		}

		if branchNext.MatchType(intsrv.GroupEnd) && (branchNext.Predecessors().Size() <= 1) {
			// Create 'if'
			basicBlock.InverseCondition()
			createIf(basicBlock, branch, branchNext, next)
			return true
		}
	}

	if next.MatchType(intsrv.TypeReturn | intsrv.TypeReturnValue | intsrv.TypeThrow) {
		// Un-optimize byte code
		next = clone(basicBlock, next)
		// Create 'if'
		createIf(basicBlock, next, next, branch)
		return true
	}

	if next.MatchType(intsrv.GroupSingleSuccessor) {
		nextLast := next
		nextNext := next.Next()

		watchdog.Clear()

		for (nextLast != nextNext) &&
			nextNext.MatchType(intsrv.GroupSingleSuccessor) &&
			(nextNext.Predecessors().Size() == 1) {
			watchdog.Check(nextNext, nextNext.Next())
			nextLast = nextNext
			nextNext = nextNext.Next()
		}

		if nextNext.MatchType(intsrv.TypeReturn | intsrv.TypeReturnValue | intsrv.TypeThrow) {
			createIf(basicBlock, next, nextNext, branch)
			return true
		}
	}

	return false
}

func createIf(basicBlock, sub, last, next intsrv.IBasicBlock) {
	condition := basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)

	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)

	toOffset := last.ToOffset()

	if toOffset == 0 {
		toOffset = basicBlock.ToOffset()
	}

	// Split sequence
	last.SetNext(cfg.End)
	next.Predecessors().Remove(last)
	// Create 'if'
	basicBlock.SetType(intsrv.TypeIf)
	basicBlock.SetToOffset(toOffset)
	basicBlock.SetCondition(condition)
	basicBlock.SetSub1(sub)
	basicBlock.SetSub2(nil)
	basicBlock.SetNext(next)
}

func createIfElse(typ int, basicBlock, sub1, last1, sub2, last2, next intsrv.IBasicBlock) {
	condition := basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)

	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)

	toOffset := last2.ToOffset()

	if toOffset == 0 {
		toOffset = last1.ToOffset()

		if toOffset == 0 {
			toOffset = basicBlock.ToOffset()
		}
	}

	// Split sequences
	last1.SetNext(cfg.End)
	next.Predecessors().Remove(last1)
	last2.SetNext(cfg.End)
	next.Predecessors().Remove(last2)
	next.Predecessors().Add(basicBlock)
	// Create 'if-else'
	basicBlock.SetType(typ)
	basicBlock.SetToOffset(toOffset)
	basicBlock.SetCondition(condition)
	basicBlock.SetSub1(sub1)
	basicBlock.SetSub2(sub2)
	basicBlock.SetNext(next)
}

func aggregateConditionalBranches(basicBlock intsrv.IBasicBlock) bool {
	change := false

	next := basicBlock.Next()
	branch := basicBlock.Branch()

	if (next.Type() == intsrv.TypeGotoInTernaryOperator) && (next.Predecessors().Size() == 1) {
		nextNext := next.Next()

		if nextNext.MatchType(intsrv.TypeConditionalBranch | intsrv.TypeCondition) {
			if branch.MatchType(intsrv.TypeStatements|intsrv.TypeGotoInTernaryOperator) &&
				(nextNext == branch.Next()) && (branch.Predecessors().Size() == 1) &&
				(nextNext.Predecessors().Size() == 2) {
				if MinDepth(nextNext) == -1 {
					updateConditionTernaryOperator(basicBlock, nextNext)
					return true
				}

				nextNextNext := nextNext.Next()
				nextNextBranch := nextNext.Branch()

				if (nextNextNext.Type() == intsrv.TypeGotoInTernaryOperator) &&
					(nextNextNext.Predecessors().Size() == 1) {
					nextNextNextNext := nextNextNext.Next()

					if nextNextNextNext.MatchType(intsrv.TypeConditionalBranch | intsrv.TypeCondition) {
						if nextNextBranch.MatchType(intsrv.TypeStatements|intsrv.TypeGotoInTernaryOperator) && (nextNextNextNext == nextNextBranch.Next()) && (nextNextBranch.Predecessors().Size() == 1) && (nextNextNextNext.Predecessors().Size() == 2) {
							if MinDepth(nextNextNextNext) == -2 {
								updateCondition(basicBlock, nextNext, nextNextNextNext)
								return true
							}
						}
					}
				}
			}
			if (nextNext.Next() == branch) && checkJdk118TernaryOperatorPattern(next, nextNext, 153) { // IFEQ
				convertConditionalBranchToGotoInTernaryOperator(basicBlock, next, nextNext)
				return true
			}
			if (nextNext.Branch() == branch) && checkJdk118TernaryOperatorPattern(next, nextNext, 154) { // IFNE
				convertConditionalBranchToGotoInTernaryOperator(basicBlock, next, nextNext)
				return true
			}
			if nextNext.Predecessors().Size() == 1 {
				convertGotoInTernaryOperatorToCondition(next, nextNext)
				return true
			}
		}
	}

	if next.MatchType(intsrv.TypeConditionalBranch | intsrv.GroupCondition) {
		// Test line numbers
		lineNumber1 := basicBlock.LastLineNumber()
		lineNumber2 := next.FirstLineNumber()

		if (lineNumber2 - lineNumber1) <= 1 {
			change = aggregateConditionalBranches(next)

			if next.MatchType(intsrv.TypeConditionalBranch|intsrv.GroupCondition) && (next.Predecessors().Size() == 1) {
				// Aggregate conditional branches
				if next.Next() == branch {
					updateConditionalBranches(basicBlock, createLeftCondition(basicBlock), intsrv.TypeConditionOr, next)
					return true
				} else if next.Branch() == branch {
					updateConditionalBranches(basicBlock, createLeftInverseCondition(basicBlock), intsrv.TypeConditionAnd, next)
					return true
				} else if branch.MatchType(intsrv.TypeConditionalBranch | intsrv.GroupCondition) {
					change = aggregateConditionalBranches(branch)

					if branch.MatchType(intsrv.TypeConditionalBranch | intsrv.GroupCondition) {
						if (next.Next() == branch.Next()) && (next.Branch() == branch.Branch()) {
							updateConditionTernaryOperator2(basicBlock)
							return true
						} else if (next.Branch() == branch.Next()) && (next.Next() == branch.Branch()) {
							updateConditionTernaryOperator2(basicBlock)
							branch.InverseCondition()
							return true
						}
					}
				}
			}
		}
	}

	if branch.MatchType(intsrv.TypeConditionalBranch | intsrv.GroupCondition) {
		// Test line numbers
		lineNumber1 := basicBlock.LastLineNumber()
		lineNumber2 := branch.FirstLineNumber()

		if (lineNumber2 - lineNumber1) <= 1 {
			change = aggregateConditionalBranches(branch)

			if branch.MatchType(intsrv.TypeConditionalBranch|intsrv.GroupCondition) && (branch.Predecessors().Size() == 1) {
				// Aggregate conditional branches
				if branch.Branch() == next {
					updateConditionalBranches(basicBlock, createLeftCondition(basicBlock), intsrv.TypeConditionAnd, branch)
					return true
				} else if branch.Next() == next {
					updateConditionalBranches(basicBlock, createLeftInverseCondition(basicBlock), intsrv.TypeConditionOr, branch)
					return true
				}
			}
		}
	}

	if basicBlock.Type() == intsrv.TypeConditionalBranch {
		basicBlock.SetType(intsrv.TypeCondition)
		return true
	}

	return change
}

func createLeftCondition(basicBlock intsrv.IBasicBlock) intsrv.IBasicBlock {
	if basicBlock.Type() == intsrv.TypeConditionalBranch {
		return basicBlock.ControlFlowGraph().NewBasicBlock4(intsrv.TypeCondition,
			basicBlock.FromOffset(), basicBlock.ToOffset(), false)
	} else {
		left := basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)
		left.InverseCondition()
		return left
	}
}

func createLeftInverseCondition(basicBlock intsrv.IBasicBlock) intsrv.IBasicBlock {
	if basicBlock.Type() == intsrv.TypeConditionalBranch {
		return basicBlock.ControlFlowGraph().NewBasicBlock3(intsrv.TypeCondition, basicBlock.FromOffset(), basicBlock.ToOffset())
	} else {
		return basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)
	}
}

func updateConditionalBranches(basicBlock, leftBasicBlock intsrv.IBasicBlock, operator int, subBasicBlock intsrv.IBasicBlock) {
	basicBlock.SetType(operator)
	basicBlock.SetToOffset(subBasicBlock.ToOffset())
	basicBlock.SetNext(subBasicBlock.Next())
	basicBlock.SetBranch(subBasicBlock.Branch())
	basicBlock.SetCondition(cfg.End)
	basicBlock.SetSub1(leftBasicBlock)
	basicBlock.SetSub2(subBasicBlock)

	subBasicBlock.Next().Replace(subBasicBlock, basicBlock)
	subBasicBlock.Branch().Replace(subBasicBlock, basicBlock)
}

func updateConditionTernaryOperator(basicBlock, nextNext intsrv.IBasicBlock) {
	fromOffset := nextNext.FromOffset()
	toOffset := nextNext.ToOffset()
	next := nextNext.Next()
	branch := nextNext.Branch()

	if basicBlock.Type() == intsrv.TypeConditionalBranch {
		basicBlock.SetType(intsrv.TypeCondition)
	}
	if (nextNext.Type() == intsrv.TypeCondition) && !nextNext.IsInverseCondition() {
		basicBlock.InverseCondition()
	}

	condition := nextNext

	condition.SetType(basicBlock.Type())
	condition.SetFromOffset(basicBlock.FromOffset())
	condition.SetToOffset(basicBlock.ToOffset())
	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)
	condition.SetCondition(basicBlock.Condition())
	condition.SetSub1(basicBlock.Sub1())
	condition.SetSub2(basicBlock.Sub2())
	condition.Predecessors().Clear()

	basicBlock.SetType(intsrv.TypeConditionTernaryOperator)
	basicBlock.SetFromOffset(fromOffset)
	basicBlock.SetToOffset(toOffset)
	basicBlock.SetCondition(condition)
	basicBlock.SetSub1(basicBlock.Next())
	basicBlock.SetSub2(basicBlock.Branch())
	basicBlock.SetNext(next)
	basicBlock.SetBranch(branch)
	basicBlock.Sub1().SetNext(cfg.End)
	basicBlock.Sub2().SetNext(cfg.End)

	next.Replace(nextNext, basicBlock)
	branch.Replace(nextNext, basicBlock)

	basicBlock.Sub1().Predecessors().Clear()
	basicBlock.Sub2().Predecessors().Clear()
}

func updateCondition(basicBlock, nextNext, nextNextNextNext intsrv.IBasicBlock) {
	fromOffset := nextNextNextNext.FromOffset()
	toOffset := nextNextNextNext.ToOffset()
	next := nextNextNextNext.Next()
	branch := nextNextNextNext.Branch()

	condition := basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)
	condition.SetType(intsrv.TypeCondition)

	basicBlock.Next().SetNext(cfg.End)
	basicBlock.Next().Predecessors().Clear()
	basicBlock.Branch().SetNext(cfg.End)
	basicBlock.Branch().Predecessors().Clear()

	nextNextNextNext.SetType(intsrv.TypeConditionTernaryOperator)
	nextNextNextNext.SetFromOffset(condition.ToOffset())
	nextNextNextNext.SetToOffset(condition.ToOffset())
	nextNextNextNext.SetCondition(condition)
	nextNextNextNext.SetSub1(basicBlock.Next())
	nextNextNextNext.SetSub2(basicBlock.Branch())
	nextNextNextNext.SetNext(cfg.End)
	nextNextNextNext.SetBranch(cfg.End)
	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)

	condition = nextNext.ControlFlowGraph().NewBasicBlock1(nextNext)
	condition.SetType(intsrv.TypeCondition)

	nextNext.Next().SetNext(cfg.End)
	nextNext.Next().Predecessors().Clear()
	nextNext.Branch().SetNext(cfg.End)
	nextNext.Branch().Predecessors().Clear()

	nextNext.SetType(intsrv.TypeConditionTernaryOperator)
	nextNext.SetFromOffset(condition.ToOffset())
	nextNext.SetToOffset(condition.ToOffset())
	nextNext.SetCondition(condition)
	nextNext.SetSub1(nextNext.Next())
	nextNext.SetSub2(nextNext.Branch())
	nextNext.SetNext(cfg.End)
	nextNext.SetBranch(cfg.End)
	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)

	basicBlock.SetType(intsrv.TypeCondition)
	basicBlock.SetFromOffset(fromOffset)
	basicBlock.SetToOffset(toOffset)
	basicBlock.SetSub1(nextNextNextNext)
	basicBlock.SetSub2(nextNext)
	basicBlock.SetNext(next)
	basicBlock.SetBranch(branch)

	next.Replace(nextNextNextNext, basicBlock)
	branch.Replace(nextNextNextNext, basicBlock)
}

func updateConditionTernaryOperator2(basicBlock intsrv.IBasicBlock) {
	next := basicBlock.Next()
	branch := basicBlock.Branch()

	cfg1 := basicBlock.ControlFlowGraph()
	condition := cfg1.NewBasicBlock3(intsrv.TypeCondition, basicBlock.FromOffset(), basicBlock.ToOffset())

	condition.SetNext(cfg.End)
	condition.SetBranch(cfg.End)

	basicBlock.SetType(intsrv.TypeConditionTernaryOperator)
	basicBlock.SetToOffset(basicBlock.FromOffset())
	basicBlock.SetCondition(condition)
	basicBlock.SetSub1(next)
	basicBlock.SetSub2(branch)
	basicBlock.SetNext(next.Next())
	basicBlock.SetBranch(next.Branch())

	next.Next().Replace(next, basicBlock)
	next.Branch().Replace(next, basicBlock)
	branch.Next().Replace(branch, basicBlock)
	branch.Branch().Replace(branch, basicBlock)

	next.Predecessors().Clear()
	branch.Predecessors().Clear()
}

func convertGotoInTernaryOperatorToCondition(basicBlock, next intsrv.IBasicBlock) {
	basicBlock.SetType(intsrv.TypeCondition)
	basicBlock.SetNext(next.Next())
	basicBlock.SetBranch(next.Branch())

	next.Next().Replace(next, basicBlock)
	next.Branch().Replace(next, basicBlock)

	next.SetType(intsrv.TypeDeleted)
}

func convertConditionalBranchToGotoInTernaryOperator(basicBlock, next, nextNext intsrv.IBasicBlock) {
	basicBlock.SetType(intsrv.TypeGotoInTernaryOperator)
	basicBlock.SetNext(nextNext)
	basicBlock.Branch().Predecessors().Remove(basicBlock)
	basicBlock.SetBranch(cfg.End)
	basicBlock.SetInverseCondition(false)

	nextNext.Replace(next, basicBlock)

	next.SetType(intsrv.TypeDeleted)
}

func checkJdk118TernaryOperatorPattern(next, nextNext intsrv.IBasicBlock, ifByteCode int) bool {
	if (nextNext.ToOffset() - nextNext.FromOffset()) == 3 {
		code := next.ControlFlowGraph().Method().Attribute("Code").(intcls.IAttributeCode).Code()
		nextFromOffset := next.FromOffset()
		nextNextFromOffset := nextNext.FromOffset()
		return (code[nextFromOffset] == 3) && // ICONST_0
			(((code[nextFromOffset+1] & 255) == 167) || ((code[nextFromOffset+1] & 255) == 200)) && // GOTO or GOTO_W
			(int(code[nextNextFromOffset]&255) == ifByteCode) && // IFEQ or IFNE
			(nextNextFromOffset+3 == nextNext.ToOffset())
	}

	return false
}

func reduceSwitchDeclaration(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	var defaultSC, lastSC intsrv.ISwitchCase
	maxOffset := -1

	for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
		if maxOffset < switchCase.Offset() {
			maxOffset = switchCase.Offset()
		}

		if switchCase.IsDefaultCase() {
			defaultSC = switchCase
		} else {
			lastSC = switchCase
		}
	}

	if lastSC == nil {
		lastSC = defaultSC
	}

	var lastSwitchCaseBasicBlock intsrv.IBasicBlock
	v := util.NewBitSet()
	ends := util.NewSet[intsrv.IBasicBlock]()

	for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
		bb := switchCase.BasicBlock()

		if switchCase.Offset() == maxOffset {
			lastSwitchCaseBasicBlock = bb
		} else {
			visit(v, bb, maxOffset, ends)
		}
	}

	end := cfg.End

	for _, bb := range ends.ToSlice() {
		if (end == cfg.End) || (end.FromOffset() < bb.FromOffset()) {
			end = bb
		}
	}

	if end == cfg.End {
		if (lastSC.BasicBlock() == lastSwitchCaseBasicBlock) && searchLoopStart(basicBlock, maxOffset) {
			replaceLoopStartWithSwitchBreak(util.NewBitSet(), basicBlock)
			end = cfg.LoopStart
			defaultSC.SetBasicBlock(end)
		} else {
			end = lastSwitchCaseBasicBlock
		}
	} else {
		visit(v, lastSwitchCaseBasicBlock, end.FromOffset(), ends)
	}

	endPredecessors := end.Predecessors()
	endPredecessorIterator := endPredecessors.Iterator()

	for endPredecessorIterator.HasNext() {
		endPredecessor := endPredecessorIterator.Next()

		if v.Get(endPredecessor.Index()) {
			endPredecessor.Replace(end, cfg.SwitchBreak)
			_ = endPredecessorIterator.Remove()
		}
	}

	if defaultSC.BasicBlock() == end {
		iterator := basicBlock.SwitchCases().Iterator()
		for iterator.HasNext() {
			if iterator.Next().BasicBlock() == end {
				_ = iterator.Remove()
			}
		}
	} else {
		for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
			if switchCase.BasicBlock() == end {
				switchCase.SetBasicBlock(cfg.SwitchBreak)
			}
		}
	}

	reduced := true

	for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
		reduced = reduced && ReduceGraphGotoReducer2(visited, switchCase.BasicBlock(), jsrTargets)
	}

	for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
		bb := switchCase.BasicBlock()

		// assert bb != end;

		predecessors := bb.Predecessors()

		if predecessors.Size() > 1 {
			predecessorIterator := predecessors.Iterator()

			for predecessorIterator.HasNext() {
				predecessor := predecessorIterator.Next()

				if predecessor != basicBlock {
					predecessor.Replace(bb, cfg.End)
					_ = predecessorIterator.Remove()
				}
			}
		}
	}

	// Change type
	basicBlock.SetType(intsrv.TypeSwitch)
	basicBlock.SetNext(cfg.End)
	endPredecessors.Add(basicBlock)

	return reduced && ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets)
}

func searchLoopStart(basicBlock intsrv.IBasicBlock, maxOffset int) bool {
	watchdog := NewWatchDog()

	for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
		bb := switchCase.BasicBlock()

		watchdog.Clear()

		for bb.FromOffset() < maxOffset {
			if bb == cfg.LoopStart {
				return true
			}

			if bb.MatchType(intsrv.GroupEnd | intsrv.GroupCondition) {
				break
			}

			var next intsrv.IBasicBlock

			if bb.MatchType(intsrv.GroupSingleSuccessor) {
				next = bb.Next()
			} else if bb.Type() == intsrv.TypeConditionalBranch {
				next = bb.Branch()
			} else if bb.Type() == intsrv.TypeSwitchDeclaration {
				maximum := bb.FromOffset()

				for _, sc := range bb.SwitchCases().ToSlice() {
					if maximum < sc.BasicBlock().FromOffset() {
						next = sc.BasicBlock()
						maximum = next.FromOffset()
					}
				}
			}

			if bb == next {
				break
			}

			watchdog.Check(bb, next)
			bb = next
		}
	}

	return false
}

func reduceTryDeclaration(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	reduced := true
	var finallyBB intsrv.IBasicBlock = nil

	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		if exceptionHandler.InternalThrowableName() == "" {
			reduced = ReduceGraphGotoReducer2(visited, exceptionHandler.BasicBlock(), jsrTargets)
			finallyBB = exceptionHandler.BasicBlock()
			break
		}
	}

	jsrTarget := searchJsrTarget(basicBlock, jsrTargets)

	reduced = reduced && ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets)

	tryBB := basicBlock.Next()

	if tryBB.MatchType(intsrv.GroupSynthetic) {
		return false
	}

	maxOffset := basicBlock.FromOffset()
	tryWithResourcesFlag := true
	var tryWithResourcesBB intsrv.IBasicBlock

	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		if exceptionHandler.InternalThrowableName() != "" {
			reduced = reduced && ReduceGraphGotoReducer2(visited, exceptionHandler.BasicBlock(), jsrTargets)
		}

		bb := exceptionHandler.BasicBlock()

		if bb.MatchType(intsrv.GroupSynthetic) {
			return false
		}

		if maxOffset < bb.FromOffset() {
			maxOffset = bb.FromOffset()
		}

		if tryWithResourcesFlag {
			predecessors := bb.Predecessors()

			if predecessors.Size() == 1 {
				tryWithResourcesFlag = false
			} else {
				// assert predecessors.Size() == 2;

				if tryWithResourcesBB == nil {
					for _, predecessor := range predecessors.ToSlice() {
						if predecessor != basicBlock {
							// assert predecessor.Type() == intsrv.TypeTryDeclaration;
							tryWithResourcesBB = predecessor
							break
						}
					}
				} else if !predecessors.Contains(tryWithResourcesBB) {
					tryWithResourcesFlag = false
				}
			}
		}
	}

	if tryWithResourcesFlag {
		// One of 'try-with-resources' patterns
		for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
			exceptionHandler.BasicBlock().Predecessors().Remove(basicBlock)
		}
		for _, predecessor := range basicBlock.Predecessors().ToSlice() {
			predecessor.Replace(basicBlock, tryBB)
			tryBB.Replace(basicBlock, predecessor)
		}
		basicBlock.SetType(intsrv.TypeDeleted)
	} else if reduced {
		end := searchEndBlock(basicBlock, maxOffset)

		updateBlock(tryBB, end, maxOffset)

		if (finallyBB != nil) && (basicBlock.ExceptionHandlers().Size() == 1) &&
			(tryBB.Type() == intsrv.TypeTry) && (tryBB.Next() == cfg.End) &&
			(basicBlock.FromOffset() == tryBB.FromOffset()) && !containsFinally(tryBB) {
			// Merge inner try
			_ = basicBlock.ExceptionHandlers().AddAllAt(0, tryBB.ExceptionHandlers().ToSlice())

			for _, exceptionHandler := range tryBB.ExceptionHandlers().ToSlice() {
				predecessors := exceptionHandler.BasicBlock().Predecessors()
				predecessors.Clear()
				predecessors.Add(basicBlock)
			}

			tryBB.SetType(intsrv.TypeDeleted)
			tryBB = tryBB.Sub1()
			predecessors := tryBB.Predecessors()
			predecessors.Clear()
			predecessors.Add(basicBlock)
		}

		// Update blocks
		toOffset := maxOffset
		for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
			bb := exceptionHandler.BasicBlock()

			if bb == end {
				exceptionHandler.SetBasicBlock(cfg.End)
			} else {
				offset := maxOffset
				if bb.FromOffset() == maxOffset {
					offset = end.FromOffset()
				}

				if offset == 0 {
					offset = math.MaxInt
				}

				last := updateBlock(bb, end, offset)

				if toOffset < last.ToOffset() {
					toOffset = last.ToOffset()
				}
			}
		}

		basicBlock.SetSub1(tryBB)
		basicBlock.SetNext(cfg.End)
		end.Predecessors().Add(basicBlock)

		if jsrTarget == nil {
			// Change type
			if (finallyBB != nil) && checkEclipseFinallyPattern(basicBlock, finallyBB, maxOffset) {
				basicBlock.SetType(intsrv.TypeTryEclipse)
			} else {
				basicBlock.SetType(intsrv.TypeTry)
			}
		} else {
			// Change type
			basicBlock.SetType(intsrv.TypeTryJsr)
			// Merge 1.1 to 1.4 sub try block
			removeJsrAndMergeSubTry(basicBlock)
		}

		basicBlock.SetToOffset(toOffset)
	}

	return reduced
}

func containsFinally(basicBlock intsrv.IBasicBlock) bool {
	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		if exceptionHandler.InternalThrowableName() == "" {
			return true
		}
	}

	return false
}

func checkEclipseFinallyPattern(basicBlock, finallyBB intsrv.IBasicBlock, maxOffset int) bool {
	nextOpcode := SearchNextOpcode(basicBlock, maxOffset)

	if (nextOpcode == 0) ||
		(nextOpcode == 167) || // GOTO
		(nextOpcode == 200) { // GOTO_W
		return true
	}

	next := basicBlock.Next()

	if !next.MatchType(intsrv.GroupEnd) && (finallyBB.FromOffset() < next.FromOffset()) {
		cfg1 := finallyBB.ControlFlowGraph()
		toLineNumber := cfg1.LineNumber(finallyBB.ToOffset() - 1)
		fromLineNumber := cfg1.LineNumber(next.FromOffset())

		if fromLineNumber < toLineNumber {
			return true
		}
	}

	return false
}

func searchJsrTarget(basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) intsrv.IBasicBlock {
	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		if exceptionHandler.InternalThrowableName() == "" {
			bb := exceptionHandler.BasicBlock()

			if bb.Type() == intsrv.TypeStatements {
				bb = bb.Next()

				if (bb.Type() == intsrv.TypeJsr) && (bb.Next().Type() == intsrv.TypeThrow) {
					// Java 1.1 to 1.4 finally pattern found
					jsrTarget := bb.Branch()
					jsrTargets.Set(jsrTarget.Index())
					return jsrTarget
				}
			}
		}
	}

	return nil
}

func searchEndBlock(basicBlock intsrv.IBasicBlock, maxOffset int) intsrv.IBasicBlock {
	var end intsrv.IBasicBlock
	last := splitSequence(basicBlock.Next(), maxOffset)

	if !last.MatchType(intsrv.GroupEnd) {
		next := last.Next()

		if (next.FromOffset() >= maxOffset) ||
			(!next.MatchType(intsrv.TypeEnd|intsrv.TypeReturn|intsrv.TypeSwitch|
				intsrv.TypeLoopStart|intsrv.TypeLoopContinue|intsrv.TypeLoopEnd) &&
				(next.ToOffset() < basicBlock.FromOffset())) {
			return next
		}

		end = next
	}

	for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
		bb := exceptionHandler.BasicBlock()

		if bb.FromOffset() < maxOffset {
			last = splitSequence(bb, maxOffset)

			if !last.MatchType(intsrv.GroupEnd) {
				next := last.Next()

				if (next.FromOffset() >= maxOffset) ||
					(!next.MatchType(intsrv.TypeEnd|intsrv.TypeReturn|intsrv.TypeSwitch|
						intsrv.TypeLoopStart|intsrv.TypeLoopContinue|intsrv.TypeLoopEnd) &&
						(next.ToOffset() < basicBlock.FromOffset())) {
					return next
				}

				if end == nil {
					end = next
				} else if end != next {
					end = cfg.End
				}
			}
		} else {
			// Last handler block
			cfg1 := bb.ControlFlowGraph()
			lineNumber := cfg1.LineNumber(bb.FromOffset())
			watchdog := NewWatchDog()
			next := bb.Next()

			last = bb

			for (last != next) && last.MatchType(intsrv.GroupSingleSuccessor) &&
				(next.Predecessors().Size() == 1) && (lineNumber <= cfg1.LineNumber(next.FromOffset())) {
				watchdog.Check(next, next.Next())
				last = next
				next = next.Next()
			}

			if !last.MatchType(intsrv.GroupEnd) {
				if (last != next) && ((next.Predecessors().Size() > 1) || !next.MatchType(intsrv.GroupEnd)) {
					return next
				}

				if (end != next) && (exceptionHandler.InternalThrowableName() != "") {
					end = cfg.End
				}
			}
		}
	}

	if (end != nil) && end.MatchType(intsrv.TypeSwitch|intsrv.TypeLoopStart|intsrv.TypeLoopContinue|intsrv.TypeLoopEnd) {
		return end
	}

	return cfg.End
}

func splitSequence(basicBlock intsrv.IBasicBlock, maxOffset int) intsrv.IBasicBlock {
	next := basicBlock.Next()
	watchdog := NewWatchDog()

	for (next.FromOffset() < maxOffset) && next.MatchType(intsrv.GroupSingleSuccessor) {
		watchdog.Check(next, next.Next())
		basicBlock = next
		next = next.Next()
	}

	if (basicBlock.ToOffset() > maxOffset) && (basicBlock.Type() == intsrv.TypeTry) {
		// Split last try block
		exceptionHandlers := basicBlock.ExceptionHandlers()
		bb := exceptionHandlers.Get(exceptionHandlers.Size() - 1).BasicBlock()
		last := splitSequence(bb, maxOffset)

		next = last.Next()
		last.SetNext(cfg.End)

		basicBlock.SetToOffset(last.ToOffset())
		basicBlock.SetNext(next)

		next.Predecessors().Remove(last)
		next.Predecessors().Add(basicBlock)
	}

	return basicBlock
}

func updateBlock(basicBlock, end intsrv.IBasicBlock, maxOffset int) intsrv.IBasicBlock {
	watchdog := NewWatchDog()

	for basicBlock.MatchType(intsrv.GroupSingleSuccessor) {
		watchdog.Check(basicBlock, basicBlock.Next())
		next := basicBlock.Next()

		if (next == end) || (next.FromOffset() > maxOffset) {
			next.Predecessors().Remove(basicBlock)
			basicBlock.SetNext(cfg.End)
			break
		}

		basicBlock = next
	}

	return basicBlock
}

func removeJsrAndMergeSubTry(basicBlock intsrv.IBasicBlock) {
	if basicBlock.ExceptionHandlers().Size() == 1 {
		subTry := basicBlock.Sub1()

		if subTry.MatchType(intsrv.TypeTry | intsrv.TypeTryJsr | intsrv.TypeTryEclipse) {
			for _, exceptionHandler := range subTry.ExceptionHandlers().ToSlice() {
				if exceptionHandler.InternalThrowableName() == "" {
					return
				}
			}

			// Append 'catch' handlers
			for _, exceptionHandler := range subTry.ExceptionHandlers().ToSlice() {
				bb := exceptionHandler.BasicBlock()
				basicBlock.AddExceptionHandler(exceptionHandler.InternalThrowableName(), bb)
				bb.Replace(subTry, basicBlock)
			}

			// Move 'try' clause to parent 'try' block
			basicBlock.SetSub1(subTry.Sub1())
			subTry.Sub1().Replace(subTry, basicBlock)
		}
	}
}

func reduceJsr(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	branch := basicBlock.Branch()
	reduced := ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets) && ReduceGraphGotoReducer2(visited, branch, jsrTargets)

	if (branch.Index() >= 0) && jsrTargets.Get(branch.Index()) {
		// ReduceGraphGotoReducer JSR
		delta := basicBlock.ToOffset() - basicBlock.FromOffset()

		if delta > 3 {
			opcode := LastOpcode(basicBlock)

			if opcode == 168 { // JSR
				basicBlock.SetType(intsrv.TypeStatements)
				basicBlock.SetToOffset(basicBlock.ToOffset() - 3)
				branch.Predecessors().Remove(basicBlock)
				return true
			} else if delta > 5 { // JSR_W
				basicBlock.SetType(intsrv.TypeStatements)
				basicBlock.SetToOffset(basicBlock.ToOffset() - 5)
				branch.Predecessors().Remove(basicBlock)
				return true
			}
		}

		// Delete JSR
		basicBlock.SetType(intsrv.TypeDeleted)
		branch.Predecessors().Remove(basicBlock)
		nextPredecessors := basicBlock.Next().Predecessors()
		nextPredecessors.Remove(basicBlock)

		for _, predecessor := range basicBlock.Predecessors().ToSlice() {
			predecessor.Replace(basicBlock, basicBlock.Next())
			nextPredecessors.Add(predecessor)
		}

		return true
	}

	if basicBlock.Branch().Predecessors().Size() > 1 {
		// Aggregate JSR
		next := basicBlock.Next()
		iterator := basicBlock.Branch().Predecessors().Iterator()

		for iterator.HasNext() {
			predecessor := iterator.Next()

			if (predecessor != basicBlock) &&
				(predecessor.Type() == intsrv.TypeJsr) &&
				(predecessor.Next() == next) {
				for _, predecessorPredecessor := range predecessor.Predecessors().ToSlice() {
					predecessorPredecessor.Replace(predecessor, basicBlock)
					basicBlock.Predecessors().Add(predecessorPredecessor)
				}
				next.Predecessors().Remove(predecessor)
				_ = iterator.Remove()
				reduced = true
			}
		}
	}

	return reduced
}

func reduceLoop2(visited util.IBitSet, basicBlock intsrv.IBasicBlock, jsrTargets util.IBitSet) bool {
	clone1 := visited.Clone()
	reduced := ReduceGraphGotoReducer2(visited, basicBlock.Sub1(), jsrTargets)

	if reduced == false {
		visitedMembers := util.NewBitSet()
		updateBasicBlock := searchUpdateBlockAndCreateContinueLoop(visitedMembers, basicBlock.Sub1())

		visited = clone1.Clone()
		reduced = ReduceGraphGotoReducer2(visited, basicBlock.Sub1(), jsrTargets)

		if updateBasicBlock != nil {
			removeLastContinueLoop(basicBlock.Sub1().Sub1())

			ifBasicBlock := basicBlock.ControlFlowGraph().NewBasicBlock3(intsrv.TypeIf, basicBlock.Sub1().FromOffset(), basicBlock.ToOffset())

			ifBasicBlock.SetCondition(cfg.End)
			ifBasicBlock.SetSub1(basicBlock.Sub1())
			ifBasicBlock.SetNext(updateBasicBlock)
			updateBasicBlock.Predecessors().Add(ifBasicBlock)
			basicBlock.SetSub1(ifBasicBlock)
		}

		if reduced == false {
			visitedMembers.ClearAll()

			conditionalBranch := getLastConditionalBranch(visitedMembers, basicBlock.Sub1())

			if (conditionalBranch != nil) && (conditionalBranch.Next() == cfg.LoopStart) {
				visitedMembers.ClearAll()
				visitedMembers.Set(conditionalBranch.Index())
				changeEndLoopToJump(visitedMembers, basicBlock.Next(), basicBlock.Sub1())

				newLoopBB := basicBlock.ControlFlowGraph().NewBasicBlock1(basicBlock)
				predecessors := conditionalBranch.Predecessors()

				for _, predecessor := range predecessors.ToSlice() {
					predecessor.Replace(conditionalBranch, cfg.LoopEnd)
				}

				newLoopBB.SetNext(conditionalBranch)
				predecessors.Clear()
				predecessors.Add(newLoopBB)
				basicBlock.SetSub1(newLoopBB)

				visitedMembers.ClearAll()
				reduced = ReduceGraphGotoReducer2(visitedMembers, newLoopBB, jsrTargets)
			}
		}
	}

	return reduced && ReduceGraphGotoReducer2(visited, basicBlock.Next(), jsrTargets)
}

func removeLastContinueLoop(basicBlock intsrv.IBasicBlock) {
	visited := util.NewBitSet()
	next := basicBlock.Next()

	for !next.MatchType(intsrv.GroupEnd) && (visited.Get(next.Index()) == false) {
		visited.Set(next.Index())
		basicBlock = next
		next = basicBlock.Next()
	}

	if next == cfg.LoopContinue {
		basicBlock.SetNext(cfg.End)
	}
}

func getLastConditionalBranch(visited util.IBitSet, basicBlock intsrv.IBasicBlock) intsrv.IBasicBlock {
	if !basicBlock.MatchType(intsrv.GroupEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeSwitchDeclaration,
			intsrv.TypeTryDeclaration, intsrv.TypeJsr, intsrv.TypeLoop, intsrv.TypeIfElse,
			intsrv.TypeSwitch, intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			return getLastConditionalBranch(visited, basicBlock.Next())
		case intsrv.TypeIf, intsrv.TypeConditionalBranch, intsrv.TypeCondition,
			intsrv.TypeConditionOr, intsrv.TypeConditionAnd:
			bb := getLastConditionalBranch(visited, basicBlock.Branch())
			if bb != nil {
				return bb
			}
			bb = getLastConditionalBranch(visited, basicBlock.Next())
			if bb != nil {
				return bb
			}
			return basicBlock
		}
	}

	return nil
}

func visit(visited util.IBitSet, basicBlock intsrv.IBasicBlock, maxOffset int, ends util.ISet[intsrv.IBasicBlock]) {
	if basicBlock.FromOffset() >= maxOffset {
		ends.Add(basicBlock)
	} else if (basicBlock.Index() >= 0) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeConditionalBranch, intsrv.TypeJsr, intsrv.TypeCondition:
			visit(visited, basicBlock.Branch(), maxOffset, ends)
			fallthrough
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeGoto,
			intsrv.TypeGotoInTernaryOperator, intsrv.TypeLoop:
			visit(visited, basicBlock.Next(), maxOffset, ends)
			break
		case intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			visit(visited, basicBlock.Sub1(), maxOffset, ends)
			fallthrough
		case intsrv.TypeTryDeclaration:
			for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
				visit(visited, exceptionHandler.BasicBlock(), maxOffset, ends)
			}
			visit(visited, basicBlock.Next(), maxOffset, ends)
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			visit(visited, basicBlock.Sub2(), maxOffset, ends)
			fallthrough
		case intsrv.TypeIf:
			visit(visited, basicBlock.Sub1(), maxOffset, ends)
			visit(visited, basicBlock.Next(), maxOffset, ends)
			break
		case intsrv.TypeConditionOr, intsrv.TypeConditionAnd:
			visit(visited, basicBlock.Sub1(), maxOffset, ends)
			visit(visited, basicBlock.Sub2(), maxOffset, ends)
			break
		case intsrv.TypeSwitch:
			visit(visited, basicBlock.Next(), maxOffset, ends)
			fallthrough
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
				visit(visited, switchCase.BasicBlock(), maxOffset, ends)
			}
			break
		}
	}
}

func replaceLoopStartWithSwitchBreak(visited util.IBitSet, basicBlock intsrv.IBasicBlock) {
	if !basicBlock.MatchType(intsrv.GroupEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())
		basicBlock.Replace(cfg.LoopStart, cfg.SwitchBreak)

		switch basicBlock.Type() {
		case intsrv.TypeConditionalBranch, intsrv.TypeJsr, intsrv.TypeCondition:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Branch())
			fallthrough
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeGoto,
			intsrv.TypeGotoInTernaryOperator, intsrv.TypeLoop:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Next())
			break
		case intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Sub1())
			fallthrough
		case intsrv.TypeTryDeclaration:
			for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
				replaceLoopStartWithSwitchBreak(visited, exceptionHandler.BasicBlock())
			}
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Sub2())
			fallthrough
		case intsrv.TypeIf:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Sub1())
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Next())
			break
		case intsrv.TypeConditionOr, intsrv.TypeConditionAnd:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Sub1())
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Sub2())
			break
		case intsrv.TypeSwitch:
			replaceLoopStartWithSwitchBreak(visited, basicBlock.Next())
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
				replaceLoopStartWithSwitchBreak(visited, switchCase.BasicBlock())
			}
			break
		}
	}
}

func searchUpdateBlockAndCreateContinueLoop(visited util.IBitSet, basicBlock intsrv.IBasicBlock) intsrv.IBasicBlock {
	var updateBasicBlock intsrv.IBasicBlock = nil

	if !basicBlock.MatchType(intsrv.GroupEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeConditionalBranch, intsrv.TypeJsr, intsrv.TypeCondition,
			intsrv.TypeConditionTernaryOperator:
			updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Branch())
			fallthrough
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeGoto,
			intsrv.TypeGotoInTernaryOperator, intsrv.TypeLoop:
			if updateBasicBlock == nil {
				updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Next())
			}
			break
		case intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Sub1())
			fallthrough
		case intsrv.TypeTryDeclaration:
			for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
				if updateBasicBlock == nil {
					updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, exceptionHandler.BasicBlock())
				}
			}
			if updateBasicBlock == nil {
				updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Next())
			}
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Sub2())
			fallthrough
		case intsrv.TypeIf:
			if updateBasicBlock == nil {
				updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Sub1())
			}
			if updateBasicBlock == nil {
				updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Next())
			}
			break
		case intsrv.TypeConditionOr, intsrv.TypeConditionAnd:
			updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Sub1())
			if updateBasicBlock == nil {
				updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Sub2())
			}
			break
		case intsrv.TypeSwitch:
			updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, basicBlock.Next())
			fallthrough
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
				if updateBasicBlock == nil {
					updateBasicBlock = searchUpdateBlockAndCreateContinueLoop2(visited, basicBlock, switchCase.BasicBlock())
				}
			}
			break
		}
	}

	return updateBasicBlock
}

func searchUpdateBlockAndCreateContinueLoop2(visited util.IBitSet, basicBlock, subBasicBlock intsrv.IBasicBlock) intsrv.IBasicBlock {
	if subBasicBlock != nil {
		if basicBlock.FromOffset() < subBasicBlock.FromOffset() {

			if basicBlock.FirstLineNumber() == intmod.UnknownLineNumber {
				if subBasicBlock.MatchType(intsrv.GroupSingleSuccessor) &&
					(subBasicBlock.Next().Type() == intsrv.TypeLoopStart) {
					stackDepth := EvalStackDepth(subBasicBlock)

					for stackDepth != 0 {
						predecessors := subBasicBlock.Predecessors()
						if predecessors.Size() != 1 {
							break
						}

						subBasicBlock = predecessors.Iterator().Next()
						stackDepth += EvalStackDepth(subBasicBlock)
					}

					removePredecessors(subBasicBlock)
					return subBasicBlock
				}
			} else if basicBlock.FirstLineNumber() > subBasicBlock.FirstLineNumber() {
				removePredecessors(subBasicBlock)
				return subBasicBlock
			}
		}

		return searchUpdateBlockAndCreateContinueLoop(visited, subBasicBlock)
	}

	return nil
}

func removePredecessors(basicBlock intsrv.IBasicBlock) {
	predecessors := basicBlock.Predecessors()
	iterator := predecessors.Iterator()

	for iterator.HasNext() {
		iterator.Next().Replace(basicBlock, cfg.LoopContinue)
	}

	predecessors.Clear()
}

func changeEndLoopToJump(visited util.IBitSet, target, basicBlock intsrv.IBasicBlock) {
	if !basicBlock.MatchType(intsrv.GroupEnd) && (visited.Get(basicBlock.Index()) == false) {
		visited.Set(basicBlock.Index())

		switch basicBlock.Type() {
		case intsrv.TypeConditionalBranch, intsrv.TypeJsr, intsrv.TypeCondition:
			if basicBlock.Branch() == cfg.LoopEnd {
				basicBlock.SetBranch(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Branch())
			}
			fallthrough
		case intsrv.TypeStart, intsrv.TypeStatements, intsrv.TypeGoto,
			intsrv.TypeGotoInTernaryOperator, intsrv.TypeLoop:
			if basicBlock.Next() == cfg.LoopEnd {
				basicBlock.SetNext(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Next())
			}
			break
		case intsrv.TypeTry, intsrv.TypeTryJsr, intsrv.TypeTryEclipse:
			if basicBlock.Sub1() == cfg.LoopEnd {
				basicBlock.SetSub1(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Sub1())
			}
			fallthrough
		case intsrv.TypeTryDeclaration:
			for _, exceptionHandler := range basicBlock.ExceptionHandlers().ToSlice() {
				if exceptionHandler.BasicBlock() == cfg.LoopEnd {
					exceptionHandler.SetBasicBlock(newJumpBasicBlockGraphReducer(basicBlock, target))
				} else {
					changeEndLoopToJump(visited, target, exceptionHandler.BasicBlock())
				}
			}
			break
		case intsrv.TypeIfElse, intsrv.TypeTernaryOperator:
			if basicBlock.Sub2() == cfg.LoopEnd {
				basicBlock.SetSub2(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Sub2())
			}
			fallthrough
		case intsrv.TypeIf:
			if basicBlock.Sub1() == cfg.LoopEnd {
				basicBlock.SetSub1(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Sub1())
			}
			if basicBlock.Next() == cfg.LoopEnd {
				basicBlock.SetNext(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Next())
			}
			break
		case intsrv.TypeConditionOr, intsrv.TypeConditionAnd:
			if basicBlock.Sub1() == cfg.LoopEnd {
				basicBlock.SetSub1(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Sub1())
			}
			if basicBlock.Sub2() == cfg.LoopEnd {
				basicBlock.SetSub2(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Sub2())
			}
			break
		case intsrv.TypeSwitch:
			if basicBlock.Next() == cfg.LoopEnd {
				basicBlock.SetNext(newJumpBasicBlockGraphReducer(basicBlock, target))
			} else {
				changeEndLoopToJump(visited, target, basicBlock.Next())
			}
			fallthrough
		case intsrv.TypeSwitchDeclaration:
			for _, switchCase := range basicBlock.SwitchCases().ToSlice() {
				if switchCase.BasicBlock() == cfg.LoopEnd {
					switchCase.SetBasicBlock(newJumpBasicBlockGraphReducer(basicBlock, target))
				} else {
					changeEndLoopToJump(visited, target, switchCase.BasicBlock())
				}
			}
			break
		}
	}
}

func newJumpBasicBlockGraphReducer(bb, target intsrv.IBasicBlock) intsrv.IBasicBlock {
	predecessors := util.NewSet[intsrv.IBasicBlock]()

	predecessors.Add(bb)
	target.Predecessors().Remove(bb)

	return bb.ControlFlowGraph().NewBasicBlock5(intsrv.TypeJump, bb.FromOffset(), target.FromOffset(), predecessors)
}

func clone(bb, next intsrv.IBasicBlock) intsrv.IBasicBlock {
	clone1 := next.ControlFlowGraph().NewBasicBlock3(next.Type(), next.FromOffset(), next.ToOffset())
	clone1.SetNext(cfg.End)
	clone1.Predecessors().Add(bb)
	next.Predecessors().Remove(bb)
	bb.SetNext(clone1)
	return clone1
}
