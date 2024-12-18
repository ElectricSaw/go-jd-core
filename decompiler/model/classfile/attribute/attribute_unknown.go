package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

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
