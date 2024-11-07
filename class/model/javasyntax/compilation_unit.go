package javasyntax

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"

type CompilationUnit struct {
	typeDeclarations declaration.ITypeDeclaration
}

func (u *CompilationUnit) GetTypeDeclaration() declaration.ITypeDeclaration {
	return u.typeDeclarations
}
