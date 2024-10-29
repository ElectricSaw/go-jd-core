package constant

func NewConstantClass(nameIndex int) ConstantClass {
	return ConstantClass{
		tag:       ConstTagClass,
		nameIndex: nameIndex,
	}
}

type ConstantClass struct {
	tag       ACC
	nameIndex int
}

func (c *ConstantClass) Tag() ACC {
	return c.tag
}

func (c *ConstantClass) NameIndex() int {
	return c.nameIndex
}
