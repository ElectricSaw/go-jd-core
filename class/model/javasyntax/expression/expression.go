package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/util"
)

var UnknownLineNumber = 0

type AbstractExpression struct {
	util.DefaultBase[intsyn.IExpression]
}

func (e *AbstractExpression) LineNumber() int    { return UnknownLineNumber }
func (e *AbstractExpression) Type() intsyn.IType { return nil }
func (e *AbstractExpression) Priority() int      { return -1 }

func (e *AbstractExpression) Accept(visitor intsyn.IExpressionVisitor) {}

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

func (e *AbstractExpression) DimensionExpressionList() intsyn.IExpression { return NeNoExpression }
func (e *AbstractExpression) Parameters() intsyn.IExpression              { return NeNoExpression }
func (e *AbstractExpression) Condition() intsyn.IExpression               { return NeNoExpression }
func (e *AbstractExpression) Expression() intsyn.IExpression              { return NeNoExpression }
func (e *AbstractExpression) TrueExpression() intsyn.IExpression          { return NeNoExpression }
func (e *AbstractExpression) FalseExpression() intsyn.IExpression         { return NeNoExpression }
func (e *AbstractExpression) Index() intsyn.IExpression                   { return NeNoExpression }
func (e *AbstractExpression) LeftExpression() intsyn.IExpression          { return NeNoExpression }
func (e *AbstractExpression) RightExpression() intsyn.IExpression         { return NeNoExpression }
func (e *AbstractExpression) Descriptor() string                          { return "" }
func (e *AbstractExpression) DoubleValue() float64                        { return 0 }
func (e *AbstractExpression) FloatValue() float32                         { return 0 }
func (e *AbstractExpression) IntegerValue() int                           { return 0 }
func (e *AbstractExpression) InternalTypeName() string                    { return "" }
func (e *AbstractExpression) LongValue() int64                            { return 0 }
func (e *AbstractExpression) Name() string                                { return "" }
func (e *AbstractExpression) ObjectType() intsyn.IObjectType              { return _type.OtTypeUndefinedObject }
func (e *AbstractExpression) Operator() string                            { return "" }
func (e *AbstractExpression) StringValue() string                         { return "" }
