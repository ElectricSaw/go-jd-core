package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewMemberDeclarations() intmod.IMemberDeclarations {
	return &MemberDeclarations{
		DefaultList: *util.NewDefaultList[intmod.IMemberDeclaration](0),
	}
}

func NewMemberDeclarationsWithSize(size int) intmod.IMemberDeclarations {
	return &MemberDeclarations{
		DefaultList: *util.NewDefaultList[intmod.IMemberDeclaration](size),
	}
}

type MemberDeclarations struct {
	AbstractMemberDeclaration
	util.DefaultList[intmod.IMemberDeclaration]
}

func (d *MemberDeclarations) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
