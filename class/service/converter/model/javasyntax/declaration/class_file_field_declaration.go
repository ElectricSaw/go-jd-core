package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileFieldDeclaration(flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator) intsrv.IClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration).(*declaration.FieldDeclaration),
	}
}

func NewClassFileFieldDeclaration2(flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator, firstLineNumber int) intsrv.IClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration).(*declaration.FieldDeclaration),
		firstLineNumber:  firstLineNumber,
	}
}

func NewClassFileFieldDeclaration3(annotationReferences intmod.IAnnotationReference, flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator) intsrv.IClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(annotationReferences, flags, typ, fieldDeclaration).(*declaration.FieldDeclaration),
	}
}

func NewClassFileFieldDeclaration4(annotationReferences intmod.IAnnotationReference, flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator, firstLineNumber int) intsrv.IClassFileFieldDeclaration {
	return &ClassFileFieldDeclaration{
		FieldDeclaration: *declaration.NewFieldDeclarationWithAll(annotationReferences, flags, typ, fieldDeclaration).(*declaration.FieldDeclaration),
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
