package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileEnumDeclaration(annotationReferences reference.IAnnotationReference, flags int, internalTypeName, name string, interfaces _type.IType, bodyDeclaration *ClassFileBodyDeclaration) *ClassFileEnumDeclaration {
	d := &ClassFileEnumDeclaration{
		EnumDeclaration: *declaration.NewEnumDeclarationWithAll(annotationReferences, flags, internalTypeName, name, interfaces, nil, bodyDeclaration),
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

func NewClassFileConstant(lineNumber int, name string, index int, arguments expression.IExpression, bodyDeclaration *declaration.BodyDeclaration) *ClassFileConstant {
	return &ClassFileConstant{
		Constant: *declaration.NewConstant5(lineNumber, name, arguments, bodyDeclaration),
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
