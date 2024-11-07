package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/reference"

func NewTypeDeclaration(annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration) *TypeDeclaration {
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
	bodyDeclaration      *BodyDeclaration
}

func (d *TypeDeclaration) GetAnnotationReferences() reference.IAnnotationReference {
	return d.annotationReferences
}

func (d *TypeDeclaration) GetFlags() int {
	return d.flags
}

func (d *TypeDeclaration) GetInternalTypeName() string {
	return d.internalTypeName
}

func (d *TypeDeclaration) GetName() string {
	return d.name
}

func (d *TypeDeclaration) GetBodyDeclaration() *BodyDeclaration {
	return d.bodyDeclaration
}
