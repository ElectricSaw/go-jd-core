package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewAnnotationElementValue(reference *AnnotationReference) intsyn.IAnnotationElementValue {
	return &AnnotationElementValue{
		AnnotationReference: *NewAnnotationReferenceWithAll(reference.Type(),
			reference.ElementValue(), reference.ElementValuePairs()).(*AnnotationReference),
	}
}

type AnnotationElementValue struct {
	AnnotationReference
}

func (r *AnnotationElementValue) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitAnnotationElementValue(r)
}

func (r *AnnotationElementValue) String() string {
	return fmt.Sprintf("AnnotationElementValue{type=%v, elementValue=%s, elementValuePairs=%v", r.Type(), r.ElementValue(), r.ElementValuePairs())
}
