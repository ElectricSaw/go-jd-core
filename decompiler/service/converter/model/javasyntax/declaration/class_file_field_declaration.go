package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileFieldDeclaration(flags int, typ model.IType,
	fieldDeclaration *model.FieldDeclarator) ClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(nil, flags, typ, fieldDeclaration, -1)
}

func NewClassFileFieldDeclaration2(flags int, typ model.IType,
	fieldDeclaration *model.FieldDeclarator, firstLineNumber int) ClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(nil, flags, typ, fieldDeclaration, firstLineNumber)
}

func NewClassFileFieldDeclaration3(annotationReferences *model.AnnotationReference,
	flags int, typ model.IType, fieldDeclaration *model.FieldDeclarator) ClassFileFieldDeclaration {
	return NewClassFileFieldDeclaration4(annotationReferences, flags, typ, fieldDeclaration, -1)
}

func NewClassFileFieldDeclaration4(annotationReferences *model.AnnotationReference,
	flags int, typ model.IType, fieldDeclaration *model.FieldDeclarator, firstLineNumber int) ClassFileFieldDeclaration {
	d := ClassFileFieldDeclaration{
		DefaultBase:          *util.NewDefaultBase[model.IMemberDeclaration]().(*util.DefaultBase[model.IMemberDeclaration]),
		AnnotationReferences: annotationReferences,
		Flags:                flags,
		Type:                 typ,
		FieldDeclarators:     fieldDeclaration,
		FirstLineNumber:      firstLineNumber,
	}
	d.SetValue(&d)
	return d
}

type ClassFileFieldDeclaration struct {
	util.DefaultBase[model.IMemberDeclaration]

	AnnotationReferences *model.AnnotationReference
	Flags                int
	Type                 model.IType
	FieldDeclarators     *model.FieldDeclarator
	FirstLineNumber      int
}

func (d *ClassFileFieldDeclaration) GetAnnotationReferences() *model.AnnotationReference {
	return d.AnnotationReferences
}

func (d *ClassFileFieldDeclaration) GetFlags() int {
	return d.Flags
}

func (d *ClassFileFieldDeclaration) GetType() model.IType {
	return d.Type
}

func (d *ClassFileFieldDeclaration) GetFieldDeclarators() *model.FieldDeclarator {
	return d.FieldDeclarators
}

func (d *ClassFileFieldDeclaration) GetFirstLineNumber() int {
	return d.FirstLineNumber
}

func (d *ClassFileFieldDeclaration) IsClassDeclaration() bool {
	return false
}

func (d *ClassFileFieldDeclaration) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}

func (d *ClassFileFieldDeclaration) HashCode() int {
	result := 327494460 + d.Flags
	if d.AnnotationReferences != nil {
		result = 31*result + (d.AnnotationReferences.HashCode())
	} else {
		result = 31*result + 0
	}
	result = 31*result + d.Type.HashCode()
	result = 31*result + d.FieldDeclarators.HashCode()
	return result
}

func (d *ClassFileFieldDeclaration) String() string {
	return fmt.Sprintf("ClassFileFieldDeclaration{%s %s, firstLineNumber=%d}", d.Type, d.FieldDeclarators, d.FirstLineNumber)
}
