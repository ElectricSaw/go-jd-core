package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	"fmt"
)

func NewClassFileAnnotationDeclaration(annotationDeclarators declaration.IFieldDeclarator, annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *ClassFileBodyDeclaration) *ClassFileAnnotationDeclaration {
	d := &ClassFileAnnotationDeclaration{
		AnnotationDeclaration: *declaration.NewAnnotationDeclaration(annotationDeclarators, annotationReferences, flags, internalTypeName, name, bodyDeclaration),
	}
	if bodyDeclaration == nil {
		d.firstLineNumber = 0
	} else {
		d.firstLineNumber = bodyDeclaration.FirstLineNumber()
	}

	return d
}

type ClassFileAnnotationDeclaration struct {
	declaration.AnnotationDeclaration

	firstLineNumber int
}

func (d *ClassFileAnnotationDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileAnnotationDeclaration) String() string {
	return fmt.Sprintf("ClassFileAnnotationDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
