package _type

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

type IType interface {
	TypeVisitable
	
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType
	Size() int

	IsGenericType() bool
	IsInnerObjectType() bool
	IsObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	OuterType() IObjectType
	InternalName() string

	///////////////////////////////////////////////////////////////

	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentList() bool
	TypeArgumentFirst() ITypeArgument  // ITypeArgument
	TypeArgumentList() []ITypeArgument // ITypeArgument
	TypeArgumentSize() int
	IsGenericTypeArgument() bool
	IsInnerObjectTypeArgument() bool
	IsObjectTypeArgument() bool
	IsPrimitiveTypeArgument() bool
	IsWildcardExtendsTypeArgument() bool
	IsWildcardSuperTypeArgument() bool
	IsWildcardTypeArgument() bool
	Type() IType

	///////////////////////////////////////////////////////////////

	HashCode() int
}

type TypeVisitor interface {
	VisitPrimitiveType(y *PrimitiveType)
	VisitObjectType(y *ObjectType)
	VisitInnerObjectType(y *InnerObjectType)
	VisitTypes(types *Types)
	VisitGenericType(y *GenericType)
}

type TypeVisitable interface {
	AcceptTypeVisitor(visitor TypeVisitor)
}

type IObjectType interface {
	Dimension() int
	CreateType(dimension int) IType
	QualifiedName() string
	OuterType() IObjectType
	InternalName() string
	TypeArguments() ITypeArgument
	CreateTypeWithArgs(typeArguments ITypeArgument) IObjectType
	AcceptTypeVisitor(visitor TypeVisitor)
	AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor)
	HashCode() int
}

type AbstractType struct {
	AbstractTypeArgument
}

func (t *AbstractType) Name() string {
	return ""
}

func (t *AbstractType) Descriptor() string {
	return ""
}

func (t *AbstractType) Dimension() int {
	return -1
}

func (t *AbstractType) CreateType(dimension int) IType {
	return nil
}

func (t *AbstractType) Size() int {
	return 1
}

func (t *AbstractType) IsGenericType() bool {
	return false
}

func (t *AbstractType) IsInnerObjectType() bool {
	return false
}

func (t *AbstractType) IsObjectType() bool {
	return false
}

func (t *AbstractType) IsPrimitiveType() bool {
	return false
}

func (t *AbstractType) IsTypes() bool {
	return false
}

func (t *AbstractType) OuterType() IObjectType {
	return OtTypeUndefinedObject
}

func (t *AbstractType) InternalName() string {
	return ""
}

func (t *AbstractType) AcceptTypeVisitor(visitor TypeVisitor) {
}

func hashCodeWithString(str string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return -1
	}
	return int(h.Sum32())
}

func hashCodeWithStruct(data any) int {
	bytes := toBytes(data)
	if bytes == nil {
		return -1
	}

	h := fnv.New32a()
	_, err := h.Write(bytes)
	if err != nil {
		return -1
	}

	return int(h.Sum32())
}

func toBytes(data any) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil
	}
	return buf.Bytes()
}
