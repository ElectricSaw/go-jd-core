package constant

func NewConstantInteger(value int32) ConstantInteger {
	return ConstantInteger{
		tag:   ConstTagInteger,
		value: value,
	}
}

type ConstantInteger struct {
	tag   TAG
	value int32
}

func (c ConstantInteger) Tag() TAG {
	return c.tag
}

func (c ConstantInteger) Value() int32 {
	return c.value
}

func (c ConstantInteger) constantValueIgnoreFunc() {}
