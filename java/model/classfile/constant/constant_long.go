package constant

func NewConstantLong(value int64) ConstantLong {
	return ConstantLong{
		tag:   ConstTagLong,
		value: value,
	}
}

type ConstantLong struct {
	tag   TAG
	value int64
}

func (c ConstantLong) Tag() TAG {
	return c.tag
}

func (c ConstantLong) Value() int64 {
	return c.value
}

func (c ConstantLong) constantValueIgnoreFunc() {}
