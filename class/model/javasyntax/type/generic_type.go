package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewGenericType(name string) intmod.IGenericType {
	return NewGenericTypeWithAll(name, 0)
}

func NewGenericTypeWithAll(name string, dimension int) intmod.IGenericType {
	t := &GenericType{
		name:      name,
		dimension: dimension,
	}
	t.SetValue(t)
	return t
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

func (t *GenericType) CreateType(dimension int) intmod.IType {
	if t.dimension == dimension {
		return t
	} else {
		return NewGenericTypeWithAll(t.name, dimension).(intmod.IType)
	}
}

func (t *GenericType) IsGenericType() bool {
	return true
}

func (t *GenericType) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) IsTypeArgumentAssignableFrom(_ map[string]intmod.IType, typeArgument intmod.ITypeArgument) bool {
	if o, ok := typeArgument.(*GenericType); ok {
		return t.Equals(o)
	}

	return false
}

func (t *GenericType) IsGenericTypeArgument() bool {
	return true
}

func (t *GenericType) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitGenericType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *GenericType) Equals(o intmod.IGenericType) bool {
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
