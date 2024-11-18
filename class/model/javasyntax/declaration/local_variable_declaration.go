package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewLocalVariableDeclaration(typ intsyn.IType, localVariableDeclarators intsyn.ILocalVariableDeclarator) intsyn.ILocalVariableDeclaration {
	return &LocalVariableDeclaration{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclaration struct {
	final                    bool
	typ                      intsyn.IType
	localVariableDeclarators intsyn.ILocalVariableDeclarator
}

func (d *LocalVariableDeclaration) IsFinal() bool {
	return d.final
}

func (d *LocalVariableDeclaration) SetFinal(final bool) {
	d.final = final
}

func (d *LocalVariableDeclaration) Type() intsyn.IType {
	return d.typ
}

func (d *LocalVariableDeclaration) LocalVariableDeclarators() intsyn.ILocalVariableDeclarator {
	return d.localVariableDeclarators
}

func (d *LocalVariableDeclaration) SetLocalVariableDeclarators(localVariableDeclarators intsyn.ILocalVariableDeclarator) {
	d.localVariableDeclarators = localVariableDeclarators
}

func (d *LocalVariableDeclaration) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclaration(d)
}
