package _type

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewDiamondTypeArgument() intsyn.IDiamondTypeArgument {
	return &DiamondTypeArgument{}
}

type DiamondTypeArgument struct {
	AbstractTypeArgument
}

func (a *DiamondTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	return true
}

func (a *DiamondTypeArgument) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitDiamondTypeArgument(a)
}
