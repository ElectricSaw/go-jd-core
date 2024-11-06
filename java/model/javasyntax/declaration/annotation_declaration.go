package declaration

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/reference"
	"fmt"
)

func NewAnnotationDeclaration(annotationDeclarators IFieldDeclarator, annotationReferences reference.IAnnotationReference, flags int, internalTypeName string, name string, bodyDeclaration *BodyDeclaration) *AnnotationDeclaration {
	return &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration),
		annotationDeclaratiors: annotationDeclarators,
	}
}

type AnnotationDeclaration struct {
	TypeDeclaration

	annotationDeclaratiors IFieldDeclarator
}

func (d *AnnotationDeclaration) GetAnnotationDeclarator() IFieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
