package reference

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"fmt"
)

func NewExpressionElementValue(expression expression.Expression) *ExpressionElementValue {
	return &ExpressionElementValue{
		expression: expression,
	}
}

type ExpressionElementValue struct {
	IElementValue

	expression expression.Expression
}

func (e *ExpressionElementValue) GetExpression() expression.Expression {
	return e.expression
}

func (e *ExpressionElementValue) SetExpression(expression expression.Expression) {
	e.expression = expression
}

func (e *ExpressionElementValue) Accept(visitor ReferenceVisitor) {
	visitor.VisitExpressionElementValue(e)
}

func (e *ExpressionElementValue) String() string {
	return fmt.Sprintf("ExpressionElementValue{%s}", *e)
}
