package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewAttributeInnerClasses(classes []intcls.IInnerClass) intcls.IAttributeInnerClasses {
	return &AttributeInnerClasses{classes}
}

type AttributeInnerClasses struct {
	classes []intcls.IInnerClass
}

func (a AttributeInnerClasses) InnerClasses() []intcls.IInnerClass {
	return a.classes
}

func (a AttributeInnerClasses) IsAttribute() bool {
	return true
}
