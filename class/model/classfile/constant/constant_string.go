package constant

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewConstantString(stringIndex int) intcls.IConstantString {
	return &ConstantString{
		tag:         intcls.ConstTagString,
		stringIndex: stringIndex,
	}
}

type ConstantString struct {
	tag         intcls.TAG
	stringIndex int
}

func (c ConstantString) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantString) StringIndex() int {
	return c.stringIndex
}
