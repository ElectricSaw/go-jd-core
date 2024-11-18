package reference

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
	"fmt"
)

func NewElementValues() intsyn.IElementValues {
	return &ElementValues{}
}

type ElementValues struct {
	util.DefaultList[intsyn.IElementValue]
}

func (e *ElementValues) Name() string {
	return ""
}

func (e *ElementValues) ElementValue() intsyn.IElementValue {
	return nil
}

func (e *ElementValues) Accept(visitor intsyn.IReferenceVisitor) {
	visitor.VisitElementValues(e)
}

func (e *ElementValues) String() string {
	return fmt.Sprintf("ElementValues{%s}", *e)
}
