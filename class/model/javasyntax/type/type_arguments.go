package _type

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeArguments() *TypeArguments {
	return &TypeArguments{}
}

type TypeArguments struct {
	AbstractTypeArgument
	util.DefaultList[intsyn.ITypeArgument]
}

func (t *TypeArguments) IsTypeArgumentAssignableFrom(typeBounds map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	ata, ok := typeArgument.(*TypeArguments)
	if !ok {
		return false
	}

	if t.Size() != ata.Size() {
		return false
	}

	for i := 0; i < t.Size(); i++ {

		if !t.Get(i).IsTypeArgumentAssignableFrom(typeBounds, ata.Get(i)) {
			return false
		}
	}

	return true
}

func (t *TypeArguments) IsTypeArgumentList() bool {
	return true
}

func (t *TypeArguments) TypeArgumentFirst() intsyn.ITypeArgument {
	return t.Get(0)
}

func (t *TypeArguments) TypeArgumentList() []intsyn.ITypeArgument {
	return t.Elements()
}

func (t *TypeArguments) TypeArgumentSize() int {
	return t.Size()
}

func (t *TypeArguments) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitTypeArguments(t)
}
