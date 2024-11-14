package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewAbstractLineNumberTypeExpression(typ _type.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		typ: typ,
	}
}

func NewAbstractLineNumberTypeExpressionWithAll(lineNumber int, typ _type.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
}

type AbstractLineNumberTypeExpression struct {
	AbstractLineNumberExpression

	typ _type.IType
}

func (e *AbstractLineNumberTypeExpression) Type() _type.IType {
	return e.typ
}

func (e *AbstractLineNumberTypeExpression) SetType(typ _type.IType) {
	e.typ = typ
}
