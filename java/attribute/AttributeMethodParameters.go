package attribute

func NewAttributeMethodParameters(parameters []MethodParameter) AttributeMethodParameters {
	return AttributeMethodParameters{parameters}
}

type AttributeMethodParameters struct {
	parameters []MethodParameter
}

func (a AttributeMethodParameters) Parameters() []MethodParameter {
	return a.parameters
}

func (a AttributeMethodParameters) attributeIgnoreFunc() {}
