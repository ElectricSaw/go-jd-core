package _type

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeParameters() intsyn.ITypeParameters {
	return &TypeParameters{}
}

type TypeParameters struct {
	util.DefaultList[intsyn.ITypeParameter]
}

func (t *TypeParameters) AcceptTypeParameterVisitor(visitor intsyn.ITypeParameterVisitor) {
	visitor.VisitTypeParameters(t)
}
