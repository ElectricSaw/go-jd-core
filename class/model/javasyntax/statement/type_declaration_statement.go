package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewTypeDeclarationStatement(typeDeclaration intmod.ITypeDeclaration) intmod.ITypeDeclarationStatement {
	return &TypeDeclarationStatement{
		typeDeclaration: typeDeclaration,
	}
}

type TypeDeclarationStatement struct {
	AbstractStatement

	typeDeclaration intmod.ITypeDeclaration
}

func (s *TypeDeclarationStatement) TypeDeclaration() intmod.ITypeDeclaration {
	return s.typeDeclaration
}

func (s *TypeDeclarationStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitTypeDeclarationStatement(s)
}
