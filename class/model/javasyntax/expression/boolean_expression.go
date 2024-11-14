package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewBooleanExpression(value bool) *BooleanExpression {
	return &BooleanExpression{
		value: value,
	}
}

func NewBooleanExpressionWithLineNumber(lineNumber int, value bool) *BooleanExpression {
	return &BooleanExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		value: value,
	}
}

type BooleanExpression struct {
	AbstractLineNumberExpression

	value bool
}

func (e *BooleanExpression) Type() _type.IType {
	return _type.PtTypeBoolean
}

func (e *BooleanExpression) IsTrue() bool {
	return e.value
}

func (e *BooleanExpression) IsFalse() bool {
	return !e.value
}

func (e *BooleanExpression) IsBooleanExpression() bool {
	return true
}

func (e *BooleanExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitBooleanExpression(e)
}

func (e *BooleanExpression) String() string {
	value := "false"
	if e.value {
		value = "true"
	}
	return fmt.Sprintf("BooleanExpression{%s}", value)
}
