package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewMemberDeclarations() intsyn.IMemberDeclarations {
	return &MemberDeclarations{}
}

type MemberDeclarations struct {
	AbstractMemberDeclaration
	util.DefaultList[intsyn.IMemberDeclaration]
}

func (d *MemberDeclarations) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
