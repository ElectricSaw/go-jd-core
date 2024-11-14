package expression

type Expressions struct {
	AbstractExpression

	Expressions []Expression
}

func (e *Expressions) List() []Expression {
	return e.Expressions
}

func (e *Expressions) Accept(visitor ExpressionVisitor) {
	visitor.VisitExpressions(e)
}
