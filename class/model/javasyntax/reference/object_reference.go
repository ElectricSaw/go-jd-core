package reference

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewObjectReference(internalName, qualifiedName, name string) *ObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectType(internalName, qualifiedName, name),
	}
}

func NewObjectReferenceWithDim(internalName, qualifiedName, name string, dimension int) *ObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithDim(internalName, qualifiedName, name, dimension),
	}
}

func NewObjectReferenceWithArgs(internalName, qualifiedName, name string, typeArguments _type.ITypeArgument) *ObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments),
	}
}

func NewObjectReferenceWithAll(internalName, qualifiedName, name string, typeArguments _type.ITypeArgument, dimension int) *ObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension),
	}
}

type ObjectReference struct {
	_type.ObjectType
}

func (e *ObjectReference) Accept(visitor ReferenceVisitor) {
	visitor.VisitObjectReference(e)
}