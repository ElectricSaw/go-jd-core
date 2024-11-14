package declaration

type FieldDeclarators struct {
	FieldDeclarators []FieldDeclarator
}

func (d *FieldDeclarators) List() []Declaration {
	ret := make([]Declaration, 0, len(d.FieldDeclarators))
	for _, f := range d.FieldDeclarators {
		ret = append(ret, &f)
	}
	return ret
}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration *FieldDeclaration) {
	for _, declarator := range d.FieldDeclarators {
		declarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) Accept(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}
