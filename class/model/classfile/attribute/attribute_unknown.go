package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewUnknownAttribute() intcls.IUnknownAttribute {
	return &UnknownAttribute{}
}

type UnknownAttribute struct {
}

func (a UnknownAttribute) IsUnknownAttribute() bool {
	return true
}

func (a UnknownAttribute) IsAttribute() bool {
	return true
}
