package attribute

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
)

func NewAttributeModuleMainClass(mainClass intcls.IConstantClass) intcls.IAttributeModuleMainClass {
	return &AttributeModuleMainClass{mainClass: mainClass}
}

type AttributeModuleMainClass struct {
	mainClass intcls.IConstantClass
}

func (a AttributeModuleMainClass) MainClass() intcls.IConstantClass {
	return a.mainClass
}

func (a AttributeModuleMainClass) IsAttribute() bool {
	return true
}
