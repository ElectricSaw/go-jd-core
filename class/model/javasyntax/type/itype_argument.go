package _type

type ITypeArgument interface {
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentList() bool
	TypeArgumentFirst() ITypeArgument  // ITypeArgument
	TypeArgumentList() []ITypeArgument // ITypeArgument
	TypeArgumentSize() int
	IsGenericTypeArgument() bool
	IsInnerObjectTypeArgument() bool
	IsObjectTypeArgument() bool
	IsPrimitiveTypeArgument() bool
	IsWildcardExtendsTypeArgument() bool
	IsWildcardSuperTypeArgument() bool
	IsWildcardTypeArgument() bool
	Type() IType
}

type TypeArgumentVisitable interface {
	AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor)
}

type TypeArgumentVisitor interface {
	VisitTypeArguments(arguments *TypeArguments)
	VisitDiamondTypeArgument(argument *DiamondTypeArgument)
	VisitWildcardExtendsTypeArgument(argument *WildcardExtendsTypeArgument)
	VisitWildcardSuperTypeArgument(argument *WildcardSuperTypeArgument)
	VisitWildcardTypeArgument(argument *WildcardTypeArgument)
	VisitPrimitiveType(t *PrimitiveType)
	VisitObjectType(t *ObjectType)
	VisitInnerObjectType(t *InnerObjectType)
	VisitGenericType(t *GenericType)
}

type AbstractTypeArgument struct {
}

func (t *AbstractTypeArgument) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *AbstractTypeArgument) TypeArgumentList() []ITypeArgument {
	return nil
}

func (t *AbstractTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *AbstractTypeArgument) Type() IType {
	return TypeUndefinedObject
}

func (t *AbstractTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]IType, _ ITypeArgument) bool {
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

func (t *AbstractTypeArgument) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {

}