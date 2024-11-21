package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewAnnotationDeclaration(annotationDeclarators intmod.IFieldDeclarator,
	annotationReferences intmod.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intmod.IDeclaration) intmod.IAnnotationDeclaration {
	return &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		annotationDeclaratiors: annotationDeclarators,
	}
}

type AnnotationDeclaration struct {
	TypeDeclaration
	util.DefaultBase[intmod.IMemberDeclaration]

	annotationDeclaratiors intmod.IFieldDeclarator
}

func (d *AnnotationDeclaration) AnnotationDeclarators() intmod.IFieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
