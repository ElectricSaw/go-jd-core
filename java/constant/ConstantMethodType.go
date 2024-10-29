package constant

func NewConstantMethodType(descriptorIndex int) ConstantMethodType {
	return ConstantMethodType{
		tag:             ConstTagMethodType,
		descriptorIndex: descriptorIndex,
	}
}

type ConstantMethodType struct {
	tag             ACC
	descriptorIndex int
}

func (c *ConstantMethodType) Tag() ACC {
	return c.tag
}

func (c *ConstantMethodType) DescriptorIndex() int {
	return c.descriptorIndex
}
