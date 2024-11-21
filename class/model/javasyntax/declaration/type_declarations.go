package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewTypeDeclarations() intmod.ITypeDeclarations {
	return &TypeDeclarations{}
}

type TypeDeclarations struct {
	AbstractTypeDeclaration
	util.DefaultList[intmod.ITypeDeclaration]
}

func (d *TypeDeclarations) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitTypeDeclarations(d)
}
