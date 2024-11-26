package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileClassDeclaration(annotationReferences intmod.IAnnotationReference,
	flags int, internalName string, name string, typeParameters intmod.ITypeParameter,
	superType intmod.IObjectType, interfaces intmod.IType, bodyDeclaration intsrv.IClassFileBodyDeclaration,
) intsrv.IClassFileClassDeclaration {
	d := &ClassFileClassDeclaration{
		ClassDeclaration: *declaration.NewClassDeclarationWithAll(annotationReferences,
			flags, internalName, name, bodyDeclaration, typeParameters, interfaces, superType).(*declaration.ClassDeclaration),
	}

	if bodyDeclaration != nil {
		d.firstLineNumber = bodyDeclaration.FirstLineNumber()
	}

	return d
}

type ClassFileClassDeclaration struct {
	declaration.ClassDeclaration

	firstLineNumber int
}

func (d *ClassFileClassDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileClassDeclaration) String() string {
	return fmt.Sprintf("ClassFileClassDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
