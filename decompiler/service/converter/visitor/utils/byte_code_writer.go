package utils

import (
	"fmt"
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
)

/**
* Example:
// Byte code:
//   0: aconst_null
//   1: astore_2
//   2: aconst_null
//   3: astore_3
//   4: aload_0
//   5: aload_1
//   6: invokeinterface 142 2 0
//   11: astore_2
//   12: aload_2
//   13: ifnull +49 -> 62
//   16: aload_2
//   17: invokestatic 146	jd/core/process/deserializer/ClassFileDeserializer:Deserialize	(Ljava/io/DataInput;)Ljd/core/model/classfile/ClassFile;
//   20: astore_3
//   21: goto +41 -> 62
//   24: astore 4
//   26: aconst_null
//   27: astore_3
//   28: aload_2
//   29: ifnull +46 -> 75
//   32: aload_2
//   33: invokevirtual 149	java/io/DataInputStream:close	()V
//   36: goto +39 -> 75
//   39: astore 6
//   41: goto +34 -> 75
//   44: astore 5
//   46: aload_2
//   47: ifnull +12 -> 59
//   50: aload_2
//   51: invokevirtual 149	java/io/DataInputStream:close	()V
//   54: goto +5 -> 59
//   57: astore 6
//   59: aload 5
//   61: athrow
//   62: aload_2
//   63: ifnull +12 -> 75
//   66: aload_2
//   67: invokevirtual 149	java/io/DataInputStream:close	()V
//   70: goto +5 -> 75
//   73: astore 6
//   75: aload_3
//   76: areturn
// Line number table:
//   #Java source line	-> byte code offset
//   #112	-> 0
//   #113	-> 2
//   #117	-> 4
//   #118	-> 12
//   #119	-> 16
//   #120	-> 21
//   #121	-> 24
//   #123	-> 26
//   #128	-> 28
//   #129	-> 32
//   #127	-> 44
//   #128	-> 46
//   #129	-> 50
//   #130	-> 59
//   #128	-> 62
//   #129	-> 66
//   #132	-> 75
// Local variable table:
//   start	length	slot	name	signature
//   0	77	0	loader	Loader
//   0	77	1	internalClassPath	String
//   1	66	2	dis	java.io.DataInputStream
//   3	73	3	classFile	ClassFile
//   24	3	4	e	IOException
//   44	16	5	localObject	Object
//   39	1	6	localIOException1	IOException
//   57	1	6	localIOException2	IOException
//   73	1	6	localIOException3	IOException
// Exception table:
//   from	to	target	type
//   4	21	24	java/io/IOException
//   32	36	39	java/io/IOException
//   4	28	44	finally
//   50	54	57	java/io/IOException
//   66	70	73	java/io/IOException
*/

