package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewElementValueArrayInitializerElementValue(elementValueArrayInitializer intmod.IElementValue) intmod.IElementValueArrayInitializerElementValue {
	v := &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
	v.SetValue(v)
	return v
}

func NewElementValueArrayInitializerElementValues(elementValueArrayInitializer intmod.IElementValues) intmod.IElementValueArrayInitializerElementValue {
	v := &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: elementValueArrayInitializer,
	}
	v.SetValue(v)
	return v
}

func NewElementValueArrayInitializerElementValueEmpty() intmod.IElementValueArrayInitializerElementValue {
	v := &ElementValueArrayInitializerElementValue{
		elementValueArrayInitializer: nil,
	}
	v.SetValue(v)
	return v
}

type ElementValueArrayInitializerElementValue struct {
	util.DefaultBase[intmod.IElementValueArrayInitializerElementValue]

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
