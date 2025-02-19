package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileClassDeclaration(annotationReferences intmod.IAnnotationReference,
	flags int, internalName string, name string, typeParameters intmod.ITypeParameter,
	superType intmod.IObjectType, interfaces intmod.IType, bodyDeclaration intsrv.IClassFileBodyDeclaration,
) intsrv.IClassFileClassDeclaration {
	d := &ClassFileClassDeclaration{
		ClassDeclaration: *model.NewClassDeclarationWithAll(annotationReferences,
			flags, internalName, name, bodyDeclaration, typeParameters, interfaces, superType).(*model.ClassDeclaration),
	}

	if bodyDeclaration != nil {
		d.firstLineNumber = bodyDeclaration.FirstLineNumber()
	}
	d.SetValue(d)

	return d
}

type ClassFileClassDeclaration struct {
	model.ClassDeclaration

	firstLineNumber int
}

func (d *ClassFileClassDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileClassDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitClassDeclaration(d)
}

func (d *ClassFileClassDeclaration) String() string {
	return fmt.Sprintf("ClassFileClassDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
