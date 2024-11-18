package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileInterfaceDeclaration(
	annotationReferences reference.IAnnotationReference,
	flags int,
	internalTypeName string,
	name string,
	typeParameters _type.ITypeParameter,
	interfaces _type.IType,
	bodyDeclaration *ClassFileBodyDeclaration,
) *ClassFileInterfaceDeclaration {
	return &ClassFileInterfaceDeclaration{
		InterfaceDeclaration: *declaration.NewInterfaceDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, &bodyDeclaration.BodyDeclaration, typeParameters, interfaces),
		firstLineNumber: bodyDeclaration.FirstLineNumber(),
	}
}

type ClassFileInterfaceDeclaration struct {
	declaration.InterfaceDeclaration

	firstLineNumber int
}

func (d *ClassFileInterfaceDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileInterfaceDeclaration) String() string {
	return fmt.Sprintf("ClassFileInterfaceDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
