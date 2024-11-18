package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewElementValueArrayInitializerElementValue(elementValueArrayInitializer intsyn.IElementValue) intsyn.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValueList(elementValueArrayInitializer intsyn.IElementValuePair) intsyn.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
}

func NewElementValueArrayInitializerElementValueEmpty() intsyn.IElementValueArrayInitializerElementValue {
	return &ElementValueArrayInitializerElementValue{elementValueArrayInitializer: nil}
}

type ElementValueArrayInitializerElementValue struct {
	elementValueArrayInitializer intsyn.IElementValue
}

func (e *ElementValueArrayInitializerElementValue) ElementValueArrayInitializer() intsyn.IElementValue {
	return e.elementValueArrayInitializer
}

func (e *ElementValueArrayInitializerElementValue) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitElementValueArrayInitializerElementValue(e)
}

func (e *ElementValueArrayInitializerElementValue) String() string {
	return fmt.Sprintf("ElementValueArrayInitializerElementValue{%s}", e.elementValueArrayInitializer)
}
