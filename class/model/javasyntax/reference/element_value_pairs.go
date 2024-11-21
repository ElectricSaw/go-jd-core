package reference

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewElementValuePairs() intmod.IElementValuePairs {
	return &ElementValuePairs{
		DefaultList: *util.NewDefaultList[intmod.IElementValuePair](0),
	}
}

type ElementValuePairs struct {
	util.DefaultList[intmod.IElementValuePair]
}

func (e *ElementValuePairs) Name() string {
	return ""
}

func (e *ElementValuePairs) ElementValue() intmod.IElementValue {
	return nil
}

func (e *ElementValuePairs) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitElementValuePairs(e)
}

func (e *ElementValuePairs) String() string {
	return fmt.Sprintf("ElementValuePairs{%v}", *e)
}
