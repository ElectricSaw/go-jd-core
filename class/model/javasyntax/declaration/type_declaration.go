package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewTypeDeclaration(annotationReferences intsyn.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intsyn.IDeclaration) intsyn.ITypeDeclaration {
	return &TypeDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		internalTypeName:     internalTypeName,
		name:                 name,
		bodyDeclaration:      bodyDeclaration,
	}
}

type TypeDeclaration struct {
	AbstractTypeDeclaration

	annotationReferences intsyn.IAnnotationReference
	flags                int
	internalTypeName     string
	name                 string
	bodyDeclaration      intsyn.IDeclaration
}

func (d *TypeDeclaration) AnnotationReferences() intsyn.IAnnotationReference {
	return d.annotationReferences
}

func (d *TypeDeclaration) Flags() int {
	return d.flags
}

func (d *TypeDeclaration) SetFlags(flags int) {
	d.flags = flags
}

func (d *TypeDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *TypeDeclaration) Name() string {
	return d.name
}

func (d *TypeDeclaration) BodyDeclaration() intsyn.IBodyDeclaration {
	return d.bodyDeclaration.(intsyn.IBodyDeclaration)
}
