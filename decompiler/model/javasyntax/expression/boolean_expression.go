package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

var True = NewBooleanExpression(true)
var False = NewBooleanExpression(false)

func NewBooleanExpression(value bool) intmod.IBooleanExpression {
	return NewBooleanExpressionWithLineNumber(0, value)
}

func NewBooleanExpressionWithLineNumber(lineNumber int, value bool) intmod.IBooleanExpression {
	e := &BooleanExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		value: value,
	}
	e.SetValue(e)
	return e
}

type BooleanExpression struct {
	AbstractLineNumberExpression

	value bool
}

func (e *BooleanExpression) Type() intmod.IType {
	return _type.PtTypeBoolean.(intmod.IType)
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

func (e *BooleanExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitBooleanExpression(e)
}

func (e *BooleanExpression) String() string {
	value := "false"
	if e.value {
		value = "true"
	}
	return fmt.Sprintf("BooleanExpression{%s}", value)
}
