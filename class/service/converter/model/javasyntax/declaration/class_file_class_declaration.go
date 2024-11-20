package declaration

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"fmt"
)

func NewClassFileClassDeclaration() intsrv.IClassFileClassDeclaration {
	return &ClassFileClassDeclaration{}
}

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
