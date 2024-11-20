package declaration

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLocalVariableDeclarator(name string) intsyn.ILocalVariableDeclarator {
	return &LocalVariableDeclarator{
		name: name,
	}
}

func NewLocalVariableDeclarator2(name string, variableInitializer intsyn.IVariableInitializer) intsyn.ILocalVariableDeclarator {
	return &LocalVariableDeclarator{
		name:                name,
		variableInitializer: variableInitializer,
	}
}

func NewLocalVariableDeclarator3(lineNumber int, name string, variableInitializer intsyn.IVariableInitializer) intsyn.ILocalVariableDeclarator {
	return &LocalVariableDeclarator{
		lineNumber:          lineNumber,
		name:                name,
		variableInitializer: variableInitializer,
	}
}

type LocalVariableDeclarator struct {
	util.DefaultBase[intsyn.ILocalVariableDeclarator]

	lineNumber          int
	name                string
	dimension           int
	variableInitializer intsyn.IVariableInitializer
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

func (d *LocalVariableDeclarator) VariableInitializer() intsyn.IVariableInitializer {
	return d.variableInitializer
}

func (d *LocalVariableDeclarator) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarator(d)
}

func (d *LocalVariableDeclarator) String() string {
	return fmt.Sprintf("LocalVariableDeclarator{name=%s, dimension=%d, variableInitializer=%v}", d.name, d.dimension, d.variableInitializer)
}
