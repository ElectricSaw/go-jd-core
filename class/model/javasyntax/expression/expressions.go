package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewExpressions() intmod.IExpressions {
	return &Expressions{}
}

type Expressions struct {
	AbstractExpression
	util.DefaultList[intmod.IExpression]
}

func (e *Expressions) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitExpressions(e)
}
