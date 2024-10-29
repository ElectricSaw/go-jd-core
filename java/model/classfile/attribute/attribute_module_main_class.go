package attribute

import (
	"bitbucket.org/coontec/javaClass/java/model/classfile/constant"
)

func NewAttributeModuleMainClass(mainClass constant.ConstantClass) AttributeModuleMainClass {
	return AttributeModuleMainClass{mainClass: mainClass}
}

type AttributeModuleMainClass struct {
	mainClass constant.ConstantClass
}

func (a AttributeModuleMainClass) MainClass() constant.ConstantClass {
	return a.mainClass
}

func (a AttributeModuleMainClass) attributeIgnoreFunc() {}
