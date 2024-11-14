package reference

type AnnotationReferences struct {
	IAnnotationReference

	AnnotationReferences []IAnnotationReference
}

func (r *AnnotationReferences) List() []IReference {
	ret := make([]IReference, 0, len(r.AnnotationReferences))
	for _, v := range r.AnnotationReferences {
		ret = append(ret, v)
	}
	return ret
}

func (r *AnnotationReferences) Accept(visitor ReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}
