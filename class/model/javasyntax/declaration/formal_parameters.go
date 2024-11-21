package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewFormalParameters() intmod.IFormalParameters {
	return &FormalParameters{}
}

type FormalParameters struct {
	FormalParameter
	util.DefaultList[intmod.IFormalParameter]
}

func (d *FormalParameters) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}
