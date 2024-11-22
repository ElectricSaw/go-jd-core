package attribute

import (
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/constant"
)

func NewAttributeConstantValue(constantValue constant.ConstantValue) *AttributeConstantValue {
	return &AttributeConstantValue{constantValue}
}

type AttributeConstantValue struct {
	constantValue constant.ConstantValue
}

func (a AttributeConstantValue) ConstantValue() constant.ConstantValue {
	return a.constantValue
}

func (a AttributeConstantValue) attributeIgnoreFunc() {}
