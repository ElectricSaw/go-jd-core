package expression

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

type AbstractExpression struct {
	util.DefaultBase[IExpression]
}

func (e *AbstractExpression) Accept(visitor IExpressionVisitor) {}

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

func (e *AbstractExpression) DimensionExpressionList() IExpression { return NeNoExpression }
func (e *AbstractExpression) Parameters() IExpression              { return NeNoExpression }
func (e *AbstractExpression) Condition() IExpression               { return NeNoExpression }
func (e *AbstractExpression) Expression() IExpression              { return NeNoExpression }
func (e *AbstractExpression) TrueExpression() IExpression          { return NeNoExpression }
func (e *AbstractExpression) FalseExpression() IExpression         { return NeNoExpression }
func (e *AbstractExpression) Index() IExpression                   { return NeNoExpression }
func (e *AbstractExpression) LeftExpression() IExpression          { return NeNoExpression }
func (e *AbstractExpression) RightExpression() IExpression         { return NeNoExpression }
func (e *AbstractExpression) Descriptor() string                   { return "" }
func (e *AbstractExpression) DoubleValue() float64                 { return 0 }
func (e *AbstractExpression) FloatValue() float32                  { return 0 }
func (e *AbstractExpression) IntegerValue() int                    { return 0 }
func (e *AbstractExpression) InternalTypeName() string             { return "" }
func (e *AbstractExpression) LongValue() int64                     { return 0 }
func (e *AbstractExpression) Name() string                         { return "" }
func (e *AbstractExpression) ObjectType() IObjectType              { return model.OtTypeUndefinedObject }
func (e *AbstractExpression) Operator() string                     { return "" }
func (e *AbstractExpression) StringValue() string                  { return "" }
