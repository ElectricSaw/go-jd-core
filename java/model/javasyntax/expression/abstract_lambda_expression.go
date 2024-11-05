package expression

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
)

func NewAbstractLambdaExpression(typ _type.IType, statements statement.Statement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		statements:                       statements,
	}
}

func NewAbstractLambdaExpressionWithAll(lineNumber int, typ _type.IType, statements statement.Statement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		statements:                       statements,
	}
}

type AbstractLambdaExpression struct {
	AbstractLineNumberTypeExpression

	statements statement.Statement
}

func (e *AbstractLambdaExpression) GetPriority() int {
	return 17
}

func (e *AbstractLambdaExpression) GetStatements() statement.Statement {
	return e.statements
}