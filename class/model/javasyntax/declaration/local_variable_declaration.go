package declaration

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewLocalVariableDeclaration(typ _type.IType, localVariableDeclarators ILocalVariableDeclarator) *LocalVariableDeclaration {
	return &LocalVariableDeclaration{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclaration struct {
	final                    bool
	typ                      _type.IType
	localVariableDeclarators ILocalVariableDeclarator
}

func (d *LocalVariableDeclaration) IsFinal() bool {
	return d.final
}

func (d *LocalVariableDeclaration) SetFinal(final bool) {
	d.final = final
}

func (d *LocalVariableDeclaration) GetType() _type.IType {
	return d.typ
}

func (d *LocalVariableDeclaration) GetLocalVariableDeclarators() ILocalVariableDeclarator {
	return d.localVariableDeclarators
}

func (d *LocalVariableDeclaration) SetLocalVariableDeclarators(localVariableDeclarators ILocalVariableDeclarator) {
	d.localVariableDeclarators = localVariableDeclarators
}

func (d *LocalVariableDeclaration) Accept(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclaration(d)
}
