package _type

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

var WildcardTypeArgumentEmpty = NewWildcardTypeArgument()

func NewWildcardTypeArgument() intsyn.IWildcardTypeArgument {
	return &WildcardTypeArgument{}
}

type WildcardTypeArgument struct {
	AbstractTypeArgument
}

func (t *WildcardTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	return true
}

func (t *WildcardTypeArgument) IsWildcardTypeArgument() bool {
	return true
}

func (t *WildcardTypeArgument) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitWildcardTypeArgument(t)
}

func (t *WildcardTypeArgument) Equals(o intsyn.ITypeArgument) bool {
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
