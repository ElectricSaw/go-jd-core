package reference

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewAnnotationReferences() intsyn.IAnnotationReferences {
	return &AnnotationReferences{
		DefaultList: *util.NewDefaultList[intsyn.IAnnotationReference](0),
	}
}

type AnnotationReferences struct {
	util.DefaultList[intsyn.IAnnotationReference]
}

func (r *AnnotationReferences) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
