package _type

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewInnerObjectType(internalName, qualifiedName, name string, outerType intsyn.IObjectType) intsyn.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectTypeWithDim(internalName, qualifiedName, name string, dimension int, outerType intsyn.IObjectType) intsyn.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument, outerType intsyn.IObjectType) intsyn.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewInnerObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments intsyn.ITypeArgument, dimension int, outerType intsyn.IObjectType) intsyn.IInnerObjectType {
	return &InnerObjectType{
		ObjectType: ObjectType{
			internalName:  internalName,
			qualifiedName: qualifiedName,
			name:          name,
			typeArguments: typeArguments,
			dimension:     dimension,
			descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
		},
		outerType: outerType,
	}
}

type InnerObjectType struct {
	ObjectType

	outerType intsyn.IObjectType
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) HashCode() int {
	result := 111476860 + t.ObjectType.HashCode()
	result = 31*result + t.outerType.HashCode()
	return result
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) CreateType(dimension int) intsyn.IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, t.typeArguments, dimension, t.outerType).(intsyn.IType)
}

func (t *InnerObjectType) CreateTypeWithArg(typeArguments intsyn.ITypeArgument) intsyn.IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, typeArguments, t.dimension, t.outerType).(intsyn.IType)
}

func (t *InnerObjectType) IsInnerObjectType() bool {
	return true
}

func (t *InnerObjectType) OuterType() intsyn.IObjectType {
	return t.outerType
}

func (t *InnerObjectType) AcceptTypeVisitor(visitor intsyn.ITypeVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) IsInnerObjectTypeArgument() bool {
	return true
}

func (t *InnerObjectType) AcceptTypeArgumentVisitor(visitor intsyn.ITypeArgumentVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) String() string {
	if t.typeArguments == nil {
		return fmt.Sprintf("InnerObjectType { %s.%s }", t.outerType, t.descriptor)
	} else {
		return fmt.Sprintf("InnerObjectType { %s.%s<%s> }", t.outerType, t.descriptor, t.typeArguments)
	}
}
