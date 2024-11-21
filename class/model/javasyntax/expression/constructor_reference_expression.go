package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewConstructorReferenceExpression(typ intmod.IType, objectType intmod.IObjectType,
	descriptor string) intmod.IConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

func NewConstructorReferenceExpressionWithAll(lineNumber int, typ intmod.IType,
	objectType intmod.IObjectType, descriptor string) intmod.IConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

type ConstructorReferenceExpression struct {
	AbstractLineNumberTypeExpression

	objectType intmod.IObjectType
	descriptor string
}

func (e *ConstructorReferenceExpression) ObjectType() intmod.IObjectType {
	return e.objectType
}

func (e *ConstructorReferenceExpression) Descriptor() string {
	return e.descriptor
}

func (e *ConstructorReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitConstructorReferenceExpression(e)
}
