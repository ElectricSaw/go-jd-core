package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewLengthExpression(expression intmod.IExpression) intmod.ILengthExpression {
	return NewLengthExpressionWithAll(0, expression)
}

func NewLengthExpressionWithAll(lineNumber int, expression intmod.IExpression) intmod.ILengthExpression {
	e := &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		expression:                   expression,
	}
	e.SetValue(e)
	return e
}

type LengthExpression struct {
	AbstractLineNumberExpression

	expression intmod.IExpression
}

func (e *LengthExpression) Type() intmod.IType {
	return _type.PtTypeInt.(intmod.IType)
}

func (e *LengthExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *LengthExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *LengthExpression) IsLengthExpression() bool {
	return true
}

func (e *LengthExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLengthExpression(e)
}

func (e *LengthExpression) String() string {
	return fmt.Sprintf("LengthExpression{%s}", e.expression)
}
