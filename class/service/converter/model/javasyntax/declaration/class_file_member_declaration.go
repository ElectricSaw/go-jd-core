package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"

type ClassFileMemberDeclaration interface {
	declaration.IMemberDeclaration

	FirstLineNumber() int
}
