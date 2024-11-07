package _type

type IType interface {
	Name() string
	Descriptor() string
	Dimension() int
	CreateType(dimension int) IType

	IsGenericType() bool
	IsInnerObjectType() bool
	IsObjectType() bool
	IsPrimitiveType() bool
	IsTypes() bool

	//OuterType() IObjectType
	//InternalName() string

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

type TypeVisitor interface {
	VisitPrimitiveType(y *PrimitiveType)
	VisitObjectType(y *ObjectType)
	VisitInnerObjectType(y *InnerObjectType)
	VisitTypes(types *Types)
	VisitGenericType(y *GenericType)
}

type TypeVisitable interface {
	AcceptTypeVisitor(visitor TypeVisitor)
}

type IObjectType interface {
	Dimension() int
	QualifiedName() string
	OuterType() IObjectType
	InternalName() string
	TypeArguments() ITypeArgument
	AcceptTypeVisitor(visitor TypeVisitor)
	AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor)
}

type AbstractType struct {
}

func (t *AbstractType) GetName() string {
	return ""
}

func (t *AbstractType) GetDescriptor() string {
	return ""
}

func (t *AbstractType) GetDimension() int {
	return -1
}

func (t *AbstractType) CreateType(dimension int) IType {
	return nil
}

func (t *AbstractType) IsGenericType() bool {
	return false
}

func (t *AbstractType) IsInnerObjectType() bool {
	return false
}

func (t *AbstractType) IsObjectType() bool {
	return false
}

func (t *AbstractType) IsPrimitiveType() bool {
	return false
}

func (t *AbstractType) IsTypes() bool {
	return false
}

func (t *AbstractType) OuterType() IObjectType {
	return TypeUndefinedObject
}

func (t *AbstractType) InternalName() string {
	return ""
}

func (t *AbstractType) AcceptTypeVisitor(visitor TypeVisitor) {
}
