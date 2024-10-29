package constant

func NewConstantNameAndType(nameIndex int, descriptorIndex int) ConstantNameAndType {
	return ConstantNameAndType{
		tag:             ConstTagNameAndType,
		nameIndex:       nameIndex,
		descriptorIndex: descriptorIndex,
	}
}

type ConstantNameAndType struct {
	tag             TAG
	nameIndex       int
	descriptorIndex int
}

func (c ConstantNameAndType) Tag() TAG {
	return c.tag
}

func (c ConstantNameAndType) NameIndex() int {
	return c.nameIndex
}

func (c ConstantNameAndType) DescriptorIndex() int {
	return c.descriptorIndex
}
