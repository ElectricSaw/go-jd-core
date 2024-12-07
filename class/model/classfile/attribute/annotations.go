package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAnnotations(annotations []intcls.IAnnotation) intcls.IAnnotations {
	return &Annotations{
		annotations: annotations,
	}
}

type Annotations struct {
	annotations []intcls.IAnnotation
}

func (a Annotations) Annotations() []intcls.IAnnotation {
	return a.annotations
}

func (a *Annotations) IsAttribute() bool {
	return true
}
