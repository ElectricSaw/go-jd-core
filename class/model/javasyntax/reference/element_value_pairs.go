package reference

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewElementValuePairs() intmod.IElementValuePairs {
	return &ElementValuePairs{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IElementValuePair](0).(*util.DefaultList[intmod.IElementValuePair]),
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
