package constant

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewConstantDouble(value float64) intcls.IConstantDouble {
	return &ConstantDouble{
		tag:   intcls.ConstTagDouble,
		value: value,
	}
}

type ConstantDouble struct {
	tag   intcls.TAG
	value float64
}

func (c ConstantDouble) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantDouble) Value() float64 {
	return c.value
}

func (c ConstantDouble) IsConstantValue() bool {
	return true
}
