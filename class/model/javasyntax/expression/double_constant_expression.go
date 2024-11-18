package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewDoubleConstantExpression(value float64) intsyn.IDoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeDouble.(intsyn.IType)),
		value:                            value,
	}
}

func NewDoubleConstantExpressionWithAll(lineNumber int, value float64) intsyn.IDoubleConstantExpression {
	return &DoubleConstantExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, _type.PtTypeDouble.(intsyn.IType)),
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

func (e *DoubleConstantExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitDoubleConstantExpression(e)
}

func (e *DoubleConstantExpression) String() string {
	return fmt.Sprintf("DoubleConstantExpression{%f}", e.value)
}
