package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewExpressionStatement(expression intmod.IExpression) intmod.IExpressionStatement {
	return &ExpressionStatement{
		expression: expression,
	}
}

type ExpressionStatement struct {
	AbstractStatement

	expression intmod.IExpression
}

func (s *ExpressionStatement) Expression() intmod.IExpression {
	return s.expression
}

func (s *ExpressionStatement) SetExpression(expression intmod.IExpression) {
	s.expression = expression
}

func (s *ExpressionStatement) IsExpressionStatement() bool {
	return true
}

func (s *ExpressionStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{%s}", s.expression)
}
