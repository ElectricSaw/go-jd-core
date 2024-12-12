package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewElementValuePair(elementName string, elementValue intcls.IElementValue) intcls.IElementValuePair {
	return &ElementValuePair{
		elementName:  elementName,
		elementValue: elementValue,
	}
}

type ElementValuePair struct {
	elementName  string
	elementValue intcls.IElementValue
}

func (e *ElementValuePair) Name() string {
	return e.elementName
}

func (e *ElementValuePair) SetName(name string) {
	e.elementName = name
}

func (e *ElementValuePair) Value() intcls.IElementValue {
	return e.elementValue
}

func (e *ElementValuePair) SetValue(value intcls.IElementValue) {
	e.elementValue = value
}
