package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLocalVariableReferenceExpression(typ intmod.IType, name string) intmod.ILocalVariableReferenceExpression {
	return &LocalVariableReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		name:                             name,
	}
}

func NewLocalVariableReferenceExpressionWithAll(lineNumber int, typ intmod.IType, name string) intmod.ILocalVariableReferenceExpression {
	return &LocalVariableReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		name:                             name,
	}
}

type LocalVariableReferenceExpression struct {
	AbstractLineNumberTypeExpression

	name string
}

func (e *LocalVariableReferenceExpression) Name() string {
	return e.name
}

func (e *LocalVariableReferenceExpression) IsLocalVariableReferenceExpression() bool {
	return true
}

func (e *LocalVariableReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLocalVariableReferenceExpression(e)
}

func (e *LocalVariableReferenceExpression) String() string {
	return fmt.Sprintf("LocalVariableReferenceExpression{type=%s, name=%s}", e.typ, e.name)
}
