package visitor

import (
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
	modexp "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	modsts "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvexp "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/expression"
	srvsts "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor/utils"
	"github.com/ElectricSaw/go-jd-core/class/util"
	"log"
	"math"
	"reflect"
	"strings"
)

var GlobalJsrReturnAddressExpression = NewJsrReturnAddressExpression()

func NewByteCodeParser(typeMaker intsrv.ITypeMaker,
	localVariableMaker intsrv.ILocalVariableMaker,
	classFile intcls.IClassFile,
	bodyDeclaration intsrv.IClassFileBodyDeclaration,
	comd intsrv.IClassFileConstructorOrMethodDeclaration) intsrv.IByteCodeParser {
	var attributeBootstrapMethods intcls.IAttributeBootstrapMethods
	if tmp := classFile.Attribute("BootstrapMethods"); tmp != nil {
		attributeBootstrapMethods = tmp.(intcls.IAttributeBootstrapMethods)
	}

	p := &ByteCodeParser{
		memberVisitor:                NewByteCodeParserMemberVisitor(),
		searchFirstLineNumberVisitor: NewSearchFirstLineNumberVisitor(),
		eraseTypeArgumentVisitor:     NewEraseTypeArgumentVisitor(),
		lambdaParameterNamesVisitor:  NewLambdaParameterNamesVisitor(),
		renameLocalVariablesVisitor:  NewRenameLocalVariablesVisitor(),

		typeMaker:                 typeMaker,
		localVariableMaker:        localVariableMaker,
		genericTypesSupported:     classFile.MajorVersion() >= 49,
		internalTypeName:          classFile.InternalTypeName(),
		attributeBootstrapMethods: attributeBootstrapMethods,
		bodyDeclaration:           bodyDeclaration,
		returnedType:              comd.ReturnedType(),
		typeBounds:                comd.TypeBounds(),
	}

	if p.genericTypesSupported {
		p.typeParametersToTypeArgumentsBinder = NewJava5TypeParametersToTypeArgumentsBinder(typeMaker, p.internalTypeName, comd)
	} else {
		p.typeParametersToTypeArgumentsBinder = NewJavaTypeParametersToTypeArgumentsBinder()
	}
	return p
}

type ByteCodeParser struct {
	memberVisitor                *ByteCodeParserMemberVisitor
	searchFirstLineNumberVisitor intsrv.ISearchFirstLineNumberVisitor
	eraseTypeArgumentVisitor     intsrv.IEraseTypeArgumentVisitor
	lambdaParameterNamesVisitor  *LambdaParameterNamesVisitor
	renameLocalVariablesVisitor  intsrv.IRenameLocalVariablesVisitor

	typeMaker                           intsrv.ITypeMaker
	localVariableMaker                  intsrv.ILocalVariableMaker
	genericTypesSupported               bool
	internalTypeName                    string
	typeParametersToTypeArgumentsBinder intsrv.ITypeParametersToTypeArgumentsBinder
	attributeBootstrapMethods           intcls.IAttributeBootstrapMethods
	bodyDeclaration                     intsrv.IClassFileBodyDeclaration
	typeBounds                          map[string]intmod.IType
	returnedType                        intmod.IType
}

