package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewPreOperatorExpression(operator string, expression intsyn.IExpression) intsyn.IPreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		operator:                     operator,
		expression:                   expression,
	}
}

func NewPreOperatorExpressionWithAll(lineNumber int, operator string, expression intsyn.IExpression) intsyn.IPreOperatorExpression {
	return &PreOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
}

type PreOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression intsyn.IExpression
}

func (e *PreOperatorExpression) Operator() string {
	return e.operator
}

func (e *PreOperatorExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *PreOperatorExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *PreOperatorExpression) Type() intsyn.IType {
	return e.expression.Type()
}

func (e *PreOperatorExpression) Priority() int {
	return 2
}

func (e *PreOperatorExpression) IsPreOperatorExpression() bool {
	return true
}

func (e *PreOperatorExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitPreOperatorExpression(e)
}

func (e *PreOperatorExpression) String() string {
	return fmt.Sprintf("PreOperatorExpression{%s %s}", e.operator, e.expression)
}
