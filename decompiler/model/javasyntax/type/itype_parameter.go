package _type

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

type AbstractTypeParameter struct {
}

func (p *AbstractTypeParameter) AcceptTypeParameterVisitor(visitor intmod.ITypeParameterVisitor) {
}