func (p *ByteCodeParser) Parse(basicBlock intsrv.IBasicBlock, statements intmod.IStatements, stack util.IStack[intmod.IExpression]) {
	cfg := basicBlock.ControlFlowGraph()
	fromOffset := basicBlock.FromOffset()
	toOffset := basicBlock.ToOffset()

	method := cfg.Method()
	constants := method.Constants()
	code := method.Attribute("Code").(intcls.IAttributeCode).Code()
	syntheticFlag := (method.AccessFlags() & intmod.FlagSynthetic) != 0

	var indexRef, arrayRef, valueRef, expression1, expression2, expression3 intmod.IExpression
	var type1, type2, type3 intmod.IType
	var constantMemberRef intcls.IConstantMemberRef
	var constantNameAndType intcls.IConstantNameAndType
	var typeName, name, descriptor string
	var ot intmod.IObjectType
	var i, count, value, lineNumber int
	var localVariable intsrv.ILocalVariable

	for offset := fromOffset; offset < toOffset; offset++ {
		opcode := int(code[offset] & 255)
		if syntheticFlag {
			lineNumber = intmod.UnknownLineNumber
		} else {
			lineNumber = cfg.LineNumber(offset)
		}

		switch opcode {
		case 0: // NOP
		case 1: // ACONST_nil
			stack.Push(modexp.NewNullExpressionWithAll(lineNumber, _type.OtTypeUndefinedObject))
		case 2: // ICONST_M1
			stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, _type.PtMaybeNegativeByteType, -1))
		case 3, 4: // ICONST_0, ICONST_1
			stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, _type.PtMaybeBooleanType, opcode-3))
		case 5, 6, 7, 8: // ICONST_2 ... ICONST_5
			stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, _type.PtMaybeByteType, opcode-3))
		case 9, 10: // LCONST_0, LCONST_1
			stack.Push(modexp.NewLongConstantExpressionWithAll(lineNumber, int64(opcode-9)))
		case 11, 12, 13: // FCONST_0, FCONST_1, FCONST_2
			stack.Push(modexp.NewFloatConstantExpressionWithAll(lineNumber, float32(opcode-11)))
		case 14, 15: // DCONST_0, DCONST_1
			stack.Push(modexp.NewDoubleConstantExpressionWithAll(lineNumber, float64(opcode-14)))
		case 16: // BIPUSH
			value, offset = utils.PrefixReadInt8(code, offset)
			stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, utils.GetPrimitiveTypeFromValue(value), value))
		case 17: // SIPUSH
			value, offset = utils.PrefixReadInt16(code, offset)
			stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, utils.GetPrimitiveTypeFromValue(value), value))
		case 18: // LDC
			value, offset = utils.PrefixReadInt8(code, offset)
			p.parseLDC(stack, constants, lineNumber, constants.Constant(value))
		case 19, 20: // LDC_W, LDC2_W
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parseLDC(stack, constants, lineNumber, constants.Constant(value))
		case 21: // ILOAD
			value, offset = utils.PrefixReadInt8(code, offset)
			localVariable = p.localVariableMaker.LocalVariable(value, offset)
			parseILOAD(statements, stack, lineNumber, offset, localVariable)
		case 22, 23, 24: // LLOAD, FLOAD, DLOAD
			value, offset = utils.PrefixReadInt8(code, offset)
			localVariable = p.localVariableMaker.LocalVariable(value, offset)
			stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
		case 25: // ALOAD
			i, offset = utils.PrefixReadInt8(code, offset)
			localVariable = p.localVariableMaker.LocalVariable(i, offset)
			if (i == 0) && ((method.AccessFlags() & intmod.FlagStatic) == 0) {
				stack.Push(modexp.NewThisExpressionWithAll(lineNumber, localVariable.Type()))
			} else {
				stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
			}
		case 26, 27, 28, 29: // ILOAD_0 ... ILOAD_3
			localVariable = p.localVariableMaker.LocalVariable(opcode-26, offset)
			parseILOAD(statements, stack, lineNumber, offset, localVariable)
		case 30, 31, 32, 33: // LLOAD_0 ... LLOAD_3
			localVariable = p.localVariableMaker.LocalVariable(opcode-30, offset)
			stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
		case 34, 35, 36, 37: // FLOAD_0 ... FLOAD_3
			localVariable = p.localVariableMaker.LocalVariable(opcode-34, offset)
			stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
		case 38, 39, 40, 41: // DLOAD_0 ... DLOAD_3
			localVariable = p.localVariableMaker.LocalVariable(opcode-38, offset)
			stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
		case 42: // ALOAD_0
			localVariable = p.localVariableMaker.LocalVariable(0, offset)
			if (method.AccessFlags() & intmod.FlagStatic) == 0 {
				stack.Push(modexp.NewThisExpressionWithAll(lineNumber, localVariable.Type()))
			} else {
				stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
			}
			break
		case 43, 44, 45: // ALOAD_1 ... ALOAD_3
			localVariable = p.localVariableMaker.LocalVariable(opcode-42, offset)
			stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
			break
		case 46, 47, 48, 49, 50, 51, 52, 53: // IALOAD, LALOAD, FALOAD, DALOAD, AALOAD, BALOAD, CALOAD, SALOAD
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			stack.Push(modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef))
			break
		case 54, 55, 56, 57: // ISTORE, LSTORE, FSTORE, DSTORE
			value, offset = utils.PrefixReadInt8(code, offset)
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(value, offset+2, valueRef)
			p.parseSTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 58: // ASTORE
			value, offset = utils.PrefixReadInt8(code, offset)
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(value, offset+1, valueRef)
			p.parseASTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 59, 60, 61, 62: // ISTORE_0 ... ISTORE_3
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(opcode-59, offset+1, valueRef)
			p.parseSTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 63, 64, 65, 66: // LSTORE_0 ... LSTORE_3
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(opcode-63, offset+1, valueRef)
			p.parseSTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 67, 68, 69, 70: // FSTORE_0 ... FSTORE_3
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(opcode-67, offset+1, valueRef)
			p.parseSTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 71, 72, 73, 74: // DSTORE_0 ... DSTORE_3
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(opcode-71, offset+1, valueRef)
			p.parseSTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 75, 76, 77, 78: // ASTORE_0 ... ASTORE_3
			valueRef = stack.Pop()
			localVariable = p.getLocalVariableInAssignment(opcode-75, offset+1, valueRef)
			p.parseASTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
			break
		case 79: // IASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			type1 = arrayRef.Type()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, type1.CreateType(type1.Dimension()-1),
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 80: // LASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 81: // FASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 82: // DASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 83: // AASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			type1 = arrayRef.Type()
			if type1.Dimension() > 0 {
				type2 = type1.CreateType(type1.Dimension() - 1)
			} else {
				type2 = type1.CreateType(0)
			}
			p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(type2, valueRef)
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, type2,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 84: // BASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeByte,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 85: // CASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeChar,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 86: // SASTORE
			valueRef = stack.Pop()
			indexRef = stack.Pop()
			arrayRef = stack.Pop()
			statements.Add(modsts.NewExpressionStatement(
				modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeShort,
					modexp.NewArrayExpressionWithAll(lineNumber, arrayRef, indexRef), "=", valueRef, 16)))
			break
		case 87, 88: // POP, POP2
			expression1 = stack.Pop()
			if !expression1.IsLocalVariableReferenceExpression() && !expression1.IsFieldReferenceExpression() {
				p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(_type.OtTypeObject, expression1)
				statements.Add(modsts.NewExpressionStatement(expression1))
			}
			break
		case 89: // DUP : ..., value => ..., value, value
			expression1 = stack.Pop()
			stack.Push(expression1)
			stack.Push(expression1)
			break
		case 90: // DUP_X1 : ..., value2, value1 => ..., value1, value2, value1
			expression1 = stack.Pop()
			expression2 = stack.Pop()
			stack.Push(expression1)
			stack.Push(expression2)
			stack.Push(expression1)
			break
		case 91: // DUP_X2
			expression1 = stack.Pop()
			expression2 = stack.Pop()
			type2 = expression2.Type()
			if _type.PtTypeLong == type2 || _type.PtTypeDouble == type2 {
				// ..., value2, value1 => ..., value1, value2, value1
				stack.Push(expression1)
				stack.Push(expression2)
				stack.Push(expression1)
			} else {
				// ..., value3, value2, value1 => ..., value1, value3, value2, value1
				expression3 = stack.Pop()
				stack.Push(expression1)
				stack.Push(expression3)
				stack.Push(expression2)
				stack.Push(expression1)
			}
			break
		case 92: // DUP2
			expression1 = stack.Pop()
			type1 = expression1.Type()
			if _type.PtTypeLong == type1 || _type.PtTypeDouble == type1 {
				// ..., value => ..., value, value
				stack.Push(expression1)
				stack.Push(expression1)
			} else {
				// ..., value2, value1 => ..., value2, value1, value2, value1
				expression2 = stack.Pop()
				stack.Push(expression2)
				stack.Push(expression1)
				stack.Push(expression2)
				stack.Push(expression1)
			}
			break
		case 93: // DUP2_X1
			expression1 = stack.Pop()
			expression2 = stack.Pop()
			type1 = expression1.Type()
			if _type.PtTypeLong == type1 || _type.PtTypeDouble == type1 {
				// ..., value2, value1 => ..., value1, value2, value1
				stack.Push(expression1)
				stack.Push(expression2)
				stack.Push(expression1)
			} else {
				// ..., value3, value2, value1 => ..., value2, value1, value3, value2, value1
				expression3 = stack.Pop()
				stack.Push(expression2)
				stack.Push(expression1)
				stack.Push(expression3)
				stack.Push(expression2)
				stack.Push(expression1)
			}
			break
		case 94: // DUP2_X2
			expression1 = stack.Pop()
			expression2 = stack.Pop()
			type1 = expression1.Type()
			if _type.PtTypeLong == type1 || _type.PtTypeDouble == type1 {
				type2 = expression2.Type()

				if _type.PtTypeLong == type2 || _type.PtTypeDouble == type2 {
					// ..., value2, value1 => ..., value1, value2, value1
					stack.Push(expression1)
					stack.Push(expression2)
					stack.Push(expression1)
				} else {
					// ..., value3, value2, value1 => ..., value1, value3, value2, value1
					expression3 = stack.Pop()
					stack.Push(expression1)
					stack.Push(expression3)
					stack.Push(expression2)
					stack.Push(expression1)
				}
			} else {
				expression3 = stack.Pop()
				type3 = expression3.Type()

				if _type.PtTypeLong == type3 || _type.PtTypeDouble == type3 {
					// ..., value3, value2, value1 => ..., value2, value1, value3, value2, value1
					stack.Push(expression2)
					stack.Push(expression1)
					stack.Push(expression3)
					stack.Push(expression2)
					stack.Push(expression1)
				} else {
					// ..., value4, value3, value2, value1 => ..., value2, value1, value4, value3, value2, value1
					expression4 := stack.Pop()
					stack.Push(expression2)
					stack.Push(expression1)
					stack.Push(expression4)
					stack.Push(expression3)
					stack.Push(expression2)
					stack.Push(expression1)
				}
			}
			break
		case 95: // SWAP : ..., value2, value1 => ..., value1, value2
			expression1 = stack.Pop()
			expression2 = stack.Pop()
			stack.Push(expression1)
			stack.Push(expression2)
			break
		case 96: // IADD
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "+", expression2, 6))
			break
		case 97: // LADD
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "+", expression2, 6))
			break
		case 98: // FADD
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat, expression1, "+", expression2, 6))
			break
		case 99: // DADD
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble, expression1, "+", expression2, 6))
			break
		case 100: // ISUB
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "-", expression2, 6))
			break
		case 101: // LSUB
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "-", expression2, 6))
			break
		case 102: // FSUB
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat, expression1, "-", expression2, 6))
			break
		case 103: // DSUB
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble, expression1, "-", expression2, 6))
			break
		case 104: // IMUL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "*", expression2, 5))
			break
		case 105: // LMUL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "*", expression2, 5))
			break
		case 106: // FMUL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat, expression1, "*", expression2, 5))
			break
		case 107: // DMUL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble, expression1, "*", expression2, 5))
			break
		case 108: // IDIV
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "/", expression2, 5))
			break
		case 109: // LDIV
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "/", expression2, 5))
			break
		case 110: // FDIV
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat, expression1, "/", expression2, 5))
			break
		case 111: // DDIV
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble, expression1, "/", expression2, 5))
			break
		case 112: // IREM
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "%", expression2, 5))
			break
		case 113: // LREM
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "%", expression2, 5))
			break
		case 114: // FREM
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat, expression1, "%", expression2, 5))
			break
		case 115: // DREM
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble, expression1, "%", expression2, 5))
			break
		case 116, 117, 118, 119: // INEG, LNEG, FNEG, DNEG
			stack.Push(p.newPreArithmeticOperatorExpression(lineNumber, "-", stack.Pop()))
			break
		case 120: // ISHL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, "<<", expression2, 7))
			break
		case 121: // LSHL
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "<<", expression2, 7))
			break
		case 122: // ISHR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeInt, expression1, ">>", expression2, 7))
			break
		case 123: // LSHR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, ">>", expression2, 7))
			break
		case 124: // IUSHR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerBinaryOperatorExpression(lineNumber, expression1, ">>>", expression2, 7))
			break
		case 125: // LUSHR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, ">>>", expression2, 7))
			break
		case 126: // IAND
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerOrBooleanBinaryOperatorExpression(lineNumber, expression1, "&", expression2, 10))
			break
		case 127: // LAND
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "&", expression2, 10))
			break
		case 128: // IOR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerOrBooleanBinaryOperatorExpression(lineNumber, expression1, "|", expression2, 12))
			break
		case 129: // LOR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "|", expression2, 12))
			break
		case 130: // IXOR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(p.newIntegerOrBooleanBinaryOperatorExpression(lineNumber, expression1, "^", expression2, 11))
			break
		case 131: // LXOR
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong, expression1, "^", expression2, 11))
			break
		case 132: // IINC
			value, offset = utils.PrefixReadInt8(code, offset)
			localVariable = p.localVariableMaker.LocalVariable(value, offset)
			value, offset = utils.PrefixReadInt8(code, offset)
			p.parseIINC(statements, stack, lineNumber, offset, localVariable, value)
			break
		case 133: // I2L
			stack.Push(modexp.NewCastExpressionWithAll(lineNumber, _type.PtTypeLong, stack.Pop(), false))
			break
		case 134: // I2F
			stack.Push(modexp.NewCastExpressionWithAll(lineNumber, _type.PtTypeFloat, stack.Pop(), false))
			break
		case 135: // I2D
			stack.Push(modexp.NewCastExpressionWithAll(lineNumber, _type.PtTypeDouble, stack.Pop(), false))
			break
		case 136: // L2I
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeInt, forceExplicitCastExpression(stack.Pop())))
			break
		case 137: // L2F
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeFloat, forceExplicitCastExpression(stack.Pop())))
			break
		case 138: // L2D
			stack.Push(modexp.NewCastExpressionWithAll(lineNumber, _type.PtTypeDouble, stack.Pop(), false))
			break
		case 139: // F2I
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeInt, forceExplicitCastExpression(stack.Pop())))
			break
		case 140: // F2L
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeLong, forceExplicitCastExpression(stack.Pop())))
			break
		case 141: // F2D
			stack.Push(modexp.NewCastExpressionWithAll(lineNumber, _type.PtTypeDouble, stack.Pop(), false))
			break
		case 142: // D2I
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeInt, forceExplicitCastExpression(stack.Pop())))
			break
		case 143: // D2L
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeLong, forceExplicitCastExpression(stack.Pop())))
			break
		case 144: // D2F
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeFloat, forceExplicitCastExpression(stack.Pop())))
			break
		case 145: // I2B
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeByte, forceExplicitCastExpression(stack.Pop())))
			break
		case 146: // I2C
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeChar, forceExplicitCastExpression(stack.Pop())))
			break
		case 147: // I2S
			stack.Push(modexp.NewCastExpressionWithLineNumber(lineNumber, _type.PtTypeShort, forceExplicitCastExpression(stack.Pop())))
			break
		case 148, 149, 150, 151, 152: // LCMP, FCMPL, FCMPG, DCMPL, DCMPG
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			stack.Push(srvexp.NewClassFileCmpExpression(lineNumber, expression1, expression2))
			break
		case 153: // IFEQ
			p.parseIF(stack, lineNumber, basicBlock, "!=", "==", 8)
			offset += 2 // Skip branch offset
			break
		case 154: // IFNE
			p.parseIF(stack, lineNumber, basicBlock, "==", "!=", 8)
			offset += 2 // Skip branch offset
			break
		case 155: // IFLT
			p.parseIF(stack, lineNumber, basicBlock, ">=", "<", 7)
			offset += 2 // Skip branch offset
			break
		case 156: // IFGE
			p.parseIF(stack, lineNumber, basicBlock, "<", ">=", 7)
			offset += 2 // Skip branch offset
			break
		case 157: // IFGT
			p.parseIF(stack, lineNumber, basicBlock, "<=", ">", 7)
			offset += 2 // Skip branch offset
			break
		case 158: // IFLE
			p.parseIF(stack, lineNumber, basicBlock, ">", "<=", 7)
			offset += 2 // Skip branch offset
			break
		case 159: // IF_ICMPEQ
		case 165: // IF_ACMPEQ
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := "=="
			if basicBlock.IsInverseCondition() {
				op = "!="
			}
			stack.Push(p.newIntegerOrBooleanComparisonOperatorExpression(lineNumber, expression1, op, expression2, 9))
			offset += 2 // Skip branch offset
			break
		case 160: // IF_ICMPNE
		case 166: // IF_ACMPNE
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := "=="
			if basicBlock.IsInverseCondition() {
				op = "!="
			}
			stack.Push(p.newIntegerOrBooleanComparisonOperatorExpression(lineNumber, expression1, op, expression2, 9))
			offset += 2 // Skip branch offset
			break
		case 161: // IF_ICMPLT
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := "<"
			if basicBlock.IsInverseCondition() {
				op = ">="
			}
			stack.Push(p.newIntegerComparisonOperatorExpression(lineNumber, expression1, op, expression2, 8))
			offset += 2 // Skip branch offset
			break
		case 162: // IF_ICMPGE
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := ">="
			if basicBlock.IsInverseCondition() {
				op = "<"
			}
			stack.Push(p.newIntegerComparisonOperatorExpression(lineNumber, expression1, op, expression2, 8))
			offset += 2 // Skip branch offset
			break
		case 163: // IF_ICMPGT
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := ">"
			if basicBlock.IsInverseCondition() {
				op = "<="
			}
			stack.Push(p.newIntegerComparisonOperatorExpression(lineNumber, expression1, op, expression2, 8))
			offset += 2 // Skip branch offset
			break
		case 164: // IF_ICMPLE
			expression2 = stack.Pop()
			expression1 = stack.Pop()
			op := "<="
			if basicBlock.IsInverseCondition() {
				op = ">"
			}
			stack.Push(p.newIntegerComparisonOperatorExpression(lineNumber, expression1, op, expression2, 8))
			offset += 2 // Skip branch offset
			break
		case 168: // JSR
			stack.Push(GlobalJsrReturnAddressExpression)
		case 167: // GOTO
			offset += 2 // Skip branch offset
			break
		case 169: // RET
			offset++ // Skip index
			break
		case 170: // TABLESWITCH
			offset = (offset + 4) & 0xFFFC // Skip padding
			offset += 4                    // Skip default offset

			var low, high int
			low, offset = utils.SuffixReadInt32(code, offset)
			high, offset = utils.SuffixReadInt32(code, offset)
			offset += (4 * (high - low + 1)) - 1
			statements.Add(modsts.NewSwitchStatement(stack.Pop(),
				util.NewDefaultListWithCapacity[intmod.IBlock](high-low+2)))
		case 171: // LOOKUPSWITCH
			offset = (offset + 4) & 0xFFFC // Skip padding
			offset += 4                    // Skip default offset
			count, offset = utils.SuffixReadInt32(code, offset)
			offset += (8 * count) - 1
			statements.Add(modsts.NewSwitchStatement(stack.Pop(),
				util.NewDefaultListWithCapacity[intmod.IBlock](count+1)))
		case 172, 173, 174, 175, 176: // IRETURN, LRETURN, FRETURN, DRETURN, ARETURN
			p.parseXRETURN(statements, stack, lineNumber)
		case 177: // RETURN
			statements.Add(modsts.Return)
		case 178: // GETSTATIC
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parseGetStatic(stack, constants, lineNumber, value)
		case 179: // PUTSTATIC
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parsePutStatic(statements, stack, constants, lineNumber, value)
			break
		case 180: // GETFIELD
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parseGetField(stack, constants, lineNumber, value)
			break
		case 181: // PUTFIELD
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parsePutField(statements, stack, constants, lineNumber, value)
			break
		case 182, 183, 184, 185: // INVOKEVIRTUAL, INVOKESPECIAL, INVOKESTATIC, INVOKEINTERFACE
			value, offset = utils.PrefixReadInt16(code, offset)
			constantMemberRef = constants.Constant(value).(intcls.IConstantMemberRef)
			typeName, _ = constants.ConstantTypeName(constantMemberRef.ClassIndex())
			ot = p.typeMaker.MakeFromDescriptorOrInternalTypeName(typeName)
			constantNameAndType = constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
			name, _ = constants.ConstantUtf8(constantNameAndType.NameIndex())
			descriptor, _ = constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			methodTypes := p.makeMethodTypes(ot.InternalName(), name, descriptor)
			parameters := p.extractParametersFromStack(statements, stack, methodTypes.ParameterTypes())
			if opcode == 184 { // INVOKESTATIC
				expression1 = p.typeParametersToTypeArgumentsBinder.NewMethodInvocationExpression(
					lineNumber, modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, ot),
					ot, name, descriptor, methodTypes, parameters)
				if _type.PtTypeVoid == methodTypes.ReturnedType() {
					p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(_type.OtTypeObject, expression1)
					statements.Add(modsts.NewExpressionStatement(expression1))
				} else {
					stack.Push(expression1)
				}
			} else {
				expression1 = stack.Pop()
				if expression1.IsLocalVariableReferenceExpression() {
					expression1.(intsrv.IClassFileLocalVariableReferenceExpression).
						LocalVariable().(intsrv.ILocalVariable).TypeOnLeft(p.typeBounds, ot)
				}
				if opcode == 185 { // INVOKEINTERFACE
					offset += 2 // Skip 'count' and one byte
				}
				if _type.PtTypeVoid == methodTypes.ReturnedType() {
					if opcode == 183 && // INVOKESPECIAL
						"<init>" == name {

						if expression1.IsNewExpression() {
							p.typeParametersToTypeArgumentsBinder.UpdateNewExpression(
								expression1.(intsrv.IClassFileNewExpression), descriptor, methodTypes, parameters)
						} else if ot.Descriptor() == expression1.Type().Descriptor() {
							statements.Add(modsts.NewExpressionStatement(
								p.typeParametersToTypeArgumentsBinder.NewConstructorInvocationExpression(
									lineNumber, ot, descriptor, methodTypes, parameters)))
						} else {
							statements.Add(modsts.NewExpressionStatement(
								p.typeParametersToTypeArgumentsBinder.NewSuperConstructorInvocationExpression(
									lineNumber, ot, descriptor, methodTypes, parameters)))
						}
					} else {
						expression1 = p.typeParametersToTypeArgumentsBinder.NewMethodInvocationExpression(
							lineNumber, p.getMethodInstanceReference(expression1, ot, name, descriptor),
							ot, name, descriptor, methodTypes, parameters)
						p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(_type.OtTypeObject, expression1)
						statements.Add(modsts.NewExpressionStatement(expression1))
					}
				} else {
					if opcode == 182 { // INVOKEVIRTUAL
						if "toString" == name && "()Ljava/lang/String;" == descriptor {
							typeName, _ = constants.ConstantTypeName(constantMemberRef.ClassIndex())
							if "java/lang/StringBuilder" == typeName || "java/lang/StringBuffer" == typeName {
								stack.Push(utils.StringConcatenationUtilCreate1(expression1, lineNumber, typeName))
								break
							}
						}
					}
					stack.Push(p.typeParametersToTypeArgumentsBinder.
						NewMethodInvocationExpression(lineNumber,
							p.getMethodInstanceReference(expression1, ot, name, descriptor),
							ot, name, descriptor, methodTypes, parameters))
				}
			}
			break
		case 186: // INVOKEDYNAMIC
			value, offset = utils.PrefixReadInt16(code, offset)
			p.parseInvokeDynamic(statements, stack, constants, lineNumber, value)
			offset += 2 // Skip 2 bytes
			break
		case 187: // NEW
			value, offset = utils.PrefixReadInt16(code, offset)
			typeName, _ = constants.ConstantTypeName(value)
			stack.Push(p.newNewExpression(lineNumber, typeName))
			break
		case 188: // NEWARRAY
			value, offset = utils.PrefixReadInt8(code, offset)
			type1 = utils.GetPrimitiveTypeFromTag(value).CreateType(1)
			stack.Push(modexp.NewNewArray(lineNumber, type1, stack.Pop()))
			break
		case 189: // ANEWARRAY
			value, offset = utils.PrefixReadInt16(code, offset)
			typeName, _ = constants.ConstantTypeName(value)
			if typeName[0] == '[' {
				type1 = p.typeMaker.MakeFromDescriptor(typeName).(intmod.IType)
				type1 = type1.CreateType(type1.Dimension() + 1)
			} else {
				type1 = p.typeMaker.MakeFromInternalTypeName(typeName).CreateType(1)
			}
			if strings.HasSuffix(typeName, _type.OtTypeClass.InternalName()) {
				ot = type1.(intmod.IObjectType)
				if ot.TypeArguments() == nil {
					type1 = ot.CreateTypeWithArgs(_type.WildcardTypeArgumentEmpty)
				}
			}
			stack.Push(modexp.NewNewArray(lineNumber, type1, stack.Pop()))
			break
		case 190: // ARRAYLENGTH
			stack.Push(modexp.NewLengthExpressionWithAll(lineNumber, stack.Pop()))
			break
		case 191: // ATHROW
			statements.Add(modsts.NewThrowStatement(stack.Pop()))
			break
		case 192: // CHECKCAST
			value, offset = utils.PrefixReadInt16(code, offset)
			typeName, _ = constants.ConstantTypeName(value)
			type1 = p.typeMaker.MakeFromDescriptorOrInternalTypeName(typeName)
			expression1 = stack.Peek()
			if type1.IsObjectType() && expression1.Type().IsObjectType() &&
				p.typeMaker.IsRawTypeAssignable(type1.(intmod.IObjectType),
					expression1.Type().(intmod.IObjectType)) {
				// Ignore cast
			} else if expression1.IsCastExpression() {
				// Skip double cast
				expression1.(intmod.ICastExpression).SetType(type1)
			} else {
				p.searchFirstLineNumberVisitor.Init()
				expression1.Accept(p.searchFirstLineNumberVisitor)
				stack.Push(modexp.NewCastExpressionWithLineNumber(p.searchFirstLineNumberVisitor.LineNumber(),
					type1, forceExplicitCastExpression(stack.Pop())))
			}
		case 193: // INSTANCEOF
			value, offset = utils.PrefixReadInt16(code, offset)
			typeName, _ = constants.ConstantTypeName(value)
			type1 = p.typeMaker.MakeFromDescriptorOrInternalTypeName(typeName)
			if type1 == nil {
				type1 = utils.GetPrimitiveTypeFromDescriptor(typeName)
			}
			stack.Push(modexp.NewInstanceOfExpressionWithAll(lineNumber, stack.Pop(), type1.(intmod.IObjectType)))
		case 194: // MONITORENTER
			statements.Add(srvsts.NewClassFileMonitorEnterStatement(stack.Pop()))
		case 195: // MONITOREXIT
			statements.Add(srvsts.NewClassFileMonitorExitStatement(stack.Pop()))
		case 196: // WIDE
			opcode, offset = utils.PrefixReadInt8(code, offset)
			i, offset = utils.PrefixReadInt16(code, offset)

			if opcode == 132 { // IINC
				count, offset = utils.PrefixReadInt16(code, offset)
				p.parseIINC(statements, stack, lineNumber, offset, p.localVariableMaker.LocalVariable(i, offset), count)
			} else {
				switch opcode {
				case 21: // ILOAD
					localVariable = p.localVariableMaker.LocalVariable(i, offset+4)
					parseILOAD(statements, stack, offset, lineNumber, localVariable)
				case 22, 23, 24, 25: // LLOAD, FLOAD, DLOAD, ALOAD
					stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(
						lineNumber, offset, p.localVariableMaker.LocalVariable(i, offset)))
				case 54: // ISTORE
					valueRef = stack.Pop()
					localVariable = p.getLocalVariableInAssignment(i, offset+4, valueRef)
					statements.Add(modsts.NewExpressionStatement(
						modexp.NewBinaryOperatorExpression(lineNumber, localVariable.Type(),
							srvexp.NewClassFileLocalVariableReferenceExpression(
								lineNumber, offset, localVariable), "=", valueRef, 16)))
				case 55: // LSTORE
					valueRef = stack.Pop()
					localVariable = p.getLocalVariableInAssignment(i, offset+4, valueRef)
					statements.Add(modsts.NewExpressionStatement(
						modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeLong,
							srvexp.NewClassFileLocalVariableReferenceExpression(
								lineNumber, offset, localVariable), "=", valueRef, 16)))
				case 56: // FSTORE
					valueRef = stack.Pop()
					localVariable = p.getLocalVariableInAssignment(i, offset+4, valueRef)
					statements.Add(modsts.NewExpressionStatement(
						modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeFloat,
							srvexp.NewClassFileLocalVariableReferenceExpression(
								lineNumber, offset, localVariable), "=", valueRef, 16)))
				case 57: // DSTORE
					valueRef = stack.Pop()
					localVariable = p.getLocalVariableInAssignment(i, offset+4, valueRef)
					statements.Add(modsts.NewExpressionStatement(
						modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeDouble,
							srvexp.NewClassFileLocalVariableReferenceExpression(
								lineNumber, offset, localVariable), "=", valueRef, 16)))
				case 58: // ASTORE
					valueRef = stack.Pop()
					localVariable = p.getLocalVariableInAssignment(i, offset+4, valueRef)
					p.parseASTORE(statements, stack, lineNumber, offset, localVariable, valueRef)
				case 169: // RET
				}
			}
		case 197: // MULTIANEWARRAY
			value, offset = utils.PrefixReadInt16(code, offset)
			typeName, _ = constants.ConstantTypeName(value)
			type1 = p.typeMaker.MakeFromDescriptor(typeName)
			i, offset = utils.PrefixReadInt8(code, offset)

			dimensions := modexp.NewExpressions()
			for ; i > 0; i-- {
				dimensions.Add(stack.Pop())
			}
			dimensions.Reverse()
			stack.Push(modexp.NewNewArray(lineNumber, type1, dimensions))
		case 198: // IFnil
			expression1 = stack.Pop()
			p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(_type.OtTypeObject, expression1)
			op := "=="
			if basicBlock.IsInverseCondition() {
				op = "!="
			}
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression1, op,
				modexp.NewNullExpressionWithAll(expression1.LineNumber(), expression1.Type()), 9))
			offset += 2 // Skip branch offset
			checkStack(stack, code, offset)
		case 199: // IFNONnil
			expression1 = stack.Pop()
			p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(_type.OtTypeObject, expression1)
			op := "=="
			if basicBlock.IsInverseCondition() {
				op = "!="
			}
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression1, op,
				modexp.NewNullExpressionWithAll(expression1.LineNumber(), expression1.Type()), 9))
			offset += 2 // Skip branch offset
			checkStack(stack, code, offset)
		case 201: // JSR_W
			stack.Push(GlobalJsrReturnAddressExpression)
			fallthrough
		case 200: // GOTO_W
			offset += 4 // Skip branch offset
		}
	}
}

