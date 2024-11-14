package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"fmt"
)

type ClassFileClassDeclaration struct {
	declaration.ClassDeclaration

	firstLineNumber int
}

func (d *ClassFileClassDeclaration) FirstLineNumber() int {
	return d.firstLineNumber
}

func (d *ClassFileClassDeclaration) String() string {
	return fmt.Sprintf("ClassFileClassDeclaration{%s, firstLineNumber=%d}", d.InternalTypeName(), d.firstLineNumber)
}
