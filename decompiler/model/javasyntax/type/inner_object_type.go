package _type

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewInnerObjectType(internalName, qualifiedName, name string, outerType intmod.IObjectType) intmod.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectTypeWithDim(internalName, qualifiedName, name string, dimension int, outerType intmod.IObjectType) intmod.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments intmod.ITypeArgument, outerType intmod.IObjectType) intmod.IInnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewInnerObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments intmod.ITypeArgument, dimension int, outerType intmod.IObjectType) intmod.IInnerObjectType {
	t := &InnerObjectType{
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
	t.SetValue(t)
	return t
}

type InnerObjectType struct {
	ObjectType

	outerType intmod.IObjectType
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) HashCode() int {
	result := 111476860 + t.ObjectType.HashCode()
	result = 31*result + t.outerType.HashCode()
	return result
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) CreateType(dimension int) intmod.IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, t.typeArguments, dimension, t.outerType).(intmod.IType)
}

func (t *InnerObjectType) CreateTypeWithArg(typeArguments intmod.ITypeArgument) intmod.IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, typeArguments, t.dimension, t.outerType).(intmod.IType)
}

func (t *InnerObjectType) IsInnerObjectType() bool {
	return true
}

func (t *InnerObjectType) OuterType() intmod.IObjectType {
	return t.outerType
}

func (t *InnerObjectType) AcceptTypeVisitor(visitor intmod.ITypeVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) IsInnerObjectTypeArgument() bool {
	return true
}

func (t *InnerObjectType) AcceptTypeArgumentVisitor(visitor intmod.ITypeArgumentVisitor) {
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