func (p *ByteCodeParser) extractParametersFromStack(statements intmod.IStatements,
	stack util.IStack[intmod.IExpression], parameterTypes intmod.IType) intmod.IExpression {
	if parameterTypes == nil {
		return nil
	}

	switch parameterTypes.Size() {
	case 0:
		return nil
	case 1:
		parameter := stack.Pop()
		if parameter.IsNewArray() {
			parameter = utils.MakeNewArrayMaker(statements, parameter)
		}
		return p.checkIfLastStatementIsAMultiAssignment(statements, parameter)
	default:
		parameters := modexp.NewExpressions()
		count := parameterTypes.Size() - 1
		for i := count; i >= 0; i-- {
			parameter := stack.Pop()
			if parameter.IsNewArray() {
				parameter = utils.MakeNewArrayMaker(statements, parameter)
			}
			parameters.Add(p.checkIfLastStatementIsAMultiAssignment(statements, parameter))
		}

		parameters.Reverse()
		return parameters
	}
}

func (p *ByteCodeParser) checkIfLastStatementIsAMultiAssignment(statements intmod.IStatements,
	parameter intmod.IExpression) intmod.IExpression {
	if !statements.IsEmpty() {
		expr := statements.Last().Expression()

		if expr.IsBinaryOperatorExpression() && (getLastRightExpression(expr) == parameter) {
			// Return multi assignment expression
			statements.RemoveLast()
			return expr
		}
	}

	return parameter
}

