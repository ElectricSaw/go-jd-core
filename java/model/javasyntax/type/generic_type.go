package _type

import "fmt"

func NewGenericType(name string, dimension int) *GenericType {
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

func (t *GenericType) GetName() string {
	return t.name
}

func (t *GenericType) GetDescriptor() string {
	return t.name
}

func (t *GenericType) GetDimension() int {
	return t.dimension
}

func (t *GenericType) CreateType(dimension int) IType {
	if t.dimension == dimension {
		return t
	} else {
		return NewGenericType(t.name, dimension)
	}
}

func (t *GenericType) IsGenericType() bool {
	return true
}

func (t *GenericType) AcceptTypeVisitor(visitor TypeVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) IsTypeArgumentAssignableFrom(_ map[string]IType, typeArgument ITypeArgument) bool {
	if o, ok := typeArgument.(*GenericType); ok {
		return t.Equals(o)
	}

	return false
}

func (t *GenericType) IsGenericTypeArgument() bool {
	return true
}

func (t *GenericType) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) Equals(o *GenericType) bool {
	if o == nil {
		return false
	}

	if o == t {
		return true
	}

	if t.dimension != o.dimension {
		return false
	}
	if t.name != o.name {
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
