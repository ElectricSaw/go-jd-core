package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypes() intmod.ITypes {
	return NewTypesWithSlice(make([]intmod.IType, 0))
}

func NewTypesWithSlice(types []intmod.IType) intmod.ITypes {
	return &Types{
		DefaultList: *util.NewDefaultListWithSlice[intmod.IType](types).(*util.DefaultList[intmod.IType]),
	}
}

type Types struct {
	AbstractTypeArgument
	util.DefaultList[intmod.IType]
}

func (t *Types) Name() string {
	return ""
}

func (t *Types) Descriptor() string {
	return ""
}

func (t *Types) Dimension() int {
	return -1
}

func (t *Types) CreateType(dimension int) intmod.IType {
	return nil
}

func (t *Types) IsGenericType() bool {
	return false
}

func (t *Types) IsInnerObjectType() bool {
	return false
}

func (t *Types) IsObjectType() bool {
	return false
}

func (t *Types) IsPrimitiveType() bool {
	return false
}

func (t *Types) IsTypes() bool {
	return true
}

func (t *Types) OuterType() intmod.IObjectType {
	return OtTypeUndefinedObject
}

func (t *Types) InternalName() string {
	return ""
}

func (t *Types) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
	visitor.VisitTypes(t)
}
