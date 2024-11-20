package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLocalVariableReferenceExpression(typ intsyn.IType, name string) intsyn.ILocalVariableReferenceExpression {
	return &LocalVariableReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		name:                             name,
	}
}

func NewLocalVariableReferenceExpressionWithAll(lineNumber int, typ intsyn.IType, name string) intsyn.ILocalVariableReferenceExpression {
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

func (e *LocalVariableReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitLocalVariableReferenceExpression(e)
}

func (e *LocalVariableReferenceExpression) String() string {
	return fmt.Sprintf("LocalVariableReferenceExpression{type=%s, name=%s}", e.typ, e.name)
}
