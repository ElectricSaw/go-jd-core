package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileEnumDeclaration(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName, name string, interfaces intmod.IType,
	bodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileEnumDeclaration {
	d := &ClassFileEnumDeclaration{
		EnumDeclaration: *declaration.NewEnumDeclarationWithAll(annotationReferences, flags,
			internalTypeName, name, interfaces, nil, bodyDeclaration).(*declaration.EnumDeclaration),
	}
	if bodyDeclaration != nil {
		d.firstLineNumber = bodyDeclaration.FirstLineNumber()
	}
	return d
}

type ClassFileEnumDeclaration struct {
	declaration.EnumDeclaration

	firstLineNumber int
}

func (d *ClassFileEnumDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileEnumDeclaration) String() string {
	return fmt.Sprintf("ClassFileEnumDeclaration{%s, firstLineNumber:%d}", d.InternalTypeName(), d.firstLineNumber)
}

func NewClassFileConstant(lineNumber int, name string, index int, arguments intmod.IExpression,
	bodyDeclaration intmod.IBodyDeclaration) intsrv.IClassFileConstant {
	return &ClassFileConstant{
		Constant: *declaration.NewConstant5(lineNumber, name, arguments, bodyDeclaration).(*declaration.Constant),
		index:    index,
	}
}

type ClassFileConstant struct {
	declaration.Constant

	index int
}

func (d *ClassFileConstant) Index() int {
	return d.index
}

func (d *ClassFileConstant) String() string {
	return fmt.Sprintf("ClassFileConstant{%s : %d}", d.Name(), d.Index())
}
