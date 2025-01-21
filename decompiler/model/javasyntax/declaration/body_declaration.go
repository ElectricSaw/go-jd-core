package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewBodyDeclaration(internalTypeName string, memberDeclaration MemberDeclaration) BodyDeclaration {
	d := &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
	return d
}

type BodyDeclaration struct {
	internalTypeName   string
	memberDeclarations MemberDeclaration
}

func (d *BodyDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) MemberDeclarations() MemberDeclaration {
	return d.memberDeclarations
}

func (d *BodyDeclaration) SetMemberDeclarations(memberDeclaration MemberDeclaration) {
	d.memberDeclarations = memberDeclaration
}

func (d *BodyDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
