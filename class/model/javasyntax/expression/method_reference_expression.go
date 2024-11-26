package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewMethodReferenceExpression(typ intmod.IType, expression intmod.IExpression,
	internalTypeName, name, descriptor string) intmod.IMethodReferenceExpression {
	return NewMethodReferenceExpressionWithAll(0, typ, expression, internalTypeName, name, descriptor)
}

func NewMethodReferenceExpressionWithAll(lineNumber int, typ intmod.IType, expression intmod.IExpression,
	internalTypeName, name, descriptor string) intmod.IMethodReferenceExpression {
	e := &MethodReferenceExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		internalTypeName:                 internalTypeName,
		name:                             name,
		descriptor:                       descriptor,
	}
	e.SetValue(e)
	return e
}

type MethodReferenceExpression struct {
	AbstractLineNumberTypeExpression
	util.DefaultBase[intmod.IMethodReferenceExpression]

	expression       intmod.IExpression
	internalTypeName string
	name             string
	descriptor       string
}

func (e *MethodReferenceExpression) Expression() intmod.IExpression {
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

func (e *MethodReferenceExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *MethodReferenceExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitMethodReferenceExpression(e)
}
