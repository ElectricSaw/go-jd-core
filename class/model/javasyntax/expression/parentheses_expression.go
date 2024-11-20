package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewParenthesesExpression(expression intsyn.IExpression) intsyn.IParenthesesExpression {
	return &ParenthesesExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(expression.LineNumber()),
		expression:                   expression,
	}
}

type ParenthesesExpression struct {
	AbstractLineNumberExpression

	expression intsyn.IExpression
}

func (e *ParenthesesExpression) Type() intsyn.IType {
	return e.expression.Type()
}

func (e *ParenthesesExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *ParenthesesExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *ParenthesesExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitParenthesesExpression(e)
}
