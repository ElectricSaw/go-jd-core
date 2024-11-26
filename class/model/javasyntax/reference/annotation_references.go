package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewAnnotationReferences() intmod.IAnnotationReferences {
	return &AnnotationReferences{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IAnnotationReference](0),
	}
}

type AnnotationReferences struct {
	AnnotationReference
	util.DefaultList[intmod.IAnnotationReference]
}

func (r *AnnotationReferences) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
