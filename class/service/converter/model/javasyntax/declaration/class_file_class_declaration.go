package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileClassDeclaration(annotationReferences intmod.IAnnotationReference,
	flags int, internalName string, name string, typeParameters intmod.ITypeParameter,
	superType intmod.IType, interfaces intmod.IType, bodyDeclaration intsrv.IClassFileBodyDeclaration,
) intsrv.IClassFileClassDeclaration {
	return &ClassFileClassDeclaration{}
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
