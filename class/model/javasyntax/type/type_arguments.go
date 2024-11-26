package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeArguments() intmod.ITypeArguments {
	return NewTypeArgumentsWithSize(0)
}

func NewTypeArgumentsWithSize(size int) intmod.ITypeArguments {
	return &TypeArguments{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.ITypeArgument](size),
	}
}

type TypeArguments struct {
	AbstractType
	AbstractTypeArgument
	util.DefaultList[intmod.ITypeArgument]
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
