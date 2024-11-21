package javasyntax

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewCompilationUnit(typeDeclarations intmod.ITypeDeclaration) *CompilationUnit {
	return &CompilationUnit{
		typeDeclarations: typeDeclarations,
	}
}

type CompilationUnit struct {
	typeDeclarations intmod.ITypeDeclaration
}

func (u *CompilationUnit) TypeDeclarations() intmod.ITypeDeclaration {
	return u.typeDeclarations
}
