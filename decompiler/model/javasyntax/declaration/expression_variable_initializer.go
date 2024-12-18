package declaration

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewExpressionVariableInitializer(expression intmod.IExpression) intmod.IExpressionVariableInitializer {
	return &ExpressionVariableInitializer{
		expression: expression,
	}
}

type ExpressionVariableInitializer struct {
	AbstractVariableInitializer

	expression intmod.IExpression
}

func (i *ExpressionVariableInitializer) Expression() intmod.IExpression {
	return i.expression
}

func (i *ExpressionVariableInitializer) LineNumber() int {
	return i.expression.LineNumber()
}

func (i *ExpressionVariableInitializer) SetExpression(expression intmod.IExpression) {
	i.expression = expression
}

func (i *ExpressionVariableInitializer) IsExpressionVariableInitializer() bool {
	return true
}

func (i *ExpressionVariableInitializer) AcceptDeclaration(visitor intmod.IDeclarationVisitor) {
	visitor.VisitExpressionVariableInitializer(i)
}
