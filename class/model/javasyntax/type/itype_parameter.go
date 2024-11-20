package _type

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type AbstractTypeParameter struct {
	util.DefaultBase[intsyn.ITypeParameter]
}

func (p *AbstractTypeParameter) AcceptTypeParameterVisitor(visitor intsyn.ITypeParameterVisitor) {
}