var OpcodeNames = []string{
	"nop", "aconst_null", "iconst_m1", "iconst_0", "iconst_1",
	"iconst_2", "iconst_3", "iconst_4", "iconst_5", "lconst_0",
	"lconst_1", "fconst_0", "fconst_1", "fconst_2", "dconst_0",
	"dconst_1", "bipush", "sipush", "ldc", "ldc_w", "ldc2_w", "iload",
	"lload", "fload", "dload", "aload", "iload_0", "iload_1", "iload_2",
	"iload_3", "lload_0", "lload_1", "lload_2", "lload_3", "fload_0",
	"fload_1", "fload_2", "fload_3", "dload_0", "dload_1", "dload_2",
	"dload_3", "aload_0", "aload_1", "aload_2", "aload_3", "iaload",
	"laload", "faload", "daload", "aaload", "baload", "caload", "saload",
	"istore", "lstore", "fstore", "dstore", "astore", "istore_0",
	"istore_1", "istore_2", "istore_3", "lstore_0", "lstore_1",
	"lstore_2", "lstore_3", "fstore_0", "fstore_1", "fstore_2",
	"fstore_3", "dstore_0", "dstore_1", "dstore_2", "dstore_3",
	"astore_0", "astore_1", "astore_2", "astore_3", "iastore", "lastore",
	"fastore", "dastore", "aastore", "bastore", "castore", "sastore",
	"pop", "pop2", "dup", "dup_x1", "dup_x2", "dup2", "dup2_x1",
	"dup2_x2", "swap", "iadd", "ladd", "fadd", "dadd", "isub", "lsub",
	"fsub", "dsub", "imul", "lmul", "fmul", "dmul", "idiv", "ldiv",
	"fdiv", "ddiv", "irem", "lrem", "frem", "drem", "ineg", "lneg",
	"fneg", "dneg", "ishl", "lshl", "ishr", "lshr", "iushr", "lushr",
	"iand", "land", "ior", "lor", "ixor", "lxor", "iinc", "i2l", "i2f",
	"i2d", "l2i", "l2f", "l2d", "f2i", "f2l", "f2d", "d2i", "d2l", "d2f",
	"i2b", "i2c", "i2s", "lcmp", "fcmpl", "fcmpg",
	"dcmpl", "dcmpg", "ifeq", "ifne", "iflt", "ifge", "ifgt", "ifle",
	"if_icmpeq", "if_icmpne", "if_icmplt", "if_icmpge", "if_icmpgt",
	"if_icmple", "if_acmpeq", "if_acmpne", "goto", "jsr", "ret",
	"tableswitch", "lookupswitch", "ireturn", "lreturn", "freturn",
	"dreturn", "areturn", "return", "getstatic", "putstatic", "getfield",
	"putfield", "invokevirtual", "invokespecial", "invokestatic",
	"invokeinterface", "<illegal opcode>", "new", "newarray", "anewarray",
	"arraylength", "athrow", "checkcast", "instanceof", "monitorenter",
	"monitorexit", "wide", "multianewarray", "ifnull", "ifnonnull",
	"goto_w", "jsr_w", "breakpoint", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "<illegal opcode>", "<illegal opcode>", "<illegal opcode>",
	"<illegal opcode>", "impdep1", "impdep2",
}

func Write(linePrefix string, method intcls.IMethod) string {
	attributeCode := method.Attribute("Code").(intcls.IAttributeCode)

	if attributeCode == nil {
		return ""
	} else {
		constants := method.Constants()
		sb := ""

		writeByteCode(linePrefix, &sb, constants, attributeCode)
		writeLineNumberTable(linePrefix, &sb, attributeCode)
		writeLocalVariableTable(linePrefix, &sb, attributeCode)
		writeExceptionTable(linePrefix, &sb, constants, attributeCode)

		return sb
	}
}

func Write2(linePrefix string, method intcls.IMethod, fromOffset, toOffset int) string {
	attributeCode := method.Attribute("Code").(intcls.IAttributeCode)

	if attributeCode == nil {
		return ""
	} else {
		constants := method.Constants()
		sb := ""
		code := attributeCode.Code()

		writeByteCode2(linePrefix, &sb, constants, code, fromOffset, toOffset)

		return sb
	}
}

func writeByteCode(linePrefix string, sb *string, constants intcls.IConstantPool, attributeCode intcls.IAttributeCode) {
	code := attributeCode.Code()
	length := len(code)

	*sb += fmt.Sprintf("%sByte code:\n", linePrefix)
	writeByteCode2(linePrefix, sb, constants, code, 0, length)
}

