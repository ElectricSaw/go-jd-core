package statement

import "bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"

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

func (s *AssertStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *AssertStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *AssertStatement) GetMessage() expression.Expression {
	return s.message
}

func (s *AssertStatement) Accept(visitor StatementVisitor) {
	visitor.VisitAssertStatement(s)
}
