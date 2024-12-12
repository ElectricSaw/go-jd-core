package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewTypeParameter(identifier string) intmod.ITypeParameter {
	p := &TypeParameter{identifier: identifier}
	p.SetValue(p)
	return p
}

type TypeParameter struct {
	AbstractTypeParameter
	util.DefaultBase[intmod.ITypeParameter]

	identifier string
}

func (t *TypeParameter) Identifier() string {
	return t.identifier
}

func (t *TypeParameter) AcceptTypeParameterVisitor(visitor intmod.ITypeParameterVisitor) {
	visitor.VisitTypeParameter(t)
}

func (t *TypeParameter) String() string {
	return "TypeParameter{ identifier=" + t.identifier + " }"
}
