package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewAnnotationDeclaration(annotationDeclarators FieldDeclarator,
	annotationReferences AnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration Declaration) AnnotationDeclaration {
	d := &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		annotationDeclaratiors: annotationDeclarators,
	}
	d.SetValue(d)
	return d
}

type AnnotationDeclaration struct {
	TypeDeclaration

	annotationDeclaratiors FieldDeclarator
}

func (d *AnnotationDeclaration) AnnotationDeclarators() FieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) AcceptDeclaration(visitor IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
