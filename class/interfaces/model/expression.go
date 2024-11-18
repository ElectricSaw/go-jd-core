package model

import (
	"bitbucket.org/coontec/javaClass/class/api"
	"bitbucket.org/coontec/javaClass/class/util"
)

var UnknownLineNumber = api.UnknownLineNumber

type IExpression interface {
	util.Base[IExpression]

	LineNumber() int
	Type() IType
	Priority() int

	Accept(visitor IExpressionVisitor)

	IsArrayExpression() bool
	IsBinaryOperatorExpression() bool
	IsBooleanExpression() bool
	IsCastExpression() bool
	IsConstructorInvocationExpression() bool
	IsDoubleConstantExpression() bool
	IsFieldReferenceExpression() bool
	IsFloatConstantExpression() bool
	IsIntegerConstantExpression() bool
	IsLengthExpression() bool
	IsLocalVariableReferenceExpression() bool
	IsLongConstantExpression() bool
	IsMethodInvocationExpression() bool
	IsNewArray() bool
	IsNewExpression() bool
	IsNewInitializedArray() bool
	IsNullExpression() bool
	IsObjectTypeReferenceExpression() bool
	IsPostOperatorExpression() bool
	IsPreOperatorExpression() bool
	IsStringConstantExpression() bool
	IsSuperConstructorInvocationExpression() bool
	IsSuperExpression() bool
	IsTernaryOperatorExpression() bool
	IsThisExpression() bool

	DimensionExpressionList() IExpression
	Parameters() IExpression
	Condition() IExpression
	Expression() IExpression
	TrueExpression() IExpression
	FalseExpression() IExpression
	Index() IExpression
	LeftExpression() IExpression
	RightExpression() IExpression
	Descriptor() string
	DoubleValue() float64
	FloatValue() float32
	IntegerValue() int
	InternalTypeName() string
	LongValue() int64
	Name() string
	ObjectType() IObjectType
	Operator() string
	StringValue() string
}

type IExpressionVisitor interface {
	VisitArrayExpression(expression IArrayExpression)
	VisitBinaryOperatorExpression(expression IBinaryOperatorExpression)
	VisitBooleanExpression(expression IBooleanExpression)
	VisitCastExpression(expression ICastExpression)
	VisitCommentExpression(expression ICommentExpression)
	VisitConstructorInvocationExpression(expression IConstructorInvocationExpression)
	VisitConstructorReferenceExpression(expression IConstructorReferenceExpression)
	VisitDoubleConstantExpression(expression IDoubleConstantExpression)
	VisitEnumConstantReferenceExpression(expression IEnumConstantReferenceExpression)
	VisitExpressions(expression IExpressions)
	VisitFieldReferenceExpression(expression IFieldReferenceExpression)
	VisitFloatConstantExpression(expression IFloatConstantExpression)
	VisitIntegerConstantExpression(expression IIntegerConstantExpression)
	VisitInstanceOfExpression(expression IInstanceOfExpression)
	VisitLambdaFormalParametersExpression(expression ILambdaFormalParametersExpression)
	VisitLambdaIdentifiersExpression(expression ILambdaIdentifiersExpression)
	VisitLengthExpression(expression ILengthExpression)
	VisitLocalVariableReferenceExpression(expression ILocalVariableReferenceExpression)
	VisitLongConstantExpression(expression ILongConstantExpression)
	VisitMethodInvocationExpression(expression IMethodInvocationExpression)
	VisitMethodReferenceExpression(expression IMethodReferenceExpression)
	VisitNewArray(expression INewArray)
	VisitNewExpression(expression INewExpression)
	VisitNewInitializedArray(expression INewInitializedArray)
	VisitNoExpression(expression INoExpression)
	VisitNullExpression(expression INullExpression)
	VisitObjectTypeReferenceExpression(expression IObjectTypeReferenceExpression)
	VisitParenthesesExpression(expression IParenthesesExpression)
	VisitPostOperatorExpression(expression IPostOperatorExpression)
	VisitPreOperatorExpression(expression IPreOperatorExpression)
	VisitStringConstantExpression(expression IStringConstantExpression)
	VisitSuperConstructorInvocationExpression(expression ISuperConstructorInvocationExpression)
	VisitSuperExpression(expression ISuperExpression)
	VisitTernaryOperatorExpression(expression ITernaryOperatorExpression)
	VisitThisExpression(expression IThisExpression)
	VisitTypeReferenceDotClassExpression(expression ITypeReferenceDotClassExpression)
}

