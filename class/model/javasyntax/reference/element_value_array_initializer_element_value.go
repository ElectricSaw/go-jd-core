package reference

import "fmt"

func NewElementValueArrayInitializerElementValue(elementValueArrayInitializer IElementValue) *ElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValueList(elementValueArrayInitializer IElementValuePair) *ElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValueEmpty() *ElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{elementValueArrayInitializer: nil}
}

type ElementValueArrayInitializerElementValue struct {
	IElementValue

	elementValueArrayInitializer IElementValue
}

func (e *ElementValueArrayInitializerElementValue) ElementValueArrayInitializer() IElementValue {
	return e.elementValueArrayInitializer
}

func (e *ElementValueArrayInitializerElementValue) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValueArrayInitializerElementValue(e)
}

func (e *ElementValueArrayInitializerElementValue) String() string {
	return fmt.Sprintf("ElementValueArrayInitializerElementValue{%s}", e.elementValueArrayInitializer)
}
