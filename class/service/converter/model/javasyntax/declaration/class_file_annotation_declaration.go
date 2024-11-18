package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileAnnotationDeclaration(annotationDeclarators intsyn.IFieldDeclarator, annotationReferences intsyn.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *ClassFileBodyDeclaration) intsyn.IAnnotationDeclaration {
	d := &ClassFileAnnotationDeclaration{
		AnnotationDeclaration: *declaration.NewAnnotationDeclaration(annotationDeclarators,
			annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*declaration.AnnotationDeclaration),
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
