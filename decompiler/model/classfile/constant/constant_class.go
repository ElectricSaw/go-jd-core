package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantClass(nameIndex int) intcls.IConstantClass {
	return &ConstantClass{
		tag:       intcls.ConstTagClass,
		nameIndex: nameIndex,
	}
}

type ConstantClass struct {
	tag       intcls.TAG
	nameIndex int
}

func (c ConstantClass) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantClass) NameIndex() int {
	return c.nameIndex
}
