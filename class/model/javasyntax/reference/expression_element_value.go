package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewExpressionElementValue(expression intsyn.IExpression) intsyn.IExpressionElementValue {
	return &ExpressionElementValue{
		expression: expression,
	}
}

type ExpressionElementValue struct {
	expression intsyn.IExpression
}

func (e *ExpressionElementValue) Expression() intsyn.IExpression {
	return e.expression
}

func (e *ExpressionElementValue) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *ExpressionElementValue) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitExpressionElementValue(e)
}

func (e *ExpressionElementValue) String() string {
	return fmt.Sprintf("ExpressionElementValue{%s}", *e)
}
