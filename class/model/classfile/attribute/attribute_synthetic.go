package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

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
