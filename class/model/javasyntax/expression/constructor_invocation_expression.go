package expression

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewConstructorInvocationExpression(objectType intsyn.IObjectType, descriptor string,
	parameters intsyn.IExpression) intsyn.IConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpression(_type.PtTypeVoid.(intsyn.IType),
			objectType, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

func NewConstructorInvocationExpressionWithAll(lineNumber int, objectType intsyn.IObjectType,
	descriptor string, parameters intsyn.IExpression) intsyn.IConstructorInvocationExpression {
	return &ConstructorInvocationExpression{
		ConstructorReferenceExpression: *NewConstructorReferenceExpressionWithAll(lineNumber,
			_type.PtTypeVoid.(intsyn.IType), objectType, descriptor).(*ConstructorReferenceExpression),
		parameters: parameters,
	}
}

type ConstructorInvocationExpression struct {
	ConstructorReferenceExpression

	parameters intsyn.IExpression
}

func (e *ConstructorInvocationExpression) Parameters() intsyn.IExpression {
	return e.parameters
}

func (e *ConstructorInvocationExpression) Priority() int {
	return 1
}

func (e *ConstructorInvocationExpression) SetParameters(params intsyn.IExpression) {
	e.parameters = params
}

func (e *ConstructorInvocationExpression) IsConstructorInvocationExpression() bool {
	return true
}

func (e *ConstructorInvocationExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitConstructorInvocationExpression(e)
}

func (e *ConstructorInvocationExpression) String() string {
	return fmt.Sprintf("ConstructorInvocationExpression{call this(%s)}", e.descriptor)
}
