package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypeArguments() intmod.ITypeArguments {
	return NewTypeArgumentsWithCapacity(0)
}

func NewTypeArgumentsWithCapacity(capacity int) intmod.ITypeArguments {
	return &TypeArguments{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.ITypeArgument](capacity).(*util.DefaultList[intmod.ITypeArgument]),
	}
}

type TypeArguments struct {
	AbstractType
	AbstractTypeArgument
	util.DefaultList[intmod.ITypeArgument]
}

func (t *TypeArguments) IsList() bool {
	return t.DefaultList.IsList()
}

func (t *TypeArguments) Size() int {
	return t.DefaultList.Size()
}

func (t *TypeArguments) ToSlice() []intmod.ITypeArgument {
	return t.DefaultList.ToSlice()
}

func (t *TypeArguments) ToList() *util.DefaultList[intmod.ITypeArgument] {
	return t.DefaultList.ToList()
}

func (t *TypeArguments) First() intmod.ITypeArgument {
	return t.DefaultList.First()
}

func (t *TypeArguments) Last() intmod.ITypeArgument {
	return t.DefaultList.Last()
}

func (t *TypeArguments) Iterator() util.IIterator[intmod.ITypeArgument] {
	return t.DefaultList.Iterator()
}

func (t *TypeArguments) IsTypeArgumentAssignableFrom(typeBounds map[string]intmod.IType, typeArgument intmod.ITypeArgument) bool {
	ata, ok := typeArgument.(*TypeArguments)
	if !ok {
		return false
	}

	if t.Size() != ata.Size() {
		return false
	}

	for i := 0; i < t.DefaultList.Size(); i++ {

		if !t.Get(i).IsTypeArgumentAssignableFrom(typeBounds, ata.Get(i)) {
			return false
		}
	}

	return true
}

func (t *TypeArguments) IsTypeArgumentList() bool {
	return true
}

func (t *TypeArguments) TypeArgumentFirst() intmod.ITypeArgument {
	return t.Get(0)
}

func (t *TypeArguments) TypeArgumentList() []intmod.ITypeArgument {
	return t.ToSlice()
}

func (t *TypeArguments) TypeArgumentSize() int {
	return t.DefaultList.Size()
}

func (t *TypeArguments) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitTypeArguments(t)
}
