package constant

func NewConstantDouble(value float64) *ConstantDouble {
	return &ConstantDouble{
		tag:   ConstTagDouble,
		value: value,
	}
}

type ConstantDouble struct {
	tag   TAG
	value float64
}

func (c ConstantDouble) Tag() TAG {
	return c.tag
}

func (c ConstantDouble) Value() float64 {
	return c.value
}

func (c ConstantDouble) constantValueIgnoreFunc() {}
