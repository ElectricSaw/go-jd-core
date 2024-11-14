package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
)

func NewClassFileLocalVariableReferenceExpression(lineNumber, offset int,
	localVariable localvariable.ILocalVariableReference) *ClassFileLocalVariableReferenceExpression {
	e := &ClassFileLocalVariableReferenceExpression{
		LocalVariableReferenceExpression: *expression.NewLocalVariableReferenceExpressionWithAll(lineNumber, nil, ""),
		offset:                           offset,
		localVariable:                    localVariable,
	}
	e.localVariable.AddReference(e)

	return e
}

type ClassFileLocalVariableReferenceExpression struct {
	expression.LocalVariableReferenceExpression

	offset        int
	localVariable localvariable.ILocalVariableReference
}

func (e *ClassFileLocalVariableReferenceExpression) Offset() int {
	return e.offset
}

func (e *ClassFileLocalVariableReferenceExpression) LocalVariable() localvariable.ILocalVariableReference {
	return e.localVariable
}

func (e *ClassFileLocalVariableReferenceExpression) SetLocalVariable(localVariable localvariable.ILocalVariableReference) {
	e.localVariable = localVariable
}
