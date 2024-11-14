package attribute

func NewAnnotations(annotations []Annotation) *Annotations {
	return &Annotations{
		annotations: annotations,
	}
}

type Annotations struct {
	annotations []Annotation
}

func (a Annotations) Annotations() []Annotation {
	return a.annotations
}

func (a *Annotations) attributeIgnoreFunc() {}
