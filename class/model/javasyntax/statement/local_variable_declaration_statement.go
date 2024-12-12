package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
)

func NewLocalVariableDeclarationStatement(typ intmod.IType,
	localVariableDeclarators intmod.ILocalVariableDeclarator) intmod.ILocalVariableDeclarationStatement {
	return &LocalVariableDeclarationStatement{
		typ:                      typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclarationStatement struct {
	AbstractStatement
	declaration.LocalVariableDeclaration

	final                    bool
	typ                      intmod.IType
	localVariableDeclarators intmod.ILocalVariableDeclarator
}

func (s *LocalVariableDeclarationStatement) IsFinal() bool {
	return s.final
}

func (s *LocalVariableDeclarationStatement) SetFinal(final bool) {
	s.final = final
}

func (s *LocalVariableDeclarationStatement) Type() intmod.IType {
	return s.typ
}

func (s *LocalVariableDeclarationStatement) LocalVariableDeclarators() intmod.ILocalVariableDeclarator {
	return s.localVariableDeclarators
}

func (s *LocalVariableDeclarationStatement) SetLocalVariableDeclarators(declarators intmod.ILocalVariableDeclarator) {
	s.localVariableDeclarators = declarators
}

func (s *LocalVariableDeclarationStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitLocalVariableDeclarationStatement(s)
}
