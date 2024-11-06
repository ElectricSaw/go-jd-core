package reference

type AbstractNopReferenceVisitor struct {
}

func (e *AbstractNopReferenceVisitor) VisitAnnotationElementValue(reference *AnnotationElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReference(reference *AnnotationReference)    {}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReferences(references *AnnotationReferences) {}
func (e *AbstractNopReferenceVisitor) VisitElementValueArrayInitializerElementValue(reference *ElementValueArrayInitializerElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitElementValues(references *ElementValues)         {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePair(reference *ElementValuePair)    {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePairs(references *ElementValuePairs) {}
func (e *AbstractNopReferenceVisitor) VisitExpressionElementValue(reference *ExpressionElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitInnerObjectReference(reference *InnerObjectReference) {}
func (e *AbstractNopReferenceVisitor) VisitObjectReference(reference *ObjectReference)           {}
