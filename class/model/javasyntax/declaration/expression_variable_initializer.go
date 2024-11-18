package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
)

func NewExpressionVariableInitializer(expression expression.Expression) intsyn.IExpressionVariableInitializer {
	return &ExpressionVariableInitializer{
		expression: expression,
	}
}

type ExpressionVariableInitializer struct {
	AbstractVariableInitializer

	expression expression.Expression
}

func (i *ExpressionVariableInitializer) Expression() expression.Expression {
	return i.expression
}

func (i *ExpressionVariableInitializer) LineNumber() int {
	return i.expression.LineNumber()
}

func (i *ExpressionVariableInitializer) SetExpression(expression expression.Expression) {
	i.expression = expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}
