package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewDoubleConstantExpression(value float64) intmod.IDoubleConstantExpression {
	return NewDoubleConstantExpressionWithAll(0, value)
}

func NewDoubleConstantExpressionWithAll(lineNumber int, value float64) intmod.IDoubleConstantExpression {
	e := &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeDouble.(intmod.IType)),
		value:                            value,
	}
	e.SetValue(e)
	return e
}

type DoubleConstantExpression struct {
	AbstractLineNumberTypeExpression

	value float64
}

func (e *DoubleConstantExpression) DoubleValue() float64 {
	return e.value
}

func (e *DoubleConstantExpression) IsDoubleConstantExpression() bool {
	return true
}

func (e *DoubleConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitDoubleConstantExpression(e)
}

func (e *DoubleConstantExpression) String() string {
	return fmt.Sprintf("DoubleConstantExpression{%f}", e.value)
}
