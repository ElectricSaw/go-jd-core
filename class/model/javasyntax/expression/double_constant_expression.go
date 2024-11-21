package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewDoubleConstantExpression(value float64) intmod.IDoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeDouble.(intmod.IType)),
		value:                            value,
	}
}

func NewDoubleConstantExpressionWithAll(lineNumber int, value float64) intmod.IDoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeDouble.(intmod.IType)),
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

func (e *DoubleConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitDoubleConstantExpression(e)
}

func (e *DoubleConstantExpression) String() string {
	return fmt.Sprintf("DoubleConstantExpression{%f}", e.value)
}