func (p *ByteCodeParser) getLocalVariableInAssignment(index, offset int, value intmod.IExpression) intsrv.ILocalVariable {
	valueType := value.Type()

	if value.IsNullExpression() {
		return p.localVariableMaker.LocalVariableInNullAssignment(index, offset, valueType)
	} else if value.IsLocalVariableReferenceExpression() {
		valueLocalVariable := value.(intsrv.IClassFileLocalVariableReferenceExpression).
			LocalVariable().(intsrv.ILocalVariable)
		lv := p.localVariableMaker.LocalVariableInAssignmentWithLocalVariable(p.typeBounds, index, offset, valueLocalVariable)
		valueLocalVariable.VariableOnLeft(p.typeBounds, lv)
		return lv
	} else if value.IsMethodInvocationExpression() {
		if valueType.IsObjectType() {
			// Remove type arguments
			valueType = valueType.(intmod.IObjectType).CreateTypeWithArgs(nil)
		} else if valueType.IsGenericType() {
			valueType = _type.OtTypeUndefinedObject
		}
		return p.localVariableMaker.LocalVariableInAssignment(p.typeBounds, index, offset, valueType)
	} else {
		return p.localVariableMaker.LocalVariableInAssignment(p.typeBounds, index, offset, valueType)
	}
}

func (p *ByteCodeParser) parseLDC(stack util.IStack[intmod.IExpression], constants intcls.IConstantPool,
	lineNumber int, constant intcls.IConstant) {
	switch constant.Tag() {
	case intcls.ConstTagInteger:
		i := constant.(intcls.IConstantInteger).Value()
		stack.Push(modexp.NewIntegerConstantExpressionWithAll(lineNumber, utils.GetPrimitiveTypeFromValue(i), i))
	case intcls.ConstTagFloat:
		f := constant.(intcls.IConstantFloat).Value()
		if f == -math.MaxFloat32 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeFloat,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeFloat),
				"java/lang/Float", "MIN_VALUE", "F"))
		} else if f == math.MaxFloat32 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeFloat,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeFloat),
				"java/lang/Float", "MAX_VALUE", "F"))
		} else if f == float32(math.Inf(-1)) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeFloat,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeFloat),
				"java/lang/Float", "NEGATIVE_INFINITY", "F"))
		} else if f == float32(math.Inf(1)) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeFloat,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeFloat),
				"java/lang/Float", "POSITIVE_INFINITY", "F"))
		} else if math.IsNaN(float64(f)) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeFloat,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeFloat),
				"java/lang/Float", "NaN", "F"))
		} else {
			stack.Push(modexp.NewFloatConstantExpressionWithAll(lineNumber, f))
		}
	case intcls.ConstTagClass:
		typeNameIndex := constant.(intcls.IConstantClass).NameIndex()
		typeName := constants.Constant(typeNameIndex).(intcls.IConstantUtf8).Value()
		typ := p.typeMaker.MakeFromDescriptorOrInternalTypeName(typeName).(intmod.IType)
		if typ == nil {
			typ = utils.GetPrimitiveTypeFromDescriptor(typeName)
		}
		stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, typ))
	case intcls.ConstTagLong:
		l := constant.(intcls.IConstantLong).Value()
		if l == math.MinInt64 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeLong,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeLong),
				"java/lang/Long", "MIN_VALUE", "J"))
		} else if l == math.MaxInt64 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeLong,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeLong),
				"java/lang/Long", "MAX_VALUE", "J"))
		} else {
			stack.Push(modexp.NewLongConstantExpressionWithAll(lineNumber, l))
		}
	case intcls.ConstTagDouble:
		d := constant.(intcls.IConstantDouble).Value()
		if d == math.SmallestNonzeroFloat64 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeDouble),
				"java/lang/Double", "MIN_VALUE", "D"))
		} else if d == math.MaxFloat64 {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeDouble),
				"java/lang/Double", "MAX_VALUE", "D"))
		} else if d == math.Inf(-1) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeDouble),
				"java/lang/Double", "NEGATIVE_INFINITY", "D"))
		} else if d == math.Inf(1) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeDouble),
				"java/lang/Double", "POSITIVE_INFINITY", "D"))
		} else if math.IsNaN(d) {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeDouble),
				"java/lang/Double", "NaN", "D"))
		} else if d == math.E {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeMath),
				"java/lang/Math", "E", "D"))
		} else if d == math.Pi {
			stack.Push(modexp.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeDouble,
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, _type.OtTypeMath),
				"java/lang/Math", "PI", "D"))
		} else {
			stack.Push(modexp.NewDoubleConstantExpressionWithAll(lineNumber, d))
		}
		break
	case intcls.ConstTagString:
		stringIndex := constant.(intcls.IConstantString).StringIndex()
		str, _ := constants.ConstantUtf8(stringIndex)
		stack.Push(modexp.NewStringConstantExpressionWithAll(lineNumber, str))
		break
	}
}

func parseILOAD(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber, offset int, localVariable intsrv.ILocalVariable) {
	if !statements.IsEmpty() {
		expr := statements.Last().Expression()

		if expr.LineNumber() == lineNumber && expr.IsPreOperatorExpression() && expr.Expression().IsLocalVariableReferenceExpression() {
			cflvre := expr.Expression().(intsrv.IClassFileLocalVariableReferenceExpression)

			if cflvre.LocalVariable() == localVariable {
				// IINC pattern found -> Remove last statement and create a pre-incrementation
				statements.RemoveLast()
				stack.Push(expr)
				return
			}
		}
	}

	stack.Push(srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable))
}

