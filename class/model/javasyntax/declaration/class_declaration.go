package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassDeclaration(flags int, internalTypeName string, name string, bodyDeclaration *intsyn.IBodyDeclaration) intsyn.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(nil, flags,
			internalTypeName, name, bodyDeclaration, nil, nil).(*InterfaceDeclaration),
	}
}

func NewClassDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *intsyn.IBodyDeclaration, typeParameters _type.ITypeParameter, interfaces _type.IType, superType _type.IObjectType) intsyn.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*InterfaceDeclaration),
		superType: superType,
	}
}

type ClassDeclaration struct {
	InterfaceDeclaration

	superType _type.IObjectType
}

func (d *ClassDeclaration) SuperType() _type.IObjectType {
	return d.superType
}

func (d *ClassDeclaration) IsClassDeclaration() bool {
	return true
}

func (d *ClassDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassDeclaration) String() string {
	return fmt.Sprintf("ClassDeclaration{%s}", d.internalTypeName)
}
