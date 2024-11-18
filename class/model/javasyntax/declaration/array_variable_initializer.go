package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewArrayVariableInitializer(typ _type.IType) intsyn.IArrayVariableInitializer {
	return &ArrayVariableInitializer{
		typ: typ,
	}
}

type ArrayVariableInitializer struct {
	AbstractVariableInitializer
	util.DefaultList[intsyn.IVariableInitializer]

	typ _type.IType
}

func (i *ArrayVariableInitializer) Type() _type.IType {
	return i.typ
}

func (i *ArrayVariableInitializer) LineNumber() int {
	if i.Size() == 0 {
		return 0
	}
	return i.Get(0).LineNumber()
}

func (i *ArrayVariableInitializer) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}
