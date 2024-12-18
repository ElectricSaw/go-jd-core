package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewParenthesesExpression(expression intmod.IExpression) intmod.IParenthesesExpression {
	e := &ParenthesesExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(expression.LineNumber()),
		expression:                   expression,
	}
	e.SetValue(e)
	return e
}

type ParenthesesExpression struct {
	AbstractLineNumberExpression

	expression intmod.IExpression
}

func (e *ParenthesesExpression) Type() intmod.IType {
	return e.expression.Type()
}

func (e *ParenthesesExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *ParenthesesExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *ParenthesesExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitParenthesesExpression(e)
}
