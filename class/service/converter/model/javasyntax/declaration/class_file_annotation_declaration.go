package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileAnnotationDeclaration(annotationReferences intmod.IAnnotationReference,
	flags int, internalTypeName string, name string, bodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileAnnotationDeclaration {
	d := &ClassFileAnnotationDeclaration{
		AnnotationDeclaration: *declaration.NewAnnotationDeclaration(nil, annotationReferences,
			flags, internalTypeName, name, bodyDeclaration).(*declaration.AnnotationDeclaration),
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
