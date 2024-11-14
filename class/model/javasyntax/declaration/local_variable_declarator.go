package declaration

import "fmt"

func NewLocalVariableDeclarator(name string) *LocalVariableDeclarator {
	return &LocalVariableDeclarator{
		name: name,
	}
}

func NewLocalVariableDeclarator2(name string, variableInitializer VariableInitializer) *LocalVariableDeclarator {
	return &LocalVariableDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
	}
}

func NewLocalVariableDeclarator3(lineNumber int, name string, variableInitializer VariableInitializer) *LocalVariableDeclarator {
	return &LocalVariableDeclarator{
		lineNumber:          lineNumber,
		name:                name,
		variableInitializer: variableInitializer,
	}
}

type LocalVariableDeclarator struct {
	lineNumber          int
	name                string
	dimension           int
	variableInitializer VariableInitializer
}

func (d *LocalVariableDeclarator) Name() string {
	return d.name
}

func (d *LocalVariableDeclarator) SetName(name string) {
	d.name = name
}

func (d *LocalVariableDeclarator) Dimension() int {
	return d.dimension
}

func (d *LocalVariableDeclarator) SetDimension(dimension int) {
	d.dimension = dimension
}

func (d *LocalVariableDeclarator) LineNumber() int {
	return d.lineNumber
}

func (d *LocalVariableDeclarator) VariableInitializer() VariableInitializer {
	return d.variableInitializer
}

func (d *LocalVariableDeclarator) Accept(visitor DeclarationVisitor) {
	visitor.VisitLocalVariableDeclarator(d)
}

func (d *LocalVariableDeclarator) String() string {
	return fmt.Sprintf("LocalVariableDeclarator{name=%s, dimension=%d, variableInitializer=%v}", d.name, d.dimension, d.variableInitializer)
}
