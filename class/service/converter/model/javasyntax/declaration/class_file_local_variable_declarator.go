package declaration

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
)

func NewClassFileLocalVariableDeclarator(localVariable localvariable.ILocalVariable) *ClassFileLocalVariableDeclarator {
	return &ClassFileLocalVariableDeclarator{
		LocalVariableDeclarator: *declaration.NewLocalVariableDeclarator(""),
		localVariable:           localVariable,
	}
}

func NewClassFileLocalVariableDeclarator2(lineNumber int, localVariable localvariable.ILocalVariable, initializer declaration.VariableInitializer) *ClassFileLocalVariableDeclarator {
	return &ClassFileLocalVariableDeclarator{
		LocalVariableDeclarator: *declaration.NewLocalVariableDeclarator3(lineNumber, "", initializer),
		localVariable:           localVariable,
	}
}

type ClassFileLocalVariableDeclarator struct {
	declaration.LocalVariableDeclarator

	localVariable localvariable.ILocalVariable
}

func (d *ClassFileLocalVariableDeclarator) Name() string {
	return d.localVariable.Name()
}

func (d *ClassFileLocalVariableDeclarator) SetName(name string) {
	d.localVariable.SetName(name)
}

func (d *ClassFileLocalVariableDeclarator) LocalVariable() localvariable.ILocalVariableReference {
	return d.localVariable
}

func (d *ClassFileLocalVariableDeclarator) SetLocalVariable(localVariable localvariable.ILocalVariableReference) {
	d.localVariable = localVariable.(localvariable.ILocalVariable)
}
