package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewTypeDeclarationStatement(typeDeclaration intsyn.ITypeDeclaration) intsyn.ITypeDeclarationStatement {
	return &TypeDeclarationStatement{
		typeDeclaration: typeDeclaration,
	}
}

type TypeDeclarationStatement struct {
	AbstractStatement

	typeDeclaration intsyn.ITypeDeclaration
}

func (s *TypeDeclarationStatement) TypeDeclaration() intsyn.ITypeDeclaration {
	return s.typeDeclaration
}

func (s *TypeDeclarationStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitTypeDeclarationStatement(s)
}
