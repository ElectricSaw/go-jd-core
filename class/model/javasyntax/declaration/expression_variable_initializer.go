package declaration

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewExpressionVariableInitializer(expression intsyn.IExpression) intsyn.IExpressionVariableInitializer {
	return &ExpressionVariableInitializer{
		expression: expression,
	}
}

type ExpressionVariableInitializer struct {
	AbstractVariableInitializer

	expression intsyn.IExpression
}

func (i *ExpressionVariableInitializer) Expression() intsyn.IExpression {
	return i.expression
}

func (i *ExpressionVariableInitializer) LineNumber() int {
	return i.expression.LineNumber()
}

func (i *ExpressionVariableInitializer) SetExpression(expression intsyn.IExpression) {
	i.expression = expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) Accept(visitor intsyn.IDeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}
