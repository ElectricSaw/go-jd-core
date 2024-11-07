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

func (e *AbstractLineNumberExpression) GetLineNumber() int {
	return e.lineNumber
}

func (e *AbstractLineNumberExpression) GetPriority() int {
	return 0
}
