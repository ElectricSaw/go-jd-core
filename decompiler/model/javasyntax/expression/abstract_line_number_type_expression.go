package expression

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLineNumberTypeExpression(typ model.IType) LineNumberTypeExpression {
	return NewLineNumberTypeExpressionWithAll(model.UnknownLineNumber, typ)
}

func NewLineNumberTypeExpressionWithAll(lineNumber int, typ model.IType) LineNumberTypeExpression {
	v := LineNumberTypeExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    -1,
	}
	v.SetValue(&v)
	return v
}

type LineNumberTypeExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       model.IType
	Priority   int
}

func (e *LineNumberTypeExpression) Accept(visitor IExpressionVisitor) {}

func (e *LineNumberTypeExpression) IsArrayExpression() bool                      { return false }
func (e *LineNumberTypeExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LineNumberTypeExpression) IsBooleanExpression() bool                    { return false }
func (e *LineNumberTypeExpression) IsCastExpression() bool                       { return false }
func (e *LineNumberTypeExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LineNumberTypeExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LineNumberTypeExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LineNumberTypeExpression) IsFloatConstantExpression() bool              { return false }
func (e *LineNumberTypeExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LineNumberTypeExpression) IsLengthExpression() bool                     { return false }
func (e *LineNumberTypeExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LineNumberTypeExpression) IsLongConstantExpression() bool               { return false }
func (e *LineNumberTypeExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LineNumberTypeExpression) IsNewArray() bool                             { return false }
func (e *LineNumberTypeExpression) IsNewExpression() bool                        { return false }
func (e *LineNumberTypeExpression) IsNewInitializedArray() bool                  { return false }
func (e *LineNumberTypeExpression) IsNullExpression() bool                       { return false }
func (e *LineNumberTypeExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LineNumberTypeExpression) IsPostOperatorExpression() bool               { return false }
func (e *LineNumberTypeExpression) IsPreOperatorExpression() bool                { return false }
func (e *LineNumberTypeExpression) IsStringConstantExpression() bool             { return false }
func (e *LineNumberTypeExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LineNumberTypeExpression) IsSuperExpression() bool                      { return false }
func (e *LineNumberTypeExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LineNumberTypeExpression) IsThisExpression() bool                       { return false }

func (e *LineNumberTypeExpression) DimensionExpressionList() IExpression { return NeNoExpression }
func (e *LineNumberTypeExpression) Parameters() IExpression              { return NeNoExpression }
func (e *LineNumberTypeExpression) Condition() IExpression               { return NeNoExpression }
func (e *LineNumberTypeExpression) Expression() IExpression              { return NeNoExpression }
func (e *LineNumberTypeExpression) TrueExpression() IExpression          { return NeNoExpression }
func (e *LineNumberTypeExpression) FalseExpression() IExpression         { return NeNoExpression }
func (e *LineNumberTypeExpression) Index() IExpression                   { return NeNoExpression }
func (e *LineNumberTypeExpression) LeftExpression() IExpression          { return NeNoExpression }
func (e *LineNumberTypeExpression) RightExpression() IExpression         { return NeNoExpression }
func (e *LineNumberTypeExpression) Descriptor() string                   { return "" }
func (e *LineNumberTypeExpression) DoubleValue() float64                 { return 0 }
func (e *LineNumberTypeExpression) FloatValue() float32                  { return 0 }
func (e *LineNumberTypeExpression) IntegerValue() int                    { return 0 }
func (e *LineNumberTypeExpression) InternalTypeName() string             { return "" }
func (e *LineNumberTypeExpression) LongValue() int64                     { return 0 }
func (e *LineNumberTypeExpression) Name() string                         { return "" }
func (e *LineNumberTypeExpression) ObjectType() *model.ObjectType {
	return &model.OtTypeUndefinedObject
}
func (e *LineNumberTypeExpression) Operator() string    { return "" }
func (e *LineNumberTypeExpression) StringValue() string { return "" }

func (e *LineNumberTypeExpression) String() string {
	return fmt.Sprintf("LineNumberTypeExpression{ line-number: %d, priority: %d }", e.LineNumber, e.Priority)
}
