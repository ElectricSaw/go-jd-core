package reference

import "fmt"

func NewElementValuePair(name string, elementValue IElementValue) *ElementValuePair {
	return &ElementValuePair{
		name:         name,
		elementValue: elementValue,
	}
}

type ElementValuePair struct {
	IElementValuePair

	name         string
	elementValue IElementValue
}

func (e *ElementValuePair) GetName() string {
	return e.name
}

func (e *ElementValuePair) GetElementValue() IElementValue {
	return e.elementValue
}

func (e *ElementValuePair) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValuePair(e)
}

func (e *ElementValuePair) String() string {
	return fmt.Sprintf("ElementValuePair{name=%s, elementValue=%s}", e.name, e.elementValue)
}
