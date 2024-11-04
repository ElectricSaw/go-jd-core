package _type

var WildcardTypeArgumentEmpty = NewWildcardTypeArgument()

func NewWildcardTypeArgument() *WildcardTypeArgument {
	return &WildcardTypeArgument{}
}

type WildcardTypeArgument struct {
	AbstractTypeArgument
}

func (t *WildcardTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	return true
}

func (t *WildcardTypeArgument) IsWildcardTypeArgument() bool {
	return true
}

func (t *WildcardTypeArgument) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitWildcardTypeArgument(t)
}

func (t *WildcardTypeArgument) equals(o ITypeArgument) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	that, ok := o.(*WildcardSuperTypeArgument)
	if !ok {
		return false
	}

	return that.typ == nil
}

func (t *WildcardTypeArgument) String() string {
	return "Wildcard{?}"
}
