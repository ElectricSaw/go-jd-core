package constant

func NewConstantLong(value int64) ConstantLong {
	return ConstantLong{
		tag:   ConstTagLong,
		value: value,
	}
}

type ConstantLong struct {
	tag   ACC
	value int64
}

func (c *ConstantLong) Tag() ACC {
	return c.tag
}

func (c *ConstantLong) Value() int64 {
	return c.value
}

func (c *ConstantLong) constantValueIgnoreFunc() {}
