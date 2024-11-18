package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
)

func NewLocalVariableDeclarationStatement(typ intsyn.IType,
	localVariableDeclarators intsyn.ILocalVariableDeclarator) intsyn.ILocalVariableDeclarationStatement {
	return &LocalVariableDeclarationStatement{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclarationStatement struct {
	AbstractStatement
	declaration.LocalVariableDeclaration

	final                    bool
	typ                      intsyn.IType
	localVariableDeclarators intsyn.ILocalVariableDeclarator
}

func (s *LocalVariableDeclarationStatement) IsFinal() bool {
	return s.final
}

func (s *LocalVariableDeclarationStatement) SetFinal(final bool) {
	s.final = final
}

func (s *LocalVariableDeclarationStatement) Type() intsyn.IType {
	return s.typ
}

func (s *LocalVariableDeclarationStatement) LocalVariableDeclarators() intsyn.ILocalVariableDeclarator {
	return s.localVariableDeclarators
}

func (s *LocalVariableDeclarationStatement) SetLocalVariableDeclarators(declarators intsyn.ILocalVariableDeclarator) {
	s.localVariableDeclarators = declarators
}

func (s *LocalVariableDeclarationStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitLocalVariableDeclarationStatement(s)
}