type IArrayExpression interface {
	Type() IType
	SetType(typ IType)
	Expression() IExpression
	Index() IExpression
	Priority() int
	SetExpression(expression IExpression)
	SetIndex(index IExpression)
	IsArrayExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IBinaryOperatorExpression interface {
	LeftExpression() IExpression
	Operator() string
	RightExpression() IExpression
	Priority() int
	SetLeftExpression(leftExpression IExpression)
	SetOperator(operator string)
	SetRightExpression(rightExpression IExpression)
	SetPriority(priority int)
	IsBinaryOperatorExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IBooleanExpression interface {
	Type() IType
	IsTrue() bool
	IsFalse() bool
	IsBooleanExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ICastExpression interface {
	Type() IType
	SetType(typ IType)
	Expression() IExpression
	IsExplicit() bool
	Priority() int
	SetExpression(expression IExpression)
	SetExplicit(explicit bool)
	IsCastExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ICommentExpression interface {
	LineNumber() int
	Type() IType
	Priority() int
	Text() string
	Accept(visitor IExpressionVisitor)
	String() string
}

type IConstructorInvocationExpression interface {
	Type() IType
	SetType(typ IType)
	Parameters() IExpression
	Priority() int
	SetParameters(params IExpression)
	IsConstructorInvocationExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IConstructorReferenceExpression interface {
	Type() IType
	SetType(typ IType)
	ObjectType() IObjectType
	Descriptor() string
	Accept(visitor IExpressionVisitor)
}

type IDoubleConstantExpression interface {
	Type() IType
	SetType(typ IType)
	DoubleValue() float64
	IsDoubleConstantExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IEnumConstantReferenceExpression interface {
	Type() IType
	SetType(typ IType)
	ObjectType() IObjectType
	Name() string
	Accept(visitor IExpressionVisitor)
	String() string
}

type IExpressions interface {
	util.IList[IExpression]
	Accept(visitor IExpressionVisitor)
}

type IFieldReferenceExpression interface {
	Type() IType
	SetType(typ IType)
	Expression() IExpression
	InternalTypeName() string
	Name() string
	Descriptor() string
	SetExpression(expression IExpression)
	SetName(name string)
	IsFieldReferenceExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IFloatConstantExpression interface {
	Type() IType
	SetType(typ IType)
	FloatValue() float32
	IsFloatConstantExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IInstanceOfExpression interface {
	Type() IType
	Expression() IExpression
	InstanceOfType() IType
	Priority() int
	SetExpression(expression IExpression)
	Accept(visitor IExpressionVisitor)
}

type IIntegerConstantExpression interface {
	Type() IType
	SetType(typ IType)
	IntegerValue() int
	IsIntegerConstantExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ILambdaFormalParametersExpression interface {
	Priority() int
	Statements() IStatement
	FormalParameters() IFormalParameter
	SetFormalParameters(formalParameters IFormalParameter)
	Accept(visitor IExpressionVisitor)
	String() string
}

type ILambdaIdentifiersExpression interface {
	Priority() int
	Statements() IStatement
	ReturnedType() IType
	ParameterNames() []string
	Accept(visitor IExpressionVisitor)
	String() string
}

type ILengthExpression interface {
	Type() IType
	Expression() IExpression
	SetExpression(expression IExpression)
	IsLengthExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ILocalVariableReferenceExpression interface {
	Type() IType
	SetType(typ IType)
	Name() string
	IsLocalVariableReferenceExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ILongConstantExpression interface {
	Type() IType
	SetType(typ IType)
	LongValue() int64
	IsLongConstantExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IMethodInvocationExpression interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	NonWildcardTypeArguments() ITypeArgument
	SetNonWildcardTypeArguments(arguments ITypeArgument)
	Parameters() IExpression
	SetParameters(params IExpression)
	Priority() int
	Accept(visitor IExpressionVisitor)
	String() string
}

type IMethodReferenceExpression interface {
	Expression() IExpression
	InternalTypeName() string
	Name() string
	Descriptor() string
	SetExpression(expression IExpression)
	Accept(visitor IExpressionVisitor)
}

type INewArray interface {
	Type() IType
	SetType(typ IType)
	DimensionExpressionList() IExpression
	SetDimensionExpressionList(dimensionExpressionList IExpression)
	Priority() int
	IsNewArray() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type INewExpression interface {
	ObjectType() IObjectType
	SetObjectType(objectType IObjectType)
	Type() IType
	SetType(objectType IObjectType)
	Priority() int
	Descriptor() string
	SetDescriptor(descriptor string)
	Parameters() IExpression
	SetParameters(params IExpression)
	BodyDeclaration() IBodyDeclaration
	IsNewExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type INewInitializedArray interface {
	Type() IType
	SetType(typ IType)
	ArrayInitializer() IArrayVariableInitializer
	Priority() int
	IsNewInitializedArray() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type INoExpression interface {
	Accept(visitor IExpressionVisitor)
	String() string
}

type INullExpression interface {
	Type() IType
	SetType(typ IType)
	IsNullExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IObjectTypeReferenceExpression interface {
	LineNumber() int
	ObjectType() IObjectType
	Type() IType
	IsExplicit() bool
	SetExplicit(explicit bool)
	Priority() int
	IsObjectTypeReferenceExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IParenthesesExpression interface {
	Type() IType
	Expression() IExpression
	SetExpression(expression IExpression)
	Accept(visitor IExpressionVisitor)
}

type IPostOperatorExpression interface {
	Operator() string
	Expression() IExpression
	SetExpression(expression IExpression)
	Type() IType
	Priority() int
	IsPostOperatorExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IPreOperatorExpression interface {
	Operator() string
	Expression() IExpression
	SetExpression(expression IExpression)
	Type() IType
	Priority() int
	IsPreOperatorExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IStringConstantExpression interface {
	StringValue() string
	Type() IType
	IsStringConstantExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ISuperConstructorInvocationExpression interface {
	Type() IType
	SetType(typ IType)
	Parameters() IExpression
	SetParameters(expression IExpression)
	Priority() int
	IsSuperConstructorInvocationExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ISuperExpression interface {
	Type() IType
	IsSuperExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ITernaryOperatorExpression interface {
	Condition() IExpression
	SetCondition(expression IExpression)
	TrueExpression() IExpression
	SetTrueExpression(expression IExpression)
	FalseExpression() IExpression
	SetFalseExpression(expression IExpression)
	Priority() int
	IsTernaryOperatorExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type IThisExpression interface {
	Type() IType
	IsExplicit() bool
	SetExplicit(explicit bool)
	IsThisExpression() bool
	Accept(visitor IExpressionVisitor)
	String() string
}

type ITypeReferenceDotClassExpression interface {
	LineNumber() int
	TypeDotClass() IType
	Type() IType
	Priority() int
	Accept(visitor IExpressionVisitor)
	String() string
}
