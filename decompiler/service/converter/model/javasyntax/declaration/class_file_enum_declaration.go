package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileEnumDeclaration(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName, name string, interfaces intmod.IType,
	bodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileEnumDeclaration {
	d := &ClassFileEnumDeclaration{
		EnumDeclaration: *model.NewEnumDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, interfaces, nil, bodyDeclaration).(*model.EnumDeclaration),
	}
	if bodyDeclaration != nil {
		d.firstLineNumber = bodyDeclaration.FirstLineNumber()
	}
	d.SetValue(d)
	return d
}

type ClassFileEnumDeclaration struct {
	model.EnumDeclaration

	firstLineNumber int
}

func (d *ClassFileEnumDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileEnumDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitEnumDeclaration(d)
}

func (d *ClassFileEnumDeclaration) String() string {
	return fmt.Sprintf("ClassFileEnumDeclaration{%s, firstLineNumber:%d}", d.InternalTypeName(), d.firstLineNumber)
}

func NewClassFileConstant(lineNumber int, name string, index int, arguments intmod.IExpression,
	bodyDeclaration intmod.IBodyDeclaration) intsrv.IClassFileConstant {
	c := &ClassFileConstant{
		Constant: *model.NewConstant5(lineNumber, name, arguments, bodyDeclaration).(*model.Constant),
		index:    index,
	}
	c.SetValue(c)
	return c
}

type ClassFileConstant struct {
	model.Constant

	index int
}

func (d *ClassFileConstant) Index() int {
	return d.index
}

func (d *ClassFileConstant) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitEnumDeclarationConstant(d)
}

func (d *ClassFileConstant) String() string {
	return fmt.Sprintf("ClassFileConstant{%s : %d}", d.Name(), d.Index())
}
