package constant

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewConstantMethodHandle(referenceKind int, referenceIndex int) intcls.IConstantMethodHandle {
	return &ConstantMethodHandle{
		tag:            intcls.ConstTagMethodHandle,
		referenceKind:  referenceKind,
		referenceIndex: referenceIndex,
	}
}

type ConstantMethodHandle struct {
	tag            intcls.TAG
	referenceKind  int
	referenceIndex int
}

func (c ConstantMethodHandle) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantMethodHandle) ReferenceKind() int {
	return c.referenceKind
}

func (c ConstantMethodHandle) ReferenceIndex() int {
	return c.referenceIndex
}
