package declaration

type LocalVariableDeclarators struct {
	LocalVariableDeclarators []LocalVariableDeclarator
}

func (d *LocalVariableDeclarators) GetLineNumber() int {
	if len(d.LocalVariableDeclarators) == 0 {
		return 0
	}

	return d.LocalVariableDeclarators[0].GetLineNumber()
}

func (d *LocalVariableDeclarators) Accept(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclarators(d)
}
