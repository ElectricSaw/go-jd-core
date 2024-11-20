package _type

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type AbstractTypeArgument struct {
}

func (t *AbstractTypeArgument) TypeArgumentFirst() intsyn.ITypeArgument {
	return t
}

func (t *AbstractTypeArgument) TypeArgumentList() []intsyn.ITypeArgument {
	return nil
}

func (t *AbstractTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *AbstractTypeArgument) Type() intsyn.IType {
	return OtTypeUndefinedObject.(intsyn.IType)
}

func (t *AbstractTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]intsyn.IType, _ intsyn.ITypeArgument) bool {
	return false
}

func (t *AbstractTypeArgument) IsTypeArgumentList() bool {
	return false
}

func (t *AbstractTypeArgument) IsGenericTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsObjectTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) IsWildcardTypeArgument() bool {
	return false
}

func (t *AbstractTypeArgument) AcceptTypeArgumentVisitor(_ intsyn.ITypeArgumentVisitor) {

}

func (t *AbstractTypeArgument) HashCode() int {
	return hashCodeWithStruct(t)
}
