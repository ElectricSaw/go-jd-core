package declaration

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/model/localvariable"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewClassFileLocalVariableDeclarator(localVariable localvariable.ILocalVariable) ClassFileLocalVariableDeclarator {
	return NewClassFileLocalVariableDeclarator2(-1, localVariable, nil)
}

func NewClassFileLocalVariableDeclarator2(lineNumber int, localVariable localvariable.ILocalVariable,
	initializer model.IVariableInitializer) ClassFileLocalVariableDeclarator {
	d := ClassFileLocalVariableDeclarator{
		DefaultBase:         *util.NewDefaultBase[model.ILocalVariableDeclarator]().(*util.DefaultBase[model.ILocalVariableDeclarator]),
		LineNumber:          lineNumber,
		Name:                "",
		VariableInitializer: initializer,
		LocalVariable:       localVariable,
	}
	d.SetValue(&d)
	return d
}

type ClassFileLocalVariableDeclarator struct {
	util.DefaultBase[model.ILocalVariableDeclarator]

	LineNumber          int
	Name                string
	Dimension           int
	VariableInitializer model.IVariableInitializer
	LocalVariable       localvariable.ILocalVariable
}

func (d *ClassFileLocalVariableDeclarator) GetLineNumber() int {
	return d.LineNumber
}

func (d *ClassFileLocalVariableDeclarator) GetName() string {
	return d.Name
}

func (d *ClassFileLocalVariableDeclarator) GetDimension() int {
	return d.Dimension
}

func (d *ClassFileLocalVariableDeclarator) GetVariableInitializer() model.IVariableInitializer {
	return d.VariableInitializer
}

func (d *ClassFileLocalVariableDeclarator) SetName(name string) {
	d.LocalVariable.SetName(name)
}

func (d *ClassFileLocalVariableDeclarator) GetLocalVariable() localvariable.ILocalVariableReference {
	return d.LocalVariable
}

func (d *ClassFileLocalVariableDeclarator) SetLocalVariable(localVariable localvariable.ILocalVariableReference) {
	d.LocalVariable = localVariable.(localvariable.ILocalVariable)
}

func (d *ClassFileLocalVariableDeclarator) AcceptDeclaration(visitor model.IDeclarationVisitor) {
	visitor.VisitLocalVariableDeclarator(d)
}

func (d *ClassFileLocalVariableDeclarator) String() string {
	return fmt.Sprintf("ClassFileLocalVariableDeclarator{ name=%s, dimension=%d, variable-initializer=%v }", d.Name, d.Dimension, d.VariableInitializer)
}
