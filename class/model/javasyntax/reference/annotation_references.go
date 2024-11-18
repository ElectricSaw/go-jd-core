package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewAnnotationReferences() intsyn.IAnnotationReferences {
	return &AnnotationReferences{}
}

type AnnotationReferences struct {
	util.DefaultList[intsyn.IAnnotationReference]
}

func (r *AnnotationReferences) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
