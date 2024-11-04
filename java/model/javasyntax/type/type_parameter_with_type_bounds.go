package _type

import "fmt"

func NewTypeParameterWithTypeBounds(identifier string, typeBounds AbstractType) *TypeParameterWithTypeBounds {
	return &TypeParameterWithTypeBounds{
		TypeParameter: TypeParameter{
			identifier: identifier,
		},
		typeBounds: typeBounds,
	}
}

type TypeParameterWithTypeBounds struct {
	TypeParameter

	typeBounds AbstractType
}

func (t *TypeParameterWithTypeBounds) TypeBounds() AbstractType {
	return t.typeBounds
}

func (t *TypeParameterWithTypeBounds) AcceptTypeParameterVisitor(visitor TypeParameterVisitor) {
	visitor.VisitTypeParameterWithTypeBounds(t)
}

func (t *TypeParameterWithTypeBounds) String() string {
	return fmt.Sprintf("TypeParameter{ identifier=%s, typeBounds=%s }", t.identifier, t.TypeBounds())
}
