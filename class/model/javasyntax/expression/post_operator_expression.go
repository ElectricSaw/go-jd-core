package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewPostOperatorExpression(operator string, expression Expression) *PostOperatorExpression {
	return &PostOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		operator:                     operator,
		expression:                   expression,
	}
}

func NewPostOperatorExpressionWithAll(lineNumber int, operator string, expression Expression) *PostOperatorExpression {
	return &PostOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
}

type PostOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression Expression
}

func (e *PostOperatorExpression) Operator() string {
	return e.operator
}

func (e *PostOperatorExpression) Expression() Expression {
	return e.expression
}

func (e *PostOperatorExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *PostOperatorExpression) Type() _type.IType {
	return e.expression.Type()
}

func (e *PostOperatorExpression) Priority() int {
	return 1
}

func (e *PostOperatorExpression) IsPostOperatorExpression() bool {
	return true
}

func (e *PostOperatorExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitPostOperatorExpression(e)
}

func (e *PostOperatorExpression) String() string {
	return fmt.Sprintf("PostOperatorExpression{%s %s}", e.expression, e.operator)
}
