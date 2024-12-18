package service

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

type IAnnotationConverter interface {
	ConvertWithAnnotations2(visibles, invisibles intcls.IAnnotations) intmod.IAnnotationReference
	ConvertWithAnnotations(annotations intcls.IAnnotations) intmod.IAnnotationReference
	ConvertWithAnnotation(annotation intcls.IAnnotation) intmod.IAnnotationReference
	ConvertWithElementValue(ev intcls.IElementValue) intmod.IElementValue
	VisitPrimitiveType(elementValue intcls.IElementValuePrimitiveType)
	VisitClassInfo(elementValue intcls.IElementValueClassInfo)
	VisitAnnotationValue(elementValue intcls.IElementValueAnnotationValue)
	VisitEnumConstValue(elementValue intcls.IElementValueEnumConstValue)
	VisitArrayValue(elementValue intcls.IElementValueArrayValue)
}

type IStatementMaker interface {
	Make(cfg IControlFlowGraph) intmod.IStatements
}

type IByteCodeParser interface {
	Parse(basicBlock IBasicBlock, statements intmod.IStatements, stack util.IStack[intmod.IExpression])
}

type ITypeParametersToTypeArgumentsBinder interface {
	NewConstructorInvocationExpression(lineNumber int, objectType intmod.IObjectType, descriptor string,
		methodTypes IMethodTypes, parameters intmod.IExpression) IClassFileConstructorInvocationExpression
	NewSuperConstructorInvocationExpression(lineNumber int, objectType intmod.IObjectType, descriptor string,
		methodTypes IMethodTypes, parameters intmod.IExpression) IClassFileSuperConstructorInvocationExpression
	NewMethodInvocationExpression(lineNumber int, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
		methodTypes IMethodTypes, parameters intmod.IExpression) IClassFileMethodInvocationExpression
	NewFieldReferenceExpression(lineNumber int, typ intmod.IType, expression intmod.IExpression,
		objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression
	BindParameterTypesWithArgumentTypes(typ intmod.IType, expression intmod.IExpression)
	UpdateNewExpression(ne IClassFileNewExpression, descriptor string, methodTypes IMethodTypes, parameters intmod.IExpression)
}

type ILocalVariableMaker interface {
	LocalVariable(index, offset int) ILocalVariable
	IsCompatible(lv ILocalVariable, valueType intmod.IType) bool
	LocalVariableInAssignment(typeBounds map[string]intmod.IType,
		index, offset int, valueType intmod.IType) ILocalVariable
	LocalVariableInNullAssignment(index, offset int, valueType intmod.IType) ILocalVariable
	LocalVariableInAssignmentWithLocalVariable(typeBounds map[string]intmod.IType,
		index, offset int, valueLocalVariable ILocalVariable) ILocalVariable
	ExceptionLocalVariable(index, offset int, t intmod.IObjectType) ILocalVariable
	RemoveLocalVariable(lv ILocalVariable)
	ContainsName(name string) bool
	Make(containsLineNumber bool, typeMaker ITypeMaker)
	FormalParameters() intmod.IFormalParameter
	PushFrame(statements intmod.IStatements)
	PopFrame()
	ChangeFrame(localVariable ILocalVariable)
}

type ITypeMaker interface {
	ParseClassFileSignature(classFile intcls.IClassFile) ITypeTypes
	ParseMethodSignature(classFile intcls.IClassFile, method intcls.IMethod) IMethodTypes
	ParseFieldSignature(classFile intcls.IClassFile, field intcls.IField) intmod.IType
	MakeFromSignature(signature string) intmod.IType
	MakeFromDescriptorOrInternalTypeName(descriptorOrInternalTypeName string) intmod.IObjectType
	MakeFromDescriptor(descriptor string) intmod.IObjectType
	MakeFromInternalTypeName(internalTypeName string) intmod.IObjectType
	SearchSuperParameterizedType(superObjectType, objectType intmod.IObjectType) intmod.IObjectType
	IsAssignable(typeBounds map[string]intmod.IType, left, right intmod.IObjectType) bool
	IsRawTypeAssignable(left, right intmod.IObjectType) bool
	MakeTypeTypes(internalTypeName string) ITypeTypes
	SetFieldType(internalTypeName, fieldName string, typ intmod.IType)
	MakeFieldType(internalTypeName, fieldName, descriptor string) intmod.IType
	SetMethodReturnedType(internalTypeName, methodName, descriptor string, typ intmod.IType)
	MakeMethodTypes(descriptor string) IMethodTypes
	MakeMethodTypes2(internalTypeName, methodName, descriptor string) IMethodTypes
	MatchCount(internalTypeName, name string, parameterCount int, constructor bool) int
	MatchCount2(typeBounds map[string]intmod.IType, internalTypeName, name string, parameters intmod.IExpression, constructor bool) int
}

type IClassPathLoader interface {
	Load(internalName string) ([]byte, error)
	CanLoad(internalName string) bool
}

type ISignatureReader interface {
	Signature() string
	Array() []byte
	Length() int
	Index() int
	Inc()
	Dec()
	Read() byte
	NextEqualsTo(c byte) bool
	Search(c byte) bool
	SearchEndMarker() byte
	Available() bool
	Substring(beginIndex int) string
	String() string
}

type ITypeTypes interface {
	ThisType() intmod.IObjectType
	SetThisType(thisType intmod.IObjectType)
	TypeParameters() intmod.ITypeParameter
	SetTypeParameters(typeParameters intmod.ITypeParameter)
	SuperType() intmod.IObjectType
	SetSuperType(superType intmod.IObjectType)
	Interfaces() intmod.IType
	SetInterfaces(interfaces intmod.IType)
}

type IMethodTypes interface {
	TypeParameters() intmod.ITypeParameter
	SetTypeParameters(typeParameters intmod.ITypeParameter)
	ParameterTypes() intmod.IType
	SetParameterTypes(parameterTypes intmod.IType)
	ReturnedType() intmod.IType
	SetReturnedType(returnedType intmod.IType)
	ExceptionTypes() intmod.IType
	SetExceptionTypes(exceptionTypes intmod.IType)
}

type IWatchDog interface {
	Clear()
	Check(parent, child IBasicBlock)
}

type ILink interface {
	ParentIndex() int
	SetParentIndex(parentIndex int)
	ChildIndex() int
	SetChildIndex(childIndex int)
	HashCode() int
	Equals(o ILink) bool
}
