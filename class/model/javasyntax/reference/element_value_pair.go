package reference

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewElementValuePair(name string, elementValue intmod.IElementValue) intmod.IElementValuePair {
	p := &ElementValuePair{
		name:         name,
		elementValue: elementValue,
	}
	p.SetValue(p)
	return p
}

type ElementValuePair struct {
	util.DefaultBase[intmod.IElementValuePair]

	name         string
	elementValue intmod.IElementValue
}

func (e *ElementValuePair) Name() string {
	return e.name
}

func (e *ElementValuePair) ElementValue() intmod.IElementValue {
	return e.elementValue
}

func (e *ElementValuePair) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitElementValuePair(e)
}

func (e *ElementValuePair) String() string {
	return fmt.Sprintf("ElementValuePair{name=%s, elementValue=%s}", e.name, e.elementValue)
}
