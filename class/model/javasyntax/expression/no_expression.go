package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() intmod.IExpression {
	return &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeVoid.(intmod.IType)),
	}
}

type NoExpression struct {
	AbstractLineNumberTypeExpression
	util.DefaultBase[intmod.IExpression]
}

func (e *NoExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) String() string {
	return "NoExpression"
}
