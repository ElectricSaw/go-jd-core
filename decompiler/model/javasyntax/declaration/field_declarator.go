package declaration

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewFieldDeclarator(name string) FieldDeclarator {
	return NewFieldDeclarator3(name, 0, nil)
}

func NewFieldDeclarator2(name string, variableInitializer VariableInitializer) FieldDeclarator {
	return NewFieldDeclarator3(name, 0, variableInitializer)
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer VariableInitializer) FieldDeclarator {
	d := &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
		dimension:           dimension,
	}
	d.SetValue(d)
	return d
}

type FieldDeclarator struct {
	util.DefaultBase[FieldDeclarator]

	fieldDeclaration    FieldDeclaration
	name                string
	dimension           int
	variableInitializer VariableInitializer
}

func (d *FieldDeclarator) SetFieldDeclaration(fieldDeclaration FieldDeclaration) {
	d.fieldDeclaration = fieldDeclaration
}

func (d *FieldDeclarator) FieldDeclaration() FieldDeclaration {
	return d.fieldDeclaration
}

func (d *FieldDeclarator) Name() string {
	return d.name
}

func (d *FieldDeclarator) Dimension() int {
	return d.dimension
}

func (d *FieldDeclarator) VariableInitializer() VariableInitializer {
	return d.variableInitializer
}

func (d *FieldDeclarator) SetVariableInitializer(variableInitializer VariableInitializer) {
	d.variableInitializer = variableInitializer
}

func (d *FieldDeclarator) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{%s}", d.name)
}
