package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewAbstractLambdaExpression(typ intsyn.IType, statements intsyn.IStatement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		statements:                       statements,
	}
}

func NewAbstractLambdaExpressionWithAll(lineNumber int, typ intsyn.IType, statements intsyn.IStatement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		statements:                       statements,
	}
}

type AbstractLambdaExpression struct {
	AbstractLineNumberTypeExpression

	statements intsyn.IStatement
}

func (e *AbstractLambdaExpression) Priority() int {
	return 17
}

func (e *AbstractLambdaExpression) Statements() intsyn.IStatement {
	return e.statements
}
