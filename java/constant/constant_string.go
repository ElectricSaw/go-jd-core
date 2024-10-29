package constant

func NewConstantString(stringIndex int) ConstantString {
	return ConstantString{
		tag:         ConstTagString,
		stringIndex: stringIndex,
	}
}

type ConstantString struct {
	tag         ACC
	stringIndex int
}

func (c *ConstantString) Tag() ACC {
	return c.tag
}

func (c *ConstantString) StringIndex() int {
	return c.stringIndex
}
