package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewConstructorInvocationExpression(objectType _type.ObjectType, descriptor string, parameters Expression) *ConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid, objectType, descriptor),
		parameters:                     parameters,
	}
}

func NewConstructorInvocationExpressionWithAll(lineNumber int, objectType _type.ObjectType, descriptor string, parameters Expression) *ConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber, _type.PtTypeVoid, objectType, descriptor),
		parameters:                     parameters,
	}
}

type ConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters Expression
}

func (e *ConstructorInvocationExpression) GetParameters() Expression {
	return e.parameters
}

func (e *ConstructorInvocationExpression) GetPriority() int {
	return 1
}

func (e *ConstructorInvocationExpression) SetParameters(params Expression) {
	e.parameters = params
}

func (e *ConstructorInvocationExpression) IsConstructorInvocationExpression() bool {
	return true
}

func (e *ConstructorInvocationExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitConstructorInvocationExpression(e)
}

func (e *ConstructorInvocationExpression) String() string {
	return fmt.Sprintf("ConstructorInvocationExpression{call this(%s)}", e.descriptor)
}
