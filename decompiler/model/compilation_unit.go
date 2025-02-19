package model

func NewCompilationUnit(typeDeclarations ITypeDeclaration) CompilationUnit {
	return CompilationUnit{
		TypeDeclarations: typeDeclarations,
	}
}

type CompilationUnit struct {
	TypeDeclarations ITypeDeclaration
}
