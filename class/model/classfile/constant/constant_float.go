package constant

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewConstantFloat(value float32) intcls.IConstantFloat {
	return &ConstantFloat{
		tag:   intcls.ConstTagFloat,
		value: value,
	}
}

type ConstantFloat struct {
	tag   intcls.TAG
	value float32
}

func (c ConstantFloat) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantFloat) Value() float32 {
	return c.value
}

func (c ConstantFloat) IsConstantValue() bool {
	return true
}
