package expression

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewAbstractLineNumberExpression(lineNumber int) *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: lineNumber,
	}
}

func NewAbstractLineNumberExpressionEmpty() *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: intmod.UnknownLineNumber,
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
