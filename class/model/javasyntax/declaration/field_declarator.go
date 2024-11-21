package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewFieldDeclarator(name string) intmod.IFieldDeclarator {
	return &FieldDeclarator{
		name: name,
	}
}

func NewFieldDeclarator2(name string, variableInitializer intmod.IVariableInitializer) intmod.IFieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
	}
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer intmod.IVariableInitializer) intmod.IFieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
		dimension:           dimension,
	}
}

type FieldDeclarator struct {
	util.DefaultBase[intmod.IFieldDeclarator]

	fieldDeclaration    intmod.IFieldDeclaration
	name                string
	dimension           int
	variableInitializer intmod.IVariableInitializer
}

func (d *FieldDeclarator) SetFieldDeclaration(fieldDeclaration intmod.IFieldDeclaration) {
	d.fieldDeclaration = fieldDeclaration
}

func (d *FieldDeclarator) FieldDeclaration() intmod.IFieldDeclaration {
	return d.fieldDeclaration
}

func (d *FieldDeclarator) Name() string {
	return d.name
}

func (d *FieldDeclarator) Dimension() int {
	return d.dimension
}

func (d *FieldDeclarator) VariableInitializer() intmod.IVariableInitializer {
	return d.variableInitializer
}

func (d *FieldDeclarator) SetVariableInitializer(variableInitializer intmod.IVariableInitializer) {
	d.variableInitializer = variableInitializer
}

func (d *FieldDeclarator) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{%s}", d.name)
}
