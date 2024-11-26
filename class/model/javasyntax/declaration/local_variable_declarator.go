package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLocalVariableDeclarator(name string) intmod.ILocalVariableDeclarator {
	return NewLocalVariableDeclarator3(0, name, nil)
}

func NewLocalVariableDeclarator2(name string, variableInitializer intmod.IVariableInitializer) intmod.ILocalVariableDeclarator {
	return NewLocalVariableDeclarator3(0, name, variableInitializer)
}

func NewLocalVariableDeclarator3(lineNumber int, name string, variableInitializer intmod.IVariableInitializer) intmod.ILocalVariableDeclarator {
	d := &LocalVariableDeclarator{
		lineNumber:          lineNumber,
		name:                name,
		variableInitializer: variableInitializer,
	}
	d.SetValue(d)
	return d
}

type LocalVariableDeclarator struct {
	util.DefaultBase[intmod.ILocalVariableDeclarator]

	lineNumber          int
	name                string
	dimension           int
	variableInitializer intmod.IVariableInitializer
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

func (d *LocalVariableDeclarator) VariableInitializer() intmod.IVariableInitializer {
	return d.variableInitializer
}

func (d *LocalVariableDeclarator) Accept(visitor intmod.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarator(d)
}

func (d *LocalVariableDeclarator) String() string {
	return fmt.Sprintf("LocalVariableDeclarator{name=%s, dimension=%d, variableInitializer=%v}", d.name, d.dimension, d.variableInitializer)
}
