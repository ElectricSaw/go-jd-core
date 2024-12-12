package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewNullExpression(typ intmod.IType) intmod.INullExpression {
	return NewNullExpressionWithAll(0, typ)
}

func NewNullExpressionWithAll(lineNumber int, typ intmod.IType) intmod.INullExpression {
	e := &NullExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
	}
	e.SetValue(e)
	return e
}

type NullExpression struct {
	AbstractLineNumberTypeExpression
}

func (e *NullExpression) IsNullExpression() bool {
	return true
}

func (e *NullExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitNullExpression(e)
}

func (e *NullExpression) String() string {
	return fmt.Sprintf("NullExpression{type=%s}", e.typ)
}
