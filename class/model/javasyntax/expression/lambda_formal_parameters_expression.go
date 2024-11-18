package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewLambdaFormalParametersExpression(typ intsyn.IType, formalParameters intsyn.IFormalParameter,
	statements intsyn.IStatement) intsyn.ILambdaFormalParametersExpression {
	return &LambdaFormalParametersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpression(typ, statements),
		formalParameters:         formalParameters,
	}
}

func NewLambdaFormalParametersExpressionWithAll(lineNumber int, typ intsyn.IType,
	formalParameters intsyn.IFormalParameter, statements intsyn.IStatement) intsyn.ILambdaFormalParametersExpression {
	return &LambdaFormalParametersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		formalParameters:         formalParameters,
	}
}

type LambdaFormalParametersExpression struct {
	AbstractLambdaExpression

	formalParameters intsyn.IFormalParameter
}

func (e *LambdaFormalParametersExpression) FormalParameters() intsyn.IFormalParameter {
	return e.formalParameters
}

func (e *LambdaFormalParametersExpression) SetFormalParameters(formalParameters intsyn.IFormalParameter) {
	e.formalParameters = formalParameters
}

func (e *LambdaFormalParametersExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitLambdaFormalParametersExpression(e)
}

func (e *LambdaFormalParametersExpression) String() string {
	return fmt.Sprintf("LambdaFormalParametersExpression{%s -> %d}", e.formalParameters, e.statements)
}
