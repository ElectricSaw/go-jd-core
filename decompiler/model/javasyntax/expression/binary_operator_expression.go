package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewBinaryOperatorExpression(lineNumber int, typ intmod.IType, leftExpression intmod.IExpression,
	operator string, rightExpression intmod.IExpression, priority int) intmod.IBinaryOperatorExpression {
	e := &BinaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		leftExpression:                   leftExpression,
		operator:                         operator,
		rightExpression:                  rightExpression,
		priority:                         priority,
	}
	e.SetValue(e)
	return e
}

type BinaryOperatorExpression struct {
	AbstractLineNumberTypeExpression

	leftExpression  intmod.IExpression
	operator        string
	rightExpression intmod.IExpression
	priority        int
}

func (e *BinaryOperatorExpression) LeftExpression() intmod.IExpression {
	return e.leftExpression
}

func (e *BinaryOperatorExpression) Operator() string {
	return e.operator
}

func (e *BinaryOperatorExpression) RightExpression() intmod.IExpression {
	return e.rightExpression
}

func (e *BinaryOperatorExpression) Priority() int {
	return e.priority
}

func (e *BinaryOperatorExpression) SetLeftExpression(leftExpression intmod.IExpression) {
	e.leftExpression = leftExpression
}

func (e *BinaryOperatorExpression) SetOperator(operator string) {
	e.operator = operator
}

func (e *BinaryOperatorExpression) SetRightExpression(rightExpression intmod.IExpression) {
	e.rightExpression = rightExpression
}

func (e *BinaryOperatorExpression) SetPriority(priority int) {
	e.priority = priority
}

func (e *BinaryOperatorExpression) IsBinaryOperatorExpression() bool {
	return true
}

func (e *BinaryOperatorExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitBinaryOperatorExpression(e)
}

func (e *BinaryOperatorExpression) String() string {
	return fmt.Sprintf("BinaryOperatorExpression{%s %s %s}", e.leftExpression, e.operator, e.rightExpression)
}
