package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewExpressionVariableInitializer(expression Expression) ExpressionVariableInitializer {
	return &ExpressionVariableInitializer{
		expression: expression,
	}
}

type ExpressionVariableInitializer struct {
	AbstractVariableInitializer

	expression Expression
}

func (i *ExpressionVariableInitializer) Expression() Expression {
	return i.expression
}

func (i *ExpressionVariableInitializer) LineNumber() int {
	return i.expression.LineNumber()
}

func (i *ExpressionVariableInitializer) SetExpression(expression Expression) {
	i.expression = expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) AcceptDeclaration(visitor DeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}
