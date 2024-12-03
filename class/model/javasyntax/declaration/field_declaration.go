package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewFieldDeclaration(flags int, typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator) intmod.IFieldDeclaration {
	return NewFieldDeclarationWithAll(nil, flags, typ, fieldDeclaration)
}

func NewFieldDeclarationWithAll(annotationReferences intmod.IAnnotationReference, flags int,
	typ intmod.IType, fieldDeclaration intmod.IFieldDeclarator) intmod.IFieldDeclaration {
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
	util.DefaultBase[intmod.IMemberDeclaration]

	annotationReferences intmod.IAnnotationReference
	flags                int
	typ                  intmod.IType
	fieldDeclarators     intmod.IFieldDeclarator
}

func (d *FieldDeclaration) Flags() int {
	return d.flags
}

func (d *FieldDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *FieldDeclaration) AnnotationReferences() intmod.IAnnotationReference {
	return d.annotationReferences
}

func (d *FieldDeclaration) Type() intmod.IType {
	return d.typ
}

func (d *FieldDeclaration) SetType(t intmod.IType) {
	d.typ = t
}

func (d *FieldDeclaration) FieldDeclarators() intmod.IFieldDeclarator {
	return d.fieldDeclarators
}

func (d *FieldDeclaration) SetFieldDeclarators(fd intmod.IFieldDeclarator) {
	d.fieldDeclarators = fd
}

func (d *FieldDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}
