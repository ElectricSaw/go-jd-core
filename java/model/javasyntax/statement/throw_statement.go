package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	"fmt"
)

func NewThrowStatement(expression expression.Expression) *ThrowStatement {
	return &ThrowStatement{
		expression: expression,
	}
}

type ThrowStatement struct {
	AbstractStatement

	expression expression.Expression
}

func (s *ThrowStatement) GetExpression() expression.Expression {
	return s.expression
}

func (s *ThrowStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *ThrowStatement) IsThrowStatement() bool {
	return true
}

func (s *ThrowStatement) Accept(visitor StatementVisitor) {
	visitor.VisitThrowStatement(s)
}

func (s *ThrowStatement) String() string {
	return fmt.Sprintf("ThrowStatement{throw %s}", s.expression)
}