func (p *ByteCodeParser) parseSTORE(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber, offset int, localVariable intsrv.ILocalVariable, valueRef intmod.IExpression) {
	vre := srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable)

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(vre.Type(), valueRef)

	if valueRef.LineNumber() == lineNumber && valueRef.IsBinaryOperatorExpression() && valueRef.LeftExpression().IsLocalVariableReferenceExpression() {
		lvr := valueRef.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)

		if lvr.LocalVariable() == localVariable {
			boe := valueRef.(intmod.IBinaryOperatorExpression)
			var expr intmod.IExpression

			switch boe.Operator() {
			case "*":
				expr = createAssignment(boe, "*=")
			case "/":
				expr = createAssignment(boe, "/=")
			case "%":
				expr = createAssignment(boe, "%=")
			case "<<":
				expr = createAssignment(boe, "<<=")
			case ">>":
				expr = createAssignment(boe, ">>=")
			case ">>>":
				expr = createAssignment(boe, ">>>=")
			case "&":
				expr = createAssignment(boe, "&=")
			case "^":
				expr = createAssignment(boe, "^=")
			case "|":
				expr = createAssignment(boe, "|=")
			case "=":
				expr = boe
			case "+":
				if isPositiveOne(boe.RightExpression()) {
					if stackContainsLocalVariableReference(stack, localVariable) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "++")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "++", boe.LeftExpression())
					}
				} else if isNegativeOne(boe.RightExpression()) {
					if stackContainsLocalVariableReference(stack, localVariable) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "--")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "--", boe.LeftExpression())
					}
				} else {
					expr = createAssignment(boe, "+=")
				}
			case "-":
				if isPositiveOne(boe.RightExpression()) {
					if stackContainsLocalVariableReference(stack, localVariable) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "--")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "--", boe.LeftExpression())
					}
				} else if isNegativeOne(boe.RightExpression()) {
					if stackContainsLocalVariableReference(stack, localVariable) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "++")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "++", boe.LeftExpression())
					}
				} else {
					expr = createAssignment(boe, "-=")
				}
				break
			default:
				log.Println("Unexpected value expression")
			}

			if !stack.IsEmpty() && (stack.Peek() == valueRef) {
				stack.Replace(valueRef, expr)
			} else {
				statements.Add(modsts.NewExpressionStatement(expr))
			}
			return
		}
	}

	p.createAssignment(statements, stack, lineNumber, vre, valueRef)
}

func stackContainsLocalVariableReference(stack util.IStack[intmod.IExpression], localVariable intsrv.ILocalVariable) bool {
	if !stack.IsEmpty() {
		expr := stack.Peek()

		if expr.IsLocalVariableReferenceExpression() {
			lvr := expr.(intsrv.IClassFileLocalVariableReferenceExpression)
			return lvr.LocalVariable() == localVariable
		}
	}

	return false
}

func (p *ByteCodeParser) parsePUT(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber int, fr intmod.IFieldReferenceExpression, valueRef intmod.IExpression) {
	if valueRef.IsNewArray() {
		valueRef = utils.MakeNewArrayMaker(statements, valueRef)
	}

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(fr.Type(), valueRef)

	if valueRef.LineNumber() == lineNumber && valueRef.IsBinaryOperatorExpression() && valueRef.LeftExpression().IsFieldReferenceExpression() {
		boefr := valueRef.LeftExpression().(intmod.IFieldReferenceExpression)

		if boefr.Name() == fr.Name() && boefr.Expression().Type() == fr.Expression().Type() {
			boe := valueRef.(intmod.IBinaryOperatorExpression)
			var expr intmod.IExpression

			switch boe.Operator() {
			case "*":
				expr = createAssignment(boe, "*=")
			case "/":
				expr = createAssignment(boe, "/=")
			case "%":
				expr = createAssignment(boe, "%=")
			case "<<":
				expr = createAssignment(boe, "<<=")
			case ">>":
				expr = createAssignment(boe, ">>=")
			case ">>>":
				expr = createAssignment(boe, ">>>=")
			case "&":
				expr = createAssignment(boe, "&=")
			case "^":
				expr = createAssignment(boe, "^=")
			case "|":
				expr = createAssignment(boe, "|=")
			case "=":
				expr = boe
			case "+":
				if isPositiveOne(boe.RightExpression()) {
					if stackContainsFieldReference(stack, fr) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "++")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "++", boe.LeftExpression())
					}
				} else if isNegativeOne(boe.RightExpression()) {
					if stackContainsFieldReference(stack, fr) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "--")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "--", boe.LeftExpression())
					}
				} else {
					expr = createAssignment(boe, "+=")
				}
				break
			case "-":
				if isPositiveOne(boe.RightExpression()) {
					if stackContainsFieldReference(stack, fr) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "--")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "--", boe.LeftExpression())
					}
				} else if isPositiveOne(boe.RightExpression()) {
					if stackContainsFieldReference(stack, fr) {
						stack.Pop()
						stack.Push(valueRef)
						expr = p.newPostArithmeticOperatorExpression(boe.LineNumber(), boe.LeftExpression(), "++")
					} else {
						expr = p.newPreArithmeticOperatorExpression(boe.LineNumber(), "++", boe.LeftExpression())
					}
				} else {
					expr = createAssignment(boe, "-=")
				}
				break
			default:
				log.Println("Unexpected value expression")
			}

			if !stack.IsEmpty() && (stack.Peek() == valueRef) {
				stack.Replace(valueRef, expr)
			} else {
				statements.Add(modsts.NewExpressionStatement(expr))
			}
			return
		}
	}

	p.createAssignment(statements, stack, lineNumber, fr, valueRef)
}

func (p *ByteCodeParser) parseInvokeDynamic(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	constants intcls.IConstantPool, lineNumber, index int) {
	// Remove previous 'getClass()' or cast if exists
	if !statements.IsEmpty() {
		expression := statements.Last().Expression()

		if expression.IsMethodInvocationExpression() {
			mie := expression.(intmod.IMethodInvocationExpression)

			if mie.Name() == "getClass" && mie.Descriptor() == "()Ljava/lang/Class;" && mie.InternalTypeName() == "java/lang/Object" {
				statements.RemoveLast()
			}
		}
	}

	// Create expression
	constantMemberRef := constants.Constant(index).(intcls.IConstantMemberRef)
	indyCnat := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	indyMethodName, _ := constants.ConstantUtf8(indyCnat.NameIndex())
	indyDescriptor, _ := constants.ConstantUtf8(indyCnat.DescriptorIndex())
	indyMethodTypes := p.typeMaker.MakeMethodTypes(indyDescriptor)

	indyParameters := p.extractParametersFromStack(statements, stack, indyMethodTypes.ParameterTypes())
	bootstrapMethod := p.attributeBootstrapMethods.BootstrapMethods()[constantMemberRef.ClassIndex()]
	bootstrapArguments := bootstrapMethod.BootstrapArguments()

	if "makeConcatWithConstants" == indyMethodName {
		// Create Java 9+ string concatenation
		recipe, _ := constants.ConstantString(bootstrapArguments[0])
		stack.Push(utils.StringConcatenationUtilCreate2(recipe, indyParameters))
		return
	} else if "makeConcat" == indyMethodName {
		// Create Java 9+ string concatenation
		stack.Push(utils.StringConcatenationUtilCreate3(indyParameters))
		return
	}

	cmt0 := constants.Constant(bootstrapArguments[0]).(intcls.IConstantMethodType)
	descriptor0, _ := constants.ConstantUtf8(cmt0.DescriptorIndex())
	methodTypes0 := p.typeMaker.MakeMethodTypes(descriptor0)
	parameterCount := 0
	if methodTypes0.ParameterTypes() != nil {
		parameterCount = methodTypes0.ParameterTypes().Size()
	}
	constantMethodHandle1 := constants.Constant(bootstrapArguments[1]).(intcls.IConstantMethodHandle)
	cmr1 := constants.Constant(constantMethodHandle1.ReferenceIndex()).(intcls.IConstantMemberRef)
	typeName, _ := constants.ConstantTypeName(cmr1.ClassIndex())
	cnat1 := constants.Constant(cmr1.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name1, _ := constants.ConstantUtf8(cnat1.NameIndex())
	descriptor1, _ := constants.ConstantUtf8(cnat1.DescriptorIndex())

	if typeName == p.internalTypeName {
		for _, methodDeclaration := range p.bodyDeclaration.MethodDeclarations() {
			if ((methodDeclaration.Flags() & (intmod.FlagSynthetic | intmod.FlagPrivate)) ==
				(intmod.FlagSynthetic | intmod.FlagPrivate)) &&
				methodDeclaration.Method().Name() == name1 &&
				methodDeclaration.Method().Descriptor() == descriptor1 {
				// Create lambda expression
				cfmd := methodDeclaration.(intsrv.IClassFileMethodDeclaration)
				stack.Push(modexp.NewLambdaIdentifiersExpressionWithAll(
					lineNumber, indyMethodTypes.ReturnedType(), indyMethodTypes.ReturnedType(),
					p.prepareLambdaParameterNames(cfmd.FormalParameters(), parameterCount),
					p.prepareLambdaStatements(cfmd.FormalParameters(), indyParameters, cfmd.Statements())))
				return
			}
		}
	}

	if indyParameters == nil {
		// Create static method reference
		ot := p.typeMaker.MakeFromInternalTypeName(typeName)

		if name1 == "<init>" {
			stack.Push(modexp.NewConstructorReferenceExpressionWithAll(lineNumber,
				indyMethodTypes.ReturnedType(), ot, descriptor1))
		} else {
			stack.Push(modexp.NewMethodReferenceExpressionWithAll(lineNumber, indyMethodTypes.ReturnedType(),
				modexp.NewObjectTypeReferenceExpressionWithLineNumber(lineNumber, ot), typeName, name1, descriptor1))
		}
		return
	}

	// Create method reference
	stack.Push(modexp.NewMethodReferenceExpressionWithAll(lineNumber, indyMethodTypes.ReturnedType(),
		indyParameters, typeName, name1, descriptor1))
}

func (p *ByteCodeParser) prepareLambdaParameterNames(formalParameters intmod.IFormalParameter,
	parameterCount int) util.IList[string] {
	if (formalParameters == nil) || (parameterCount == 0) {
		return nil
	} else {
		p.lambdaParameterNamesVisitor.Init()
		formalParameters.AcceptDeclaration(p.lambdaParameterNamesVisitor)
		names := p.lambdaParameterNamesVisitor.Names()

		if names.Size() == parameterCount {
			return names
		} else {
			return names.SubList(names.Size()-parameterCount, names.Size())
		}
	}
}

func (p *ByteCodeParser) prepareLambdaStatements(formalParameters intmod.IFormalParameter,
	indyParameters intmod.IExpression, baseStatement intmod.IStatement) intmod.IStatement {
	if baseStatement != nil {
		if (formalParameters != nil) && (indyParameters != nil) {
			size := indyParameters.Size()

			if (size > 0) && (size <= formalParameters.Size()) {
				mapping := make(map[string]string)
				expression := indyParameters.First()

				if expression.IsLocalVariableReferenceExpression() {
					name := formalParameters.First().Name()
					newName := expression.Name()

					if name != newName {
						mapping[name] = newName
					}
				}

				if size > 1 {
					formalParameterList := formalParameters.ToList()
					list := indyParameters.ToList()

					for i := 1; i < size; i++ {
						expression = list.Get(i)

						if expression.IsLocalVariableReferenceExpression() {
							name := formalParameterList.Get(i).Name()
							newName := expression.Name()

							if name != newName {
								mapping[name] = newName
							}
						}
					}
				}

				if len(mapping) != 0 {
					p.renameLocalVariablesVisitor.Init(mapping)
					baseStatement.AcceptStatement(p.renameLocalVariablesVisitor)
				}
			}
		}

		if baseStatement.Size() == 1 {
			statement := baseStatement.First()

			if statement.IsReturnExpressionStatement() {
				return modsts.NewLambdaExpressionStatement(statement.Expression())
			} else if statement.IsExpressionStatement() {
				return modsts.NewLambdaExpressionStatement(statement.Expression())
			}
		}
	}

	return baseStatement
}

func stackContainsFieldReference(stack util.IStack[intmod.IExpression], fr intmod.IFieldReferenceExpression) bool {
	if !stack.IsEmpty() {
		expression := stack.Peek()

		if expression.IsFieldReferenceExpression() {
			return expression.Name() == fr.Name() && expression.Expression().Type() == fr.Expression().Type()
		}
	}

	return false
}

func createAssignment(boe intmod.IBinaryOperatorExpression, operator string) intmod.IExpression {
	boe.SetOperator(operator)
	boe.SetPriority(16)
	return boe
}

func isPositiveOne(expression intmod.IExpression) bool {
	if expression.IsIntegerConstantExpression() && expression.IntegerValue() == 1 {
		return true
	}
	if expression.IsLongConstantExpression() && expression.LongValue() == 1 {
		return true
	}
	if expression.IsFloatConstantExpression() && expression.FloatValue() == 1.0 {
		return true
	}
	return expression.IsDoubleConstantExpression() && expression.DoubleValue() == 1.0
}

func isNegativeOne(expression intmod.IExpression) bool {
	if expression.IsIntegerConstantExpression() && expression.IntegerValue() == -1 {
		return true
	}
	if expression.IsLongConstantExpression() && expression.LongValue() == -1 {
		return true
	}
	if expression.IsFloatConstantExpression() && expression.FloatValue() == -1.0 {
		return true
	}
	return expression.IsDoubleConstantExpression() && expression.DoubleValue() == -1.0
}

func (p *ByteCodeParser) parseASTORE(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber, offset int, localVariable intsrv.ILocalVariable, valueRef intmod.IExpression) {
	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(localVariable.Type(), valueRef)
	localVariable.TypeOnRight(p.typeBounds, valueRef.Type())

	vre := srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable)
	oldValueRef := valueRef

	if valueRef.IsNewArray() {
		valueRef = utils.MakeNewArrayMaker(statements, valueRef)
	}

	if oldValueRef != valueRef {
		stack.Replace(oldValueRef, valueRef)
	}

	p.createAssignment(statements, stack, lineNumber, vre, valueRef)
}

