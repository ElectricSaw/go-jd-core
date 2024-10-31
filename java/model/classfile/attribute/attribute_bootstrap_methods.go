package attribute

func NewAttributeBootstrapMethods(bootstrapMethod []BootstrapMethod) *AttributeBootstrapMethods {
	return &AttributeBootstrapMethods{bootstrapMethod: bootstrapMethod}
}

type AttributeBootstrapMethods struct {
	bootstrapMethod []BootstrapMethod
}

func (a AttributeBootstrapMethods) BootstrapMethods() []BootstrapMethod {
	return a.bootstrapMethod
}

func (a AttributeBootstrapMethods) attributeIgnoreFunc() {}
