package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewExpressions() intsyn.IExpressions {
	return &Expressions{}
}

type Expressions struct {
	AbstractExpression
	util.DefaultList[intsyn.IExpression]
}

func (e *Expressions) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitExpressions(e)
}
