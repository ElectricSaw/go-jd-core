package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassDeclaration(flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration) *ClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(nil, flags, internalTypeName, name, bodyDeclaration, nil, nil),
	}
}

func NewClassDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration, typeParameters _type.ITypeParameter, interfaces _type.IType, superType _type.ObjectType) *ClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(annotationReferences, flags, internalTypeName, name, bodyDeclaration, typeParameters, interfaces),
		superType:            superType,
	}
}

type ClassDeclaration struct {
	InterfaceDeclaration

	superType _type.ObjectType
}

func (d *ClassDeclaration) SuperType() *_type.ObjectType {
	return &d.superType
}

func (d *ClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.internalTypeName)
}
