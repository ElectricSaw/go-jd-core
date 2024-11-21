package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewArrayVariableInitializer(typ intmod.IType) intmod.IArrayVariableInitializer {
	return &ArrayVariableInitializer{
		typ: typ,
	}
}

type ArrayVariableInitializer struct {
	AbstractVariableInitializer
	util.DefaultList[intmod.IVariableInitializer]

	typ intmod.IType
}

func (i *ArrayVariableInitializer) Type() intmod.IType {
	return i.typ
}

func (i *ArrayVariableInitializer) LineNumber() int {
	if i.Size() == 0 {
		return 0
	}
	return i.Get(0).LineNumber()
}

func (i *ArrayVariableInitializer) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}
