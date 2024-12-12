package service

import (
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

type IUpdateExpressionVisitor interface {
	intmod.IJavaSyntaxVisitor

	UpdateExpression(_ intmod.IExpression) intmod.IExpression
	UpdateBaseExpression(baseExpression intmod.IExpression) intmod.IExpression
}

type IAddCastExpressionVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IAggregateFieldsVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IAutoboxingVisitor interface {
	IUpdateExpressionVisitor
}

type IBaseTypeToTypeArgumentVisitor interface {
	intmod.ITypeVisitor

	Init()
	TypeArgument() intmod.ITypeArgument
}

type IBindTypeArgumentsToTypeArgumentsVisitor interface {
	intmod.ITypeArgumentVisitor

	Init()
	SetBindings(bindings map[string]intmod.ITypeArgument)
	TypeArgument() intmod.ITypeArgument
}

type IBindTypeParametersToNonWildcardTypeArgumentsVisitor interface {
	intmod.ITypeArgumentVisitor
	intmod.ITypeParameterVisitor

	Init(bindings map[string]intmod.ITypeArgument)
	TypeArgument() intmod.ITypeArgument
}

type IBindTypesToTypesVisitor interface {
	intmod.ITypeVisitor

	Init()
	Type() intmod.IType
	SetBindings(bindings map[string]intmod.ITypeArgument)
}

type IChangeFrameOfLocalVariablesVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type ICreateInstructionsVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type ICreateLocalVariableVisitor interface {
	intmod.ITypeArgumentVisitor
	ILocalVariableVisitor

	Init(index, offset int)
	LocalVariable() ILocalVariable
}

type ICreateParameterVisitor interface {
	intmod.ITypeArgumentVisitor

	Init(index int, name string)
	LocalVariable() ILocalVariable
}

type ICreateTypeFromTypeArgumentVisitor interface {
	intmod.ITypeArgumentVisitor

	Init()
}

type IDeclaredSyntheticLocalVariableVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init()
}

type IEraseTypeArgumentVisitor interface {
	intmod.ITypeVisitor

	Init()
	Type() intmod.IType
}

type IGenerateParameterSuffixNameVisitor interface {
	intmod.ITypeArgumentVisitor

	Suffix() string
}

type IGetTypeArgumentVisitor interface {
	intmod.ITypeVisitor

	Init()
	TypeArguments() intmod.ITypeArgument
}

type IInitEnumVisitor interface {
	intmod.IJavaSyntaxVisitor

	Constants() util.IList[intmod.IConstant]
}

type IInitInnerClassVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateNewExpressionVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateParametersAndLocalVariablesVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IAddLocalClassDeclarationVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IInitInstanceFieldVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IData interface {
	Declaration() IClassFileConstructorDeclaration
	SetDeclaration(declaration IClassFileConstructorDeclaration)
	Statements() intmod.IStatements
	SetStatements(statements intmod.IStatements)
	Index() int
	SetIndex(index int)
}

type IInitStaticFieldVisitor interface {
	intmod.IJavaSyntaxVisitor

	SetInternalTypeName(internalTypeName string)
}

type IMergeTryWithResourcesStatementVisitor interface {
	intmod.IStatementVisitor

	SafeAccept(list intmod.IStatement)
	AcceptSlice(list []intmod.IStatement)
	SafeAcceptSlice(list []intmod.IStatement)
}

type IPopulateBindingsWithTypeArgumentVisitor interface {
	intmod.ITypeArgumentVisitor

	Init(contextualTypeBounds map[string]intmod.IType, bindings map[string]intmod.ITypeArgument,
		typeBounds map[string]intmod.IType, typeArgument intmod.ITypeArgument)
}

type IPopulateBindingsWithTypeParameterVisitor interface {
	intmod.ITypeParameterVisitor

	Init(bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType)
}

type IPopulateBlackListNamesVisitor interface {
	intmod.ITypeArgumentVisitor
}

type IRemoveBinaryOpReturnStatementsVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IRemoveDefaultConstructorVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IRemoveFinallyStatementsVisitor interface {
	intmod.IStatementVisitor

	Init()
}

type IRemoveLastContinueStatementVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IRenameLocalVariablesVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init(nameMapping map[string]string)
}

type ISearchFirstLineNumberVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init()
	LineNumber() int
}

type ISearchFromOffsetVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init()
	Offset() int
}

type ISearchInTypeArgumentVisitor interface {
	intmod.ITypeArgumentVisitor

	Init()
	ContainsWildcard() bool
	ContainsWildcardSuperOrExtendsType() bool
	ContainsGeneric() bool
}

type ISearchLocalVariableReferenceVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init(index int)
	ContainsReference() bool
}

type ISearchLocalVariableVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init()
	Variables() []ILocalVariable
}

type ISearchUndeclaredLocalVariableVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init()
	Variables() []ILocalVariable
	RemoveAll(removal []ILocalVariable)
}

type ISortMembersVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type ITypeArgumentToTypeVisitor interface {
	intmod.ITypeArgumentVisitor

	Init()
	Type() intmod.IType
}

type IUpdateBridgeMethodTypeVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateBridgeMethodVisitor interface {
	IUpdateExpressionVisitor

	Init(bodyDeclaration IClassFileBodyDeclaration) bool
}

type IBodyDeclarationsVisitor interface {
	intmod.IJavaSyntaxVisitor

	BridgeMethodDeclarationsLink() map[string]map[string]IClassFileMethodDeclaration
	SetBridgeMethodDeclarationsLink(bridgeMethodDeclarationsLink map[string]map[string]IClassFileMethodDeclaration)
	Mapped() map[string]IClassFileMethodDeclaration
	SetMapped(mapped map[string]IClassFileMethodDeclaration)
}

type IUpdateClassTypeArgumentsVisitor interface {
	IUpdateExpressionVisitor

	Init()
	TypeArgument() intmod.ITypeArgument
}

type IUpdateIntegerConstantTypeVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateJavaSyntaxTreeStep0Visitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateJavaSyntaxTreeStep1Visitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateJavaSyntaxTreeStep2Visitor interface {
	intmod.IJavaSyntaxVisitor
}

type IUpdateOuterFieldTypeVisitor interface {
	intmod.IJavaSyntaxVisitor
}

type ISearchFieldVisitor interface {
	intmod.IJavaSyntaxVisitor

	Init(name string)
	Found() bool
}

type IUpdateTypeVisitor interface {
	intmod.ITypeArgumentVisitor

	SetLocalVariableType(localVariableType intcls.ILocalVariableType)
}
