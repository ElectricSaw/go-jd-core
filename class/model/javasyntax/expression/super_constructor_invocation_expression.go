package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewSuperConstructorInvocationExpression(typ intsyn.IObjectType, descriptor string,
	parameters intsyn.IExpression) intsyn.ISuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid.(intsyn.IType),
			typ, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

func NewSuperConstructorInvocationExpressionWithAll(lineNumber int, typ intsyn.IObjectType,
	descriptor string, parameters intsyn.IExpression) intsyn.ISuperConstructorInvocationExpression {
	return &SuperConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intsyn.IType), typ, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

type SuperConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters intsyn.IExpression
}

func (e *SuperConstructorInvocationExpression) Parameters() intsyn.IExpression {
	return e.parameters
}

func (e *SuperConstructorInvocationExpression) SetParameters(expression intsyn.IExpression) {
	e.parameters = expression
}

func (e *SuperConstructorInvocationExpression) Priority() int {
	return 1
}

func (e *SuperConstructorInvocationExpression) IsSuperConstructorInvocationExpression() bool {
	return true
}

func (e *SuperConstructorInvocationExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitSuperConstructorInvocationExpression(e)
}

func (e *SuperConstructorInvocationExpression) String() string {
	return fmt.Sprintf("SuperConstructorInvocationExpression{call super(%s)}", e.descriptor)
}
