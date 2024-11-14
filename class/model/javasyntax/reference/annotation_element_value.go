package reference

import "fmt"

func NewAnnotationElementValue(reference *AnnotationReference) *AnnotationElementValue {
	return &AnnotationElementValue{
		AnnotationReference: *NewAnnotationReferenceWithAll(reference.Type(), reference.ElementValue(), reference.ElementValuePairs()),
	}
}

type AnnotationElementValue struct {
	AnnotationReference
}

func (r *AnnotationElementValue) Accept(visitor ReferenceVisitor) {
	visitor.VisitAnnotationElementValue(r)
}

func (r *AnnotationElementValue) String() string {
	return fmt.Sprintf("AnnotationElementValue{type=%v, elementValue=%s, elementValuePairs=%v", r.Type(), r.ElementValue(), r.ElementValuePairs())
}
