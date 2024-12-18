package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypeParameters() intmod.ITypeParameters {
	return &TypeParameters{}
}

type TypeParameters struct {
	util.DefaultList[intmod.ITypeParameter]
}

func (t *TypeParameters) Identifier() string {
	return ""
}

func (t *TypeParameters) AcceptTypeParameterVisitor(visitor intmod.ITypeParameterVisitor) {
	visitor.VisitTypeParameters(t)
}

func (t *TypeParameters) String() string {
	return ""
}
