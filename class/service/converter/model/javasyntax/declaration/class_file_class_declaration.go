package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
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
	d.SetValue(d)

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
