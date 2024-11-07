package declaration

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewExpressionVariableInitializer(expression expression.Expression) *ExpressionVariableInitializer {
	return &ExpressionVariableInitializer{
		expression: expression,
	}
}

type ExpressionVariableInitializer struct {
	AbstractVariableInitializer

	expression expression.Expression
}

func (i *ExpressionVariableInitializer) GetExpression() expression.Expression {
	return i.expression
}

func (i *ExpressionVariableInitializer) GetLineNumber() int {
	return i.expression.GetLineNumber()
}

func (i *ExpressionVariableInitializer) SetExpression(expression expression.Expression) {
	i.expression = expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) Accept(visitor DeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}
