package constant

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewConstantMemberRef(classIndex int, nameAndTypeIndex int) intcls.IConstantMemberRef {
	return &ConstantMemberRef{
		tag:              intcls.ConstTagMemberRef,
		classIndex:       classIndex,
		nameAndTypeIndex: nameAndTypeIndex,
	}
}

type ConstantMemberRef struct {
	tag              intcls.TAG
	classIndex       int
	nameAndTypeIndex int
}

func (c ConstantMemberRef) Tag() intcls.TAG {
	return c.tag
}

func (c ConstantMemberRef) ClassIndex() int {
	return c.classIndex
}

func (c ConstantMemberRef) NameAndTypeIndex() int {
	return c.nameAndTypeIndex
}
