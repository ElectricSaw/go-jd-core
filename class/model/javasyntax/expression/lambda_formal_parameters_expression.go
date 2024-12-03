package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLambdaFormalParametersExpression(typ intmod.IType, formalParameters intmod.IFormalParameter,
	statements intmod.IStatement) intmod.ILambdaFormalParametersExpression {
	return NewLambdaFormalParametersExpressionWithAll(0, typ, formalParameters, statements)
}

func NewLambdaFormalParametersExpressionWithAll(lineNumber int, typ intmod.IType,
	formalParameters intmod.IFormalParameter, statements intmod.IStatement) intmod.ILambdaFormalParametersExpression {
	e := &LambdaFormalParametersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		formalParameters:         formalParameters,
	}
	e.SetValue(e)
	return e
}

type LambdaFormalParametersExpression struct {
	AbstractLambdaExpression

	formalParameters intmod.IFormalParameter
}

func (e *LambdaFormalParametersExpression) FormalParameters() intmod.IFormalParameter {
	return e.formalParameters
}

func (e *LambdaFormalParametersExpression) SetFormalParameters(formalParameters intmod.IFormalParameter) {
	e.formalParameters = formalParameters
}

func (e *LambdaFormalParametersExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLambdaFormalParametersExpression(e)
}

func (e *LambdaFormalParametersExpression) String() string {
	return fmt.Sprintf("LambdaFormalParametersExpression{%s -> %d}", e.formalParameters, e.statements)
}
