package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewSuperExpression(typ _type.IType) *SuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		typ:                          typ,
	}
}

func NewSuperExpressionWithAll(lineNumber int, typ _type.IType) *SuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
}

type SuperExpression struct {
	AbstractLineNumberExpression

	typ _type.IType
}

func (e *SuperExpression) GetType() _type.IType {
	return e.typ
}

func (e *SuperExpression) IsSuperExpression() bool {
	return true
}

func (e *SuperExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitSuperExpression(e)
}

func (e *SuperExpression) String() string {
	return fmt.Sprintf("SuperExpression{%s}", e.typ)
}
