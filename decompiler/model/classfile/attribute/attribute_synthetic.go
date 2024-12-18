package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewAttributeSynthetic() intcls.IAttributeSynthetic {
	return &AttributeSynthetic{}
}

type AttributeSynthetic struct {
}

func (a AttributeSynthetic) IsAttributeSynthetic() bool {
	return true
}

func (a AttributeSynthetic) IsAttribute() bool {
	return true
}
