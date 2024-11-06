package declaration

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/reference"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
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

func (d *FieldDeclaration) GetFlags() int {
	return d.flags
}

func (d *FieldDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *FieldDeclaration) GetAnnotationReferences() reference.IAnnotationReference {
	return d.annotationReferences
}

func (d *FieldDeclaration) GetType() _type.IType {
	return d.typ
}

func (d *FieldDeclaration) SetType(t _type.IType) {
	d.typ = t
}

func (d *FieldDeclaration) GetFieldDeclaration() IFieldDeclarator {
	return d.fieldDeclaration
}

func (d *FieldDeclaration) SetFieldDeclaration(fd IFieldDeclarator) {
	d.fieldDeclaration = fd
}

func (d *FieldDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}
