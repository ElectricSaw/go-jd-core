package declaration

func NewTypeDeclaration(annotationReferences BaseAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration) *TypeDeclaration {
	return &TypeDeclaration{
		annotationReferences: annotationReferences,
		flags:                flags,
		internalTypeName:     internalTypeName,
		name:                 name,
		BodyDeclaration:      bodyDeclaration,
	}
}

type TypeDeclaration struct {
	annotationReferences BaseAnnotationReference
	flags                int
	internalTypeName     string
	name                 string
	bodyDeclaration      *BodyDeclaration
}

func (d *TypeDeclaration) AnnotationReferences() BaseAnnotationReference {
	return d.annotationReferences
}

func (d *TypeDeclaration) Flags() int {
	return d.flags
}

func (d *TypeDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *TypeDeclaration) Name() string {
	return d.name
}

func (d *TypeDeclaration) BodyDeclaration() *BodyDeclaration {
	return d.bodyDeclaration
}

func (d *TypeDeclaration) ignoreBaseTypeDeclaration() {
}

func (d *TypeDeclaration) ignoreMemberDeclaration() {
}
