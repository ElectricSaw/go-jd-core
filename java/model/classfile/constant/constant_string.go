package constant

func NewConstantString(stringIndex int) *ConstantString {
	return &ConstantString{
		tag:         ConstTagString,
		stringIndex: stringIndex,
	}
}

type ConstantString struct {
	tag         TAG
	stringIndex int
}

func (c ConstantString) Tag() TAG {
	return c.tag
}

func (c ConstantString) StringIndex() int {
	return c.stringIndex
}
