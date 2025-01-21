package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFormalParameters() FormalParameters {
	return NewFormalParametersWithCapacity(0)
}

func NewFormalParametersWithCapacity(capacity int) FormalParameters {
	return &FormalParameters{
		DefaultList: *util.NewDefaultListWithCapacity[FormalParameter](capacity).(*util.DefaultList[FormalParameter]),
	}
}

type FormalParameters struct {
	util.DefaultList[FormalParameter]
}

func (d *FormalParameters) AnnotationReferences() AnnotationReference {
	return nil
}

func (d *FormalParameters) IsFinal() bool {
	return false
}

func (d *FormalParameters) SetFinal(final bool) {}

func (d *FormalParameters) Type() Type {
	return nil
}

func (d *FormalParameters) IsVarargs() bool {
	return false
}

func (d *FormalParameters) Name() string {
	return ""
}

func (d *FormalParameters) SetName(name string) {}

func (d *FormalParameters) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}

func (d *FormalParameters) String() string {
	return "FormalParameters{}"
}
