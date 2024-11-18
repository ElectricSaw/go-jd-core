package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewConstructorReferenceExpression(typ intsyn.IType, objectType intsyn.IObjectType,
	descriptor string) intsyn.IConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

func NewConstructorReferenceExpressionWithAll(lineNumber int, typ intsyn.IType,
	objectType intsyn.IObjectType, descriptor string) intsyn.IConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

type ConstructorReferenceExpression struct {
	AbstractLineNumberTypeExpression

	objectType intsyn.IObjectType
	descriptor string
}

func (e *ConstructorReferenceExpression) ObjectType() intsyn.IObjectType {
	return e.objectType
}

func (e *ConstructorReferenceExpression) Descriptor() string {
	return e.descriptor
}

func (e *ConstructorReferenceExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitConstructorReferenceExpression(e)
}
