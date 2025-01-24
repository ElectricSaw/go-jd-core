package expression

import "github.com/ElectricSaw/go-jd-core/decompiler/model"

type IExpression interface {
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
	ObjectType() *model.ObjectType
	Operator() string
	StringValue() string

	String() string
}

type IExpressionVisitor interface {
	VisitArrayExpression(expression *ArrayExpression)
	VisitBinaryOperatorExpression(expression *BinaryOperatorExpression)
	VisitBooleanExpression(expression *BooleanExpression)
	VisitCastExpression(expression *CastExpression)
	VisitCommentExpression(expression *CommentExpression)
	VisitConstructorInvocationExpression(expression *ConstructorInvocationExpression)
	VisitConstructorReferenceExpression(expression *ConstructorReferenceExpression)
	VisitDoubleConstantExpression(expression *DoubleConstantExpression)
	VisitEnumConstantReferenceExpression(expression *EnumConstantReferenceExpression)
	VisitExpressions(expression *Expressions)
	VisitFieldReferenceExpression(expression *FieldReferenceExpression)
	VisitFloatConstantExpression(expression *FloatConstantExpression)
	VisitIntegerConstantExpression(expression *IntegerConstantExpression)
	VisitInstanceOfExpression(expression *InstanceOfExpression)
	VisitLambdaFormalParametersExpression(expression *LambdaFormalParametersExpression)
	VisitLambdaIdentifiersExpression(expression *LambdaIdentifiersExpression)
	VisitLengthExpression(expression *LengthExpression)
	VisitLocalVariableReferenceExpression(expression *LocalVariableReferenceExpression)
	VisitLongConstantExpression(expression *LongConstantExpression)
	VisitMethodInvocationExpression(expression *MethodInvocationExpression)
	VisitMethodReferenceExpression(expression *MethodReferenceExpression)
	VisitNewArray(expression *NewArray)
	VisitNewExpression(expression *NewExpression)
	VisitNewInitializedArray(expression *NewInitializedArray)
	VisitNoExpression(expression *NoExpression)
	VisitNullExpression(expression *NullExpression)
	VisitObjectTypeReferenceExpression(expression *ObjectTypeReferenceExpression)
	VisitParenthesesExpression(expression *ParenthesesExpression)
	VisitPostOperatorExpression(expression *PostOperatorExpression)
	VisitPreOperatorExpression(expression *PreOperatorExpression)
	VisitStringConstantExpression(expression *StringConstantExpression)
	VisitSuperConstructorInvocationExpression(expression *SuperConstructorInvocationExpression)
	VisitSuperExpression(expression *SuperExpression)
	VisitTernaryOperatorExpression(expression *TernaryOperatorExpression)
	VisitThisExpression(expression *ThisExpression)
	VisitTypeReferenceDotClassExpression(expression *TypeReferenceDotClassExpression)
}
