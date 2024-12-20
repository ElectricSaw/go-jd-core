package _type

import (
	"bytes"
	"encoding/gob"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"hash/fnv"
)

type AbstractType struct {
	AbstractTypeArgument
	util.DefaultBase[intmod.IType]
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

func (t *AbstractType) CreateType(dimension int) intmod.IType {
	return nil
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

func (t *AbstractType) OuterType() intmod.IObjectType {
	return OtTypeUndefinedObject
}

func (t *AbstractType) InternalName() string {
	return ""
}

func (t *AbstractType) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
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
	byteArray := toBytes(data)
	if byteArray == nil {
		return -1
	}

	h := fnv.New32a()
	_, err := h.Write(byteArray)
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
