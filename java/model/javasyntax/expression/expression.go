package expression

import _type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"

var UnknownLineNumber = 0

type AbstractExpression struct {
}

func (e *AbstractExpression) GetLineNumber() int   { return UnknownLineNumber }
func (e *AbstractExpression) GetType() _type.IType { return nil }
func (e *AbstractExpression) GetPriority() int     { return -1 }

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

func (e *AbstractExpression) GetDimensionExpressionList() Expression { return NeNoExpression }
func (e *AbstractExpression) GetParameters() Expression              { return NeNoExpression }
func (e *AbstractExpression) GetCondition() Expression               { return NeNoExpression }
func (e *AbstractExpression) GetExpression() Expression              { return NeNoExpression }
func (e *AbstractExpression) GetTrueExpression() Expression          { return NeNoExpression }
func (e *AbstractExpression) GetFalseExpression() Expression         { return NeNoExpression }
func (e *AbstractExpression) GetIndex() Expression                   { return NeNoExpression }
func (e *AbstractExpression) GetLeftExpression() Expression          { return NeNoExpression }
func (e *AbstractExpression) GetRightExpression() Expression         { return NeNoExpression }
func (e *AbstractExpression) GetDescriptor() string                  { return "" }
func (e *AbstractExpression) GetDoubleValue() float64                { return 0 }
func (e *AbstractExpression) GetFloatValue() float32                 { return 0 }
func (e *AbstractExpression) GetIntegerValue() int                   { return 0 }
func (e *AbstractExpression) GetInternalTypeName() string            { return "" }
func (e *AbstractExpression) GetLongValue() int64                    { return 0 }
func (e *AbstractExpression) GetName() string                        { return "" }
func (e *AbstractExpression) GetObjectType() *_type.ObjectType       { return _type.TypeUndefinedObject }
func (e *AbstractExpression) GetOperator() string                    { return "" }
func (e *AbstractExpression) GetStringValue() string                 { return "" }

type Expression interface {
	GetLineNumber() int
	GetType() _type.IType
	GetPriority() int

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

	GetDimensionExpressionList() Expression
	GetParameters() Expression
	GetCondition() Expression
	GetExpression() Expression
	GetTrueExpression() Expression
	GetFalseExpression() Expression
	GetIndex() Expression
	GetLeftExpression() Expression
	GetRightExpression() Expression
	GetDescriptor() string
	GetDoubleValue() float64
	GetFloatValue() float32
	GetIntegerValue() int
	GetInternalTypeName() string
	GetLongValue() int64
	GetName() string
	GetObjectType() *_type.ObjectType
	GetOperator() string
	GetStringValue() string
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
