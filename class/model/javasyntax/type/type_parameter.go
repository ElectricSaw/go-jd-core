package _type

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewTypeParameter(identifier string) intmod.ITypeParameter {
	return &TypeParameter{identifier: identifier}
}

type TypeParameter struct {
	AbstractType
	AbstractTypeParameter

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
