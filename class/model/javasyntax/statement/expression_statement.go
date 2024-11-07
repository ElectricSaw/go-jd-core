package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"fmt"
)

func NewExpressionStatement(expression expression.Expression) *ExpressionStatement {
	return &ExpressionStatement{
		expression: expression,
	}
}

type ExpressionStatement struct {
	AbstractStatement

	expression expression.Expression
}

func (s *ExpressionStatement) GetExpression() expression.Expression {
	return s.expression
}

func (s *ExpressionStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *ExpressionStatement) IsExpressionStatement() bool {
	return true
}

func (s *ExpressionStatement) Accept(visitor StatementVisitor) {
	visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{%s}", s.expression)
}
