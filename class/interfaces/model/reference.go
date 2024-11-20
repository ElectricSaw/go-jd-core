package model

import "bitbucket.org/coontec/go-jd-core/class/util"

const (
	EvUnknown = iota
	EvPrimitiveType
	EvEnumConstValue
	EvClassInfo
	EvAnnotationValue
	EvArrayValue
)

type IAnnotationElementValue interface {
	Type() IObjectType
	ElementValue() IElementValue
	ElementValuePairs() IElementValuePair
	Accept(visitor IReferenceVisitor)
	String() string
}

type IAnnotationReferences interface {
	util.IList[IAnnotationReference]
	Accept(visitor IReferenceVisitor)
}

type IElementValue interface {
	IReference
}

type IElementValuePair interface {
	IElementValue

	Name() string
	ElementValue() IElementValue
	Accept(visitor IReferenceVisitor)
	String() string
}

type IElementValueArrayInitializerElementValue interface {
	ElementValueArrayInitializer() IElementValue
	Accept(visitor IReferenceVisitor)
	String() string
}

type IElementValuePairs interface {
	util.IList[IElementValuePair]
	IElementValuePair
}

type IElementValues interface {
	util.IList[IElementValue]
}

type IExpressionElementValue interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	Accept(visitor IReferenceVisitor)
	String() string
}

type IInnerObjectReference interface {
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
	Accept(visitor IReferenceVisitor)
}

type IObjectReference interface {
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
	Accept(visitor IReferenceVisitor)
	String() string
}

type IReference interface {
	Accept(visitor IReferenceVisitor)
}

type IAnnotationReference interface {
	IReference

	Type() IObjectType
	ElementValue() IElementValue
	ElementValuePairs() IElementValuePair
	Accept(visitor IReferenceVisitor)
}

type IReferenceVisitor interface {
	VisitAnnotationElementValue(reference IAnnotationElementValue)
	VisitAnnotationReference(reference IAnnotationReference)
	VisitAnnotationReferences(references IAnnotationReferences)
	VisitElementValueArrayInitializerElementValue(reference IElementValueArrayInitializerElementValue)
	VisitElementValues(references IElementValues)
	VisitElementValuePair(reference IElementValuePair)
	VisitElementValuePairs(references IElementValuePairs)
	VisitExpressionElementValue(reference IExpressionElementValue)
	VisitInnerObjectReference(reference IInnerObjectReference)
	VisitObjectReference(reference IObjectReference)
}
