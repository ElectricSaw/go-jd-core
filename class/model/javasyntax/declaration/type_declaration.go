package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewTypeDeclaration(annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intmod.IDeclaration) intmod.ITypeDeclaration {
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

	annotationReferences intmod.IAnnotationReference
	flags                int
	internalTypeName     string
	name                 string
	bodyDeclaration      intmod.IDeclaration
}

func (d *TypeDeclaration) AnnotationReferences() intmod.IAnnotationReference {
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

func (d *TypeDeclaration) BodyDeclaration() intmod.IBodyDeclaration {
	return d.bodyDeclaration.(intmod.IBodyDeclaration)
}
