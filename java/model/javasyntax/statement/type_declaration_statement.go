package statement

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/declaration"

func NewTypeDeclarationStatement(typeDeclaration declaration.TypeDeclaration) *TypeDeclarationStatement {
	return &TypeDeclarationStatement{
		typeDeclaration: typeDeclaration,
	}
}

type TypeDeclarationStatement struct {
	AbstractStatement

	typeDeclaration declaration.TypeDeclaration
}

func (s *TypeDeclarationStatement) GetTypeDeclaration() declaration.TypeDeclaration {
	return s.typeDeclaration
}

func (s *TypeDeclarationStatement) Accept(visitor StatementVisitor) {
	visitor.VisitTypeDeclarationStatement(s)
}
