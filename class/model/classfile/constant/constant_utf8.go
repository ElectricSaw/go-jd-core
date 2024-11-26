package constant

func NewConstantUtf8(value string) *ConstantUtf8 {
	return &ConstantUtf8{
		tag:   ConstTagUtf8,
		value: value,
	}
}

type ConstantUtf8 struct {
	tag   TAG
	value string
}

func (c ConstantUtf8) Tag() TAG {
	return c.tag
}

func (c ConstantUtf8) Value() string {
	return c.value
}

func (c ConstantUtf8) constantValueIgnoreFunc() {}
