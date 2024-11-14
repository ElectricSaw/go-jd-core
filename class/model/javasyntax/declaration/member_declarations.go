package declaration

type MemberDeclarations struct {
	AbstractMemberDeclaration

	MemberDeclarations []IMemberDeclaration
}

func (d *MemberDeclarations) List() []Declaration {
	ret := make([]Declaration, 0, len(d.MemberDeclarations))

	for _, m := range d.MemberDeclarations {
		ret = append(ret, m)
	}

	return ret
}

func (d *MemberDeclarations) Accept(visitor DeclarationVisitor) {
	visitor.VisitMemberDeclarations(d)
}
