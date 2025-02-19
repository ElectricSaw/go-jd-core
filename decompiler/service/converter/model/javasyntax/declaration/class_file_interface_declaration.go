package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
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
		InterfaceDeclaration: *model.NewInterfaceDeclarationWithAll(annotationReferences,
			flags, internalTypeName, name, bodyDeclaration, typeParameters, interfaces).(*model.InterfaceDeclaration),
		firstLineNumber: bodyDeclaration.FirstLineNumber(),
	}
	d.SetValue(d)
	return d
}

type ClassFileInterfaceDeclaration struct {
	model.InterfaceDeclaration

	firstLineNumber int
}

func (d *ClassFileInterfaceDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileInterfaceDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitInterfaceDeclaration(d)
}

func (d *ClassFileInterfaceDeclaration) String() string {
	return fmt.Sprintf("ClassFileInterfaceDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
