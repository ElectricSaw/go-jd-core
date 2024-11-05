package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewPreOperatorExpression(operator string, expression Expression) *PreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		operator:                     operator,
		expression:                   expression,
	}
}

func NewPreOperatorExpressionWithAll(lineNumber int, operator string, expression Expression) *PreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
}

type PreOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression Expression
}

func (e *PreOperatorExpression) GetOperator() string {
	return e.operator
}

func (e *PreOperatorExpression) GetExpression() Expression {
	return e.expression
}

func (e *PreOperatorExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *PreOperatorExpression) GetType() _type.IType {
	return e.expression.GetType()
}

func (e *PreOperatorExpression) GetPriority() int {
	return 2
}

func (e *PreOperatorExpression) IsPreOperatorExpression() bool {
	return true
}

func (e *PreOperatorExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitPreOperatorExpression(e)
}

func (e *PreOperatorExpression) String() string {
	return fmt.Sprintf("PreOperatorExpression{%s %s}", e.operator, e.expression)
}
