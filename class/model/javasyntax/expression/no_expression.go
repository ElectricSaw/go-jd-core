package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() intmod.INoExpression {
	e := &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeVoid.(intmod.IType)),
	}
	e.SetValue(e)
	return e
}

type NoExpression struct {
	AbstractLineNumberTypeExpression
}

func (e *NoExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) String() string {
	return "NoExpression"
}
