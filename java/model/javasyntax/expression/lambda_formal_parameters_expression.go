package expression

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/declaration"
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewLambdaFormalParametersExpression(typ _type.IType, formalParameters declaration.IFormalParameter, statements statement.Statement) *LambdaFormalParametersExpression {
	return &LambdaFormalParametersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpression(typ, statements),
		formalParameters:         formalParameters,
	}
}

func NewLambdaFormalParametersExpressionWithAll(lineNumber int, typ _type.IType, formalParameters declaration.IFormalParameter, statements statement.Statement) *LambdaFormalParametersExpression {
	return &LambdaFormalParametersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		formalParameters:         formalParameters,
	}
}

type LambdaFormalParametersExpression struct {
	AbstractLambdaExpression

	formalParameters declaration.IFormalParameter
}

func (e *LambdaFormalParametersExpression) GetFormalParameters() declaration.IFormalParameter {
	return e.formalParameters
}

func (e *LambdaFormalParametersExpression) SetFormalParameters(formalParameters declaration.IFormalParameter) {
	e.formalParameters = formalParameters
}

func (e *LambdaFormalParametersExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitLambdaFormalParametersExpression(e)
}

func (e *LambdaFormalParametersExpression) String() string {
	return fmt.Sprintf("LambdaFormalParametersExpression{%s -> %d}", e.formalParameters, e.statements)
}
