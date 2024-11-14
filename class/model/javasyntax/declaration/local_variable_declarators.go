package declaration

type LocalVariableDeclarators struct {
	LocalVariableDeclarators []LocalVariableDeclarator
}

func (d *LocalVariableDeclarators) List() []Declaration {
	ret := make([]Declaration, 0, len(d.LocalVariableDeclarators))
	for _, param := range d.LocalVariableDeclarators {
		ret = append(ret, &param)
	}
	return ret
}

func (d *LocalVariableDeclarators) LineNumber() int {
	if len(d.LocalVariableDeclarators) == 0 {
		return 0
	}

	return d.LocalVariableDeclarators[0].LineNumber()
}

func (d *LocalVariableDeclarators) Accept(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclarators(d)
}
