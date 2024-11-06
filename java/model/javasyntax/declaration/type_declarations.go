package declaration

type TypeDeclarations struct {
	AbstractTypeDeclaration
}

func (d *TypeDeclarations) Accept(visitor DeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
