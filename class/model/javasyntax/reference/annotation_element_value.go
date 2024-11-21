package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewAnnotationElementValue(reference intmod.IAnnotationReference) intmod.IAnnotationElementValue {
	return &AnnotationElementValue{
		AnnotationReference: *NewAnnotationReferenceWithAll(reference.Type(),
			reference.ElementValue(), reference.ElementValuePairs()).(*AnnotationReference),
	}
}

type AnnotationElementValue struct {
	AnnotationReference
}

func (r *AnnotationElementValue) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitAnnotationElementValue(r)
}

func (r *AnnotationElementValue) String() string {
	return fmt.Sprintf("AnnotationElementValue{type=%v, elementValue=%s, elementValuePairs=%v", r.Type(), r.ElementValue(), r.ElementValuePairs())
}
