package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewFloatConstantExpression(value float32) intmod.IFloatConstantExpression {
	return NewFloatConstantExpressionWithAll(0, value)
}

func NewFloatConstantExpressionWithAll(lineNumber int, value float32) intmod.IFloatConstantExpression {
	e := &FloatConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeFloat.(intmod.IType)),
		value:                            value,
	}
	e.SetValue(e)
	return e
}

type FloatConstantExpression struct {
	AbstractLineNumberTypeExpression

	value float32
}

func (e *FloatConstantExpression) FloatValue() float32 {
	return e.value
}

func (e *FloatConstantExpression) IsFloatConstantExpression() bool {
	return true
}

func (e *FloatConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitFloatConstantExpression(e)
}

func (e *FloatConstantExpression) String() string {
	return fmt.Sprintf("FloatConstantExpression{%.2f}", e.value)
}
