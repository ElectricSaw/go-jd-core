package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewMethodReferenceExpression(typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName, name, descriptor string) intsyn.IMethodReferenceExpression {
	return &MethodReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

func NewMethodReferenceExpressionWithAll(lineNumber int, typ intsyn.IType, expression intsyn.IExpression,
	internalTypeName, name, descriptor string) intsyn.IMethodReferenceExpression {
	return &MethodReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

type MethodReferenceExpression struct {
	AbstractLineNumberTypeExpression

	expression       intsyn.IExpression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *MethodReferenceExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *MethodReferenceExpression) InternalTypeName() string {
	return e.internalTypeName
}

func (e *MethodReferenceExpression) Name() string {
	return e.name
}

func (e *MethodReferenceExpression) Descriptor() string {
	return e.descriptor
}

func (e *MethodReferenceExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *MethodReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitMethodReferenceExpression(e)
}
