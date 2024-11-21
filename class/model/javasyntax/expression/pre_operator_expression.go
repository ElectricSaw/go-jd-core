package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewPreOperatorExpression(operator string, expression intmod.IExpression) intmod.IPreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		operator:                     operator,
		expression:                   expression,
	}
}

func NewPreOperatorExpressionWithAll(lineNumber int, operator string, expression intmod.IExpression) intmod.IPreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
}

type PreOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression intmod.IExpression
}

func (e *PreOperatorExpression) Operator() string {
	return e.operator
}

func (e *PreOperatorExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *PreOperatorExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *PreOperatorExpression) Type() intmod.IType {
	return e.expression.Type()
}

func (e *PreOperatorExpression) Priority() int {
	return 2
}

func (e *PreOperatorExpression) IsPreOperatorExpression() bool {
	return true
}

func (e *PreOperatorExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitPreOperatorExpression(e)
}

func (e *PreOperatorExpression) String() string {
	return fmt.Sprintf("PreOperatorExpression{%s %s}", e.operator, e.expression)
}
