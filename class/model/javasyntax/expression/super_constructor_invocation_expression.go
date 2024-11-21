package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewSuperConstructorInvocationExpression(typ intmod.IObjectType, descriptor string,
	parameters intmod.IExpression) intmod.ISuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid.(intmod.IType),
			typ, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

func NewSuperConstructorInvocationExpressionWithAll(lineNumber int, typ intmod.IObjectType,
	descriptor string, parameters intmod.IExpression) intmod.ISuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intmod.IType), typ, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
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
