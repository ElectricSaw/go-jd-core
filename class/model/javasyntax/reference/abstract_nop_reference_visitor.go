package reference

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"

type AbstractNopReferenceVisitor struct {
}

func (e *AbstractNopReferenceVisitor) VisitAnnotationElementValue(reference intsyn.IAnnotationElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReference(reference intsyn.IAnnotationReference) {
}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReferences(references intsyn.IAnnotationReferences) {
}
func (e *AbstractNopReferenceVisitor) VisitElementValueArrayInitializerElementValue(reference intsyn.IElementValueArrayInitializerElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitElementValues(references intsyn.IElementValues)         {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePair(reference intsyn.IElementValuePair)    {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePairs(references intsyn.IElementValuePairs) {}
func (e *AbstractNopReferenceVisitor) VisitExpressionElementValue(reference intsyn.IExpressionElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitInnerObjectReference(reference intsyn.IInnerObjectReference) {
}
func (e *AbstractNopReferenceVisitor) VisitObjectReference(reference intsyn.IObjectReference) {}
