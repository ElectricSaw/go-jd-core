package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
	"fmt"
)

func NewElementValuePair(name string, elementValue intsyn.IElementValue) intsyn.IElementValuePair {
	return &ElementValuePair{
		name:         name,
		elementValue: elementValue,
	}
}

type ElementValuePair struct {
	util.DefaultBase[intsyn.IElementValuePair]

	name         string
	elementValue intsyn.IElementValue
}

func (e *ElementValuePair) Name() string {
	return e.name
}

func (e *ElementValuePair) ElementValue() intsyn.IElementValue {
	return e.elementValue
}

func (e *ElementValuePair) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitElementValuePair(e)
}

func (e *ElementValuePair) String() string {
	return fmt.Sprintf("ElementValuePair{name=%s, elementValue=%s}", e.name, e.elementValue)
}
