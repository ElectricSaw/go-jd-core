package declaration

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewArrayVariableInitializer(typ _type.IType) *ArrayVariableInitializer {
	return &ArrayVariableInitializer{
		typ: typ,
	}
}

type ArrayVariableInitializer struct {
	AbstractVariableInitializer

	typ                  _type.IType
	VariableInitializers []VariableInitializer
}

func (i *ArrayVariableInitializer) GetType() _type.IType {
	return i.typ
}

func (i *ArrayVariableInitializer) GetLineNumber() int {
	if len(i.VariableInitializers) == 0 {
		return 0
	}
	return i.VariableInitializers[0].GetLineNumber()
}

func (i *ArrayVariableInitializer) Accept(visitor DeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}
