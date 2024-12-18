package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFieldDeclarator(name string) intmod.IFieldDeclarator {
	return NewFieldDeclarator3(name, 0, nil)
}

func NewFieldDeclarator2(name string, variableInitializer intmod.IVariableInitializer) intmod.IFieldDeclarator {
	return NewFieldDeclarator3(name, 0, variableInitializer)
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer intmod.IVariableInitializer) intmod.IFieldDeclarator {
	d := &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
		dimension:           dimension,
	}
	d.SetValue(d)
	return d
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

func (d *FieldDeclarator) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{%s}", d.name)
}
