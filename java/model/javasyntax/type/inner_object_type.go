package _type

import "fmt"

func NewInnerObjectType(internalName, qualifiedName, name string, outerType *ObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectTypeWithDim(internalName, qualifiedName, name string, dimension int, outerType *ObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument, outerType *ObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewInnerObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int, outerType *ObjectType) *InnerObjectType {
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

	outerType *ObjectType
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) CreateType(dimension int) IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, t.typeArguments, dimension, t.outerType)
}

func (t *InnerObjectType) CreateTypeWithArg(typeArguments ITypeArgument) IType {
	return NewInnerObjectTypeWithAll(t.internalName, t.qualifiedName, t.name, typeArguments, t.dimension, t.outerType)
}

func (t *InnerObjectType) IsInnerObjectType() bool {
	return true
}

func (t *InnerObjectType) GetOuterType() *ObjectType {
	return t.outerType
}

func (t *InnerObjectType) AcceptTypeVisitor(visitor TypeVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) IsInnerObjectTypeArgument() bool {
	return true
}

func (t *InnerObjectType) AcceptTypeArgumentVisitor(visitor TypeArgumentVisitor) {
	visitor.VisitInnerObjectType(t)
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) String() string {
	if t.typeArguments == nil {
		return fmt.Sprintf("InnerObjectType { %s.%s }", t.outerType.String(), t.descriptor)
	} else {
		return fmt.Sprintf("InnerObjectType { %s.%s<%s> }", t.outerType.String(), t.descriptor, t.typeArguments)
	}
}
