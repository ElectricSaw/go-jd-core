package _type

import "fmt"

func NewTypeParameterWithTypeBounds(identifier string, typeBounds IType) *TypeParameterWithTypeBounds {
	return &TypeParameterWithTypeBounds{
		TypeParameter: TypeParameter{
			identifier: identifier,
		},
		typeBounds: typeBounds,
	}
}

type TypeParameterWithTypeBounds struct {
	TypeParameter

	typeBounds IType
}

func (t *TypeParameterWithTypeBounds) TypeBounds() IType {
	return t.typeBounds
}

func (t *TypeParameterWithTypeBounds) AcceptTypeParameterVisitor(visitor TypeParameterVisitor) {
	visitor.VisitTypeParameterWithTypeBounds(t)
}

func (t *TypeParameterWithTypeBounds) String() string {
	return fmt.Sprintf("TypeParameter{ identifier=%s, typeBounds=%s }", t.identifier, t.TypeBounds())
}
