package _type

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

var WildcardTypeArgumentEmpty = NewWildcardTypeArgument()

func NewWildcardTypeArgument() intmod.IWildcardTypeArgument {
	return &WildcardTypeArgument{}
}

type WildcardTypeArgument struct {
	AbstractTypeArgument
}

func (t *WildcardTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]intmod.IType, _ intmod.ITypeArgument) bool {
	return true
}

func (t *WildcardTypeArgument) IsWildcardTypeArgument() bool {
	return true
}

func (t *WildcardTypeArgument) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
	visitor.VisitWildcardTypeArgument(t)
}

func (t *WildcardTypeArgument) Equals(o intmod.ITypeArgument) bool {
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