func (p *ByteCodeParser) createAssignment(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber int, leftExpression intmod.IExpression, rightExpression intmod.IExpression) {
	if !stack.IsEmpty() && (stack.Peek() == rightExpression) {
		stack.Push(modexp.NewBinaryOperatorExpression(lineNumber,
			leftExpression.Type(), leftExpression, "=", stack.Pop(), 16))
		return
	}

	if !statements.IsEmpty() {
		lastStatement := statements.Last()

		if lastStatement.IsExpressionStatement() {
			lastES := lastStatement.(intmod.IExpressionStatement)
			lastExpression := lastStatement.Expression()

			if lastExpression.IsBinaryOperatorExpression() {
				if getLastRightExpression(lastExpression) == rightExpression {
					// Multi assignment
					lastES.SetExpression(modexp.NewBinaryOperatorExpression(lineNumber,
						leftExpression.Type(), leftExpression, "=", lastExpression, 16))
					return
				}

				if (lineNumber > 0) && (lastExpression.LineNumber() == lineNumber) &&
					reflect.TypeOf(lastExpression.LeftExpression()) == reflect.TypeOf(rightExpression) {
					if rightExpression.IsLocalVariableReferenceExpression() {
						lvr1 := lastExpression.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
						lvr2 := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression)

						if lvr1.LocalVariable() == lvr2.LocalVariable() {
							// Multi assignment
							lastES.SetExpression(modexp.NewBinaryOperatorExpression(lineNumber,
								leftExpression.Type(), leftExpression, "=", lastExpression, 16))
							return
						}
					} else if rightExpression.IsFieldReferenceExpression() {
						fr1 := lastExpression.LeftExpression().(intmod.IFieldReferenceExpression)
						fr2 := rightExpression.(intmod.IFieldReferenceExpression)

						if fr1.Name() == fr2.Name() && fr1.Expression().Type() == fr2.Expression().Type() {
							// Multi assignment
							lastES.SetExpression(modexp.NewBinaryOperatorExpression(lineNumber,
								leftExpression.Type(), leftExpression, "=", lastExpression, 16))
							return
						}
					}
				}
			} else if lastExpression.IsPreOperatorExpression() {
				if reflect.TypeOf(lastExpression.Expression()) == reflect.TypeOf(rightExpression) {
					if rightExpression.IsLocalVariableReferenceExpression() {
						lvr1 := lastExpression.Expression().(intsrv.IClassFileLocalVariableReferenceExpression)
						lvr2 := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression)

						if lvr1.LocalVariable() == lvr2.LocalVariable() {
							rightExpression = p.newPreArithmeticOperatorExpression(lastExpression.LineNumber(),
								lastExpression.Operator(), lastExpression.Expression())
							statements.RemoveLast()
						}
					} else if rightExpression.IsFieldReferenceExpression() {
						fr1 := lastExpression.Expression().(intmod.IFieldReferenceExpression)
						fr2 := rightExpression.(intmod.IFieldReferenceExpression)

						if fr1.Name() == fr2.Name() && fr1.Expression().Type() == fr2.Expression().Type() {
							rightExpression = p.newPreArithmeticOperatorExpression(lastExpression.LineNumber(),
								lastExpression.Operator(), lastExpression.Expression())
							statements.RemoveLast()
						}
					}
				}
			} else if lastExpression.IsPostOperatorExpression() {
				poe := lastExpression.(intmod.IPostOperatorExpression)

				if poe.Expression() == rightExpression {
					rightExpression = poe
					statements.RemoveLast()
				}
			}
		}
	}

	statements.Add(modsts.NewExpressionStatement(
		modexp.NewBinaryOperatorExpression(lineNumber, leftExpression.Type(), leftExpression, "=", rightExpression, 16)))
}

func (p *ByteCodeParser) parseIINC(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	lineNumber, offset int, localVariable intsrv.ILocalVariable, count int) {
	var expression intmod.IExpression

	if !stack.IsEmpty() {
		expression = stack.Peek()

		if (expression.LineNumber() == lineNumber) && expression.IsLocalVariableReferenceExpression() {
			exp := expression.(intsrv.IClassFileLocalVariableReferenceExpression)

			if exp.LocalVariable() == localVariable {
				// ILOAD found -> Create a post-incrementation
				stack.Pop()

				if count == 1 {
					stack.Push(p.newPostArithmeticOperatorExpression(lineNumber, expression, "++"))
				} else if count == -1 {
					stack.Push(p.newPostArithmeticOperatorExpression(lineNumber, expression, "--"))
				} else {
					// assert false
				}
				return
			}
		}
	}

	expression = srvexp.NewClassFileLocalVariableReferenceExpression(lineNumber, offset, localVariable)

	if count == 1 {
		expression = p.newPreArithmeticOperatorExpression(lineNumber, "++", expression)
	} else if count == -1 {
		expression = p.newPreArithmeticOperatorExpression(lineNumber, "--", expression)
	} else if count >= 0 {
		expression = modexp.NewBinaryOperatorExpression(lineNumber, expression.Type(),
			expression, "+=", modexp.NewIntegerConstantExpressionWithAll(lineNumber, expression.Type(), count), 16)
	} else if count < 0 {
		expression = modexp.NewBinaryOperatorExpression(lineNumber, expression.Type(),
			expression, "-=", modexp.NewIntegerConstantExpressionWithAll(lineNumber, expression.Type(), -count), 16)
	} else {
		// assert false;
		expression = nil
	}

	statements.Add(modsts.NewExpressionStatement(expression))
}

func (p *ByteCodeParser) parseIF(stack util.IStack[intmod.IExpression], lineNumber int,
	basicBlock intsrv.IBasicBlock, operator1, operator2 string, priority int) {
	expression := stack.Pop()

	if reflect.TypeOf(expression) == reflect.TypeOf(&srvexp.ClassFileCmpExpression{}) {
		cmp := expression.(intsrv.IClassFileCmpExpression)
		p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(cmp.LeftExpression().Type(), cmp.LeftExpression())
		p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(cmp.RightExpression().Type(), cmp.RightExpression())
		operator := operator2
		if basicBlock.IsInverseCondition() {
			operator = operator1
		}
		stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, cmp.LeftExpression(), operator, cmp.RightExpression(), priority))
	} else if expression.Type().IsPrimitiveType() {
		pt := expression.Type().(intmod.IPrimitiveType)
		p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(pt, expression)
		operator := operator2
		if basicBlock.IsInverseCondition() {
			operator = operator1
		}

		switch pt.JavaPrimitiveFlags() {
		case intmod.FlagBoolean:
			if basicBlock.IsInverseCondition() != ("==" == operator1) {
				stack.Push(expression)
			} else {
				stack.Push(modexp.NewPreOperatorExpressionWithAll(lineNumber, "!", expression))
			}
			break
		case intmod.FlagFloat:
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression,
				operator, modexp.NewFloatConstantExpressionWithAll(lineNumber, 0), 9))
			break
		case intmod.FlagDouble:
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression,
				operator, modexp.NewDoubleConstantExpressionWithAll(lineNumber, 0), 9))
			break
		case intmod.FlagLong:
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression,
				operator, modexp.NewLongConstantExpressionWithAll(lineNumber, 0), 9))
			break
		default:
			stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression,
				operator, modexp.NewIntegerConstantExpressionWithAll(lineNumber, pt, 0), 9))
			break
		}
	} else {
		operator := operator2
		if basicBlock.IsInverseCondition() {
			operator = operator1
		}
		stack.Push(modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, expression,
			operator, modexp.NewNullExpressionWithAll(lineNumber, expression.Type()), 9))
	}
}

