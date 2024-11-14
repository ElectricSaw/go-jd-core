package javasyntax

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"

func NewCompilationUnit(typeDeclarations declaration.ITypeDeclaration) *CompilationUnit {
	return &CompilationUnit{
		typeDeclarations: typeDeclarations,
	}
}

type CompilationUnit struct {
	typeDeclarations declaration.ITypeDeclaration
}

func (u *CompilationUnit) TypeDeclarations() declaration.ITypeDeclaration {
	return u.typeDeclarations
}
