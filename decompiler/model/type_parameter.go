package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypeParameter(identifier string) TypeParameter {
	p := TypeParameter{identifier: identifier}
	p.SetValue(&p)
	return p
}

func NewTypeParameterWithTypeBounds(identifier string, typeBounds IType) TypeParameterWithTypeBounds {
	return TypeParameterWithTypeBounds{
		identifier: identifier,
		TypeBounds: typeBounds,
	}
}

func NewTypeParameters() TypeParameters {
	return TypeParameters{}
}

type ITypeParameter interface {
	Identifier() string
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
	Equals(o interface{}) bool
	HashCode() int
	String() string
}

type ITypeParameterVisitable interface {
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
}

type ITypeParameterVisitor interface {
	VisitTypeParameter(parameter *TypeParameter)
	VisitTypeParameterWithTypeBounds(parameter *TypeParameterWithTypeBounds)
	VisitTypeParameters(parameters *TypeParameters)
}

type TypeParameter struct {
	util.DefaultBase[ITypeParameter]

	identifier string
}

func (t *TypeParameter) Identifier() string {
	return t.identifier
}

func (t *TypeParameter) AcceptTypeParameterVisitor(visitor ITypeParameterVisitor) {
	visitor.VisitTypeParameter(t)
}

func (t *TypeParameter) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *TypeParameter

	switch o := o.(type) {
	case TypeParameter:
		other = &o
	case *TypeParameter:
		other = o
	default:
		return false
	}

	return t.identifier == other.identifier
}

func (t *TypeParameter) HashCode() int {
	return hashCodeWithStruct(t)
}

func (t *TypeParameter) String() string {
	return "TypeParameter{ identifier=" + t.identifier + " }"
}

type TypeParameterWithTypeBounds struct {
	identifier string
	TypeBounds IType
}

func (t *TypeParameterWithTypeBounds) Identifier() string {
	return t.identifier
}

func (t *TypeParameterWithTypeBounds) AcceptTypeParameterVisitor(visitor ITypeParameterVisitor) {
	visitor.VisitTypeParameterWithTypeBounds(t)
}

func (t *TypeParameterWithTypeBounds) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *TypeParameter

	switch o := o.(type) {
	case TypeParameter:
		other = &o
	case *TypeParameter:
		other = o
	default:
		return false
	}

	return t.identifier == other.identifier
}

func (t *TypeParameterWithTypeBounds) HashCode() int {
	return hashCodeWithStruct(t)
}

func (t *TypeParameterWithTypeBounds) String() string {
	return fmt.Sprintf("TypeParameter{ identifier=%s, typeBounds=%s }", t.identifier, t.TypeBounds)
}

type TypeParameters struct {
	util.DefaultList[ITypeParameter]
}

func (t *TypeParameters) Identifier() string {
	return ""
}

func (t *TypeParameters) AcceptTypeParameterVisitor(visitor ITypeParameterVisitor) {
	visitor.VisitTypeParameters(t)
}

func (t *TypeParameters) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *TypeParameters

	switch o := o.(type) {
	case TypeParameters:
		other = &o
	case *TypeParameters:
		other = o
	default:
		return false
	}

	size := t.Size()
	equals := size == other.Size()

	if equals {
		es := t.ToSlice()
		otherEs := other.ToSlice()
		for i := 0; i < size; i++ {
			if es[i] != otherEs[i] {
				return false
			}
		}
	}

	return equals
}

func (t *TypeParameters) HashCode() int {
	return hashCodeWithStruct(t)
}

func (t *TypeParameters) String() string {
	sb := "TypeParameters{"
	sb += t.Get(0).String()
	for i := 1; i < t.Size(); i++ {
		sb += " & "
		sb += t.Get(i).String()
	}
	sb += "}"
	return sb
}
