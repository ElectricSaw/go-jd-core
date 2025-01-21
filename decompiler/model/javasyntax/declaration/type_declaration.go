package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewTypeDeclaration(annotationReferences AnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration Declaration) TypeDeclaration {
	d := &TypeDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		internalTypeName:     internalTypeName,
		name:                 name,
		bodyDeclaration:      bodyDeclaration,
	}
	d.SetValue(d)
	return d
}

type TypeDeclaration struct {
	AbstractTypeDeclaration

	annotationReferences AnnotationReference
	flags                int
	internalTypeName     string
	name                 string
	bodyDeclaration      Declaration
}

func (d *TypeDeclaration) AnnotationReferences() AnnotationReference {
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

func (d *TypeDeclaration) BodyDeclaration() BodyDeclaration {
	return d.bodyDeclaration.(BodyDeclaration)
}
