package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewFieldDeclarators(length int) intmod.IFieldDeclarators {
	return &FieldDeclarators{}
}

type FieldDeclarators struct {
	FieldDeclarator
	util.DefaultList[intmod.IFieldDeclarator]
}

func (d *FieldDeclarators) SetFieldDeclaration(fieldDeclaration intmod.IFieldDeclaration) {
	for _, fieldDeclarator := range d.Elements() {
		fieldDeclarator.SetFieldDeclaration(fieldDeclaration)
	}
}

func (d *FieldDeclarators) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclarators(d)
}
