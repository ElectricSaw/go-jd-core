package constant

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

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
