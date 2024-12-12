package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewSuperConstructorInvocationExpression(typ intmod.IObjectType, descriptor string,
	parameters intmod.IExpression) intmod.ISuperConstructorInvocationExpression {
	return NewSuperConstructorInvocationExpressionWithAll(0, typ, descriptor, parameters)
}

func NewSuperConstructorInvocationExpressionWithAll(lineNumber int, typ intmod.IObjectType,
	descriptor string, parameters intmod.IExpression) intmod.ISuperConstructorInvocationExpression {
	e := &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intmod.IType), typ, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
	e.SetValue(e)
	return e
}

type SuperConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters intmod.IExpression
}

func (e *SuperConstructorInvocationExpression) Parameters() intmod.IExpression {
	return e.parameters
}

func (e *SuperConstructorInvocationExpression) SetParameters(expression intmod.IExpression) {
	e.parameters = expression
}

func (e *SuperConstructorInvocationExpression) Priority() int {
	return 1
}

func (e *SuperConstructorInvocationExpression) IsSuperConstructorInvocationExpression() bool {
	return true
}

func (e *SuperConstructorInvocationExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitSuperConstructorInvocationExpression(e)
}

func (e *SuperConstructorInvocationExpression) String() string {
	return fmt.Sprintf("SuperConstructorInvocationExpression{call super(%s)}", e.descriptor)
}