func writeByteCode2(linePrefix string, sb *string, constants intcls.IConstantPool, code []byte, fromOffset, toOffset int) {
	var deltaOffset int
	for offset := fromOffset; offset < toOffset; offset++ {
		opcode := int(code[offset]) & 255
		*sb += fmt.Sprintf("%s  %d: %s", linePrefix, offset, OpcodeNames[opcode])

		switch opcode {
		case 16: // BIPUSH
			offset++
			*sb += fmt.Sprintf(" #%d", int(code[offset])&255)
		case 17: // SIPUSH
			deltaOffset, offset = PrefixReadInt16(code, offset)
			*sb += fmt.Sprintf(" #%d", deltaOffset)
		case 18:
			offset++
			writeLDC(sb, constants, constants.Constant(int(code[offset])&255))
		case 19, 20: // LDC_W, LDC2_W
			deltaOffset, offset = PrefixReadInt16(code, offset)
			writeLDC(sb, constants, constants.Constant(deltaOffset))
		case 21, 22, 23, 24, 25, // ILOAD, LLOAD, FLOAD, DLOAD, ALOAD
			54, 55, 56, 57, 58, // ISTORE, LSTORE, FSTORE, DSTORE, ASTORE
			169: // RET
			offset++
			*sb += fmt.Sprintf(" #%d", int(code[offset])&255)
		case 132: // IINC
			offset++
			*sb += fmt.Sprintf(" #%d", int(code[offset])&255)
			offset++
			*sb += fmt.Sprintf(", %d", int(code[offset])&255)
		case 153, 154, 155, 156, 157, 158, // IFEQ, IFNE, IFLT, IFGE, IFGT, IFLE
			159, 160, 161, 162, 163, 164, 165, 166, // IF_ICMPEQ, IF_ICMPNE, IF_ICMPLT, IF_ICMPGE, IF_ICMPGT, IF_ICMPLE, IF_ACMPEQ, IF_ACMPNE
			167, 168: // GOTO, JSR
			deltaOffset, offset = PrefixReadInt16(code, offset)
			*sb += fmt.Sprintf(" -> %d", offset+deltaOffset)
			break
		case 170: // TABLESWITCH
			// Skip padding
			i := (offset + 4) & 0xFFFC
			deltaOffset, i = SuffixReadInt32(code, i)
			*sb += fmt.Sprintf(" default -> %d", offset+deltaOffset)

			deltaOffset, i = SuffixReadInt32(code, i)
			low := deltaOffset
			deltaOffset, i = SuffixReadInt32(code, i)
			high := deltaOffset

			for value := low; value <= high; value++ {
				deltaOffset, i = PrefixReadInt32(code, i)
				*sb += fmt.Sprintf(", -> %d", offset+deltaOffset)
			}

			offset = i - 1
		case 171: // LOOKUPSWITCH
			// Skip padding
			i := (offset + 4) & 0xFFFC
			deltaOffset, i = SuffixReadInt32(code, i)
			*sb += fmt.Sprintf(" default -> %d", offset+deltaOffset)

			var npairs int
			npairs, i = PrefixReadInt32(code, i)

			for k := 0; k < npairs; k++ {
				deltaOffset, i = SuffixReadInt32(code, i)
				*sb += fmt.Sprintf(", %d", deltaOffset)

				deltaOffset, i = SuffixReadInt32(code, i)
				*sb += fmt.Sprintf(" -> %d", offset+deltaOffset)
			}

			offset = (i - 1)
		case 178, 179: // GETSTATIC, PUTSTATIC
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef := constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
			constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
			descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			*sb += fmt.Sprintf(" %s.%s : %s", typeName, name, descriptor)
		case 180, 181, 182, 183, 184: // GETFIELD, PUTFIELD, INVOKEVIRTUAL, INVOKESPECIAL, INVOKESTATIC
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef := constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
			descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			*sb += fmt.Sprintf(" %s : %s", name, descriptor)
		case 185, 186: // INVOKEINTERFACE, INVOKEDYNAMIC
			deltaOffset, offset = PrefixReadInt16(code, offset)
			constantMemberRef := constants.Constant(deltaOffset).(intcls.IConstantMemberRef)
			constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
			descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			*sb += fmt.Sprintf(" %s : %s", name, descriptor)
			offset += 2 // Skip 2 bytes
		case 187, 189, 192, 193: // NEW, ANEWARRAY, CHECKCAST, INSTANCEOF
			deltaOffset, offset = PrefixReadInt16(code, offset)
			typeName, _ := constants.ConstantTypeName(deltaOffset)
			*sb += fmt.Sprintf(" %s", typeName)
		case 188: // NEWARRAY
			offset++
			switch code[offset] & 255 {
			case 4:
				*sb += " boolean"
			case 5:
				*sb += " char"
			case 6:
				*sb += " float"
			case 7:
				*sb += " double"
			case 8:
				*sb += " byte"
			case 9:
				*sb += " short"
			case 10:
				*sb += " int"
			case 11:
				*sb += " long"
			}
		case 196: // WIDE
			offset++
			opcode = int(code[offset]) & 255
			i := 0
			i, offset = PrefixReadInt16(code, offset)

			if opcode == 132 { // IINC
				deltaOffset, offset = PrefixReadInt16(code, offset)
				*sb += fmt.Sprintf(" iinc #%d %d", i, deltaOffset)
			} else {
				switch opcode {
				case 21:
					*sb += fmt.Sprintf(" iload #%d", i)
				case 22:
					*sb += fmt.Sprintf(" lload #%d", i)
				case 23:
					*sb += fmt.Sprintf(" fload #%d", i)
				case 24:
					*sb += fmt.Sprintf(" dload #%d", i)
				case 25:
					*sb += fmt.Sprintf(" aload #%d", i)
				case 54:
					*sb += fmt.Sprintf(" istore #%d", i)
				case 55:
					*sb += fmt.Sprintf(" lstore #%d", i)
				case 56:
					*sb += fmt.Sprintf(" fstore #%d", i)
				case 57:
					*sb += fmt.Sprintf(" dstore #%d", i)
				case 58:
					*sb += fmt.Sprintf(" astore #%d", i)
				case 169:
					*sb += fmt.Sprintf(" ret #%d", i)
				}
			}
		case 197: // MULTIANEWARRAY
			deltaOffset, offset = PrefixReadInt16(code, offset)
			typeName, _ := constants.ConstantTypeName(deltaOffset)
			offset++
			*sb += fmt.Sprintf("%s %d", typeName, code[offset])
		case 198, 199: // IFNULL, IFNONNULL
			deltaOffset, offset = PrefixReadInt16(code, offset)
			*sb += fmt.Sprintf(" -> %d", offset+deltaOffset)
		case 200, 201: // GOTO_W, JSR_W
			deltaOffset, offset = PrefixReadInt32(code, offset)
			*sb += fmt.Sprintf(" -> %d", offset+deltaOffset)
		}

		*sb += "\n"
	}
}

