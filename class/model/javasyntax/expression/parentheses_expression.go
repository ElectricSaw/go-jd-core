package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewParenthesesExpression(expression Expression) *ParenthesesExpression {
	return &ParenthesesExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(expression.LineNumber()),
		expression:                   expression,
	}
}

type ParenthesesExpression struct {
	AbstractLineNumberExpression

	expression Expression
}

func (e *ParenthesesExpression) Type() _type.IType {
	return e.expression.Type()
}

func (e *ParenthesesExpression) Expression() Expression {
	return e.expression
}

func (e *ParenthesesExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *ParenthesesExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitParenthesesExpression(e)
}
