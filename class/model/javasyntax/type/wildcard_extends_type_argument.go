package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewWildcardExtendsTypeArgument(typ intsyn.IType) intsyn.IWildcardExtendsTypeArgument {
	return &WildcardExtendsTypeArgument{
		typ: typ,
	}
}

type WildcardExtendsTypeArgument struct {
	AbstractTypeArgument

	typ intsyn.IType
}

func (t *WildcardExtendsTypeArgument) Type() intsyn.IType {
	return t.typ
}

func (t *WildcardExtendsTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]intsyn.IType, typeArgument intsyn.ITypeArgument) bool {
	if typeArgument.IsWildcardExtendsTypeArgument() {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument.Type())
	} else if _, ok := typeArgument.(intsyn.ITypeArgument); ok {
		return t.typ.IsTypeArgumentAssignableFrom(typeBounds, typeArgument)
	}
	return false
}

func (t *WildcardExtendsTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return true
}

func (t *WildcardExtendsTypeArgument) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitWildcardExtendsTypeArgument(t)
}

func (t *WildcardExtendsTypeArgument) HashCode() int {
	if t.typ == nil {
		return 957014778
	}

	return 957014778 + t.typ.HashCode()
}

func (t *WildcardExtendsTypeArgument) Equals(o intsyn.ITypeArgument) bool {
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
