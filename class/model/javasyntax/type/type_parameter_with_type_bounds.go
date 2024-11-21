package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewTypeParameterWithTypeBounds(identifier string, typeBounds intmod.IType) intmod.ITypeParameterWithTypeBounds {
	return &TypeParameterWithTypeBounds{
		TypeParameter: TypeParameter{
			identifier: identifier,
		},
		typeBounds: typeBounds,
	}
}

type TypeParameterWithTypeBounds struct {
	TypeParameter

	typeBounds intmod.IType
}

func (t *TypeParameterWithTypeBounds) TypeBounds() intmod.IType {
	return t.typeBounds
}

func (t *TypeParameterWithTypeBounds) AcceptTypeParameterVisitor(visitor intmod.ITypeParameterVisitor) {
	visitor.VisitTypeParameterWithTypeBounds(t)
}

func (t *TypeParameterWithTypeBounds) String() string {
	return fmt.Sprintf("TypeParameter{ identifier=%s, typeBounds=%s }", t.identifier, t.TypeBounds())
}
