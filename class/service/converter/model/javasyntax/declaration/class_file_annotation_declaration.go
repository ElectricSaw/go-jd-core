package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
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
	d.SetValue(d)

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
