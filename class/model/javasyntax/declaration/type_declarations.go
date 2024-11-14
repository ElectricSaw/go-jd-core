package declaration

type TypeDeclarations struct {
	AbstractTypeDeclaration

	TypeDeclaration []TypeDeclaration
}

func (d *TypeDeclarations) List() []Declaration {
	ret := make([]Declaration, 0, len(d.TypeDeclaration))

	for _, m := range d.TypeDeclaration {
		ret = append(ret, m)
	}

	return ret
}

func (d *TypeDeclarations) Accept(visitor DeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
