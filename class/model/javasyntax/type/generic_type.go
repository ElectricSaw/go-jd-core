package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewGenericType(name string, dimension int) intsyn.IGenericType {
	return &GenericType{
		name:      name,
		dimension: dimension,
	}
}

type GenericType struct {
	AbstractType
	AbstractTypeArgument

	name      string
	dimension int
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) HashCode() int {
	result := 991890290 + hashCodeWithString(t.name)
	result = 31*result + t.Dimension()
	return result
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) Name() string {
	return t.name
}

func (t *GenericType) Descriptor() string {
	return t.name
}

func (t *GenericType) Dimension() int {
	return t.dimension
}

func (t *GenericType) CreateType(dimension int) intsyn.IType {
	if t.dimension == dimension {
		return t
	} else {
		return NewGenericType(t.name, dimension).(intsyn.IType)
	}
}

func (t *GenericType) IsGenericType() bool {
	return true
}

func (t *GenericType) AcceptTypeVisitor(visitor intsyn.ITypeVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) IsTypeArgumentAssignableFrom(_ map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	if o, ok := typeArgument.(*GenericType); ok {
		return t.Equals(o)
	}

	return false
}

func (t *GenericType) IsGenericTypeArgument() bool {
	return true
}

func (t *GenericType) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) Equals(o intsyn.IGenericType) bool {
	if o == nil {
		return false
	}

	if o == t {
		return true
	}

	if t.dimension != o.Dimension() {
		return false
	}
	if t.name != o.Name() {
		return false
	}
	return true
}

func (t *GenericType) String() string {
	msg := fmt.Sprintf("GenericType{ %s", t.name)
	if t.dimension > 0 {
		msg += fmt.Sprintf(", dimension: %d", t.dimension)
	}
	msg += "}"

	return msg
}
