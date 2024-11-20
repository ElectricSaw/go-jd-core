package _type

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewTypeParameterWithTypeBounds(identifier string, typeBounds intsyn.IType) intsyn.ITypeParameterWithTypeBounds {
	return &TypeParameterWithTypeBounds{
		TypeParameter: TypeParameter{
			identifier: identifier,
		},
		typeBounds: typeBounds,
	}
}

type TypeParameterWithTypeBounds struct {
	TypeParameter

	typeBounds intsyn.IType
}

func (t *TypeParameterWithTypeBounds) TypeBounds() intsyn.IType {
	return t.typeBounds
}

func (t *TypeParameterWithTypeBounds) AcceptTypeParameterVisitor(visitor intsyn.ITypeParameterVisitor) {
	visitor.VisitTypeParameterWithTypeBounds(t)
}

func (t *TypeParameterWithTypeBounds) String() string {
	return fmt.Sprintf("TypeParameter{ identifier=%s, typeBounds=%s }", t.identifier, t.TypeBounds())
}
