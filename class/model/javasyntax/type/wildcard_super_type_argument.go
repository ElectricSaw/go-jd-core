package _type

import "fmt"

func NewWildcardSuperTypeArgument(typ IType) *WildcardSuperTypeArgument {
	return &WildcardSuperTypeArgument{
		typ: typ,
	}
}

type WildcardSuperTypeArgument struct {
	AbstractTypeArgument

	typ IType
}

func (t *WildcardSuperTypeArgument) Type() IType {
	return t.typ
}

func (t *WildcardSuperTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	if typeArgument.IsWildcardSuperTypeArgument() {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument.Type())
	} else if _, ok := typeArgument.(ITypeArgument); ok {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument)
	}
	return false
}

func (t *WildcardSuperTypeArgument) IsWildcardSuperTypeArgument() bool {
	return true
}

func (t *WildcardSuperTypeArgument) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitWildcardSuperTypeArgument(t)
}

func (t *WildcardSuperTypeArgument) HashCode() int {
	if t.typ == nil {
		return 979510081
	}

	return 979510081 + t.typ.HashCode()
}

func (t *WildcardSuperTypeArgument) Equals(o ITypeArgument) bool {
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

	if t.typ != that.typ {
		return t.typ == that.typ
	}

	return that.typ == nil
}

func (t *WildcardSuperTypeArgument) String() string {
	return fmt.Sprintf("WildcardSuperTypeArgument{? super %s }", t.typ)
}
