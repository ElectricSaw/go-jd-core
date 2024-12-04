package utils

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
)

func SearchNextOpcode(basicBlock intsrv.IBasicBlock, maxOffset int) int {
	code := basicBlock.ControlFlowGraph().Method().Attribute("Code").(intcls.IAttributeCode).Code()
	offset := basicBlock.FromOffset()
	toOffset := basicBlock.ToOffset()

	if toOffset > maxOffset {
		toOffset = maxOffset
	}

	var deltaOffset int

	for ; offset < toOffset; offset++ {
		opcode := code[offset] & 255

		switch opcode {
		case 16, 18, /* BIPUSH, LDC, */
			21, 22, 23, 24, 25, /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD, */
			54, 55, 56, 57, 58, /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE, */
			169, /* RET, */
			188 /* NEWARRAY: */ :
			offset++
		case 17, /* SIPUSH */
			19, 20, /* LDC_W, LDC2_W */
			132,      /* IINC */
			178,      /* GETSTATIC */
			179,      /* PUTSTATIC */
			187,      /* NEW */
			180,      /* GETFIELD */
			181,      /* PUTFIELD */
			182, 183, /* INVOKEVIRTUAL, INVOKESPECIAL */
			184, /* INVOKESTATIC */
			189, /* ANEWARRAY */
			192, /* CHECKCAST */
			193 /* INSTANCEOF */ :
			offset += 2
		case 153, 154, 155, 156, 157, 158, /* IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE */
			159, 160, 161, 162, 163, 164, 165, 166, /* IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE */
			167, /* GOTO */
			198, 199 /* IFNULL, IFNONNULL */ :
			deltaOffset, offset = PrefixReadInt16(code, offset)
			if deltaOffset > 0 {
				offset += deltaOffset - 2 - 1
			}
		case 200: /* GOTO_W */

			deltaOffset, offset = PrefixReadInt32(code, offset)

			if deltaOffset > 0 {
				offset += deltaOffset - 4 - 1
			}
		case 168: /* JSR */
			offset += 2
		case 197: /* MULTIANEWARRAY */
			offset += 3
		case 185 /* INVOKEINTERFACE */, 186 /* INVOKEDYNAMIC */ :
			offset += 4
		case 201: /* JSR_W */
			offset += 4
		case 170: /* TABLESWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */
			deltaOffset, offset = PrefixReadInt32(code, offset)
			low := deltaOffset
			deltaOffset, offset = PrefixReadInt32(code, offset)
			high := deltaOffset
			offset += (4 * (high - low + 1)) - 1
		case 171: /* LOOKUPSWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */
			deltaOffset, offset = PrefixReadInt32(code, offset)
			offset += (8 * deltaOffset) - 1
		case 196: /* WIDE */
			offset++
			opcode = code[offset] & 255
			if opcode == 132 { /* IINC */
				offset += 4
			} else {
				offset += 2
			}
		}
	}

	if offset <= maxOffset {
		return int(code[offset]) & 255
	} else {
		return 0
	}
}

