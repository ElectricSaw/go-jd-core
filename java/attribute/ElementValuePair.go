package attribute

func NewElementValuePair(elementName string, elementValue ElementValue) ElementValuePair {
	return ElementValuePair{
		ElementName:  elementName,
		ElementValue: elementValue,
	}
}

type ElementValuePair struct {
	ElementName  string
	ElementValue ElementValue
}
