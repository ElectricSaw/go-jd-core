package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewLengthExpression(expression intmod.IExpression) intmod.ILengthExpression {
	return NewLengthExpressionWithAll(0, expression)
}

func NewLengthExpressionWithAll(lineNumber int, expression intmod.IExpression) intmod.ILengthExpression {
	e := &LengthExpression{
		LineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		expression:           expression,
	}
	e.SetValue(e)
	return e
}

type LengthExpression struct {
	LineNumberExpression

	expression intmod.IExpression
}

func (e *LengthExpression) Type() intmod.IType {
	return model.PtTypeInt.(intmod.IType)
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
