package reference

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewInnerObjectReference(internalName, qualifiedName, name string,
	outerType intsyn.IObjectType) intsyn.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectType(internalName, qualifiedName, name, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithDim(internalName, qualifiedName, name string,
	dimension int, outerType intsyn.IObjectType) intsyn.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithDim(internalName, qualifiedName, name, dimension, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithArgs(internalName, qualifiedName, name string,
	typeArguments intsyn.ITypeArgument, outerType intsyn.IObjectType) intsyn.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithAll(internalName, qualifiedName, name string,
	typeArguments intsyn.ITypeArgument, dimension int, outerType intsyn.IObjectType) intsyn.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension, outerType).(*_type.InnerObjectType),
	}
}

type InnerObjectReference struct {
	_type.InnerObjectType
}

func (e *InnerObjectReference) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitInnerObjectReference(e)
}
