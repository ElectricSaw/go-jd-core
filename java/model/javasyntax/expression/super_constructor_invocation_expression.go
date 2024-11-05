package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewSuperConstructorInvocationExpression(typ _type.ObjectType, descriptor string, parameters Expression) *SuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid, typ, descriptor),
		parameters:                     parameters,
	}
}

func NewSuperConstructorInvocationExpressionWithAll(lineNumber int, typ _type.ObjectType, descriptor string, parameters Expression) *SuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber, _type.PtTypeVoid, typ, descriptor),
		parameters:                     parameters,
	}
}

type SuperConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters Expression
}

func (e *SuperConstructorInvocationExpression) GetParameters() Expression {
	return e.parameters
}

func (e *SuperConstructorInvocationExpression) SetParameters(expression Expression) {
	e.parameters = expression
}

func (e *SuperConstructorInvocationExpression) GetPriority() int {
	return 1
}

func (e *SuperConstructorInvocationExpression) IsSuperConstructorInvocationExpression() bool {
	return true
}

func (e *SuperConstructorInvocationExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitSuperConstructorInvocationExpression(e)
}

func (e *SuperConstructorInvocationExpression) String() string {
	return fmt.Sprintf("SuperConstructorInvocationExpression{call super(%s)}", e.descriptor)
}
