package reference

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAnnotationReferences() intmod.IAnnotationReferences {
	return &AnnotationReferences{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IAnnotationReference](0).(*util.DefaultList[intmod.IAnnotationReference]),
	}
}

type AnnotationReferences struct {
	AnnotationReference
	util.DefaultList[intmod.IAnnotationReference]
}

func (r *AnnotationReferences) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
