package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

var NeNoExpression = NewNoExpression()

func NewNoExpression() intsyn.IExpression {
	return &NoExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(_type.PtTypeVoid.(intsyn.IType)),
	}
}

type NoExpression struct {
	AbstractLineNumberTypeExpression
	util.DefaultBase[intsyn.IExpression]
}

func (e *NoExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitNoExpression(e)
}

func (e *NoExpression) String() string {
	return "NoExpression"
}
