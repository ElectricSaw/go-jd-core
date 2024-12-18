package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantInteger(value int) intcls.IConstantInteger {
	return &ConstantInteger{
		tag:   intcls.ConstTagInteger,
		value: value,
	}
}

type ConstantInteger struct {
	tag   intcls.TAG
	value int
}

func (c ConstantInteger) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantInteger) Value() int {
	return c.value
}

func (c ConstantInteger) IsConstantValue() bool {
	return true
}
