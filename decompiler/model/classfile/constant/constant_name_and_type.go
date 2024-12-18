package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantNameAndType(nameIndex int, descriptorIndex int) intcls.IConstantNameAndType {
	return &ConstantNameAndType{
		tag:             intcls.ConstTagNameAndType,
		nameIndex:       nameIndex,
		descriptorIndex: descriptorIndex,
	}
}

type ConstantNameAndType struct {
	tag             intcls.TAG
	nameIndex       int
	descriptorIndex int
}

func (c ConstantNameAndType) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantNameAndType) NameIndex() int {
	return c.nameIndex
}

func (c ConstantNameAndType) DescriptorIndex() int {
	return c.descriptorIndex
}