func writeLDC(sb *string, constants intcls.IConstantPool, cont intcls.IConstant) {
	switch cont.Tag() {
	case intcls.ConstTagInteger:
		*sb += fmt.Sprintf(" %d", cont.(intcls.IConstantInteger).Value())
	case intcls.ConstTagFloat:
		*sb += fmt.Sprintf(" %.2f", cont.(intcls.IConstantFloat).Value())
	case intcls.ConstTagClass:
		typeNameIndex := cont.(intcls.IConstantClass).NameIndex()
		*sb += fmt.Sprintf(" %s", constants.Constant(typeNameIndex).(intcls.IConstantUtf8).Value())
	case intcls.ConstTagLong:
		*sb += fmt.Sprintf(" %d", cont.(intcls.IConstantLong).Value())
	case intcls.ConstTagDouble:
		*sb += fmt.Sprintf(" %.2f", cont.(intcls.IConstantDouble).Value())
	case intcls.ConstTagString:
		*sb += " '"
		stringIndex := cont.(intcls.IConstantString).StringIndex()
		str, _ := constants.ConstantUtf8(stringIndex)

		for _, c := range str {
			switch c {
			case '\b':
				*sb += "\\\\b"
			case '\f':
				*sb += "\\\\f"
			case '\n':
				*sb += "\\\\n"
			case '\r':
				*sb += "\\\\r"
			case '\t':
				*sb += "\\\\t"
			default:
				*sb += fmt.Sprintf("%c", c)
			}
		}
		*sb += "'"
	}
}

