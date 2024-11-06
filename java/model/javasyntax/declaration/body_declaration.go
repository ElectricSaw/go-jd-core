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

func (d *BodyDeclaration) GetInternalTypeName() string {
	return d.internalTypeName
}

func (d *BodyDeclaration) GetMemberDeclaration() IMemberDeclaration {
	return d.memberDeclaration
}

func (d *BodyDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitBodyDeclaration(d)
}
