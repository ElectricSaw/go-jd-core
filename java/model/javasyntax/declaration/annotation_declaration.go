package declaration

import "fmt"

func NewAnnotationDeclaration(annotationDeclaratiors BaseFieldDeclarator, annotationReferences BaseAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration) *AnnotationDeclaration {
	return &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration),
		annotationDeclaratiors: annotationDeclaratiors,
	}
}

type AnnotationDeclaration struct {
	TypeDeclaration

	annotationDeclaratiors BaseFieldDeclarator
}

func (d *AnnotationDeclaration) AnnotationDeclarator() BaseFieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
