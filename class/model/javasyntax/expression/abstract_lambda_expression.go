package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
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

func (e *AbstractLambdaExpression) Priority() int {
	return 17
}

func (e *AbstractLambdaExpression) Statements() statement.Statement {
	return e.statements
}
