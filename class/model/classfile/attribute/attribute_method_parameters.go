package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeMethodParameters(parameters []intcls.IMethodParameter) intcls.IAttributeMethodParameters {
	return &AttributeMethodParameters{parameters}
}

type AttributeMethodParameters struct {
	parameters []intcls.IMethodParameter
}

func (a AttributeMethodParameters) Parameters() []intcls.IMethodParameter {
	return a.parameters
}

func (a AttributeMethodParameters) IsAttribute() bool {
	return true
}
