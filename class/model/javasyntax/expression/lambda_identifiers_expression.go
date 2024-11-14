package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewLambdaIdentifiersExpression(typ _type.IType, returnedType _type.IType, paramNames []string, statements statement.Statement) *LambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpression(typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

func NewLambdaIdentifiersExpressionWithAll(lineNumber int, typ _type.IType, returnedType _type.IType, paramNames []string, statements statement.Statement) *LambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

type LambdaIdentifiersExpression struct {
	AbstractLambdaExpression

	returnedType   _type.IType
	parameterNames []string
}

func (e *LambdaIdentifiersExpression) ReturnedType() _type.IType {
	return e.returnedType
}

func (e *LambdaIdentifiersExpression) ParameterNames() []string {
	return e.parameterNames
}

func (e *LambdaIdentifiersExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitLambdaIdentifiersExpression(e)
}

func (e *LambdaIdentifiersExpression) String() string {
	return fmt.Sprintf("LambdaIdentifiersExpression{%s -> %d}", e.parameterNames, e.statements)
}
