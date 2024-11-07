package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewLocalVariableReferenceExpression(typ _type.IType, name string) *LocalVariableReferenceExpression {
	return &LocalVariableReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		name:                             name,
	}
}

func NewLocalVariableReferenceExpressionWithAll(lineNumber int, typ _type.IType, name string) *LocalVariableReferenceExpression {
	return &LocalVariableReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		name:                             name,
	}
}

type LocalVariableReferenceExpression struct {
	AbstractLineNumberTypeExpression

	name string
}

func (e *LocalVariableReferenceExpression) GetName() string {
	return e.name
}

func (e *LocalVariableReferenceExpression) IsLocalVariableReferenceExpression() bool {
	return true
}

func (e *LocalVariableReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitLocalVariableReferenceExpression(e)
}

func (e *LocalVariableReferenceExpression) String() string {
	return fmt.Sprintf("LocalVariableReferenceExpression{type=%s, name=%s}", e.typ, e.name)
}
