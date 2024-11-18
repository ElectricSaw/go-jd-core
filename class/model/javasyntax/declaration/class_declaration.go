package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewClassDeclaration(flags int, internalTypeName string, name string,
	bodyDeclaration intsyn.IBodyDeclaration) intsyn.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(nil, flags,
			internalTypeName, name, bodyDeclaration, nil, nil).(*InterfaceDeclaration),
	}
}

func NewClassDeclarationWithAll(annotationReferences intsyn.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intsyn.IBodyDeclaration,
	typeParameters intsyn.ITypeParameter, interfaces intsyn.IType, superType intsyn.IObjectType) intsyn.IClassDeclaration {
	return &ClassDeclaration{
		InterfaceDeclaration: *NewInterfaceDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*InterfaceDeclaration),
		superType: superType,
	}
}

type ClassDeclaration struct {
	InterfaceDeclaration

	superType intsyn.IObjectType
}

func (d *ClassDeclaration) SuperType() intsyn.IObjectType {
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
