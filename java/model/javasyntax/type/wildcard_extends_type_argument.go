package _type

import "fmt"

func NewWildcardExtendsTypeArgument(typ IType) *WildcardExtendsTypeArgument {
	return &WildcardExtendsTypeArgument{
		typ: typ,
	}
}

type WildcardExtendsTypeArgument struct {
	AbstractTypeArgument

	typ IType
}

func (t *WildcardExtendsTypeArgument) Type() ITypeArgument {
	return t.typ
}

func (t *WildcardExtendsTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	if typeArgument.IsWildcardExtendsTypeArgument() {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument.GetType())
	} else if _, ok := typeArgument.(ITypeArgument); ok {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument)
	}
	return false
}

func (t *WildcardExtendsTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return true
}

func (t *WildcardExtendsTypeArgument) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitWildcardExtendsTypeArgument(t)
}

func (t *WildcardExtendsTypeArgument) equals(o ITypeArgument) bool {
	if t == o {
		return true
	}

	if o == nil {
		return false
	}

	that, ok := o.(*WildcardExtendsTypeArgument)
	if !ok {
		return false
	}

	if t.typ != that.typ {
		return t.typ == that.typ
	}

	return that.typ == nil
}

func (t *WildcardExtendsTypeArgument) String() string {
	return fmt.Sprintf("WildcardExtendsTypeArgument{? extends %s }", t.typ)
}