package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewBinaryOperatorExpression(lineNumber int, typ _type.IType, leftExpression Expression, operator string, rightExpression Expression, priority int) *BinaryOperatorExpression {
	return &BinaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		leftExpression:                   leftExpression,
		operator:                         operator,
		rightExpression:                  rightExpression,
		priority:                         priority,
	}
}

type BinaryOperatorExpression struct {
	AbstractLineNumberTypeExpression

	leftExpression  Expression
	operator        string
	rightExpression Expression
	priority        int
}

func (e *BinaryOperatorExpression) LeftExpression() Expression {
	return e.leftExpression
}

func (e *BinaryOperatorExpression) Operator() string {
	return e.operator
}

func (e *BinaryOperatorExpression) RightExpression() Expression {
	return e.rightExpression
}

func (e *BinaryOperatorExpression) Priority() int {
	return e.priority
}

func (e *BinaryOperatorExpression) SetLeftExpression(leftExpression Expression) {
	e.leftExpression = leftExpression
}

func (e *BinaryOperatorExpression) SetOperator(operator string) {
	e.operator = operator
}

func (e *BinaryOperatorExpression) SetRightExpression(rightExpression Expression) {
	e.rightExpression = rightExpression
}

func (e *BinaryOperatorExpression) SetPriority(priority int) {
	e.priority = priority
}

func (e *BinaryOperatorExpression) IsBinaryOperatorExpression() bool {
	return true
}

func (e *BinaryOperatorExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitBinaryOperatorExpression(e)
}

func (e *BinaryOperatorExpression) String() string {
	return fmt.Sprintf("BinaryOperatorExpression{%s %s %s}", e.leftExpression, e.operator, e.rightExpression)
}
