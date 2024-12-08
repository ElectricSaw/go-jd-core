package _type

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

var Diamond = NewDiamondTypeArgument()

func NewDiamondTypeArgument() intmod.IDiamondTypeArgument {
	return &DiamondTypeArgument{}
}

type DiamondTypeArgument struct {
	AbstractTypeArgument
}

func (a *DiamondTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]intmod.IType, typeArgument intmod.ITypeArgument) bool {
	return true
}

func (a *DiamondTypeArgument) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitDiamondTypeArgument(a)
}
