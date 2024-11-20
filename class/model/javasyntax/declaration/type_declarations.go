package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
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
