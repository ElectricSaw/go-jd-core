package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewElementValueArrayInitializerElementValue(elementValueArrayInitializer intmod.IElementValue) intmod.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValues(elementValueArrayInitializer intmod.IElementValues) intmod.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValueEmpty() intmod.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{elementValueArrayInitializer: nil}
}

type ElementValueArrayInitializerElementValue struct {
	elementValueArrayInitializer intmod.IElementValue
}

func (e *ElementValueArrayInitializerElementValue) ElementValueArrayInitializer() intmod.IElementValue {
	return e.elementValueArrayInitializer
}

func (e *ElementValueArrayInitializerElementValue) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitElementValueArrayInitializerElementValue(e)
}

func (e *ElementValueArrayInitializerElementValue) String() string {
	return fmt.Sprintf("ElementValueArrayInitializerElementValue{%s}", e.elementValueArrayInitializer)
}
