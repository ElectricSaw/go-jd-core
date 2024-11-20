package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewAnnotationDeclaration(annotationDeclarators intsyn.IFieldDeclarator,
	annotationReferences intsyn.IAnnotationReference, flags int,
	internalTypeName string, name string, bodyDeclaration intsyn.IDeclaration) intsyn.IAnnotationDeclaration {
	return &AnnotationDeclaration{
		TypeDeclaration:        *NewTypeDeclaration(annotationReferences, flags, internalTypeName, name, bodyDeclaration).(*TypeDeclaration),
		annotationDeclaratiors: annotationDeclarators,
	}
}

type AnnotationDeclaration struct {
	TypeDeclaration
	util.DefaultBase[intsyn.IMemberDeclaration]

	annotationDeclaratiors intsyn.IFieldDeclarator
}

func (d *AnnotationDeclaration) AnnotationDeclarators() intsyn.IFieldDeclarator {
	return d.annotationDeclaratiors
}

func (d *AnnotationDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitAnnotationDeclaration(d)
}

func (d *AnnotationDeclaration) String() string {
	return fmt.Sprintf("AnnotationDeclaration { %s }", d.internalTypeName)
}
