package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/util"
)

func NewFieldDeclarators(length int) intsyn.IFieldDeclarators {
	return &FieldDeclarators{}
}

type FieldDeclarators struct {
	util.DefaultList[intsyn.IFieldDeclarator]
}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration intsyn.IFieldDeclaration) {
	for _, fieldDeclarator := range d.Elements() {
		fieldDeclarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}
