package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewSuperExpression(typ intmod.IType) intmod.ISuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		typ:                          typ,
	}
}

func NewSuperExpressionWithAll(lineNumber int, typ intmod.IType) intmod.ISuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
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
