package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewConstructorReferenceExpression(typ _type.IType, objectType _type.IObjectType, descriptor string) *ConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

func NewConstructorReferenceExpressionWithAll(lineNumber int, typ _type.IType, objectType _type.IObjectType, descriptor string) *ConstructorReferenceExpression {
	return &ConstructorReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		objectType:                       objectType,
		descriptor:                       descriptor,
	}
}

type ConstructorReferenceExpression struct {
	AbstractLineNumberTypeExpression

	objectType _type.IObjectType
	descriptor string
}

func (e *ConstructorReferenceExpression) ObjectType() _type.IObjectType {
	return e.objectType
}

func (e *ConstructorReferenceExpression) Descriptor() string {
	return e.descriptor
}

func (e *ConstructorReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitConstructorReferenceExpression(e)
}
