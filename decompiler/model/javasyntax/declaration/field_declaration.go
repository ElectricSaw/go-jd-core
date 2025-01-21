package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFieldDeclaration(flags int, typ Type, fieldDeclaration FieldDeclarator) FieldDeclaration {
	return NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration)
}

func NewFieldDeclarationWithAll(annotationReferences AnnotationReference, flags int,
	typ Type, fieldDeclaration FieldDeclarator) FieldDeclaration {
	d := &FieldDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		typ:                  typ,
		fieldDeclarators:     fieldDeclaration,
	}
	d.SetValue(d)
	return d
}

type FieldDeclaration struct {
	AbstractMemberDeclaration
	util.DefaultBase[MemberDeclaration]

	annotationReferences AnnotationReference
	flags                int
	typ                  Type
	fieldDeclarators     FieldDeclarator
}

func (d *FieldDeclaration) Flags() int {
	return d.flags
}

func (d *FieldDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *FieldDeclaration) AnnotationReferences() AnnotationReference {
	return d.annotationReferences
}

func (d *FieldDeclaration) Type() Type {
	return d.typ
}

func (d *FieldDeclaration) SetType(t Type) {
	d.typ = t
}

func (d *FieldDeclaration) FieldDeclarators() FieldDeclarator {
	return d.fieldDeclarators
}

func (d *FieldDeclaration) SetFieldDeclarators(fd FieldDeclarator) {
	d.fieldDeclarators = fd
}

func (d *FieldDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}
