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

func (t *WildcardExtendsTypeArgument) Type() IType {
	return t.typ
}

func (t *WildcardExtendsTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	if typeArgument.IsWildcardExtendsTypeArgument() {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument.Type())
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

func (t *WildcardExtendsTypeArgument) HashCode() int {
	if t.typ == nil {
		return 957014778
	}

	return 957014778 + t.typ.HashCode()
}

func (t *WildcardExtendsTypeArgument) Equals(o ITypeArgument) bool {
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
