package _type

import "fmt"

func NewInnerObjectType(internalName, qualifiedName, name string, outerType IObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectTypeWithDim(internalName, qualifiedName, name string, dimension int, outerType IObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectTypeWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument, outerType IObjectType) *InnerObjectType {
	return NewInnerObjectTypeWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewInnerObjectTypeWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int, outerType IObjectType) *InnerObjectType {
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

	outerType IObjectType
}

/////////////////////////////////////////////////////////////////////

func (t *InnerObjectType) HashCode() int {
	result := 111476860 + t.ObjectType.HashCode()
	result = 31*result + t.outerType.HashCode()
	return result
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

func (t *InnerObjectType) OuterType() IObjectType {
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
		return fmt.Sprintf("InnerObjectType { %s.%s }", t.outerType, t.descriptor)
	} else {
		return fmt.Sprintf("InnerObjectType { %s.%s<%s> }", t.outerType, t.descriptor, t.typeArguments)
	}
}
