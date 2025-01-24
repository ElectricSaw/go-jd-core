package expression

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLambdaExpression(typ model.IType, statements IStatement) LambdaExpression {
	return NewLambdaExpressionWithAll(model.UnknownLineNumber, typ, statements)
}

func NewLambdaExpressionWithAll(lineNumber int, typ model.IType, statements IStatement) LambdaExpression {
	v := LambdaExpression{
		DefaultBase: *util.NewDefaultBase[IExpression]().(*util.DefaultBase[IExpression]),
		LineNumber:  lineNumber,
		Type:        typ,
		Priority:    -1,
		Statements:  statements,
	}
	v.SetValue(&v)
	return v
}

type LambdaExpression struct {
	util.DefaultBase[IExpression]

	LineNumber int
	Type       model.IType
	Priority   int
	Statements IStatement
}

func (e *LambdaExpression) Accept(visitor IExpressionVisitor) {}

func (e *LambdaExpression) IsArrayExpression() bool                      { return false }
func (e *LambdaExpression) IsBinaryOperatorExpression() bool             { return false }
func (e *LambdaExpression) IsBooleanExpression() bool                    { return false }
func (e *LambdaExpression) IsCastExpression() bool                       { return false }
func (e *LambdaExpression) IsConstructorInvocationExpression() bool      { return false }
func (e *LambdaExpression) IsDoubleConstantExpression() bool             { return false }
func (e *LambdaExpression) IsFieldReferenceExpression() bool             { return false }
func (e *LambdaExpression) IsFloatConstantExpression() bool              { return false }
func (e *LambdaExpression) IsIntegerConstantExpression() bool            { return false }
func (e *LambdaExpression) IsLengthExpression() bool                     { return false }
func (e *LambdaExpression) IsLocalVariableReferenceExpression() bool     { return false }
func (e *LambdaExpression) IsLongConstantExpression() bool               { return false }
func (e *LambdaExpression) IsMethodInvocationExpression() bool           { return false }
func (e *LambdaExpression) IsNewArray() bool                             { return false }
func (e *LambdaExpression) IsNewExpression() bool                        { return false }
func (e *LambdaExpression) IsNewInitializedArray() bool                  { return false }
func (e *LambdaExpression) IsNullExpression() bool                       { return false }
func (e *LambdaExpression) IsObjectTypeReferenceExpression() bool        { return false }
func (e *LambdaExpression) IsPostOperatorExpression() bool               { return false }
func (e *LambdaExpression) IsPreOperatorExpression() bool                { return false }
func (e *LambdaExpression) IsStringConstantExpression() bool             { return false }
func (e *LambdaExpression) IsSuperConstructorInvocationExpression() bool { return false }
func (e *LambdaExpression) IsSuperExpression() bool                      { return false }
func (e *LambdaExpression) IsTernaryOperatorExpression() bool            { return false }
func (e *LambdaExpression) IsThisExpression() bool                       { return false }

func (e *LambdaExpression) DimensionExpressionList() IExpression { return NeNoExpression }
func (e *LambdaExpression) Parameters() IExpression              { return NeNoExpression }
func (e *LambdaExpression) Condition() IExpression               { return NeNoExpression }
func (e *LambdaExpression) Expression() IExpression              { return NeNoExpression }
func (e *LambdaExpression) TrueExpression() IExpression          { return NeNoExpression }
func (e *LambdaExpression) FalseExpression() IExpression         { return NeNoExpression }
func (e *LambdaExpression) Index() IExpression                   { return NeNoExpression }
func (e *LambdaExpression) LeftExpression() IExpression          { return NeNoExpression }
func (e *LambdaExpression) RightExpression() IExpression         { return NeNoExpression }
func (e *LambdaExpression) Descriptor() string                   { return "" }
func (e *LambdaExpression) DoubleValue() float64                 { return 0 }
func (e *LambdaExpression) FloatValue() float32                  { return 0 }
func (e *LambdaExpression) IntegerValue() int                    { return 0 }
func (e *LambdaExpression) InternalTypeName() string             { return "" }
func (e *LambdaExpression) LongValue() int64                     { return 0 }
func (e *LambdaExpression) Name() string                         { return "" }
func (e *LambdaExpression) ObjectType() *model.ObjectType        { return &model.OtTypeUndefinedObject }
func (e *LambdaExpression) Operator() string                     { return "" }
func (e *LambdaExpression) StringValue() string                  { return "" }

func (e *LambdaExpression) String() string {
	return fmt.Sprintf("LambdaExpression{ line-number: %d, priority: %d }", e.LineNumber, e.Priority)
}
