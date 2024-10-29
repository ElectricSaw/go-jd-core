package attribute

func NewAttributeInnerClasses(classes []InnerClass) AttributeInnerClasses {
	return AttributeInnerClasses{classes}
}

type AttributeInnerClasses struct {
	classes []InnerClass
}

func (a AttributeInnerClasses) InnerClasses() []InnerClass {
	return a.classes
}

func (a AttributeInnerClasses) attributeIgnoreFunc() {}
