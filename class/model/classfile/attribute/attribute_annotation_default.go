package attribute

func NewAttributeAnnotationDefault(defaultValue ElementValue) *AttributeAnnotationDefault {
	return &AttributeAnnotationDefault{defaultValue: defaultValue}
}

type AttributeAnnotationDefault struct {
	defaultValue ElementValue
}

func (a AttributeAnnotationDefault) DefaultValue() ElementValue {
	return a.defaultValue
}

func (a AttributeAnnotationDefault) attributeIgnoreFunc() {}
