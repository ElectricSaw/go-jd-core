package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewInstanceOfExpression(expression intmod.IExpression,
	instanceOfType intmod.IObjectType) intmod.IInstanceOfExpression {
	return NewInstanceOfExpressionWithAll(0, expression, instanceOfType)
}

func NewInstanceOfExpressionWithAll(lineNumber int, expression intmod.IExpression,
	instanceOfType intmod.IObjectType) intmod.IInstanceOfExpression {
	e := &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType.(intmod.IType),
	}
	e.SetValue(e)
	return e
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
