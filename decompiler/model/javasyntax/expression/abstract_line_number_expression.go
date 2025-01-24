package expression

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLineNumberExpressionEmpty() LineNumberExpression {
	return NewLineNumberExpression(model.UnknownLineNumber)
}

func NewLineNumberExpression(lineNumber int) LineNumberExpression {
	v := LineNumberExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        nil,
		Priority:    -1,
	}
	v.SetValue(&v)
	return v
}

type LineNumberExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       model.IType
	Priority   int
}

func (e *LineNumberExpression) Accept(visitor IExpressionVisitor) {}

func (e *LineNumberExpression) IsArrayExpression() bool                      { return false }
func (e *LineNumberExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LineNumberExpression) IsBooleanExpression() bool                    { return false }
func (e *LineNumberExpression) IsCastExpression() bool                       { return false }
func (e *LineNumberExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LineNumberExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LineNumberExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LineNumberExpression) IsFloatConstantExpression() bool              { return false }
func (e *LineNumberExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LineNumberExpression) IsLengthExpression() bool                     { return false }
func (e *LineNumberExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LineNumberExpression) IsLongConstantExpression() bool               { return false }
func (e *LineNumberExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LineNumberExpression) IsNewArray() bool                             { return false }
func (e *LineNumberExpression) IsNewExpression() bool                        { return false }
func (e *LineNumberExpression) IsNewInitializedArray() bool                  { return false }
func (e *LineNumberExpression) IsNullExpression() bool                       { return false }
func (e *LineNumberExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LineNumberExpression) IsPostOperatorExpression() bool               { return false }
func (e *LineNumberExpression) IsPreOperatorExpression() bool                { return false }
func (e *LineNumberExpression) IsStringConstantExpression() bool             { return false }
func (e *LineNumberExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LineNumberExpression) IsSuperExpression() bool                      { return false }
func (e *LineNumberExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LineNumberExpression) IsThisExpression() bool                       { return false }

func (e *LineNumberExpression) DimensionExpressionList() IExpression { return NeNoExpression }
func (e *LineNumberExpression) Parameters() IExpression              { return NeNoExpression }
func (e *LineNumberExpression) Condition() IExpression               { return NeNoExpression }
func (e *LineNumberExpression) Expression() IExpression              { return NeNoExpression }
func (e *LineNumberExpression) TrueExpression() IExpression          { return NeNoExpression }
func (e *LineNumberExpression) FalseExpression() IExpression         { return NeNoExpression }
func (e *LineNumberExpression) Index() IExpression                   { return NeNoExpression }
func (e *LineNumberExpression) LeftExpression() IExpression          { return NeNoExpression }
func (e *LineNumberExpression) RightExpression() IExpression         { return NeNoExpression }
func (e *LineNumberExpression) Descriptor() string                   { return "" }
func (e *LineNumberExpression) DoubleValue() float64                 { return 0 }
func (e *LineNumberExpression) FloatValue() float32                  { return 0 }
func (e *LineNumberExpression) IntegerValue() int                    { return 0 }
func (e *LineNumberExpression) InternalTypeName() string             { return "" }
func (e *LineNumberExpression) LongValue() int64                     { return 0 }
func (e *LineNumberExpression) Name() string                         { return "" }
func (e *LineNumberExpression) ObjectType() *model.ObjectType        { return &model.OtTypeUndefinedObject }
func (e *LineNumberExpression) Operator() string                     { return "" }
func (e *LineNumberExpression) StringValue() string                  { return "" }

func (e *LineNumberExpression) String() string {
	return fmt.Sprintf("LineNumberExpression{ line-number: %d, priority: %d }", e.LineNumber, e.Priority)
}
