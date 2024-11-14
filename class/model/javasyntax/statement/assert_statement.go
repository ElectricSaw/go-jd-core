package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

func NewAssertStatement(condition expression.Expression, message expression.Expression) *AssertStatement {
	return &AssertStatement{
		condition: condition,
		message:   message,
	}
}

type AssertStatement struct {
	AbstractStatement

	condition expression.Expression
	message   expression.Expression
}

func (s *AssertStatement) Condition() expression.Expression {
	return s.condition
}

func (s *AssertStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *AssertStatement) Message() expression.Expression {
	return s.message
}

func (s *AssertStatement) Accept(visitor StatementVisitor) {
	visitor.VisitAssertStatement(s)
}
