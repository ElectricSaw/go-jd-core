package reference

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAnnotationElementValue(reference intmod.IAnnotationReference) intmod.IAnnotationElementValue {
	v := &AnnotationElementValue{
		AnnotationReference: *NewAnnotationReferenceWithAll(reference.Type(),
			reference.ElementValue(), reference.ElementValuePairs()).(*AnnotationReference),
	}
	v.SetValue(v)
	return v
}

type AnnotationElementValue struct {
	AnnotationReference
	util.DefaultBase[intmod.IAnnotationElementValue]
}

func (r *AnnotationElementValue) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitAnnotationElementValue(r)
}

func (r *AnnotationElementValue) String() string {
	return fmt.Sprintf("AnnotationElementValue{type=%v, elementValue=%s, elementValuePairs=%v", r.Type(), r.ElementValue(), r.ElementValuePairs())
}
