package model

import "github.com/ElectricSaw/go-jd-core/class/util"

const (
	FlagBoolean = 1 << iota
	FlagChar
	FlagFloat
	FlagDouble
	FlagByte
	FlagShort
	FlagInt
	FlagLong
	FlagVoid
)

type IDiamondTypeArgument interface {
	ITypeArgument

	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type IGenericType interface {
	IType
	ITypeArgument

	HashCode() int
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType
	IsGenericType() bool
	AcceptTypeVisitor(visitor ITypeVisitor)
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsGenericTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	Equals(o IGenericType) bool
	String() string
}

type IInnerObjectType interface {
	IObjectType

	HashCode() int
	CreateType(dimension int) IType
	CreateTypeWithArg(typeArguments ITypeArgument) IType
	IsInnerObjectType() bool
	OuterType() IObjectType
	AcceptTypeVisitor(visitor ITypeVisitor)
	IsInnerObjectTypeArgument() bool
	TypeArguments() ITypeArgument
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	String() string
}

type IType interface {
	ITypeVisitable
	ITypeArgumentVisitable
	util.IBase[IType]

	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType

	IsGenericType() bool
	IsInnerObjectType() bool
	IsObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	OuterType() IObjectType
	InternalName() string

	HashCode() int

	//////////////////////////////////////////////////////
	// ITypeArgument 추가시 아래와 같은 오류 발생.
	// Invalid recursive type: anonymous interface refers to itsel
	//////////////////////////////////////////////////////

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

type ITypeVisitor interface {
	VisitPrimitiveType(y IPrimitiveType)
	VisitObjectType(y IObjectType)
	VisitInnerObjectType(y IInnerObjectType)
	VisitTypes(types ITypes)
	VisitGenericType(y IGenericType)
}

type ITypeVisitable interface {
	AcceptTypeVisitor(visitor ITypeVisitor)
}

type IObjectType interface {
	ITypeArgument
	ITypeVisitable
	ITypeArgumentVisitable
	util.IBase[IType]

	QualifiedName() string
	HashCode() int
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType

	IsObjectType() bool
	IsGenericType() bool
	IsInnerObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	OuterType() IObjectType
	InternalName() string
	AcceptTypeVisitor(visitor ITypeVisitor)
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentAssignableFromWithObj(typeBounds map[string]IType, objectType IObjectType) bool
	IsObjectTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	TypeArguments() ITypeArgument
	CreateTypeWithArgs(typeArguments ITypeArgument) IObjectType
	String() string
	RawEquals(other IObjectType) bool
}

type ITypeArgument interface {
	ITypeArgumentVisitable

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
	HashCode() int
}

type ITypeArgumentVisitable interface {
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type ITypeArgumentVisitor interface {
	VisitTypeArguments(arguments ITypeArguments)
	VisitDiamondTypeArgument(argument IDiamondTypeArgument)
	VisitWildcardExtendsTypeArgument(argument IWildcardExtendsTypeArgument)
	VisitWildcardSuperTypeArgument(argument IWildcardSuperTypeArgument)
	VisitWildcardTypeArgument(argument IWildcardTypeArgument)
	VisitPrimitiveType(t IPrimitiveType)
	VisitObjectType(t IObjectType)
	VisitInnerObjectType(t IInnerObjectType)
	VisitGenericType(t IGenericType)
}

type ITypeParameter interface {
	ITypeParameterVisitable
	util.IBase[ITypeParameter]

	Identifier() string
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
	String() string
}

type ITypeParameterVisitable interface {
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
}

type ITypeParameterVisitor interface {
	VisitTypeParameter(parameter ITypeParameter)
	VisitTypeParameterWithTypeBounds(parameter ITypeParameterWithTypeBounds)
	VisitTypeParameters(parameters ITypeParameters)
}

type IPrimitiveType interface {
	IType
	ITypeArgument

	HashCode() int
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType
	IsPrimitiveType() bool
	AcceptTypeVisitor(visitor ITypeVisitor)
	IsTypeArgumentAssignableFrom(_ map[string]IType, typeArgument ITypeArgument) bool
	IsPrimitiveTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	Equals(o ITypeArgument) bool
	Flags() int
	LeftFlags() int
	RightFlags() int
	JavaPrimitiveFlags() int
	String() string
}

type ITypeArguments interface {
	ITypeArgument
	util.IList[ITypeArgument]

	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentList() bool
	TypeArgumentFirst() ITypeArgument
	TypeArgumentList() []ITypeArgument
	TypeArgumentSize() int
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type ITypeParameterWithTypeBounds interface {
	ITypeParameter

	TypeBounds() IType
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
	String() string
}

type ITypeParameters interface {
	ITypeParameter
	util.IList[ITypeParameter]

	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
}

type ITypes interface {
	IType
	util.IList[IType]

	IsTypes() bool
	AcceptTypeVisitor(visitor ITypeVisitor)
}

type IUnmodifiableTypes interface {
	ITypes

	IsUnmodifiableTypes() bool
}

type IWildcardExtendsTypeArgument interface {
	ITypeArgument

	Type() IType
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardExtendsTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	HashCode() int
	Equals(o ITypeArgument) bool
	String() string
}

type IWildcardSuperTypeArgument interface {
	ITypeArgument

	Type() IType
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardSuperTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	HashCode() int
	Equals(o ITypeArgument) bool
	String() string
}

type IWildcardTypeArgument interface {
	ITypeArgument

	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	Equals(o ITypeArgument) bool
	String() string
}
