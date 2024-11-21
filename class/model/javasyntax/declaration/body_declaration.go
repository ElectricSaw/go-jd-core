package declaration

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewBodyDeclaration(internalTypeName string, memberDeclaration intmod.IMemberDeclaration) intmod.IBodyDeclaration {
	return &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
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

func (d *BodyDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
