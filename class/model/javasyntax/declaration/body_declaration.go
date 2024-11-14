package declaration

func NewBodyDeclaration(internalTypeName string, memberDeclaration IMemberDeclaration) *BodyDeclaration {
	return &BodyDeclaration{
		internalTypeName:   internalTypeName,
		memberDeclarations: memberDeclaration,
	}
}

type BodyDeclaration struct {
	internalTypeName   string
	memberDeclarations IMemberDeclaration
}

func (d *BodyDeclaration) InternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) MemberDeclarations() IMemberDeclaration {
	return d.memberDeclarations
}

func (d *BodyDeclaration) SetMemberDeclarations(memberDeclaration IMemberDeclaration) {
	d.memberDeclarations = memberDeclaration
}

func (d *BodyDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
