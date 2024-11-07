package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewMethodReferenceExpression(typ _type.IType, expression Expression, internalTypeName, name, descriptor string) *MethodReferenceExpression {
	return &MethodReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
}

func NewMethodReferenceExpressionWithAll(lineNumber int, typ _type.IType, expression Expression, internalTypeName, name, descriptor string) *MethodReferenceExpression {
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

	expression       Expression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *MethodReferenceExpression) GetExpression() Expression {
	return e.expression
}

func (e *MethodReferenceExpression) GetInternalTypeName() string {
	return e.internalTypeName
}

func (e *MethodReferenceExpression) GetName() string {
	return e.name
}

func (e *MethodReferenceExpression) GetDescriptor() string {
	return e.descriptor
}

func (e *MethodReferenceExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *MethodReferenceExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitMethodReferenceExpression(e)
}