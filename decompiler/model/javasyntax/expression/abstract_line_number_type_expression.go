package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewAbstractLineNumberTypeExpression(typ intmod.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		typ: typ,
	}
}

func NewAbstractLineNumberTypeExpressionWithAll(lineNumber int, typ intmod.IType) *AbstractLineNumberTypeExpression {
	return &AbstractLineNumberTypeExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		typ:                          typ,
	}
}

type AbstractLineNumberTypeExpression struct {
	AbstractLineNumberExpression

	typ intmod.IType
}

func (e *AbstractLineNumberTypeExpression) Type() intmod.IType {
	return e.typ
}

func (e *AbstractLineNumberTypeExpression) SetType(typ intmod.IType) {
	e.typ = typ
}
