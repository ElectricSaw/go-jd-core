package constant

func NewConstantDouble(value float64) ConstantDouble {
	return ConstantDouble{
		tag:   ConstTagDouble,
		value: value,
	}
}

type ConstantDouble struct {
	tag   ACC
	value float64
}

func (c *ConstantDouble) Tag() ACC {
	return c.tag
}

func (c *ConstantDouble) Value() float64 {
	return c.value
}

func (c *ConstantDouble) constantValueIgnoreFunc() {}
