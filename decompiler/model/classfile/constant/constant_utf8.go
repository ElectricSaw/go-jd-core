package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantUtf8(value string) intcls.IConstantUtf8 {
	return &ConstantUtf8{
		tag:   intcls.ConstTagUtf8,
		value: value,
	}
}

type ConstantUtf8 struct {
	tag   intcls.TAG
	value string
}

func (c ConstantUtf8) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantUtf8) Value() string {
	return c.value
}

func (c ConstantUtf8) IsConstantValue() bool {
	return true
}
