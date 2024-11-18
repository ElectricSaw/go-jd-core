package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewAssertStatement(condition intsyn.IExpression, message intsyn.IExpression) intsyn.IAssertStatement {
	return &AssertStatement{
		condition: condition,
		message:   message,
	}
}

type AssertStatement struct {
	AbstractStatement

	condition intsyn.IExpression
	message   intsyn.IExpression
}

func (s *AssertStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *AssertStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *AssertStatement) Message() intsyn.IExpression {
	return s.message
}

func (s *AssertStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitAssertStatement(s)
}
