package reference

type AnnotationReferences struct {
	IReference

	annotationReferences []IAnnotationReference
}

func (r *AnnotationReferences) Accept(visitor ReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
