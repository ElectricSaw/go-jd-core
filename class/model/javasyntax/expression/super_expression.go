package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewSuperExpression(typ intsyn.IType) intsyn.ISuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		typ:                          typ,
	}
}

func NewSuperExpressionWithAll(lineNumber int, typ intsyn.IType) intsyn.ISuperExpression {
	return &SuperExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
}

type SuperExpression struct {
	AbstractLineNumberExpression

	typ intsyn.IType
}

func (e *SuperExpression) Type() intsyn.IType {
	return e.typ
}

func (e *SuperExpression) IsSuperExpression() bool {
	return true
}

func (e *SuperExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitSuperExpression(e)
}

func (e *SuperExpression) String() string {
	return fmt.Sprintf("SuperExpression{%s}", e.typ)
}
