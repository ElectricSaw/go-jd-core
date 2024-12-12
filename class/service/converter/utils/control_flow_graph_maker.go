package utils

import (
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/model/cfg"
	"github.com/ElectricSaw/go-jd-core/class/util"
	"sort"
)

var Mark = cfg.End

func MakeControlFlowGraph(method intcls.IMethod) intsrv.IControlFlowGraph {
	attributeCode := method.Attribute("Code").(intcls.IAttributeCode)

	if attributeCode == nil {
		return nil
	} else {
		// Parse byte-code
		constants := method.Constants()
		code := attributeCode.Code()
		length := len(code)
		mapped := make([]intsrv.IBasicBlock, length)
		types := make([]byte, length)          // 'c' for conditional instruction, 'g' for goto, 't' for throw, 's' for switch, 'r' for return
		nextOffsets := make([]int, length)     // Next instruction offsets
		branchOffsets := make([]int, length)   // Branch offsets
		switchValues := make([][]int, length)  // Default-value and switch-values
		switchOffsets := make([][]int, length) // Default-case offset and switch-case offsets

		// --- Search leaders --- //

		// The first instruction is a leader
		mapped[0] = Mark

		typ := byte(0)
		value := 0
		lastOffset := 0
		lastStatementOffset := -1

		for offset := 0; offset < length; offset++ {
			nextOffsets[lastOffset] = offset
			lastOffset = offset
			opcode := int(code[offset] & 255)

			switch opcode {
			case 16, // BIPUSH
				18,                 // LDC
				21, 22, 23, 24, 25, // ILOAD, LLOAD, FLOAD, DLOAD, ALOAD
				188: // NEWARRAY
				offset++
				break
			case 54, 55, 56, 57, 58: // ISTORE, LSTORE, FSTORE, DSTORE, ASTORE
				offset++
				lastStatementOffset = offset
				break
			case 59, 60, 61, 62, // ISTORE_0 .. ISTORE_3
				63, 64, 65, 66, // LSTORE_0 .. LSTORE_3
				67, 68, 69, 70, // FSTORE_0 .. FSTORE_3
				71, 72, 73, 74, // DSTORE_0 .. DSTORE_3
				75, 76, 77, 78, // ASTORE_0 .. ASTORE_3
				79, 80, 81, 82, 83, 84, 85, 86, // IASTORE, LASTORE, FASTORE, DASTORE, AASTORE, BASTORE, CASTORE, SASTORE
				87, 88, // POP, POP2
				194, 195: // MONITORENTER, MONITOREXIT
				lastStatementOffset = offset
				break
			case 169: // RET
				offset++
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				types[offset] = 'R'
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 179, 181: // PUTSTATIC, PUTFIELD
				offset += 2
				lastStatementOffset = offset
				break
			case 182, 183, 184: // INVOKEVIRTUAL, INVOKESPECIAL, INVOKESTATIC
				value, offset = PrefixReadInt16(code, offset)
				constantMemberRef := constants.Constant(value).(intcls.IConstantMemberRef)
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				if descriptor[len(descriptor)-1] == 'V' {
					lastStatementOffset = offset
				}
				break
			case 185, 186: // INVOKEINTERFACE, INVOKEDYNAMIC
				value, offset = PrefixReadInt16(code, offset)
				constantMemberRef := constants.Constant(value).(intcls.IConstantMemberRef)
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				offset += 2 // Skip 2 bytes
				if descriptor[len(descriptor)-1] == 'V' {
					lastStatementOffset = offset
				}
				break
			case 132: // IINC
				offset += 2
				value = int(code[offset-1] & 255)
				if (lastStatementOffset+3 == offset) && (checkILOADForIINC(code, offset, value) == false) {
					// Last instruction is a 'statement' & the next instruction is not a matching ILOAD -> IINC as a statement
					lastStatementOffset = offset
				}
				break
			case 17, // SIPUSH
				19, 20, // LDC_W, LDC2_W
				178, 180, // GETSTATIC, GETFIELD
				187, 189, // NEW, ANEWARRAY
				192, // CHECKCAST
				193: // INSTANCEOF
				offset += 2
				break
			case 167: // GOTO

				typ = 'G'
				if lastStatementOffset+1 == offset {
					typ = 'g'
				}

				if lastStatementOffset != -1 {
					mapped[lastStatementOffset+1] = Mark
				}
				// The target of a conditional or an unconditional goto/jump instruction is a leader
				types[offset] = typ // TODO debug, remove this line
				value, offset = PrefixReadInt16(code, offset)
				branchOffset := offset + value
				mapped[branchOffset] = Mark
				types[offset] = typ
				branchOffsets[offset] = branchOffset
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 168: // JSR
				if lastStatementOffset != -1 {
					mapped[lastStatementOffset+1] = Mark
				}
				types[offset] = 'j' // TODO debug, remove this line
				// The target of a conditional or an unconditional goto/jump instruction is a leader
				value, offset = PrefixReadInt16(code, offset)
				branchOffset := offset + value
				mapped[branchOffset] = Mark
				types[offset] = 'j'
				branchOffsets[offset] = branchOffset
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 153, 154, 155, 156, 157, 158, // IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE
				159, 160, 161, 162, 163, 164, 165, 166, // IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE
				198, 199: // IFnil, IFNONnil
				if lastStatementOffset != -1 {
					mapped[lastStatementOffset+1] = Mark
				}
				// The target of a conditional or an unconditional goto/jump instruction is a leader
				value, offset = PrefixReadInt16(code, offset)
				branchOffset := offset + value
				mapped[branchOffset] = Mark
				types[offset] = 'c'
				branchOffsets[offset] = branchOffset
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 170: // TABLESWITCH
				// Skip padding
				i := (offset + 4) & 0xFFFC
				value, i = SuffixReadInt32(code, i)
				defaultOffset := offset + value
				mapped[defaultOffset] = Mark
				var low, high int
				low, i = SuffixReadInt32(code, i)
				high, i = SuffixReadInt32(code, i)
				values := make([]int, high-low+2)
				offsets := make([]int, high-low+2)

				offsets[0] = defaultOffset
				length = high - low + 2
				for j := 1; j < length; j++ {
					values[j] = low + j - 1
					value, i = SuffixReadInt32(code, i)
					branchOffset := offset + value
					offsets[j] = branchOffset
					mapped[branchOffset] = Mark
				}

				offset = (i - 1)
				types[offset] = 's'
				switchValues[offset] = values
				switchOffsets[offset] = offsets
				lastStatementOffset = offset
				break
			case 171: // LOOKUPSWITCH
				// Skip padding
				i := (offset + 4) & 0xFFFC
				value, i = SuffixReadInt32(code, i)
				defaultOffset := offset + value
				mapped[defaultOffset] = Mark
				value, i = SuffixReadInt32(code, i)
				npairs := value

				values := make([]int, npairs+1)
				offsets := make([]int, npairs+1)

				offsets[0] = defaultOffset
				for j := 1; j <= npairs; j++ {
					value, i = SuffixReadInt32(code, i)
					values[j] = value
					value, i = SuffixReadInt32(code, i)
					branchOffset := offset + value
					offsets[j] = branchOffset
					mapped[branchOffset] = Mark
				}

				offset = (i - 1)
				types[offset] = 's'
				switchValues[offset] = values
				switchOffsets[offset] = offsets
				lastStatementOffset = offset
				break
			case 172, 173, 174, 175, 176: // IRETURN, LRETURN, FRETURN, DRETURN, ARETURN
				types[offset] = 'v'
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 177: // RETURN
				if lastStatementOffset != -1 {
					mapped[lastStatementOffset+1] = Mark
				}
				types[offset] = 'r'
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 191: // ATHROW
				types[offset] = 't'
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 196: // WIDE
				offset++
				opcode = int(code[offset] & 255)
				switch opcode {
				case 132: // IINC
					offset += 4
					if (lastStatementOffset+6 == offset) &&
						(checkILOADForIINC(code, offset, int(code[offset-3]&255)<<8|int(code[offset-2]&255)) == false) {
						// Last instruction is a 'statement' & the next instruction is not a matching ILOAD -> IINC as a statement
						lastStatementOffset = offset
					}
					break
				case 169: // RET
					offset += 2
					// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
					types[offset] = 'R'
					if offset+1 < length {
						mapped[offset+1] = Mark
					}
					lastStatementOffset = offset
					break
				case 54, 55, 56, 57, 58: // ISTORE, LSTORE, FSTORE, DSTORE, ASTORE
					lastStatementOffset = offset + 2
				default:
					offset += 2
					break
				}
				break
			case 197: // MULTIANEWARRAY
				offset += 3
				break
			case 200: // GOTO_W
				typ = 'G'
				if lastStatementOffset+1 == offset {
					typ = 'g'
				}

				types[offset] = typ // TODO debug, remove this line
				value, offset = PrefixReadInt32(code, offset)
				branchOffset := offset + value
				mapped[branchOffset] = Mark
				types[offset] = typ
				branchOffsets[offset] = branchOffset
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			case 201: // JSR_W
				if lastStatementOffset != -1 {
					mapped[lastStatementOffset+1] = Mark
				}
				types[offset] = 'j' // TODO debug, remove this line
				// The target of a conditional or an unconditional goto/jump instruction is a leader
				value, offset = PrefixReadInt32(code, offset)
				branchOffset := offset + value
				mapped[branchOffset] = Mark
				types[offset] = 'j'
				branchOffsets[offset] = branchOffset
				// The instruction that immediately follows a conditional or an unconditional goto/jump instruction is a leader
				if offset+1 < length {
					mapped[offset+1] = Mark
				}
				lastStatementOffset = offset
				break
			}
		}

		nextOffsets[lastOffset] = length
		codeExceptions := attributeCode.ExceptionTable()

		if codeExceptions != nil {
			for _, codeException := range codeExceptions {
				mapped[codeException.StartPc()] = Mark
				mapped[codeException.HandlerPc()] = Mark
			}
		}

		// --- Create line numbers --- //
		cfg1 := cfg.NewControlFlowGraph(method)
		attributeLineNumberTable := attributeCode.Attribute("LineNumberTable").(intcls.IAttributeLineNumberTable)

		if attributeLineNumberTable != nil {
			// Parse line numbers
			lineNumberTable := attributeLineNumberTable.LineNumberTable()

			offsetToLineNumbers := make([]int, length)
			offset := 0
			lineNumber := lineNumberTable[0].LineNumber()
			length1 := len(lineNumberTable)
			for i := 1; i < length1; i++ {
				lineNumberEntry := lineNumberTable[i]
				toIndex := lineNumberEntry.StartPc()

				for offset < toIndex {
					offsetToLineNumbers[offset] = lineNumber
					offset++
				}

				if lineNumber > lineNumberEntry.LineNumber() {
					mapped[offset] = Mark
				}

				lineNumber = lineNumberEntry.LineNumber()
			}

			for offset < length {
				offsetToLineNumbers[offset] = lineNumber
				offset++
			}

			cfg1.SetOffsetToLineNumbers(offsetToLineNumbers)
		}

		// --- Create basic blocks --- //
		lastOffset = 0

		// Add 'start'
		startBasicBlock := cfg1.NewBasicBlock3(intsrv.TypeStart, 0, 0)

		for offset := nextOffsets[0]; offset < length; offset = nextOffsets[offset] {
			if mapped[offset] != nil {
				mapped[lastOffset] = cfg1.NewBasicBlock2(lastOffset, offset)
				lastOffset = offset
			}
		}

		mapped[lastOffset] = cfg1.NewBasicBlock2(lastOffset, length)

		// --- Set type, successors and predecessors --- //
		list := cfg1.BasicBlocks()
		basicBlocks := util.NewDefaultListWithCapacity[intsrv.IBasicBlock](list.Size())
		successor := list.Get(1)
		startBasicBlock.SetNext(successor)
		successor.Predecessors().Add(startBasicBlock)

		basicBlockLength := list.Size()
		for i := 1; i < basicBlockLength; i++ {
			basicBlock := list.Get(i)
			lastInstructionOffset := basicBlock.ToOffset() - 1

			switch types[lastInstructionOffset] {
			case 'g': // Goto
				basicBlock.SetType(intsrv.TypeGoto)
				successor = mapped[branchOffsets[lastInstructionOffset]]
				basicBlock.SetNext(successor)
				successor.Predecessors().Add(basicBlock)
				break
			case 'G': // Goto in ternary operator
				basicBlock.SetType(intsrv.TypeGotoInTernaryOperator)
				successor = mapped[branchOffsets[lastInstructionOffset]]
				basicBlock.SetNext(successor)
				successor.Predecessors().Add(basicBlock)
				break
			case 't': // Throw
				basicBlock.SetType(intsrv.TypeThrow)
				basicBlock.SetNext(cfg.End)
				break
			case 'r': // Return
				basicBlock.SetType(intsrv.TypeReturn)
				basicBlock.SetNext(cfg.End)
				break
			case 'c': // Conditional
				basicBlock.SetType(intsrv.TypeConditionalBranch)
				successor = mapped[basicBlock.ToOffset()]
				basicBlock.SetNext(successor)
				successor.Predecessors().Add(basicBlock)
				successor = mapped[branchOffsets[lastInstructionOffset]]
				basicBlock.SetBranch(successor)
				successor.Predecessors().Add(basicBlock)
				break
			case 's': // Switch
				basicBlock.SetType(intsrv.TypeSwitchDeclaration)
				values := switchValues[lastInstructionOffset]
				offsets := switchOffsets[lastInstructionOffset]
				switchCases := util.NewDefaultListWithCapacity[intsrv.ISwitchCase](len(offsets))
				defaultOffset := offsets[0]
				bb := mapped[defaultOffset]
				switchCases.Add(cfg.NewSwitchCase(bb))
				bb.Predecessors().Add(basicBlock)
				length1 := len(offsets)
				for j := 1; j < length1; j++ {
					offset := offsets[j]
					if offset != defaultOffset {
						bb = mapped[offset]
						switchCases.Add(cfg.NewSwitchCaseWithValue(values[j], bb))
						bb.Predecessors().Add(basicBlock)
					}
				}

				basicBlock.SetSwitchCases(switchCases)
				break
			case 'j': // Jsr
				basicBlock.SetType(intsrv.TypeJsr)
				successor = mapped[basicBlock.ToOffset()]
				basicBlock.SetNext(successor)
				successor.Predecessors().Add(basicBlock)
				successor = mapped[branchOffsets[lastInstructionOffset]]
				basicBlock.SetBranch(successor)
				successor.Predecessors().Add(basicBlock)
				break
			case 'R': // Ret
				basicBlock.SetType(intsrv.TypeRet)
				basicBlock.SetNext(cfg.End)
				break
			case 'v': // Return value
				basicBlock.SetType(intsrv.TypeReturnValue)
				basicBlock.SetNext(cfg.End)
				break
			default:
				basicBlock.SetType(intsrv.TypeStatements)
				successor = mapped[basicBlock.ToOffset()]
				basicBlock.SetNext(successor)
				successor.Predecessors().Add(basicBlock)
				basicBlocks.Add(basicBlock)
				break
			}
		}

		// --- Create try-catch-finally basic blocks --- //
		if codeExceptions != nil {
			cache := make(map[intcls.ICodeException]intsrv.IBasicBlock)
			constantPool := method.Constants()
			// Reuse arrays
			handlePcToStartPc := branchOffsets
			handlePcMarks := types

			sort.SliceIsSorted(codeExceptions, func(i, j int) bool {
				ce1, ce2 := codeExceptions[i], codeExceptions[j]
				cmp := ce1.StartPc() - ce2.StartPc()
				if cmp == 0 {
					cmp = ce1.EndPc() - ce2.EndPc()
				}
				return cmp < 0
			})

			for _, codeException := range codeExceptions {
				startPc := codeException.StartPc()
				handlerPc := codeException.HandlerPc()

				if startPc != handlerPc {
					if (handlePcMarks[handlerPc] != 'T') ||
						(startPc <= mapped[handlePcToStartPc[handlerPc]].FromOffset()) {
						catchType := codeException.CatchType()
						tcf := cache[codeException]

						if tcf == nil {
							endPc := codeException.EndPc()
							// Check 'endPc'
							start := mapped[startPc]

							// Insert a new 'try-catch-finally' basic block
							tcf = cfg1.NewBasicBlock3(intsrv.TypeTryDeclaration, startPc, endPc)
							tcf.SetNext(start)

							// Update predecessors
							tcfPredecessors := tcf.Predecessors()
							startPredecessors := start.Predecessors()
							iterator := startPredecessors.Iterator()

							for iterator.HasNext() {
								predecessor := iterator.Next()

								if !start.Contains(predecessor) {
									predecessor.Replace(start, tcf)
									tcfPredecessors.Add(predecessor)
									_ = iterator.Remove()
								}
							}

							startPredecessors.Add(tcf)

							// Update map
							mapped[startPc] = tcf

							// Store to objectTypeCache
							cache[codeException] = tcf
						}

						internalThrowableName := ""
						if catchType != 0 {
							constantPool.ConstantTypeName(catchType)
						}
						handlerBB := mapped[handlerPc]
						tcf.AddExceptionHandler(internalThrowableName, handlerBB)
						handlerBB.Predecessors().Add(tcf)
						handlePcToStartPc[handlerPc] = startPc
						handlePcMarks[handlerPc] = 'T'
					}
				}
			}
		}

		// --- Recheck TYPE_GOTO_IN_TERNARY_OPERATOR --- //
		for _, bb := range basicBlocks.ToSlice() {
			next := bb.Next()
			var predecessors util.ISet[intsrv.IBasicBlock]

			if (bb.Type() == intsrv.TypeStatements) && (next.Predecessors().Size() == 1) {
				if (next.Type() == intsrv.TypeGoto) && (EvalStackDepth2(constants, code, bb) > 0) {
					// Transform STATEMENTS and GOTO to GOTO_IN_TERNARY_OPERATOR
					bb.SetType(intsrv.TypeGotoInTernaryOperator)
					bb.SetToOffset(next.ToOffset())
					bb.SetNext(next.Next())
					predecessors = next.Next().Predecessors()
					predecessors.Remove(next)
					predecessors.Add(bb)
					next.SetType(intsrv.TypeDeleted)
				} else if (next.Type() == intsrv.TypeConditionalBranch) && (EvalStackDepth2(constants, code, bb) > 0) {
					// Merge STATEMENTS and CONDITIONAL_BRANCH
					bb.SetType(intsrv.TypeConditionalBranch)
					bb.SetToOffset(next.ToOffset())
					bb.SetNext(next.Next())
					predecessors = next.Next().Predecessors()
					predecessors.Remove(next)
					predecessors.Add(bb)
					bb.SetBranch(next.Branch())
					predecessors = next.Branch().Predecessors()
					predecessors.Remove(next)
					predecessors.Add(bb)
					next.SetType(intsrv.TypeDeleted)
				}
			}
		}

		return cfg1
	}
}

func checkILOADForIINC(code []byte, offset, index int) bool {
	offset++
	if offset < len(code) {
		nextOpcode := code[offset] & 255

		if nextOpcode == 21 { // ILOAD
			if index == int(code[offset+1]&255) {
				return true
			}
		} else if int(nextOpcode) == 26+index { // ILOAD_0 ... ILOAD_3
			return true
		}
	}

	return false
}
