package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type AbstractTypeParameter struct {
	util.DefaultBase[intmod.ITypeParameter]
}

func (p *AbstractTypeParameter) AcceptTypeParameterVisitor(visitor intmod.ITypeParameterVisitor) {
}