func (p *ByteCodeParser) parseXRETURN(statements intmod.IStatements, stack util.IStack[intmod.IExpression], lineNumber int) {
	valueRef := stack.Pop()

	if valueRef.IsNewArray() {
		valueRef = utils.MakeNewArrayMaker(statements, valueRef)
	}

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(p.returnedType, valueRef)

	if lineNumber > valueRef.LineNumber() {
		lineNumber = valueRef.LineNumber()
	}

	if !statements.IsEmpty() && valueRef.IsLocalVariableReferenceExpression() {
		lastStatement := statements.Last()

		if lastStatement.IsExpressionStatement() {
			expression := statements.Last().Expression()

			if lineNumber <= expression.LineNumber() &&
				expression.IsBinaryOperatorExpression() &&
				expression.Operator() == "=" &&
				expression.LeftExpression().IsLocalVariableReferenceExpression() {
				vre1 := expression.LeftExpression().(intsrv.IClassFileLocalVariableReferenceExpression)
				vre2 := valueRef.(intsrv.IClassFileLocalVariableReferenceExpression)

				if vre1.LocalVariable() == vre2.LocalVariable() {
					// Remove synthetic local variable
					p.localVariableMaker.RemoveLocalVariable(vre1.LocalVariable().(intsrv.ILocalVariable))
					// Remove assignment statement
					statements.RemoveLast()
					statements.Add(modsts.NewReturnExpressionStatementWithAll(lineNumber, expression.RightExpression()))
					return
				}
			}
		}
	}

	statements.Add(modsts.NewReturnExpressionStatementWithAll(lineNumber, valueRef))
}

func (p *ByteCodeParser) parseGetStatic(stack util.IStack[intmod.IExpression],
	constants intcls.IConstantPool, lineNumber, index int) {
	constantMemberRef := constants.Constant(index).(intcls.IConstantMemberRef)
	typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
	constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())

	if (name == "TYPE") && strings.HasPrefix(typeName, "java/lang/") {
		switch typeName {
		case "java/lang/Boolean":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeBoolean))
			return
		case "java/lang/Character":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeChar))
			return
		case "java/lang/Float":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeFloat))
			return
		case "java/lang/Double":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeDouble))
			return
		case "java/lang/Byte":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeByte))
			return
		case "java/lang/Short":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeShort))
			return
		case "java/lang/Integer":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeInt))
			return
		case "java/lang/Long":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeLong))
			return
		case "java/lang/Void":
			stack.Push(modexp.NewTypeReferenceDotClassExpressionWithAll(lineNumber, _type.PtTypeVoid))
			return
		}
	}

	ot := p.typeMaker.MakeFromInternalTypeName(typeName)
	descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
	typ := p.makeFieldType(ot.InternalName(), name, descriptor)
	explicit := p.internalTypeName != typeName || p.localVariableMaker.ContainsName(name)
	objectRef := modexp.NewObjectTypeReferenceExpressionWithAll(lineNumber, ot, explicit)
	stack.Push(p.typeParametersToTypeArgumentsBinder.NewFieldReferenceExpression(lineNumber, typ, objectRef, ot, name, descriptor))
}

func (p *ByteCodeParser) parsePutStatic(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	constants intcls.IConstantPool, lineNumber, index int) {
	constantMemberRef := constants.Constant(index).(intcls.IConstantMemberRef)
	typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
	ot := p.typeMaker.MakeFromInternalTypeName(typeName)
	constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
	descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
	typ := p.makeFieldType(ot.InternalName(), name, descriptor)
	valueRef := stack.Pop()
	objectRef := modexp.NewObjectTypeReferenceExpressionWithAll(lineNumber, ot, p.internalTypeName != typeName || p.localVariableMaker.ContainsName(name))
	fieldRef := p.typeParametersToTypeArgumentsBinder.NewFieldReferenceExpression(
		lineNumber, typ, objectRef, ot, name, descriptor).(intmod.IFieldReferenceExpression)
	p.parsePUT(statements, stack, lineNumber, fieldRef, valueRef)
}

func (p *ByteCodeParser) parseGetField(stack util.IStack[intmod.IExpression],
	constants intcls.IConstantPool, lineNumber, index int) {
	constantMemberRef := constants.Constant(index).(intcls.IConstantMemberRef)
	typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
	ot := p.typeMaker.MakeFromInternalTypeName(typeName)
	constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
	descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
	typ := p.makeFieldType(ot.InternalName(), name, descriptor)
	objectRef := stack.Pop()
	stack.Push(p.typeParametersToTypeArgumentsBinder.NewFieldReferenceExpression(
		lineNumber, typ, p.getFieldInstanceReference(objectRef, ot, name), ot, name, descriptor))
}

func (p *ByteCodeParser) parsePutField(statements intmod.IStatements, stack util.IStack[intmod.IExpression],
	constants intcls.IConstantPool, lineNumber, index int) {
	constantMemberRef := constants.Constant(index).(intcls.IConstantMemberRef)
	typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
	ot := p.typeMaker.MakeFromInternalTypeName(typeName)
	constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
	descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
	typ := p.makeFieldType(ot.InternalName(), name, descriptor)
	valueRef := stack.Pop()
	objectRef := stack.Pop()
	fieldRef := p.typeParametersToTypeArgumentsBinder.NewFieldReferenceExpression(
		lineNumber, typ, p.getFieldInstanceReference(objectRef, ot, name), ot, name, descriptor)
	p.parsePUT(statements, stack, lineNumber, fieldRef, valueRef)
}

func getLastRightExpression(boe intmod.IExpression) intmod.IExpression {
	for {
		if boe.Operator() != "=" {
			return boe
		}

		if !boe.RightExpression().IsBinaryOperatorExpression() {
			return boe.RightExpression()
		}

		boe = boe.RightExpression()
	}
}

func (p *ByteCodeParser) newNewExpression(lineNumber int, internalTypeName string) intmod.IExpression {
	objectType := p.typeMaker.MakeFromInternalTypeName(internalTypeName)

	if objectType.QualifiedName() == "" && objectType.Name() == "" {
		typeDeclaration := p.bodyDeclaration.InnerTypeDeclaration(internalTypeName).(intsrv.IClassFileTypeDeclaration)

		if typeDeclaration == nil {
			return srvexp.NewClassFileNewExpression(lineNumber, _type.OtTypeObject)
		} else if typeDeclaration.IsClassDeclaration() {
			decl := typeDeclaration.(intsrv.IClassFileClassDeclaration)
			var bodyDeclaration intmod.IBodyDeclaration

			if p.internalTypeName == internalTypeName {
				bodyDeclaration = nil
			} else {
				bodyDeclaration = decl.BodyDeclaration()
			}

			if decl.Interfaces() != nil {
				return srvexp.NewClassFileNewExpression3(lineNumber, decl.Interfaces().(intmod.IObjectType), bodyDeclaration, true)
			} else if decl.SuperType() != nil {
				return srvexp.NewClassFileNewExpression3(lineNumber, decl.SuperType(), bodyDeclaration, true)
			} else {
				return srvexp.NewClassFileNewExpression3(lineNumber, _type.OtTypeObject, bodyDeclaration, true)
			}
		}
	}

	return srvexp.NewClassFileNewExpression(lineNumber, objectType)
}

/*
 * Operators = { "+", "-", "*", "/", "%", "<<", ">>", ">>>" }
 * See "Additive Operators (+ and -) for Numeric Types": https://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.18.2
 * See "Shift Operators":                                https://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.19
 */
func (p *ByteCodeParser) newIntegerBinaryOperatorExpression(lineNumber int,
	leftExpression intmod.IExpression, operator string, rightExpression intmod.IExpression, priority int) intmod.IExpression {
	if leftExpression.IsLocalVariableReferenceExpression() {
		leftVariable := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
		leftVariable.TypeOnLeft(p.typeBounds, _type.PtMaybeByteType)
	}

	if rightExpression.IsLocalVariableReferenceExpression() {
		rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
		rightVariable.TypeOnLeft(p.typeBounds, _type.PtMaybeByteType)
	}

	return modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeInt, leftExpression, operator, rightExpression, priority)
}

/*
 * Operators = { "&", "|", "^" }
 * See "Binary Numeric Promotion": https://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.22.1
 */
func (p *ByteCodeParser) newIntegerOrBooleanBinaryOperatorExpression(lineNumber int,
	leftExpression intmod.IExpression, operator string, rightExpression intmod.IExpression,
	priority int) intmod.IExpression {
	typ := _type.PtTypeInt

	if leftExpression.IsLocalVariableReferenceExpression() {
		leftVariable := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

		if rightExpression.IsLocalVariableReferenceExpression() {
			rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

			if leftVariable.IsAssignableFrom(p.typeBounds, _type.PtTypeBoolean) || rightVariable.IsAssignableFrom(p.typeBounds, _type.PtTypeBoolean) {
				leftVariable.VariableOnRight(p.typeBounds, rightVariable)
				rightVariable.VariableOnLeft(p.typeBounds, leftVariable)

				if (leftVariable.Type() == _type.PtTypeBoolean) || (rightVariable.Type() == _type.PtTypeBoolean) {
					typ = _type.PtTypeBoolean
				}
			}
		} else {
			if rightExpression.Type() == _type.PtTypeBoolean {
				typ = _type.PtTypeBoolean
				leftVariable.TypeOnRight(p.typeBounds, typ)
			}
		}
	} else if rightExpression.IsLocalVariableReferenceExpression() {
		if leftExpression.Type() == _type.PtTypeBoolean {
			rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
			typ = _type.PtTypeBoolean
			rightVariable.TypeOnRight(p.typeBounds, typ)
		}
	}

	if typ == _type.PtTypeInt {
		if leftExpression.Type().IsPrimitiveType() && rightExpression.Type().IsPrimitiveType() {
			leftFlags := leftExpression.Type().(intmod.IPrimitiveType).Flags()
			rightFlags := rightExpression.Type().(intmod.IPrimitiveType).Flags()
			leftBoolean := (leftFlags & intmod.FlagBoolean) != 0
			rightBoolean := (rightFlags & intmod.FlagBoolean) != 0
			commonflags := leftFlags | rightFlags

			if !leftBoolean || !rightBoolean {
				commonflags &= ^intmod.FlagBoolean
			}

			typ = utils.GetPrimitiveTypeFromFlags(commonflags)

			if typ == nil {
				typ = _type.PtTypeInt
			}
		}
	}

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(leftExpression.Type(), rightExpression)
	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(rightExpression.Type(), leftExpression)

	return modexp.NewBinaryOperatorExpression(lineNumber, typ, leftExpression, operator, rightExpression, priority)
}

/*
 * Operators = { "==", "!=" }
 * See "Numerical Equality Operators == and !=": https://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.21.1
 */
