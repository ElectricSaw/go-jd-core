package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewMemberDeclarations() intmod.IMemberDeclarations {
	return NewMemberDeclarationsWithCapacity(0)
}

func NewMemberDeclarationsWithCapacity(capacity int) intmod.IMemberDeclarations {
	return &MemberDeclarations{
		DefaultList: *util.NewDefaultListWithCapacity[intmod.IMemberDeclaration](capacity),
	}
}

type MemberDeclarations struct {
	AbstractMemberDeclaration
	util.DefaultList[intmod.IMemberDeclaration]
}

func (d *MemberDeclarations) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
