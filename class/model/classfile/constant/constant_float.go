package constant

func NewConstantFloat(value float32) *ConstantFloat {
	return &ConstantFloat{
		tag:   ConstTagFloat,
		value: value,
	}
}

type ConstantFloat struct {
	tag   TAG
	value float32
}

func (c ConstantFloat) Tag() TAG {
	return c.tag
}

func (c ConstantFloat) Value() float32 {
	return c.value
}

func (c ConstantFloat) constantValueIgnoreFunc() {}
