package constant

func NewConstantMethodType(descriptorIndex int) *ConstantMethodType {
	return &ConstantMethodType{
		tag:             ConstTagMethodType,
		descriptorIndex: descriptorIndex,
	}
}

type ConstantMethodType struct {
	tag             TAG
	descriptorIndex int
}

func (c ConstantMethodType) Tag() TAG {
	return c.tag
}

func (c ConstantMethodType) DescriptorIndex() int {
	return c.descriptorIndex
}
