package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewAttributeConstantValue(constantValue intcls.IConstantValue) intcls.IAttributeConstantValue {
	return &AttributeConstantValue{constantValue}
}

type AttributeConstantValue struct {
	constantValue intcls.IConstantValue
}

func (a AttributeConstantValue) ConstantValue() intcls.IConstantValue {
	return a.constantValue
}

func (a AttributeConstantValue) IsAttribute() bool {
	return true
}
