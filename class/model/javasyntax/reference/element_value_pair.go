package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewElementValuePair(name string, elementValue intmod.IElementValue) intmod.IElementValuePair {
	return &ElementValuePair{
		name:         name,
		elementValue: elementValue,
	}
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
