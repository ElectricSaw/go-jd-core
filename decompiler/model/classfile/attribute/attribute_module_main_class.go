package attribute

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
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
