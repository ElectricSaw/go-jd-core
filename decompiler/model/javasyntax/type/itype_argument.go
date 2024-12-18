package _type

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

type AbstractTypeArgument struct {
}

func (t *AbstractTypeArgument) TypeArgumentFirst() intmod.ITypeArgument {
	return t
}

func (t *AbstractTypeArgument) TypeArgumentList() []intmod.ITypeArgument {
	return nil
}

func (t *AbstractTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *AbstractTypeArgument) Type() intmod.IType {
	return OtTypeUndefinedObject.(intmod.IType)
}

func (t *AbstractTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]intmod.IType, _ intmod.ITypeArgument) bool {
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

func (t *AbstractTypeArgument) AcceptTypeArgumentVisitor(_ intmod.ITypeArgumentVisitor) {

}

func (t *AbstractTypeArgument) HashCode() int {
	return hashCodeWithStruct(t)
}
