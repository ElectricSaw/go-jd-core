package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewBodyDeclaration(internalTypeName string, memberDeclaration intmod.IMemberDeclaration) intmod.IBodyDeclaration {
	d := &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
	d.SetValue(d)
	return d
}

type BodyDeclaration struct {
	util.DefaultBase[intmod.IDeclaration]

	internalTypeName   string
	memberDeclarations intmod.IMemberDeclaration
}

func (d *BodyDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) MemberDeclarations() intmod.IMemberDeclaration {
	return d.memberDeclarations
}

func (d *BodyDeclaration) SetMemberDeclarations(memberDeclaration intmod.IMemberDeclaration) {
	d.memberDeclarations = memberDeclaration
}

func (d *BodyDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
