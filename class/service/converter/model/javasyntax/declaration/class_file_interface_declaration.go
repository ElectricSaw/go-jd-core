package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileInterfaceDeclaration(
	annotationReferences intmod.IAnnotationReference,
	flags int,
	internalTypeName string,
	name string,
	typeParameters intmod.ITypeParameter,
	interfaces intmod.IType,
	bodyDeclaration intsrv.IClassFileBodyDeclaration,
) intsrv.IClassFileInterfaceDeclaration {
	d := &ClassFileInterfaceDeclaration{
		InterfaceDeclaration: *declaration.NewInterfaceDeclarationWithAll(annotationReferences,
			flags, internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*declaration.InterfaceDeclaration),
		firstLineNumber: bodyDeclaration.FirstLineNumber(),
	}
	d.SetValue(d)
	return d
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
