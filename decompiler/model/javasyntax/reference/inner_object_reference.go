package reference

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewInnerObjectReference(internalName, qualifiedName, name string,
	outerType intmod.IObjectType) intmod.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectType(internalName, qualifiedName, name, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithDim(internalName, qualifiedName, name string,
	dimension int, outerType intmod.IObjectType) intmod.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithDim(internalName, qualifiedName, name, dimension, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithArgs(internalName, qualifiedName, name string,
	typeArguments intmod.ITypeArgument, outerType intmod.IObjectType) intmod.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments, outerType).(*_type.InnerObjectType),
	}
}

func NewInnerObjectReferenceWithAll(internalName, qualifiedName, name string,
	typeArguments intmod.ITypeArgument, dimension int, outerType intmod.IObjectType) intmod.IInnerObjectReference {
	return &InnerObjectReference{
		InnerObjectType: *_type.NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension, outerType).(*_type.InnerObjectType),
	}
}

type InnerObjectReference struct {
	_type.InnerObjectType
}

func (e *InnerObjectReference) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitInnerObjectReference(e)
}
