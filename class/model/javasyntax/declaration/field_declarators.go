package declaration

type FieldDeclarators struct {
	FieldDeclarators []FieldDeclarator
}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration *FieldDeclaration) {
	for _, declarator := range d.FieldDeclarators {
		declarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) Accept(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}
