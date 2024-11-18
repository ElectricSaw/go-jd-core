package javasyntax

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewCompilationUnit(typeDeclarations intsyn.ITypeDeclaration) *CompilationUnit {
	return &CompilationUnit{
		typeDeclarations: typeDeclarations,
	}
}

type CompilationUnit struct {
	typeDeclarations intsyn.ITypeDeclaration
}

func (u *CompilationUnit) TypeDeclarations() intsyn.ITypeDeclaration {
	return u.typeDeclarations
}
