package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewConstructorInvocationExpression(objectType intmod.IObjectType, descriptor string,
	parameters intmod.IExpression) intmod.IConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid.(intmod.IType),
			objectType, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

func NewConstructorInvocationExpressionWithAll(lineNumber int, objectType intmod.IObjectType,
	descriptor string, parameters intmod.IExpression) intmod.IConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intmod.IType), objectType, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

type ConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters intmod.IExpression
}

func (e *ConstructorInvocationExpression) Parameters() intmod.IExpression {
	return e.parameters
}

func (e *ConstructorInvocationExpression) Priority() int {
	return 1
}

func (e *ConstructorInvocationExpression) SetParameters(params intmod.IExpression) {
	e.parameters = params
}

func (e *ConstructorInvocationExpression) IsConstructorInvocationExpression() bool {
	return true
}

func (e *ConstructorInvocationExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitConstructorInvocationExpression(e)
}

func (e *ConstructorInvocationExpression) String() string {
	return fmt.Sprintf("ConstructorInvocationExpression{call this(%s)}", e.descriptor)
}
