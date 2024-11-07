package attribute

func NewAttributeParameterAnnotations(parameterAnnotations []Annotations) *AttributeParameterAnnotations {
	return &AttributeParameterAnnotations{parameterAnnotations: parameterAnnotations}
}

type AttributeParameterAnnotations struct {
	parameterAnnotations []Annotations
}

func (a AttributeParameterAnnotations) ParameterAnnotations() []Annotations {
	return a.parameterAnnotations
}

func (a AttributeParameterAnnotations) attributeIgnoreFunc() {}
