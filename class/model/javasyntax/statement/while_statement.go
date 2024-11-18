package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"

func NewWhileStatement(condition intsyn.IExpression, statements intsyn.IStatement) intsyn.IWhileStatement {
	return &WhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type WhileStatement struct {
	AbstractStatement

	condition  intsyn.IExpression
	statements intsyn.IStatement
}

func (s *WhileStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *WhileStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *WhileStatement) Statements() intsyn.IStatement {
	return s.statements
}

func (s *WhileStatement) IsWhileStatement() bool {
	return true
}

func (s *WhileStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitWhileStatement(s)
}
