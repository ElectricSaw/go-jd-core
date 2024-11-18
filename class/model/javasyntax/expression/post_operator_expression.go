package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewPostOperatorExpression(operator string, expression intsyn.IExpression) intsyn.IPostOperatorExpression {
	return &PostOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		operator:                     operator,
		expression:                   expression,
	}
}

func NewPostOperatorExpressionWithAll(lineNumber int, operator string, expression intsyn.IExpression) intsyn.IPostOperatorExpression {
	return &PostOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
}

type PostOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression intsyn.IExpression
}

func (e *PostOperatorExpression) Operator() string {
	return e.operator
}

func (e *PostOperatorExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *PostOperatorExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *PostOperatorExpression) Type() intsyn.IType {
	return e.expression.Type()
}

func (e *PostOperatorExpression) Priority() int {
	return 1
}

func (e *PostOperatorExpression) IsPostOperatorExpression() bool {
	return true
}

func (e *PostOperatorExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitPostOperatorExpression(e)
}

func (e *PostOperatorExpression) String() string {
	return fmt.Sprintf("PostOperatorExpression{%s %s}", e.expression, e.operator)
}
