package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewArrayVariableInitializer(typ Type) ArrayVariableInitializer {
	return &ArrayVariableInitializer{
		typ: typ,
	}
}

type ArrayVariableInitializer struct {
	AbstractVariableInitializer
	util.DefaultList[VariableInitializer]

	typ Type
}

func (i *ArrayVariableInitializer) Type() Type {
	return i.typ
}

func (i *ArrayVariableInitializer) LineNumber() int {
	if i.Size() == 0 {
		return 0
	}
	return i.Get(0).LineNumber()
}

func (i *ArrayVariableInitializer) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitArrayVariableInitializer(i)
}
