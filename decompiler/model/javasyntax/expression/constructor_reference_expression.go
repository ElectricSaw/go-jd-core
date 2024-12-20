package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewConstructorReferenceExpression(typ intmod.IType, objectType intmod.IObjectType,
	descriptor string) intmod.IConstructorReferenceExpression {
	return NewConstructorReferenceExpressionWithAll(0, typ, objectType, descriptor)
}

func NewConstructorReferenceExpressionWithAll(lineNumber int, typ intmod.IType,
	objectType intmod.IObjectType, descriptor string) intmod.IConstructorReferenceExpression {
	e := &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
	e.SetValue(e)
	return e
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
