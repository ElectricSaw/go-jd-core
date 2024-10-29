package constant

func NewConstantMethodHandle(referenceKind int, referenceIndex int) ConstantMethodHandle {
	return ConstantMethodHandle{
		tag:            ConstTagMethodHandle,
		referenceKind:  referenceKind,
		referenceIndex: referenceIndex,
	}
}

type ConstantMethodHandle struct {
	tag            ACC
	referenceKind  int
	referenceIndex int
}

func (c *ConstantMethodHandle) Tag() ACC {
	return c.tag
}

func (c *ConstantMethodHandle) ReferenceKind() int {
	return c.referenceKind
}

func (c *ConstantMethodHandle) ReferenceIndex() int {
	return c.referenceIndex
}
