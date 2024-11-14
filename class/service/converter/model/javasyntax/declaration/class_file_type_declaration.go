package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"

type ClassFileTypeDeclaration interface {
	ClassFileMemberDeclaration

	InternalTypeName() string
	BodyDeclaration() *declaration.BodyDeclaration
}
