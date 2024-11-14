package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewFieldDeclaration(flags int, typ _type.IType, fieldDeclaration IFieldDeclarator) *FieldDeclaration {
	return &FieldDeclaration{
		flags:            flags,
		typ:              typ,
		fieldDeclaration: fieldDeclaration,
	}
}

func NewFieldDeclarationWithAll(annotationReferences reference.IAnnotationReference, flags int, typ _type.IType, fieldDeclaration IFieldDeclarator) *FieldDeclaration {
	return &FieldDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		typ:                  typ,
		fieldDeclaration:     fieldDeclaration,
	}
}

type FieldDeclaration struct {
	AbstractMemberDeclaration

	annotationReferences reference.IAnnotationReference
	flags                int
	typ                  _type.IType
	fieldDeclaration     IFieldDeclarator
}

func (d *FieldDeclaration) Flags() int {
	return d.flags
}

func (d *FieldDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *FieldDeclaration) AnnotationReferences() reference.IAnnotationReference {
	return d.annotationReferences
}

func (d *FieldDeclaration) Type() _type.IType {
	return d.typ
}

func (d *FieldDeclaration) SetType(t _type.IType) {
	d.typ = t
}

func (d *FieldDeclaration) FieldDeclaration() IFieldDeclarator {
	return d.fieldDeclaration
}

func (d *FieldDeclaration) SetFieldDeclaration(fd IFieldDeclarator) {
	d.fieldDeclaration = fd
}

func (d *FieldDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}
