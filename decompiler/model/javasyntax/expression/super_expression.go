package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewSuperExpression(typ intmod.IType) intmod.ISuperExpression {
	return NewSuperExpressionWithAll(0, typ)
}

func NewSuperExpressionWithAll(lineNumber int, typ intmod.IType) intmod.ISuperExpression {
	e := &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
	e.SetValue(e)
	return e
}

type SuperExpression struct {
	AbstractLineNumberExpression

	typ intmod.IType
}

func (e *SuperExpression) Type() intmod.IType {
	return e.typ
}

func (e *SuperExpression) IsSuperExpression() bool {
	return true
}

func (e *SuperExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitSuperExpression(e)
}

func (e *SuperExpression) String() string {
	return fmt.Sprintf("SuperExpression{%s}", e.typ)
}
