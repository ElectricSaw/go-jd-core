package reference

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewInnerObjectReference(internalName, qualifiedName, name string, outerType *_type.ObjectType) *InnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectType(internalName, qualifiedName, name, outerType),
	}
}

func NewInnerObjectReferenceWithDim(internalName, qualifiedName, name string, dimension int, outerType *_type.ObjectType) *InnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithDim(internalName, qualifiedName, name, dimension, outerType),
	}
}

func NewInnerObjectReferenceWithArgs(internalName, qualifiedName, name string, typeArguments _type.ITypeArgument, outerType *_type.ObjectType) *InnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments, outerType),
	}
}

func NewInnerObjectReferenceWithAll(internalName, qualifiedName, name string, typeArguments _type.ITypeArgument, dimension int, outerType *_type.ObjectType) *InnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension, outerType),
	}
}

type InnerObjectReference struct {
	_type.InnerObjectType
}

func (e *InnerObjectReference) Accept(visitor ReferenceVisitor) {
	visitor.VisitInnerObjectReference(e)
}
