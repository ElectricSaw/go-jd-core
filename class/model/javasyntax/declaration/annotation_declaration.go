package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
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
