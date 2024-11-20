package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewFieldDeclaration(flags int, typ intsyn.IType, fieldDeclaration intsyn.IFieldDeclarator) intsyn.IFieldDeclaration {
	return &FieldDeclaration{
		flags:            flags,
		typ:              typ,
		fieldDeclarators: fieldDeclaration,
	}
}

func NewFieldDeclarationWithAll(annotationReferences intsyn.IAnnotationReference, flags int,
	typ intsyn.IType, fieldDeclaration intsyn.IFieldDeclarator) intsyn.IFieldDeclaration {
	return &FieldDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		typ:                  typ,
		fieldDeclarators:     fieldDeclaration,
	}
}

type FieldDeclaration struct {
	AbstractMemberDeclaration

	annotationReferences intsyn.IAnnotationReference
	flags                int
	typ                  intsyn.IType
	fieldDeclarators     intsyn.IFieldDeclarator
}

func (d *FieldDeclaration) Flags() int {
	return d.flags
}

func (d *FieldDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *FieldDeclaration) AnnotationReferences() intsyn.IAnnotationReference {
	return d.annotationReferences
}

func (d *FieldDeclaration) Type() intsyn.IType {
	return d.typ
}

func (d *FieldDeclaration) SetType(t intsyn.IType) {
	d.typ = t
}

func (d *FieldDeclaration) FieldDeclarators() intsyn.IFieldDeclarator {
	return d.fieldDeclarators
}

func (d *FieldDeclaration) SetFieldDeclarators(fd intsyn.IFieldDeclarator) {
	d.fieldDeclarators = fd
}

func (d *FieldDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitFieldDeclaration(d)
}
