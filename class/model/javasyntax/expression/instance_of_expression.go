package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewInstanceOfExpression(expression intsyn.IExpression, instanceOfType intsyn.IObjectType) intsyn.IInstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: UnknownLineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType.(intsyn.IType),
	}
}

func NewInstanceOfExpressionWithAll(lineNumber int, expression intsyn.IExpression, instanceOfType intsyn.IObjectType) intsyn.IInstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType.(intsyn.IType),
	}
}

type InstanceOfExpression struct {
	AbstractLineNumberExpression

	expression     intsyn.IExpression
	instanceOfType intsyn.IType
}

func (e *InstanceOfExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *InstanceOfExpression) InstanceOfType() intsyn.IType {
	return e.instanceOfType
}

func (e *InstanceOfExpression) Type() intsyn.IType {
	return _type.PtTypeBoolean.(intsyn.IType)
}

func (e *InstanceOfExpression) Priority() int {
	return 8
}

func (e *InstanceOfExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *InstanceOfExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitInstanceOfExpression(e)
}
