package reference

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewObjectReference(internalName, qualifiedName, name string) intmod.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectType(internalName, qualifiedName, name).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithDim(internalName, qualifiedName, name string, dimension int) intmod.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithDim(internalName, qualifiedName, name, dimension).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithArgs(internalName, qualifiedName, name string, typeArguments intmod.ITypeArgument) intmod.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithAll(internalName, qualifiedName, name string, typeArguments intmod.ITypeArgument, dimension int) intmod.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension).(*_type.ObjectType),
	}
}

type ObjectReference struct {
	_type.ObjectType
}

func (e *ObjectReference) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitObjectReference(e)
}
