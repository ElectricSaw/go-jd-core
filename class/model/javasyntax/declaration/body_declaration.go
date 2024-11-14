package declaration

func NewBodyDeclaration(internalTypeName string, memberDeclaration IMemberDeclaration) *BodyDeclaration {
	return &BodyDeclaration{
		internalTypeName:  internalTypeName,
		memberDeclaration: memberDeclaration,
	}
}

type BodyDeclaration struct {
	internalTypeName  string
	memberDeclaration IMemberDeclaration
}

func (d *BodyDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) MemberDeclaration() IMemberDeclaration {
	return d.memberDeclaration
}

func (d *BodyDeclaration) SetMemberDeclaration(memberDeclaration IMemberDeclaration) {
	d.memberDeclaration = memberDeclaration
}

func (d *BodyDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
