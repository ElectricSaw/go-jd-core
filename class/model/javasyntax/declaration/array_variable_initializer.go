package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
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

func (i *ArrayVariableInitializer) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}
