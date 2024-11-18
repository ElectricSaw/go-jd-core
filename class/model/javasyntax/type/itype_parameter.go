package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
)

type AbstractTypeParameter struct {
	util.DefaultBase[intsyn.ITypeParameter]
}

func (p *AbstractTypeParameter) AcceptTypeParameterVisitor(visitor intsyn.ITypeParameterVisitor) {
}
