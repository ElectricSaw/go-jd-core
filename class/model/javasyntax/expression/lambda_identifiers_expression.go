package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewLambdaIdentifiersExpression(typ intsyn.IType, returnedType intsyn.IType,
	paramNames []string, statements intsyn.IStatement) intsyn.ILambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpression(typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

func NewLambdaIdentifiersExpressionWithAll(lineNumber int, typ intsyn.IType,
	returnedType intsyn.IType, paramNames []string, statements intsyn.IStatement) intsyn.ILambdaIdentifiersExpression {
	return &LambdaIdentifiersExpression{
		AbstractLambdaExpression: *NewAbstractLambdaExpressionWithAll(lineNumber, typ, statements),
		returnedType:             returnedType,
		parameterNames:           paramNames,
	}
}

type LambdaIdentifiersExpression struct {
	AbstractLambdaExpression

	returnedType   intsyn.IType
	parameterNames []string
}

func (e *LambdaIdentifiersExpression) ReturnedType() intsyn.IType {
	return e.returnedType
}

func (e *LambdaIdentifiersExpression) ParameterNames() []string {
	return e.parameterNames
}

func (e *LambdaIdentifiersExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitLambdaIdentifiersExpression(e)
}

func (e *LambdaIdentifiersExpression) String() string {
	return fmt.Sprintf("LambdaIdentifiersExpression{%s -> %d}", e.parameterNames, e.statements)
}
