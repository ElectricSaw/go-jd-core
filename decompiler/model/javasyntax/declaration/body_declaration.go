package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewBodyDeclaration(internalTypeName string, memberDeclaration intmod.IMemberDeclaration) intmod.IBodyDeclaration {
	d := &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
	return d
}

type BodyDeclaration struct {
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

func (d *BodyDeclaration) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
