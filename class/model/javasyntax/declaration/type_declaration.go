package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"

func NewTypeDeclaration(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration Declaration) *TypeDeclaration {
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

	annotationReferences reference.IAnnotationReference
	flags                int
	internalTypeName     string
	name                 string
	bodyDeclaration      Declaration
}

func (d *TypeDeclaration) AnnotationReferences() reference.IAnnotationReference {
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

func (d *TypeDeclaration) BodyDeclaration() *BodyDeclaration {
	return d.bodyDeclaration.(*BodyDeclaration)
}