func writeLineNumberTable(linePrefix string, sb *string, attributeCode intcls.IAttributeCode) {
	lineNumberTable := attributeCode.Attribute("LineNumberTable").(intcls.IAttributeLineNumberTable)

	if lineNumberTable != nil {
		*sb += fmt.Sprintf("%sLine number table:\n", linePrefix)
		*sb += fmt.Sprintf("%s  Java source line number -> byte code offset\n", linePrefix)

		for _, lineNumber := range lineNumberTable.LineNumberTable() {
			*sb += fmt.Sprintf("%s  #", linePrefix)
			*sb += fmt.Sprintf("%d\t-> ", lineNumber.LineNumber())
			*sb += fmt.Sprintf("%d\n", lineNumber.StartPc())
		}
	}
}

func writeLocalVariableTable(linePrefix string, sb *string, attributeCode intcls.IAttributeCode) {
	localVariableTable := attributeCode.Attribute("LocalVariableTable").(intcls.IAttributeLocalVariableTable)

	if localVariableTable != nil {
		*sb += fmt.Sprintf("%sLocal variable table:\n", linePrefix)
		*sb += fmt.Sprintf("%s  start\tlength\tslot\tname\tdescriptor\n", linePrefix)

		for _, localVariable := range localVariableTable.LocalVariableTable() {
			*sb += fmt.Sprintf("%s  ", linePrefix)
			*sb += fmt.Sprintf("%d\t", localVariable.StartPc())
			*sb += fmt.Sprintf("%d\t", localVariable.Length())
			*sb += fmt.Sprintf("%d\t", localVariable.Index())
			*sb += fmt.Sprintf("%s\t", localVariable.Name())
			*sb += fmt.Sprintf("%s\t", localVariable.Descriptor())
		}
	}

	var localVariableTypeTable intcls.IAttributeLocalVariableTypeTable
	if tmp := attributeCode.Attribute("LocalVariableTypeTable"); tmp != nil {
		localVariableTypeTable = tmp.(intcls.IAttributeLocalVariableTypeTable)
	}

	if localVariableTypeTable != nil {
		*sb += fmt.Sprintf("%sLocal variable type table:\n", linePrefix)
		*sb += fmt.Sprintf("%s  start\tlength\tslot\tname\tdescriptor\n", linePrefix)

		for _, localVariable := range localVariableTypeTable.LocalVariableTypeTable() {
			*sb += fmt.Sprintf("%s  ", linePrefix)
			*sb += fmt.Sprintf("%d\t", localVariable.StartPc())
			*sb += fmt.Sprintf("%d\t", localVariable.Length())
			*sb += fmt.Sprintf("%d\t", localVariable.Index())
			*sb += fmt.Sprintf("%s\t", localVariable.Name())
			*sb += fmt.Sprintf("%s\t", localVariable.Signature())
		}
	}
}

func writeExceptionTable(linePrefix string, sb *string, constants intcls.IConstantPool, attributeCode intcls.IAttributeCode) {
	codeExceptions := attributeCode.ExceptionTable()

	if codeExceptions != nil {
		*sb += fmt.Sprintf("%sException table:\n", linePrefix)
		*sb += fmt.Sprintf("%s  from\tto\ttarget\ttype\n", linePrefix)

		for _, codeException := range codeExceptions {
			*sb += fmt.Sprintf("%s  ", linePrefix)
			*sb += fmt.Sprintf("%d\t", codeException.StartPc())
			*sb += fmt.Sprintf("%d\t", codeException.EndPc())
			*sb += fmt.Sprintf("%d\t", codeException.HandlerPc())

			if codeException.CatchType() == 0 {
				*sb += "finally"
			} else {
				typeName, _ := constants.ConstantTypeName(codeException.CatchType())
				*sb += typeName
			}

			*sb += "\n"
		}
	}
}
