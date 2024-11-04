package _type

func NewDiamondTypeArgument() *DiamondTypeArgument {
	return &DiamondTypeArgument{}
}

type DiamondTypeArgument struct {
	AbstractTypeArgument
}

func (a *DiamondTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	return true
}

func (a *DiamondTypeArgument) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitDiamondTypeArgument(a)
}
