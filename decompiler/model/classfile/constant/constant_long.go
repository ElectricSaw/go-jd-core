package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantLong(value int64) intcls.IConstantLong {
	return &ConstantLong{
		tag:   intcls.ConstTagLong,
		value: value,
	}
}

type ConstantLong struct {
	tag   intcls.TAG
	value int64
}

func (c ConstantLong) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantLong) Value() int64 {
	return c.value
}

func (c ConstantLong) IsConstantValue() bool {
	return true
}
