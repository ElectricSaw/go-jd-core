package attribute

import (
	"bitbucket.org/coontec/javaClass/java/model/classfile/constant"
)

func NewAttributeConstantValue(constantValue constant.ConstantValue) AttributeConstantValue {
	return AttributeConstantValue{constantValue}
}

type AttributeConstantValue struct {
	constantValue constant.ConstantValue
}

func (a AttributeConstantValue) GetValue() constant.ConstantValue {
	return a.constantValue
}

func (a AttributeConstantValue) attributeIgnoreFunc() {}
