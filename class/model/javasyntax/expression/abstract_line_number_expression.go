package expression

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewAbstractLineNumberExpression(lineNumber int) *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: lineNumber,
	}
}

func NewAbstractLineNumberExpressionEmpty() *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: intsyn.UnknownLineNumber,
	}
}

type AbstractLineNumberExpression struct {
	AbstractExpression

	lineNumber int
}

func (e *AbstractLineNumberExpression) LineNumber() int {
	return e.lineNumber
}

func (e *AbstractLineNumberExpression) Priority() int {
	return 0
}
