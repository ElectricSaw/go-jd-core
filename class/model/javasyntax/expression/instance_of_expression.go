package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
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
	util.DefaultBase[intmod.IInstanceOfExpression]

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
