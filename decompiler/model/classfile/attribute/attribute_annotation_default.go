package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewAttributeAnnotationDefault(defaultValue intcls.IElementValue) intcls.IAttributeAnnotationDefault {
	return &AttributeAnnotationDefault{defaultValue: defaultValue}
}

type AttributeAnnotationDefault struct {
	defaultValue intcls.IElementValue
}

func (a AttributeAnnotationDefault) DefaultValue() intcls.IElementValue {
	return a.defaultValue
}

func (a AttributeAnnotationDefault) IsAttribute() bool {
	return true
}
