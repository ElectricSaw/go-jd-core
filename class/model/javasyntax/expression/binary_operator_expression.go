package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewBinaryOperatorExpression(lineNumber int, typ intsyn.IType, leftExpression intsyn.IExpression,
	operator string, rightExpression intsyn.IExpression, priority int) intsyn.IBinaryOperatorExpression {
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

	leftExpression  intsyn.IExpression
	operator        string
	rightExpression intsyn.IExpression
	priority        int
}

func (e *BinaryOperatorExpression) LeftExpression() intsyn.IExpression {
	return e.leftExpression
}

func (e *BinaryOperatorExpression) Operator() string {
	return e.operator
}

func (e *BinaryOperatorExpression) RightExpression() intsyn.IExpression {
	return e.rightExpression
}

func (e *BinaryOperatorExpression) Priority() int {
	return e.priority
}

func (e *BinaryOperatorExpression) SetLeftExpression(leftExpression intsyn.IExpression) {
	e.leftExpression = leftExpression
}

func (e *BinaryOperatorExpression) SetOperator(operator string) {
	e.operator = operator
}

func (e *BinaryOperatorExpression) SetRightExpression(rightExpression intsyn.IExpression) {
	e.rightExpression = rightExpression
}

func (e *BinaryOperatorExpression) SetPriority(priority int) {
	e.priority = priority
}

func (e *BinaryOperatorExpression) IsBinaryOperatorExpression() bool {
	return true
}

func (e *BinaryOperatorExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitBinaryOperatorExpression(e)
}

func (e *BinaryOperatorExpression) String() string {
	return fmt.Sprintf("BinaryOperatorExpression{%s %s %s}", e.leftExpression, e.operator, e.rightExpression)
}