func LastOpcode(basicBlock intsrv.IBasicBlock) int {
	code := basicBlock.ControlFlowGraph().Method().Attribute("Code").(intcls.IAttributeCode).Code()
	offset := basicBlock.FromOffset()
	toOffset := basicBlock.ToOffset()

	if offset >= toOffset {
		return 0
	}
	var deltaOffset int

	lastOffset := offset

	for ; offset < toOffset; offset++ {
		opcode := code[offset] & 255

		lastOffset = offset

		switch opcode {
		case 16, 18, /* BIPUSH, LDC, */
			21, 22, 23, 24, 25, /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD, */
			54, 55, 56, 57, 58, /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE, */
			169, /* RET, */
			188 /* NEWARRAY: */ :
			offset++
		case 17, /* SIPUSH */
			19, 20, /* LDC_W, LDC2_W */
			132,      /* IINC */
			178,      /* GETSTATIC */
			179,      /* PUTSTATIC */
			187,      /* NEW */
			180,      /* GETFIELD */
			181,      /* PUTFIELD */
			182, 183, /* INVOKEVIRTUAL, INVOKESPECIAL */
			184, /* INVOKESTATIC */
			189, /* ANEWARRAY */
			192, /* CHECKCAST */
			193 /* INSTANCEOF */ :
			offset += 2
		case 153, 154, 155, 156, 157, 158, /* IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE */
			159, 160, 161, 162, 163, 164, 165, 166, /* IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE */
			167, /* GOTO */
			198, 199 /* IFNULL, IFNONNULL */ :
			deltaOffset, offset = PrefixReadInt16(code, offset)
			if deltaOffset > 0 {
				offset += deltaOffset - 2 - 1
			}
		case 200: /* GOTO_W */
			deltaOffset, offset = PrefixReadInt32(code, offset)

			if deltaOffset > 0 {
				offset += deltaOffset - 4 - 1
			}
		case 168: /* JSR */
			offset += 2
		case 197: /* MULTIANEWARRAY */
			offset += 3
		case 185: /* INVOKEINTERFACE */
		case 186: /* INVOKEDYNAMIC */
			offset += 4
		case 201: /* JSR_W */
			offset += 4
		case 170: /* TABLESWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */

			deltaOffset, offset = PrefixReadInt32(code, offset)
			low := deltaOffset
			deltaOffset, offset = PrefixReadInt32(code, offset)
			high := deltaOffset

			offset += (4 * (high - low + 1)) - 1
		case 171: /* LOOKUPSWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */

			deltaOffset, offset = PrefixReadInt32(code, offset)

			offset += (8 * deltaOffset) - 1
		case 196: /* WIDE */
			offset++
			opcode = code[offset] & 255
			if opcode == 132 { /* IINC */
				offset += 4
			} else {
				offset += 2
			}
		}
	}

	return int(code[lastOffset]) & 255
}

func EvalStackDepth(bb intsrv.IBasicBlock) int {
	method := bb.ControlFlowGraph().Method()
	constants := method.Constants()
	attributeCode := method.Attribute("Code").(intcls.IAttributeCode)
	code := attributeCode.Code()
	return EvalStackDepth2(constants, code, bb)
}

