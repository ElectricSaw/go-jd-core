package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
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
