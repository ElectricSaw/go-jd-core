package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewLengthExpression(expression Expression) *LengthExpression {
	return &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		expression:                   expression,
	}
}

func NewLengthExpressionWithAll(lineNumber int, expression Expression) *LengthExpression {
	return &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		expression:                   expression,
	}
}

type LengthExpression struct {
	AbstractLineNumberExpression

	expression Expression
}

func (e *LengthExpression) GetType() _type.IType {
	return _type.PtTypeInt
}

func (e *LengthExpression) GetExpression() Expression {
	return e.expression
}

func (e *LengthExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *LengthExpression) IsLengthExpression() bool {
	return true
}

func (e *LengthExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitLengthExpression(e)
}

func (e *LengthExpression) String() string {
	return fmt.Sprintf("LengthExpression{%s}", e.expression)
}
