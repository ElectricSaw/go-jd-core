package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLambdaIdentifiersExpression(typ intmod.IType, returnedType intmod.IType,
	paramNames []string, statements intmod.IStatement) intmod.ILambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpression(typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

func NewLambdaIdentifiersExpressionWithAll(lineNumber int, typ intmod.IType,
	returnedType intmod.IType, paramNames []string, statements intmod.IStatement) intmod.ILambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

type LambdaIdentifiersExpression struct {
	AbstractLambdaExpression

	returnedType   intmod.IType
	parameterNames []string
}

func (e *LambdaIdentifiersExpression) ReturnedType() intmod.IType {
	return e.returnedType
}

func (e *LambdaIdentifiersExpression) ParameterNames() []string {
	return e.parameterNames
}

func (e *LambdaIdentifiersExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLambdaIdentifiersExpression(e)
}

func (e *LambdaIdentifiersExpression) String() string {
	return fmt.Sprintf("LambdaIdentifiersExpression{%s -> %d}", e.parameterNames, e.statements)
}
