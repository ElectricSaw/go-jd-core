package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewLocalVariableDeclaration(typ Type, localVariableDeclarators LocalVariableDeclarator) LocalVariableDeclaration {
	return &LocalVariableDeclaration{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclaration struct {
	final                    bool
	typ                      Type
	localVariableDeclarators LocalVariableDeclarator
}

func (d *LocalVariableDeclaration) IsFinal() bool {
	return d.final
}

func (d *LocalVariableDeclaration) SetFinal(final bool) {
	d.final = final
}

func (d *LocalVariableDeclaration) Type() Type {
	return d.typ
}

func (d *LocalVariableDeclaration) LocalVariableDeclarators() LocalVariableDeclarator {
	return d.localVariableDeclarators
}

func (d *LocalVariableDeclaration) SetLocalVariableDeclarators(localVariableDeclarators LocalVariableDeclarator) {
	d.localVariableDeclarators = localVariableDeclarators
}

func (d *LocalVariableDeclaration) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclaration(d)
}
