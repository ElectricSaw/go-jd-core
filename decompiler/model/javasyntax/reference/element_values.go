package reference

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewElementValues() intmod.IElementValues {
	return &ElementValues{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IElementValue](0).(*util.DefaultList[intmod.IElementValue]),
	}
}

type ElementValues struct {
	util.DefaultList[intmod.IElementValue]
}

func (e *ElementValues) Name() string {
	return ""
}

func (e *ElementValues) ElementValue() intmod.IElementValue {
	return nil
}

func (e *ElementValues) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitElementValues(e)
}

func (e *ElementValues) String() string {
	return fmt.Sprintf("ElementValues{%s}", *e)
}
