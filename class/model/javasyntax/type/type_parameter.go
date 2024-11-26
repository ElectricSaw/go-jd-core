package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeParameter(identifier string) intmod.ITypeParameter {
	p := &TypeParameter{identifier: identifier}
	p.SetValue(p)
	return p
}

type TypeParameter struct {
	AbstractType
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
