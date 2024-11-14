package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

var UnknownLineNumber = 0

type AbstractExpression struct {
}

func (e *AbstractExpression) LineNumber() int   { return UnknownLineNumber }
func (e *AbstractExpression) Type() _type.IType { return nil }
func (e *AbstractExpression) Priority() int     { return -1 }
func (e *AbstractExpression) Size() int {
	return 1
}

func (e *AbstractExpression) Accept(visitor ExpressionVisitor) {}

func (e *AbstractExpression) IsArrayExpression() bool                      { return false }
func (e *AbstractExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *AbstractExpression) IsBooleanExpression() bool                    { return false }
func (e *AbstractExpression) IsCastExpression() bool                       { return false }
func (e *AbstractExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *AbstractExpression) IsDoubleConstantExpression() bool             { return false }
func (e *AbstractExpression) IsFieldReferenceExpression() bool             { return false }
func (e *AbstractExpression) IsFloatConstantExpression() bool              { return false }
func (e *AbstractExpression) IsIntegerConstantExpression() bool            { return false }
func (e *AbstractExpression) IsLengthExpression() bool                     { return false }
func (e *AbstractExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *AbstractExpression) IsLongConstantExpression() bool               { return false }
func (e *AbstractExpression) IsMethodInvocationExpression() bool           { return false }
func (e *AbstractExpression) IsNewArray() bool                             { return false }
func (e *AbstractExpression) IsNewExpression() bool                        { return false }
func (e *AbstractExpression) IsNewInitializedArray() bool                  { return false }
func (e *AbstractExpression) IsNullExpression() bool                       { return false }
func (e *AbstractExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *AbstractExpression) IsPostOperatorExpression() bool               { return false }
func (e *AbstractExpression) IsPreOperatorExpression() bool                { return false }
func (e *AbstractExpression) IsStringConstantExpression() bool             { return false }
func (e *AbstractExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *AbstractExpression) IsSuperExpression() bool                      { return false }
func (e *AbstractExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *AbstractExpression) IsThisExpression() bool                       { return false }

func (e *AbstractExpression) DimensionExpressionList() Expression { return NeNoExpression }
func (e *AbstractExpression) Parameters() Expression              { return NeNoExpression }
func (e *AbstractExpression) Condition() Expression               { return NeNoExpression }
func (e *AbstractExpression) Expression() Expression              { return NeNoExpression }
func (e *AbstractExpression) TrueExpression() Expression          { return NeNoExpression }
func (e *AbstractExpression) FalseExpression() Expression         { return NeNoExpression }
func (e *AbstractExpression) Index() Expression                   { return NeNoExpression }
func (e *AbstractExpression) LeftExpression() Expression          { return NeNoExpression }
func (e *AbstractExpression) RightExpression() Expression         { return NeNoExpression }
func (e *AbstractExpression) Descriptor() string                  { return "" }
func (e *AbstractExpression) DoubleValue() float64                { return 0 }
func (e *AbstractExpression) FloatValue() float32                 { return 0 }
func (e *AbstractExpression) IntegerValue() int                   { return 0 }
func (e *AbstractExpression) InternalTypeName() string            { return "" }
func (e *AbstractExpression) LongValue() int64                    { return 0 }
func (e *AbstractExpression) Name() string                        { return "" }
func (e *AbstractExpression) ObjectType() _type.IObjectType       { return _type.OtTypeUndefinedObject }
func (e *AbstractExpression) Operator() string                    { return "" }
func (e *AbstractExpression) StringValue() string                 { return "" }

type Expression interface {
	LineNumber() int
	Type() _type.IType
	Priority() int
	Size() int

	Accept(visitor ExpressionVisitor)

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

	DimensionExpressionList() Expression
	Parameters() Expression
	Condition() Expression
	Expression() Expression
	TrueExpression() Expression
	FalseExpression() Expression
	Index() Expression
	LeftExpression() Expression
	RightExpression() Expression
	Descriptor() string
	DoubleValue() float64
	FloatValue() float32
	IntegerValue() int
	InternalTypeName() string
	LongValue() int64
	Name() string
	ObjectType() _type.IObjectType
	Operator() string
	StringValue() string
}

type ExpressionVisitor interface {
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
