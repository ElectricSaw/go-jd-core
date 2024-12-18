package _type

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
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
