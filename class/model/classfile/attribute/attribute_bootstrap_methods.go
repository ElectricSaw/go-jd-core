package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeBootstrapMethods(bootstrapMethod []intcls.IBootstrapMethod) intcls.IAttributeBootstrapMethods {
	return &AttributeBootstrapMethods{bootstrapMethod: bootstrapMethod}
}

type AttributeBootstrapMethods struct {
	bootstrapMethod []intcls.IBootstrapMethod
}

func (a AttributeBootstrapMethods) BootstrapMethods() []intcls.IBootstrapMethod {
	return a.bootstrapMethod
}

func (a AttributeBootstrapMethods) IsAttribute() bool {
	return true
}
