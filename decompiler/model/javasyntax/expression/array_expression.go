package expression

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func CreateItemType(expression IExpression) model.IType {
	typ := expression.Type()
	dimension := typ.Dimension()

	if dimension > 0 {
		return typ.CreateType(dimension - 1)
	}

	return typ.CreateType(0)
}

func NewArrayExpression(expression IExpression, index IExpression) ArrayExpression {
	return NewArrayExpressionWithAll(0, expression, index)
}

func NewArrayExpressionWithAll(lineNumber int, expression IExpression, index IExpression) ArrayExpression {
	e := ArrayExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        CreateItemType(expression),
		Priority:    -1,
		expression:  expression,
		index:       index,
	}
	e.SetValue(&e)
	return e
}

type ArrayExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       model.IType
	Priority   int
	expression IExpression
	index      IExpression
}

func (e *ArrayExpression) Accept(visitor IExpressionVisitor) {
	visitor.VisitArrayExpression(e)
}

func (e *ArrayExpression) IsArrayExpression() bool                      { return true }
func (e *ArrayExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *ArrayExpression) IsBooleanExpression() bool                    { return false }
func (e *ArrayExpression) IsCastExpression() bool                       { return false }
func (e *ArrayExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *ArrayExpression) IsDoubleConstantExpression() bool             { return false }
func (e *ArrayExpression) IsFieldReferenceExpression() bool             { return false }
func (e *ArrayExpression) IsFloatConstantExpression() bool              { return false }
func (e *ArrayExpression) IsIntegerConstantExpression() bool            { return false }
func (e *ArrayExpression) IsLengthExpression() bool                     { return false }
func (e *ArrayExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *ArrayExpression) IsLongConstantExpression() bool               { return false }
func (e *ArrayExpression) IsMethodInvocationExpression() bool           { return false }
func (e *ArrayExpression) IsNewArray() bool                             { return false }
func (e *ArrayExpression) IsNewExpression() bool                        { return false }
func (e *ArrayExpression) IsNewInitializedArray() bool                  { return false }
func (e *ArrayExpression) IsNullExpression() bool                       { return false }
func (e *ArrayExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *ArrayExpression) IsPostOperatorExpression() bool               { return false }
func (e *ArrayExpression) IsPreOperatorExpression() bool                { return false }
func (e *ArrayExpression) IsStringConstantExpression() bool             { return false }
func (e *ArrayExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *ArrayExpression) IsSuperExpression() bool                      { return false }
func (e *ArrayExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *ArrayExpression) IsThisExpression() bool                       { return false }

func (e *ArrayExpression) DimensionExpressionList() IExpression { return NeNoExpression }
func (e *ArrayExpression) Parameters() IExpression              { return NeNoExpression }
func (e *ArrayExpression) Condition() IExpression               { return NeNoExpression }
func (e *ArrayExpression) Expression() IExpression              { return NeNoExpression }
func (e *ArrayExpression) TrueExpression() IExpression          { return NeNoExpression }
func (e *ArrayExpression) FalseExpression() IExpression         { return NeNoExpression }
func (e *ArrayExpression) Index() IExpression                   { return NeNoExpression }
func (e *ArrayExpression) LeftExpression() IExpression          { return NeNoExpression }
func (e *ArrayExpression) RightExpression() IExpression         { return NeNoExpression }
func (e *ArrayExpression) Descriptor() string                   { return "" }
func (e *ArrayExpression) DoubleValue() float64                 { return 0 }
func (e *ArrayExpression) FloatValue() float32                  { return 0 }
func (e *ArrayExpression) IntegerValue() int                    { return 0 }
func (e *ArrayExpression) InternalTypeName() string             { return "" }
func (e *ArrayExpression) LongValue() int64                     { return 0 }
func (e *ArrayExpression) Name() string                         { return "" }
func (e *ArrayExpression) ObjectType() *model.ObjectType {
	return &model.OtTypeUndefinedObject
}
func (e *ArrayExpression) Operator() string    { return "" }
func (e *ArrayExpression) StringValue() string { return "" }

func (e *ArrayExpression) String() string {
	return fmt.Sprintf("ArrayExpression{%v[%v]}", e.expression, e.index)
}