func EvalStackDepth2(constants intcls.IConstantPool, code []byte, bb intsrv.IBasicBlock) int {
	var constantMemberRef intcls.IConstantMemberRef
	var constantNameAndType intcls.IConstantNameAndType
	descriptor := ""
	depth := 0
	deltaOffset := 0

	for offset, toOffset := bb.FromOffset(), bb.ToOffset(); offset < toOffset; offset++ {
		opcode := int(code[offset]) & 255

		switch opcode {
		case 1, /* ACONST_NULL */
			2, 3, 4, 5, 6, 7, 8, /* ICONST_M1, ICONST_0 ... ICONST_5 */
			9, 10, 11, 12, 13, 14, 15, /* LCONST_0, LCONST_1, FCONST_0, FCONST_1, FCONST_2, DCONST_0, DCONST_1 */
			26, 27, 28, 29, /* ILOAD_0 ... ILOAD_3 */
			30, 31, 32, 33, /* LLOAD_0 ... LLOAD_3 */
			34, 35, 36, 37, /* FLOAD_0 ... FLOAD_3 */
			38, 39, 40, 41, /* DLOAD_0 ... DLOAD_3 */
			42, 43, 44, 45, /* ALOAD_0 ... ALOAD_3 */
			89, 90, 91 /* DUP, DUP_X1, DUP_X2 */ :
			depth++
		case 16, 18, /* BIPUSH, LDC */
			21, 22, 23, 24, 25 /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD */ :
			offset++
			depth++
		case 17, /* SIPUSH */
			19, 20, /* LDC_W, LDC2_W */
			168, /* JSR */
			178, /* GETSTATIC */
			187 /* NEW */ :
			offset += 2
			depth++
			break
		case 46, 47, 48, 49, 50, 51, 52, 53, /* IALOAD, LALOAD, FALOAD, DALOAD, AALOAD, BALOAD, CALOAD, SALOAD */
			59, 60, 61, 62, /* ISTORE_0 ... ISTORE_3 */
			63, 64, 65, 66, /* LSTORE_0 ... LSTORE_3 */
			67, 68, 69, 70, /* FSTORE_0 ... FSTORE_3 */
			71, 72, 73, 74, /* DSTORE_0 ... DSTORE_3 */
			75, 76, 77, 78, /* ASTORE_0 ... ASTORE_3 */
			87,             /* POP */
			96, 97, 98, 99, /* IADD, LADD, FADD, DADD */
			100, 101, 102, 103, /* ISUB, LSUB, FSUB, DSUB */
			104, 105, 106, 107, /* IMUL, LMUL, FMUL, DMUL */
			108, 109, 110, 111, /* IDIV, LDIV, FDIV, DDIV */
			112, 113, 114, 115, /* IREM, LREM, FREM, DREM */
			120, 121, /* ISHL, LSHL */
			122, 123, /* ISHR, LSHR */
			124, 125, /* IUSHR, LUSHR */
			126, 127, /* IAND, LAND */
			128, 129, /* IOR, LOR */
			130, 131, /* IXOR, LXOR */
			148, 149, 150, 151, 152, /* LCMP, FCMPL, FCMPG, DCMPL, DCMPG */
			172, 173, 174, 175, 176, /* IRETURN, LRETURN, FRETURN, DRETURN, ARETURN */
			194, 195 /* MONITORENTER, MONITOREXIT */ :
			depth--
		case 153, 154, 155, 156, 157, 158, /* IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE */
			179, /* PUTSTATIC */
			198, 199 /* IFNULL, IFNONNULL */ :
			offset += 2
			depth--
		case 54, 55, 56, 57, 58: /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE */
			offset++
			depth--
		case 79, 80, 81, 82, 83, 84, 85, 86: /* IASTORE, LASTORE, FASTORE, DASTORE, AASTORE, BASTORE, CASTORE, SASTORE */
			depth -= 3
		case 92, 93, 94: /* DUP2, DUP2_X1, DUP2_X2 */
			depth += 2
		case 132, /* IINC */
			167, /* GOTO */
			180, /* GETFIELD */
			189, /* ANEWARRAY */
			192, /* CHECKCAST */
			193 /* INSTANCEOF */ :
			offset += 2
		case 159, 160, 161, 162, 163, 164, 165, 166, /* IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE */
			181: /* PUTFIELD */
			offset += 2
			depth -= 2
		case 88: /* POP2 */
			depth -= 2
		case 169, /* RET */
			188 /* NEWARRAY */ :
			offset++
		case 170: /* TABLESWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */

			deltaOffset, offset = SuffixReadInt32(code, offset)
			low := deltaOffset
			deltaOffset, offset = SuffixReadInt32(code, offset)
			high := deltaOffset

			offset += (4 * (high - low + 1)) - 1
			depth--
		case 171: /* LOOKUPSWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */
			deltaOffset, offset = SuffixReadInt32(code, offset)

			offset += (8 * deltaOffset) - 1
			depth--
		case 182, 183: /* INVOKEVIRTUAL, INVOKESPECIAL */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= 1 + countMethodParameters(descriptor)
			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
		case 184: /* INVOKESTATIC */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= countMethodParameters(descriptor)
			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
			break
		case 185: /* INVOKEINTERFACE */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= 1 + countMethodParameters(descriptor)
			offset += 2 /* Skip 'count' and one byte */

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
			break
		case 186: /* INVOKEDYNAMIC */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= countMethodParameters(descriptor)
			offset += 2 /* Skip 2 bytes */

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
			break
		case 196: /* WIDE */
			offset++
			opcode = int(code[offset]) & 255
			if opcode == 132 { /* IINC */
				offset += 4
			} else {
				offset += 2

				switch opcode {
				case 21, 22, 23, 24, 25: /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD */
					depth++
				case 54, 55, 56, 57, 58: /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE */
					depth--
				case 169: /* RET */
				}
			}
		case 197: /* MULTIANEWARRAY */
			offset += 3
			depth += 1 - int(code[offset])&255
		case 201: /* JSR_W */
			offset += 4
			depth++
			fallthrough
		case 200: /* GOTO_W */
			offset += 4
		}
	}

	return depth
}

func MinDepth(bb intsrv.IBasicBlock) int {
	method := bb.ControlFlowGraph().Method()
	constants := method.Constants()
	attributeCode := method.Attribute("Code").(intcls.IAttributeCode)
	code := attributeCode.Code()
	return minDepth(constants, code, bb)
}

