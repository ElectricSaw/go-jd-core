package javasyntax

import "bitbucket.org/coontec/javaClass/class/util"

type IDiamondTypeArgument interface {
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type IGenericType interface {
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

	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType
	Size() int

	IsGenericType() bool
	IsInnerObjectType() bool
	IsObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	OuterType() IObjectType
	InternalName() string

	///////////////////////////////////////////////////////////////

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

	///////////////////////////////////////////////////////////////

	HashCode() int
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
	QualifiedName() string
	HashCode() int
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType
	IsObjectType() bool
	InternalName() string
	AcceptTypeVisitor(visitor ITypeVisitor)
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentAssignableFromWithObj(typeBounds map[string]IType, objectType IObjectType) bool
	IsObjectTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	TypeArguments() ITypeArgument
	CreateTypeWithArgs(typeArguments ITypeArgument) IObjectType
	String() string
}

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
	util.Base[ITypeParameter]
	ITypeParameterVisitable

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
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentList() bool
	GetTypeArgumentFirst() ITypeArgument
	GetTypeArgumentList() []ITypeArgument
	TypeArgumentSize() int
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type ITypeParameterWithTypeBounds interface {
	TypeBounds() IType
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
	String() string
}

type ITypeParameters interface {
	util.IList[ITypeParameter]
	AcceptTypeParameterVisitor(visitor ITypeParameterVisitor)
}

type ITypes interface {
	util.IList[IType]
	IsTypes() bool
	AcceptTypeVisitor(visitor ITypeVisitor)
}

type IUnmodifiableTypes interface {
	util.IList[IType]
	ListIterator(i int) []IType
}

type IWildcardExtendsTypeArgument interface {
	Type() IType
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardExtendsTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	HashCode() int
	Equals(o ITypeArgument) bool
	String() string
}

type IWildcardSuperTypeArgument interface {
	Type() IType
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardSuperTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	HashCode() int
	Equals(o ITypeArgument) bool
	String() string
}

type IWildcardTypeArgument interface {
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsWildcardTypeArgument() bool
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
	Equals(o ITypeArgument) bool
	String() string
}
