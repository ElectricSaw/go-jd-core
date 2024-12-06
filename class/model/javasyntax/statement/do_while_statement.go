package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewDoWhileStatement(condition intmod.IExpression, statements intmod.IStatement) intmod.IDoWhileStatement {
	return &DoWhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type DoWhileStatement struct {
	AbstractStatement

	condition  intmod.IExpression
	statements intmod.IStatement
}

func (s *DoWhileStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *DoWhileStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *DoWhileStatement) Statements() intmod.IStatement {
	return s.statements
}

func (s *DoWhileStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitDoWhileStatement(s)
}
