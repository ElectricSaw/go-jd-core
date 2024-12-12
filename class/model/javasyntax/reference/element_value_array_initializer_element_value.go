package reference

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
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