func minDepth(constants intcls.IConstantPool, code []byte, bb intsrv.IBasicBlock) int {
	var constantMemberRef intcls.IConstantMemberRef
	var constantNameAndType intcls.IConstantNameAndType
	descriptor := ""
	depth := 0
	minimumDepth := 0
	deltaOffset := 0

	for offset, toOffset := bb.FromOffset(), bb.ToOffset(); offset < toOffset; offset++ {
		opcode := code[offset] & 255

		switch opcode {
		case 1, /* ACONST_NULL */
			2, 3, 4, 5, 6, 7, 8, /* ICONST_M1, ICONST_0 ... ICONST_5 */
			9, 10, 11, 12, 13, 14, 15, /* LCONST_0, LCONST_1, FCONST_0, FCONST_1, FCONST_2, DCONST_0, DCONST_1 */
			26, 27, 28, 29, /* ILOAD_0 ... ILOAD_3 */
			30, 31, 32, 33, /* LLOAD_0 ... LLOAD_3 */
			34, 35, 36, 37, /* FLOAD_0 ... FLOAD_3 */
			38, 39, 40, 41, /* DLOAD_0 ... DLOAD_3 */
			42, 43, 44, 45 /* ALOAD_0 ... ALOAD_3 */ :
			depth++
		case 89: /* DUP */
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 2
		case 90: /* DUP_X1 */
			depth -= 2
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 3
		case 91: /* DUP_X2 */
			depth -= 3
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 4
		case 16, 18, /* BIPUSH, LDC */
			21, 22, 23, 24, 25 /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD */ :
			offset++
			depth++
		case 17, /* SIPUSH */
			19, 20, /* LDC_W, LDC2_W */
			168, /* JSR */
			178, /* GETSTATIC */
			187 /* NEW */ :
			offset += 2
			depth++
		case 46, 47, 48, 49, 50, 51, 52, 53, /* IALOAD, LALOAD, FALOAD, DALOAD, AALOAD, BALOAD, CALOAD, SALOAD */
			96, 97, 98, 99, /* IADD, LADD, FADD, DADD */
			100, 101, 102, 103, /* ISUB, LSUB, FSUB, DSUB */
			104, 105, 106, 107, /* IMUL, LMUL, FMUL, DMUL */
			108, 109, 110, 111, /* IDIV, LDIV, FDIV, DDIV */
			112, 113, 114, 115, /* IREM, LREM, FREM, DREM */
			120, 121, /* ISHL, LSHL */
			122, 123, /* ISHR, LSHR */
			124, 125, /* IUSHR, LUSHR */
			126, 127, /* IAND, LAND */
			128, 129, /* IOR, LOR */
			130, 131, /* IXOR, LXOR */
			148, 149, 150, 151, 152 /* LCMP, FCMPL, FCMPG, DCMPL, DCMPG */ :
			depth -= 2
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth++
		case 59, 60, 61, 62, /* ISTORE_0 ... ISTORE_3 */
			63, 64, 65, 66, /* LSTORE_0 ... LSTORE_3 */
			67, 68, 69, 70, /* FSTORE_0 ... FSTORE_3 */
			71, 72, 73, 74, /* DSTORE_0 ... DSTORE_3 */
			75, 76, 77, 78, /* ASTORE_0 ... ASTORE_3 */
			87,                      /* POP */
			172, 173, 174, 175, 176, /* IRETURN, LRETURN, FRETURN, DRETURN, ARETURN */
			194, 195: /* MONITORENTER, MONITOREXIT */
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 153, 154, 155, 156, 157, 158, /* IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE */
			179,      /* PUTSTATIC */
			198, 199: /* IFNULL, IFNONNULL */
			offset += 2
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 54, 55, 56, 57, 58: /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE */
			offset++
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 79, 80, 81, 82, 83, 84, 85, 86: /* IASTORE, LASTORE, FASTORE, DASTORE, AASTORE, BASTORE, CASTORE, SASTORE */
			depth -= 3
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 92: /* DUP2 */
			depth -= 2
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 4
		case 93: /* DUP2_X1 */
			depth -= 3
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 5
		case 94: /* DUP2_X2 */
			depth -= 4
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth += 6
		case 132: /* IINC */
		case 167: /* GOTO */
			offset += 2
		case 180: /* GETFIELD */
		case 189: /* ANEWARRAY */
		case 192: /* CHECKCAST */
		case 193: /* INSTANCEOF */
			offset += 2
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth++
		case 159, 160, 161, 162, 163, 164, 165, 166: /* IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE */
		case 181: /* PUTFIELD */
			offset += 2
			depth -= 2
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 88: /* POP2 */
			depth -= 2
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 169: /* RET */
			offset++
		case 188: /* NEWARRAY */
			offset++
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth++
		case 170: /* TABLESWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */

			deltaOffset, offset = SuffixReadInt32(code, offset)
			low := deltaOffset
			deltaOffset, offset = SuffixReadInt32(code, offset)
			high := deltaOffset

			offset += (4 * (high - low + 1)) - 1
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 171: /* LOOKUPSWITCH */
			offset = (offset + 4) & 0xFFFC /* Skip padding */
			offset += 4                    /* Skip default offset */

			deltaOffset, offset = SuffixReadInt32(code, offset)

			offset += (8 * deltaOffset) - 1
			depth--
			if minimumDepth > depth {
				minimumDepth = depth
			}
		case 182, 183: /* INVOKEVIRTUAL, INVOKESPECIAL */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= 1 + countMethodParameters(descriptor)
			if minimumDepth > depth {
				minimumDepth = depth
			}

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
		case 184: /* INVOKESTATIC */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= countMethodParameters(descriptor)
			if minimumDepth > depth {
				minimumDepth = depth
			}

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
		case 185: /* INVOKEINTERFACE */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= 1 + countMethodParameters(descriptor)
			if minimumDepth > depth {
				minimumDepth = depth
			}
			offset += 2 /* Skip 'count' and one byte */

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
		case 186: /* INVOKEDYNAMIC */
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			depth -= countMethodParameters(descriptor)
			if minimumDepth > depth {
				minimumDepth = depth
			}
			offset += 2 /* Skip 2 bytes */

			if descriptor[len(descriptor)-1] != 'V' {
				depth++
			}
		case 196: /* WIDE */
			offset++
			opcode = code[offset] & 255
			if opcode == 132 { /* IINC */
				offset += 4
			} else {
				offset += 2

				switch opcode {
				case 21, 22, 23, 24, 25: /* ILOAD, LLOAD, FLOAD, DLOAD, ALOAD */
					depth++
				case 54, 55, 56, 57, 58: /* ISTORE, LSTORE, FSTORE, DSTORE, ASTORE */
					depth--
					if minimumDepth > depth {
						minimumDepth = depth
					}
				case 169: /* RET */
				}
			}
		case 197: /* MULTIANEWARRAY */
			offset += 3
			depth -= int(code[offset]) & 255
			if minimumDepth > depth {
				minimumDepth = depth
			}
			depth++
		case 201: /* JSR_W */
			offset += 4
			depth++
		case 200: /* GOTO_W */
			offset += 4
		}
	}

	return minimumDepth
}

