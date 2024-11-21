package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewLocalVariableDeclaration(typ intmod.IType, localVariableDeclarators intmod.ILocalVariableDeclarator) intmod.ILocalVariableDeclaration {
	return &LocalVariableDeclaration{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclaration struct {
	final                    bool
	typ                      intmod.IType
	localVariableDeclarators intmod.ILocalVariableDeclarator
}

func (d *LocalVariableDeclaration) IsFinal() bool {
	return d.final
}

func (d *LocalVariableDeclaration) SetFinal(final bool) {
	d.final = final
}

func (d *LocalVariableDeclaration) Type() intmod.IType {
	return d.typ
}

func (d *LocalVariableDeclaration) LocalVariableDeclarators() intmod.ILocalVariableDeclarator {
	return d.localVariableDeclarators
}

func (d *LocalVariableDeclaration) SetLocalVariableDeclarators(localVariableDeclarators intmod.ILocalVariableDeclarator) {
	d.localVariableDeclarators = localVariableDeclarators
}

func (d *LocalVariableDeclaration) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclaration(d)
}
