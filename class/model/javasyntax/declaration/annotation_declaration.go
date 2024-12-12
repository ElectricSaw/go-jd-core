package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewAnnotationDeclaration(annotationDeclarators intmod.IFieldDeclarator,
	annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intmod.IDeclaration) intmod.IAnnotationDeclaration {
	d := &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		annotationDeclaratiors: annotationDeclarators,
	}
	d.SetValue(d)
	return d
}

type AnnotationDeclaration struct {
	TypeDeclaration

	annotationDeclaratiors intmod.IFieldDeclarator
}

func (d *AnnotationDeclaration) AnnotationDeclarators() intmod.IFieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