func countMethodParameters(descriptor string) int {
	count := 0
	i := 2
	c := descriptor[1]

	if !(len(descriptor) > 2 && descriptor[0] == '(') {
		return -1
	}

	for c != ')' {
		for ; c == '['; i++ {
			c = descriptor[i]
		}

		if c == 'L' {
			for {
				c = descriptor[i]
				i++

				if c == ';' {
					break
				}
			}
		}

		c = descriptor[i]
		i++
		count++
	}

	return count
}

func PrefixReadInt8(code []byte, offset int) (int, int) {
	offset++
	deltaOffset := int(code[offset]&255) << 8
	return deltaOffset, offset
}

func PrefixReadInt16(code []byte, offset int) (int, int) {
	offset++
	deltaOffset := int(code[offset]&255) << 8
	offset++
	deltaOffset |= int(code[offset] & 255)
	return deltaOffset, offset
}

func PrefixReadInt32(code []byte, offset int) (int, int) {
	offset++
	deltaOffset := int(code[offset]&255) << 24
	offset++
	deltaOffset |= int(code[offset]&255) << 16
	offset++
	deltaOffset |= int(code[offset]&255) << 8
	offset++
	deltaOffset |= int(code[offset] & 255)
	return deltaOffset, offset
}

func SuffixReadInt8(code []byte, offset int) (int, int) {
	deltaOffset := int(code[offset]&255) << 8
	offset++
	return deltaOffset, offset
}

func SuffixReadInt16(code []byte, offset int) (int, int) {
	deltaOffset := int(code[offset]&255) << 8
	offset++
	deltaOffset |= int(code[offset] & 255)
	offset++
	return deltaOffset, offset
}

func SuffixReadInt32(code []byte, offset int) (int, int) {
	deltaOffset := int(code[offset]&255) << 24
	offset++
	deltaOffset |= int(code[offset]&255) << 16
	offset++
	deltaOffset |= int(code[offset]&255) << 8
	offset++
	deltaOffset |= int(code[offset] & 255)
	offset++
	return deltaOffset, offset
}
