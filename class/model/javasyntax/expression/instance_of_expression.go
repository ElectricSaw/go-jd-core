package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewInstanceOfExpression(expression intmod.IExpression, instanceOfType intmod.IObjectType) intmod.IInstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: intmod.UnknownLineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType.(intmod.IType),
	}
}

func NewInstanceOfExpressionWithAll(lineNumber int, expression intmod.IExpression, instanceOfType intmod.IObjectType) intmod.IInstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType.(intmod.IType),
	}
}

type InstanceOfExpression struct {
	AbstractLineNumberExpression

	expression     intmod.IExpression
	instanceOfType intmod.IType
}

func (e *InstanceOfExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *InstanceOfExpression) InstanceOfType() intmod.IType {
	return e.instanceOfType
}

func (e *InstanceOfExpression) Type() intmod.IType {
	return _type.PtTypeBoolean.(intmod.IType)
}

func (e *InstanceOfExpression) Priority() int {
	return 8
}

func (e *InstanceOfExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *InstanceOfExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitInstanceOfExpression(e)
}
