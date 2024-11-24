package constant

func NewConstantMemberRef(classIndex int, nameAndTypeIndex int) *ConstantMemberRef {
	return &ConstantMemberRef{
		tag:              ConstTagMemberRef,
		classIndex:       classIndex,
		nameAndTypeIndex: nameAndTypeIndex,
	}
}

type ConstantMemberRef struct {
	tag              TAG
	classIndex       int
	nameAndTypeIndex int
}

func (c ConstantMemberRef) Tag() TAG {
	return c.tag
}

func (c ConstantMemberRef) ClassIndex() int {
	return c.classIndex
}

func (c ConstantMemberRef) NameAndTypeIndex() int {
	return c.nameAndTypeIndex
}
