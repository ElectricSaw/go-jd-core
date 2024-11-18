package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/util"
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
