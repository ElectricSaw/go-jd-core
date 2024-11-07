package constant

func NewConstantInteger(value int) *ConstantInteger {
	return &ConstantInteger{
		tag:   ConstTagInteger,
		value: value,
	}
}

type ConstantInteger struct {
	tag   TAG
	value int
}

func (c ConstantInteger) Tag() TAG {
	return c.tag
}

func (c ConstantInteger) Value() int {
	return c.value
}

func (c ConstantInteger) constantValueIgnoreFunc() {}
