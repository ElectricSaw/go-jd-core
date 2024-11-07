package declaration

import "fmt"

func NewFieldDeclarator(name string) *FieldDeclarator {
	return &FieldDeclarator{
		name: name,
	}
}

func NewFieldDeclarator2(name string, variableInitializer VariableInitializer) *FieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
	}
}

func NewFieldDeclarator3(name string, dimension int, variableInitializer VariableInitializer) *FieldDeclarator {
	return &FieldDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
		dimension:           dimension,
	}
}

type FieldDeclarator struct {
	fieldDeclaration    *FieldDeclaration
	name                string
	dimension           int
	variableInitializer VariableInitializer
}

func (d *FieldDeclarator) SetFieldDeclaration(fieldDeclaration *FieldDeclaration) {
	d.fieldDeclaration = fieldDeclaration
}

func (d *FieldDeclarator) GetFieldDeclaration() *FieldDeclaration {
	return d.fieldDeclaration
}

func (d *FieldDeclarator) GetName() string {
	return d.name
}

func (d *FieldDeclarator) GetDimension() int {
	return d.dimension
}

func (d *FieldDeclarator) GetVariableInitializer() VariableInitializer {
	return d.variableInitializer
}

func (d *FieldDeclarator) SetVariableInitializer(variableInitializer VariableInitializer) {
	d.variableInitializer = variableInitializer
}

func (d *FieldDeclarator) Accept(visitor DeclarationVisitor) {
	visitor.VisitFieldDeclarator(d)
}

func (d *FieldDeclarator) String() string {
	return fmt.Sprintf("FieldDeclarator{%s}", d.name)
}
