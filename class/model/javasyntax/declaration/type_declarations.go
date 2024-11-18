package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewTypeDeclarations() intsyn.ITypeDeclarations {
	return &TypeDeclarations{}
}

type TypeDeclarations struct {
	AbstractTypeDeclaration
	util.DefaultList[intsyn.ITypeDeclaration]
}

func (d *TypeDeclarations) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
