package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewDoubleConstantExpression(value float64) *DoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeDouble),
		value:                            value,
	}
}

func NewDoubleConstantExpressionWithAll(lineNumber int, value float64) *DoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeDouble),
		value:                            value,
	}
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

func (e *DoubleConstantExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitDoubleConstantExpression(e)
}

func (e *DoubleConstantExpression) String() string {
	return fmt.Sprintf("DoubleConstantExpression{%f}", e.value)
}