func (p *ByteCodeParser) newIntegerOrBooleanComparisonOperatorExpression(lineNumber int,
	leftExpression intmod.IExpression, operator string, rightExpression intmod.IExpression,
	priority int) intmod.IExpression {
	if leftExpression.IsLocalVariableReferenceExpression() {
		leftVariable := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

		if rightExpression.IsLocalVariableReferenceExpression() {
			rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

			if leftVariable.IsAssignableFrom(p.typeBounds, _type.PtTypeBoolean) || rightVariable.IsAssignableFrom(p.typeBounds, _type.PtTypeBoolean) {
				leftVariable.VariableOnRight(p.typeBounds, rightVariable)
				rightVariable.VariableOnLeft(p.typeBounds, leftVariable)
			}
		} else {
			if rightExpression.Type() == _type.PtTypeBoolean {
				leftVariable.TypeOnRight(p.typeBounds, _type.PtTypeBoolean)
			}
		}
	} else if rightExpression.IsLocalVariableReferenceExpression() {
		if leftExpression.Type() == _type.PtTypeBoolean {
			rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
			rightVariable.TypeOnRight(p.typeBounds, _type.PtTypeBoolean)
		}
	}

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(leftExpression.Type(), rightExpression)
	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(rightExpression.Type(), leftExpression)

	return modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, leftExpression, operator, rightExpression, priority)
}

/*
 * Operators = { "==", "!=" }
 * See "Numerical Equality Operators == and !=": https://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.21.1
 */
func (p *ByteCodeParser) newIntegerComparisonOperatorExpression(lineNumber int,
	leftExpression intmod.IExpression, operator string, rightExpression intmod.IExpression, priority int) intmod.IExpression {
	if leftExpression.IsLocalVariableReferenceExpression() {
		leftVariable := leftExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
		leftVariable.TypeOnLeft(p.typeBounds, _type.PtMaybeByteType)
	}

	if rightExpression.IsLocalVariableReferenceExpression() {
		rightVariable := rightExpression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
		rightVariable.TypeOnLeft(p.typeBounds, _type.PtMaybeByteType)
	}

	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(leftExpression.Type(), rightExpression)
	p.typeParametersToTypeArgumentsBinder.BindParameterTypesWithArgumentTypes(rightExpression.Type(), leftExpression)

	return modexp.NewBinaryOperatorExpression(lineNumber, _type.PtTypeBoolean, leftExpression, operator, rightExpression, priority)
}

func (p *ByteCodeParser) newPreArithmeticOperatorExpression(lineNumber int, operator string,
	expression intmod.IExpression) intmod.IExpression {
	p.reduceIntegerLocalVariableType(expression)
	return modexp.NewPreOperatorExpressionWithAll(lineNumber, operator, expression)
}

func (p *ByteCodeParser) newPostArithmeticOperatorExpression(lineNumber int,
	expression intmod.IExpression, operator string) intmod.IExpression {
	p.reduceIntegerLocalVariableType(expression)
	return modexp.NewPostOperatorExpressionWithAll(lineNumber, operator, expression)
}

func (p *ByteCodeParser) reduceIntegerLocalVariableType(expression intmod.IExpression) {
	if expression.IsLocalVariableReferenceExpression() {
		lvre := expression.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)

		if lvre.LocalVariable().(intsrv.ILocalVariable).IsPrimitiveLocalVariable() {
			plv := lvre.LocalVariable().(intsrv.ILocalVariable)
			if plv.IsAssignableFrom(p.typeBounds, _type.PtMaybeBooleanType) {
				plv.TypeOnRight(p.typeBounds, _type.PtMaybeByteType)
			}
		}
	}
}

/**
 * @return expression, 'this' or 'super'
 */
func (p *ByteCodeParser) getFieldInstanceReference(expression intmod.IExpression, ot intmod.IObjectType, name string) intmod.IExpression {
	if p.bodyDeclaration.FieldDeclarations() != nil && expression.IsThisExpression() {
		internalName := expression.Type().InternalName()

		if ot.InternalName() != internalName {
			p.memberVisitor.Init(name, "")
			for _, field := range p.bodyDeclaration.FieldDeclarations() {
				field.FieldDeclarators().AcceptDeclaration(p.memberVisitor)
				if p.memberVisitor.Found() {
					return modexp.NewSuperExpressionWithAll(expression.LineNumber(), expression.Type())
				}
			}
		}
	}

	return expression
}

/**
 * @return expression, 'this' or 'super'
 */
func (p *ByteCodeParser) getMethodInstanceReference(expression intmod.IExpression, ot intmod.IObjectType,
	name, descriptor string) intmod.IExpression {
	if p.bodyDeclaration.MethodDeclarations() != nil && expression.IsThisExpression() {
		internalName := expression.Type().InternalName()

		if ot.InternalName() != internalName {
			p.memberVisitor.Init(name, descriptor)

			for _, member := range p.bodyDeclaration.MethodDeclarations() {
				member.AcceptDeclaration(p.memberVisitor)
				if p.memberVisitor.Found() {
					return modexp.NewSuperExpressionWithAll(expression.LineNumber(), expression.Type())
				}
			}
		}
	}

	return expression
}

func checkStack(stack util.IStack[intmod.IExpression], code []byte, offset int) {
	if (stack.Size() > 1) && (offset < len(code)) {
		opcode := int(code[offset+1] & 255)

		if (opcode == 87) || (opcode == 176) { // POP || ARETURN
			// Duplicate last expression
			condition := stack.Pop()
			stack.Push(stack.Peek())
			stack.Push(condition)
		}
	}
}

func IsAssertCondition(internalTypeName string, basicBlock intsrv.IBasicBlock) bool {
	cfg := basicBlock.ControlFlowGraph()
	offset := basicBlock.FromOffset()
	toOffset := basicBlock.ToOffset()
	value := 0

	if offset+3 > toOffset {
		return false
	}

	method := cfg.Method()
	code := method.Attribute("Code").(intcls.IAttributeCode).Code()
	opcode := int(code[offset] & 255)

	if opcode != 178 { // GETSTATIC
		return false
	}

	constants := method.Constants()
	value, offset = utils.PrefixReadInt16(code, offset)
	constantMemberRef := constants.Constant(value).(intcls.IConstantMemberRef)
	constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(intcls.IConstantNameAndType)
	name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())

	if "$assertionsDisabled" != name {
		return false
	}

	descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())

	if "Z" != descriptor {
		return false
	}

	typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())

	return internalTypeName == typeName
}

func GetExceptionLocalVariableIndex(basicBlock intsrv.IBasicBlock) int {
	cfg := basicBlock.ControlFlowGraph()
	offset := basicBlock.FromOffset()
	toOffset := basicBlock.ToOffset()
	value := 0

	if offset+1 > toOffset {
		// assert false;
		return -1
	}

	method := cfg.Method()
	code := method.Attribute("Code").(intcls.IAttributeCode).Code()
	opcode := int(code[offset] & 255)

	switch opcode {
	case 58: // ASTORE
		value, offset = utils.PrefixReadInt8(code, offset)
		return value
	case 75, 76, 77, 78: // ASTORE_0 ... ASTORE_3
		return opcode - 75
	case 87, 88: // POP, POP2
		return -1
	default:
		//assert false;
		return -1
	}
}

func (p *ByteCodeParser) makeFieldType(internalTypeName, fieldName, descriptor string) intmod.IType {
	typ := p.typeMaker.MakeFieldType(internalTypeName, fieldName, descriptor)

	if !p.genericTypesSupported {
		p.eraseTypeArgumentVisitor.Init()
		typ.AcceptTypeVisitor(p.eraseTypeArgumentVisitor)
		typ = p.eraseTypeArgumentVisitor.Type()
	}

	return typ
}

func (p *ByteCodeParser) makeMethodTypes(internalTypeName, methodName, descriptor string) intsrv.IMethodTypes {
	methodTypes := p.typeMaker.MakeMethodTypes2(internalTypeName, methodName, descriptor)

	if !p.genericTypesSupported {
		mt := NewMethodTypes()

		if methodTypes.ParameterTypes() != nil {
			p.eraseTypeArgumentVisitor.Init()
			methodTypes.ParameterTypes().AcceptTypeVisitor(p.eraseTypeArgumentVisitor)
			mt.SetParameterTypes(p.eraseTypeArgumentVisitor.Type())
		}

		p.eraseTypeArgumentVisitor.Init()
		methodTypes.ReturnedType().AcceptTypeVisitor(p.eraseTypeArgumentVisitor)
		mt.SetReturnedType(p.eraseTypeArgumentVisitor.Type())

		if methodTypes.ExceptionTypes() != nil {
			p.eraseTypeArgumentVisitor.Init()
			methodTypes.ExceptionTypes().AcceptTypeVisitor(p.eraseTypeArgumentVisitor)
			mt.SetExceptionTypes(p.eraseTypeArgumentVisitor.Type())
		}

		methodTypes = mt
	}

	return methodTypes
}

func forceExplicitCastExpression(expr intmod.IExpression) intmod.IExpression {
	// In case of downcasting, set all cast expression children as explicit to prevent missing castings
	exp := expr

	for exp.IsCastExpression() {
		ce := exp.(intmod.ICastExpression)
		ce.SetExplicit(true)
		exp = ce.Expression()
	}

	return expr
}

func NewByteCodeParserMemberVisitor() *ByteCodeParserMemberVisitor {
	return &ByteCodeParserMemberVisitor{}
}

type ByteCodeParserMemberVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	name       string
	descriptor string
	found      bool
}

func (v *ByteCodeParserMemberVisitor) Init(name, descriptor string) {
	v.name = name
	v.descriptor = descriptor
	v.found = false
}

func (v *ByteCodeParserMemberVisitor) Found() bool {
	return v.found
}

func (v *ByteCodeParserMemberVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.found = v.found || decl.Name() == v.name
}

func (v *ByteCodeParserMemberVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	v.found = v.found || decl.Name() == v.name && decl.Descriptor() == v.descriptor
}

func NewLambdaParameterNamesVisitor() *LambdaParameterNamesVisitor {
	return &LambdaParameterNamesVisitor{}
}

type LambdaParameterNamesVisitor struct {
	declaration.AbstractNopDeclarationVisitor

	names util.IList[string]
}

func (v *LambdaParameterNamesVisitor) Init() {
	v.names = util.NewDefaultList[string]()
}

func (v *LambdaParameterNamesVisitor) Names() util.IList[string] {
	return v.names
}

func (v *LambdaParameterNamesVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	v.names.Add(decl.Name())
}

func (v *LambdaParameterNamesVisitor) VisitFormalParameters(decls intmod.IFormalParameters) {
	iterator := decls.Iterator()
	for iterator.HasNext() {
		iterator.Next().AcceptDeclaration(v)
	}
}

func NewJsrReturnAddressExpression() *JsrReturnAddressExpression {
	return &JsrReturnAddressExpression{
		NullExpression: *modexp.NewNullExpression(_type.PtTypeVoid).(*modexp.NullExpression),
	}
}

type JsrReturnAddressExpression struct {
	modexp.NullExpression
}

func (e *JsrReturnAddressExpression) String() string {
	return "JsrReturnAddressExpression{}"
}
