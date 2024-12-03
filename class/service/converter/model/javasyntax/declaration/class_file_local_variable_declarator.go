package declaration

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
)

func NewClassFileLocalVariableDeclarator(localVariable intsrv.ILocalVariable) intsrv.IClassFileLocalVariableDeclarator {
	return NewClassFileLocalVariableDeclarator2(-1, localVariable, nil)
}

func NewClassFileLocalVariableDeclarator2(lineNumber int, localVariable intsrv.ILocalVariable,
	initializer intmod.IVariableInitializer) intsrv.IClassFileLocalVariableDeclarator {
	d := &ClassFileLocalVariableDeclarator{
		LocalVariableDeclarator: *declaration.NewLocalVariableDeclarator3(lineNumber, "", initializer).(*declaration.LocalVariableDeclarator),
		localVariable:           localVariable,
	}
	d.SetValue(d)
	return d
}

type ClassFileLocalVariableDeclarator struct {
	declaration.LocalVariableDeclarator

	localVariable intsrv.ILocalVariable
}

func (d *ClassFileLocalVariableDeclarator) Name() string {
	return d.localVariable.Name()
}

func (d *ClassFileLocalVariableDeclarator) SetName(name string) {
	d.localVariable.SetName(name)
}

func (d *ClassFileLocalVariableDeclarator) LocalVariable() intsrv.ILocalVariableReference {
	return d.localVariable
}

func (d *ClassFileLocalVariableDeclarator) SetLocalVariable(localVariable intsrv.ILocalVariableReference) {
	d.localVariable = localVariable.(intsrv.ILocalVariable)
}
