package expression

type Expressions struct {
	AbstractExpression
}

func (e *Expressions) Accept(visitor ExpressionVisitor) {
	visitor.VisitExpressions(e)
}
