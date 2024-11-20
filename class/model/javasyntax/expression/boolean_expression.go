package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewBooleanExpression(value bool) intsyn.IBooleanExpression {
	return &BooleanExpression{
		value: value,
	}
}

func NewBooleanExpressionWithLineNumber(lineNumber int, value bool) intsyn.IBooleanExpression {
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

func (e *BooleanExpression) Type() intsyn.IType {
	return _type.PtTypeBoolean.(intsyn.IType)
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

func (e *BooleanExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitBooleanExpression(e)
}

func (e *BooleanExpression) String() string {
	value := "false"
	if e.value {
		value = "true"
	}
	return fmt.Sprintf("BooleanExpression{%s}", value)
}
