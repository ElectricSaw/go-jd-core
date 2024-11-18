package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
	"fmt"
)

func NewElementValuePairs() intsyn.IElementValuePairs {
	return &ElementValuePairs{}
}

type ElementValuePairs struct {
	util.DefaultList[intsyn.IElementValuePair]
}

func (e *ElementValuePairs) Name() string {
	return ""
}

func (e *ElementValuePairs) ElementValue() intsyn.IElementValue {
	return nil
}

func (e *ElementValuePairs) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitElementValuePairs(e)
}

func (e *ElementValuePairs) String() string {
	return fmt.Sprintf("ElementValuePairs{%v}", *e)
}
