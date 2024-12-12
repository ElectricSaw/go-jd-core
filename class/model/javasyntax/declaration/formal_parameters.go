package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewFormalParameters() intmod.IFormalParameters {
	return NewFormalParametersWithCapacity(0)
}

func NewFormalParametersWithCapacity(capacity int) intmod.IFormalParameters {
	return &FormalParameters{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IFormalParameter](capacity).(*util.DefaultList[intmod.IFormalParameter]),
	}
}

type FormalParameters struct {
	util.DefaultList[intmod.IFormalParameter]
}

func (d *FormalParameters) AnnotationReferences() intmod.IAnnotationReference {
	return nil
}

func (d *FormalParameters) IsFinal() bool {
	return false
}

func (d *FormalParameters) SetFinal(final bool) {}

func (d *FormalParameters) Type() intmod.IType {
	return nil
}

func (d *FormalParameters) IsVarargs() bool {
	return false
}

func (d *FormalParameters) Name() string {
	return ""
}

func (d *FormalParameters) SetName(name string) {}

func (d *FormalParameters) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}

func (d *FormalParameters) String() string {
	return "FormalParameters{}"
}
