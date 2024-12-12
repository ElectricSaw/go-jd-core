package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewElementValueAnnotationValue(annotationValue intcls.IAnnotation) intcls.IElementValueAnnotationValue {
	return &ElementValueAnnotationValue{
		annotationValue: annotationValue,
	}
}

type ElementValueAnnotationValue struct {
	annotationValue intcls.IAnnotation
}

func (e *ElementValueAnnotationValue) AnnotationValue() intcls.IAnnotation {
	return e.annotationValue
}

func (e *ElementValueAnnotationValue) Accept(visitor intcls.IElementValueVisitor) {
	visitor.VisitAnnotationValue(e)
}
