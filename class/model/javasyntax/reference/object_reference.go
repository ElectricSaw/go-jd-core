package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewObjectReference(internalName, qualifiedName, name string) intsyn.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectType(internalName, qualifiedName, name).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithDim(internalName, qualifiedName, name string, dimension int) intsyn.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithDim(internalName, qualifiedName, name, dimension).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithArgs(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument) intsyn.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithArgs(internalName, qualifiedName, name, typeArguments).(*_type.ObjectType),
	}
}

func NewObjectReferenceWithAll(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument, dimension int) intsyn.IObjectReference {
	return &ObjectReference{
		ObjectType: *_type.NewObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, dimension).(*_type.ObjectType),
	}
}

type ObjectReference struct {
	_type.ObjectType
}

func (e *ObjectReference) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitObjectReference(e)
}
