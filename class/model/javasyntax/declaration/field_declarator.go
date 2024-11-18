package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"bitbucket.org/coontec/javaClass/class/util"
	"fmt"
)

func NewFieldDeclarator(name string) intsyn.IFieldDeclarator {
	return &FieldDeclarator{
		name: name,
	}
}

func NewFieldDeclarator2(name string, variableInitializer intsyn.IVariableInitializer) intsyn.IFieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
	}
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer intsyn.IVariableInitializer) intsyn.IFieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
		dimension:           dimension,
	}
}

type FieldDeclarator struct {
	util.DefaultBase[intsyn.IFieldDeclarator]

	fieldDeclaration    intsyn.IFieldDeclaration
	name                string
	dimension           int
	variableInitializer intsyn.IVariableInitializer
}

func (d *FieldDeclarator) SetFieldDeclaration(fieldDeclaration intsyn.IFieldDeclaration) {
	d.fieldDeclaration = fieldDeclaration
}

func (d *FieldDeclarator) FieldDeclaration() intsyn.IFieldDeclaration {
	return d.fieldDeclaration
}

func (d *FieldDeclarator) Name() string {
	return d.name
}

func (d *FieldDeclarator) Dimension() int {
	return d.dimension
}

func (d *FieldDeclarator) VariableInitializer() intsyn.IVariableInitializer {
	return d.variableInitializer
}

func (d *FieldDeclarator) SetVariableInitializer(variableInitializer intsyn.IVariableInitializer) {
	d.variableInitializer = variableInitializer
}

func (d *FieldDeclarator) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{%s}", d.name)
}
