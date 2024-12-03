package expression

import (
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
)

func NewClassFileLocalVariableReferenceExpression(lineNumber, offset int,
	localVariable intsrv.ILocalVariable) intsrv.IClassFileLocalVariableReferenceExpression {
	e := &ClassFileLocalVariableReferenceExpression{
		LocalVariableReferenceExpression: *expression.NewLocalVariableReferenceExpressionWithAll(
			lineNumber, nil, "").(*expression.LocalVariableReferenceExpression),
		offset:        offset,
		localVariable: localVariable,
	}
	e.localVariable.AddReference(e)
	e.SetValue(e)

	return e
}

type ClassFileLocalVariableReferenceExpression struct {
	expression.LocalVariableReferenceExpression

	offset        int
	localVariable intsrv.ILocalVariable
}

func (e *ClassFileLocalVariableReferenceExpression) Offset() int {
	return e.offset
}

func (e *ClassFileLocalVariableReferenceExpression) LocalVariable() intsrv.ILocalVariableReference {
	return e.localVariable
}

func (e *ClassFileLocalVariableReferenceExpression) SetLocalVariable(localVariable intsrv.ILocalVariableReference) {
	e.localVariable = localVariable.(intsrv.ILocalVariable)
}
