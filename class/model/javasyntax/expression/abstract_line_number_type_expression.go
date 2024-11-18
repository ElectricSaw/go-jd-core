package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewAbstractLineNumberTypeExpression(typ intsyn.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		typ: typ,
	}
}

func NewAbstractLineNumberTypeExpressionWithAll(lineNumber int, typ intsyn.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
}

type AbstractLineNumberTypeExpression struct {
	AbstractLineNumberExpression

	typ intsyn.IType
}

func (e *AbstractLineNumberTypeExpression) Type() intsyn.IType {
	return e.typ
}

func (e *AbstractLineNumberTypeExpression) SetType(typ intsyn.IType) {
	e.typ = typ
}
