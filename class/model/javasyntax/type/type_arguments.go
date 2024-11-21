package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeArguments() *TypeArguments {
	return &TypeArguments{}
}

type TypeArguments struct {
	AbstractType
	AbstractTypeArgument
	util.DefaultList[intmod.ITypeArgument]
}

func (t *TypeArguments) Size() int {
	return t.DefaultList.Size()
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
	return t.Elements()
}

func (t *TypeArguments) TypeArgumentSize() int {
	return t.DefaultList.Size()
}

func (t *TypeArguments) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitTypeArguments(t)
}
