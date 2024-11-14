package expression

func NewAbstractLineNumberExpression(lineNumber int) *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: lineNumber,
	}
}

func NewAbstractLineNumberExpressionEmpty() *AbstractLineNumberExpression {
	return &AbstractLineNumberExpression{
		lineNumber: UnknownLineNumber,
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
