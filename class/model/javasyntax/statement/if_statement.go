package statement

import intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewIfStatement(condition intsyn.IExpression, statements intsyn.IStatement) intsyn.IIfStatement {
	return &IfStatement{
		condition:  condition,
		statements: statements,
	}
}

type IfStatement struct {
	AbstractStatement

	condition  intsyn.IExpression
	statements intsyn.IStatement
}

func (s *IfStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *IfStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *IfStatement) Statements() intsyn.IStatement {
	return s.statements
}

func (s *IfStatement) IsIfStatement() bool {
	return s.condition != nil
}

func (s *IfStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitIfStatement(s)
}
