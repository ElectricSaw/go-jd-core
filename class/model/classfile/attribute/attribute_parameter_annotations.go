package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeParameterAnnotations(parameterAnnotations []intcls.IAnnotations) intcls.IAttributeParameterAnnotations {
	return &AttributeParameterAnnotations{parameterAnnotations: parameterAnnotations}
}

type AttributeParameterAnnotations struct {
	parameterAnnotations []intcls.IAnnotations
}

func (a AttributeParameterAnnotations) ParameterAnnotations() []intcls.IAnnotations {
	return a.parameterAnnotations
}

func (a AttributeParameterAnnotations) IsAttribute() bool {
	return true
}
