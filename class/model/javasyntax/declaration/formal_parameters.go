package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewFormalParameters() intsyn.IFormalParameters {
	return &FormalParameters{}
}

type FormalParameters struct {
	util.DefaultList[intsyn.IFormalParameter]
}

func (d *FormalParameters) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitFormalParameters(d)
}
