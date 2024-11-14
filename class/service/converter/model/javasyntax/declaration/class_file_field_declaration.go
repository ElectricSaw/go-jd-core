package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewClassFileFieldDeclaration(flags int, typ _type.IType, fieldDeclaration declaration.IFieldDeclarator) *ClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration),
	}
}

func NewClassFileFieldDeclaration2(flags int, typ _type.IType, fieldDeclaration declaration.IFieldDeclarator, firstLineNumber int) *ClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration),
		firstLineNumber:  firstLineNumber,
	}
}

func NewClassFileFieldDeclaration3(annotationReferences reference.IAnnotationReference, flags int, typ _type.IType, fieldDeclaration declaration.IFieldDeclarator) *ClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(annotationReferences, flags, typ, fieldDeclaration),
	}
}

func NewClassFileFieldDeclaration4(annotationReferences reference.IAnnotationReference, flags int, typ _type.IType, fieldDeclaration declaration.IFieldDeclarator, firstLineNumber int) *ClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(annotationReferences, flags, typ, fieldDeclaration),
		firstLineNumber:  firstLineNumber,
	}
}

type ClassFileFieldDeclaration struct {
	declaration.FieldDeclaration

	firstLineNumber int
}

func (d *ClassFileFieldDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileFieldDeclaration) SetFirstLineNumber(firstLineNumber int) {
	d.firstLineNumber = firstLineNumber
}

func (d *ClassFileFieldDeclaration) String() string {
	return fmt.Sprintf("ClassFileFieldDeclaration{%s %s, firstLineNumber=%d}", d.Type(), d.FieldDeclarators(), d.firstLineNumber)
}
