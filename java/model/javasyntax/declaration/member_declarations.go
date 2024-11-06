package declaration

type MemberDeclarations struct {
	AbstractMemberDeclaration

	MemberDeclarations []IMemberDeclaration
}

func (d *MemberDeclarations) Accept(visitor DeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
