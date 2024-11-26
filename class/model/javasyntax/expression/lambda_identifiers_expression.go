package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLambdaIdentifiersExpression(typ intmod.IType, returnedType intmod.IType,
	paramNames []string, statements intmod.IStatement) intmod.ILambdaIdentifiersExpression {
	return NewLambdaIdentifiersExpressionWithAll(0, typ, returnedType, paramNames, statements)
}

func NewLambdaIdentifiersExpressionWithAll(lineNumber int, typ intmod.IType,
	returnedType intmod.IType, paramNames []string, statements intmod.IStatement) intmod.ILambdaIdentifiersExpression {
	e := &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
	e.SetValue(e)
	return e
}

type LambdaIdentifiersExpression struct {
	AbstractLambdaExpression
	util.DefaultBase[intmod.ILambdaIdentifiersExpression]

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
