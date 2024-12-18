package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantMethodType(descriptorIndex int) intcls.IConstantMethodType {
	return &ConstantMethodType{
		tag:             intcls.ConstTagMethodType,
		descriptorIndex: descriptorIndex,
	}
}

type ConstantMethodType struct {
	tag             intcls.TAG
	descriptorIndex int
}

func (c ConstantMethodType) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantMethodType) DescriptorIndex() int {
	return c.descriptorIndex
}
