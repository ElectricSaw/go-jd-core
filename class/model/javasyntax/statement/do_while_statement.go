package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewDoWhileStatement(condition intsyn.IExpression, statements intsyn.IStatement) intsyn.IDoWhileStatement {
	return &DoWhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type DoWhileStatement struct {
	AbstractStatement

	condition  intsyn.IExpression
	statements intsyn.IStatement
}

func (s *DoWhileStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *DoWhileStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *DoWhileStatement) Statements() intsyn.IStatement {
	return s.statements
}

func (s *DoWhileStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitDoWhileStatement(s)
}
