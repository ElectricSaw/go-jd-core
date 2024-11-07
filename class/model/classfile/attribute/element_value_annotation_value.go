package attribute

func NewElementValueAnnotationValue(annotationValue Annotation) *ElementValueAnnotationValue {
	return &ElementValueAnnotationValue{
		annotationValue: annotationValue,
	}
}

type ElementValueAnnotationValue struct {
	annotationValue Annotation
}

func (e *ElementValueAnnotationValue) AnnotationValue() Annotation {
	return e.annotationValue
}

func (e *ElementValueAnnotationValue) Accept(visitor ElementValueVisitor) {
	visitor.VisitAnnotationValue(e)
}
