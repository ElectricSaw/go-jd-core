package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewPreOperatorExpression(operator string, expression intmod.IExpression) intmod.IPreOperatorExpression {
	return NewPreOperatorExpressionWithAll(0, operator, expression)
}

func NewPreOperatorExpressionWithAll(lineNumber int, operator string, expression intmod.IExpression) intmod.IPreOperatorExpression {
	e := &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
	e.SetValue(e)
	return e
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
