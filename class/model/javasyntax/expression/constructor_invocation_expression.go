package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
)

func NewConstructorInvocationExpression(objectType intmod.IObjectType, descriptor string,
	parameters intmod.IExpression) intmod.IConstructorInvocationExpression {
	return NewConstructorInvocationExpressionWithAll(0, objectType, descriptor, parameters)
}

func NewConstructorInvocationExpressionWithAll(lineNumber int, objectType intmod.IObjectType,
	descriptor string, parameters intmod.IExpression) intmod.IConstructorInvocationExpression {
	e := &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intmod.IType), objectType, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
	e.SetValue(e)
	return e
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
