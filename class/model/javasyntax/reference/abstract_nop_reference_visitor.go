package reference

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

type AbstractNopReferenceVisitor struct {
}

func (e *AbstractNopReferenceVisitor) VisitAnnotationElementValue(reference intmod.IAnnotationElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReference(reference intmod.IAnnotationReference) {
}
func (e *AbstractNopReferenceVisitor) VisitAnnotationReferences(references intmod.IAnnotationReferences) {
}
func (e *AbstractNopReferenceVisitor) VisitElementValueArrayInitializerElementValue(reference intmod.IElementValueArrayInitializerElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitElementValues(references intmod.IElementValues)         {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePair(reference intmod.IElementValuePair)    {}
func (e *AbstractNopReferenceVisitor) VisitElementValuePairs(references intmod.IElementValuePairs) {}
func (e *AbstractNopReferenceVisitor) VisitExpressionElementValue(reference intmod.IExpressionElementValue) {
}
func (e *AbstractNopReferenceVisitor) VisitInnerObjectReference(reference intmod.IInnerObjectReference) {
}
func (e *AbstractNopReferenceVisitor) VisitObjectReference(reference intmod.IObjectReference) {}
