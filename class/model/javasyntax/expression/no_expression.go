package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() *NoExpression {
	return &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeVoid),
	}
}

type NoExpression struct {
	AbstractLineNumberTypeExpression
}

func (e *NoExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) String() string {
	return "NoExpression"
}
