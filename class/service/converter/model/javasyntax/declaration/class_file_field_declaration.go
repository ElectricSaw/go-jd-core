package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
)

func NewClassFileFieldDeclaration(flags int, typ intmod.IType,
	fieldDeclaration intmod.IFieldDeclarator) intsrv.IClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(nil, flags, typ, fieldDeclaration, -1)
}

func NewClassFileFieldDeclaration2(flags int, typ intmod.IType,
	fieldDeclaration intmod.IFieldDeclarator, firstLineNumber int) intsrv.IClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(nil, flags, typ, fieldDeclaration, firstLineNumber)
}

func NewClassFileFieldDeclaration3(annotationReferences intmod.IAnnotationReference,
	flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator) intsrv.IClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(annotationReferences, flags, typ, fieldDeclaration, -1)
}

func NewClassFileFieldDeclaration4(annotationReferences intmod.IAnnotationReference,
	flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator, firstLineNumber int) intsrv.IClassFileFieldDeclaration {
	d := &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(annotationReferences, flags, typ, fieldDeclaration).(*declaration.FieldDeclaration),
		firstLineNumber:  firstLineNumber,
	}
	d.SetValue(d)
	return d
}

type ClassFileFieldDeclaration struct {
	ClassFileMethodDeclaration
	declaration.FieldDeclaration

	firstLineNumber int
}

func (d *ClassFileFieldDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileFieldDeclaration) SetFirstLineNumber(firstLineNumber int) {
	d.firstLineNumber = firstLineNumber
}

func (d *ClassFileFieldDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}

func (d *ClassFileFieldDeclaration) String() string {
	return fmt.Sprintf("ClassFileFieldDeclaration{%s %s, firstLineNumber=%d}", d.Type(), d.FieldDeclarators(), d.firstLineNumber)
}
