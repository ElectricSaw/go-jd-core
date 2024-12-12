package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeDeprecated() intcls.IAttributeDeprecated {
	return &AttributeDeprecated{}
}

type AttributeDeprecated struct {
}

func (a *AttributeDeprecated) IsAttributeDeprecated() bool {
	return true
}

func (a AttributeDeprecated) IsAttribute() bool {
	return true
}
