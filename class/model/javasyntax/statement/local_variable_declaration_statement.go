package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewLocalVariableDeclarationStatement(typ _type.IType, localVariableDeclarators declaration.ILocalVariableDeclarator) *LocalVariableDeclarationStatement {
	return &LocalVariableDeclarationStatement{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclarationStatement struct {
	AbstractStatement
	declaration.LocalVariableDeclaration

	final                    bool
	typ                      _type.IType
	localVariableDeclarators declaration.ILocalVariableDeclarator
}

func (s *LocalVariableDeclarationStatement) IsFinal() bool {
	return s.final
}

func (s *LocalVariableDeclarationStatement) SetFinal(final bool) {
	s.final = final
}

func (s *LocalVariableDeclarationStatement) GetType() _type.IType {
	return s.typ
}

func (s *LocalVariableDeclarationStatement) GetLocalVariableDeclarators() declaration.ILocalVariableDeclarator {
	return s.localVariableDeclarators
}

func (s *LocalVariableDeclarationStatement) SetLocalVariableDeclarators(declarators declaration.ILocalVariableDeclarator) {
	s.localVariableDeclarators = declarators
}

func (s *LocalVariableDeclarationStatement) Accept(visitor StatementVisitor) {
	visitor.VisitLocalVariableDeclarationStatement(s)
}
