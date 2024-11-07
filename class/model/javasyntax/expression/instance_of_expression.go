package expression

import _type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"

func NewInstanceOfExpression(expression Expression, instanceOfType *_type.ObjectType) *InstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: UnknownLineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType,
	}
}

func NewInstanceOfExpressionWithAll(lineNumber int, expression Expression, instanceOfType *_type.ObjectType) *InstanceOfExpression {
	return &InstanceOfExpression{
		AbstractLineNumberExpression: AbstractLineNumberExpression{
			lineNumber: lineNumber,
		},
		expression:     expression,
		instanceOfType: instanceOfType,
	}
}

type InstanceOfExpression struct {
	AbstractLineNumberExpression

	expression     Expression
	instanceOfType _type.IType
}

func (e *InstanceOfExpression) GetExpression() Expression {
	return e.expression
}

func (e *InstanceOfExpression) GetInstanceOfType() _type.IType {
	return e.instanceOfType
}

func (e *InstanceOfExpression) GetType() _type.IType {
	return _type.PtTypeBoolean
}

func (e *InstanceOfExpression) GetPriority() int {
	return 8
}

func (e *InstanceOfExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *InstanceOfExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitInstanceOfExpression(e)
}
