package statement

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

func NewWhileStatement(condition intmod.IExpression, statements intmod.IStatement) intmod.IWhileStatement {
	return &WhileStatement{
		condition:  condition,
		statements: statements,
	}
}

type WhileStatement struct {
	AbstractStatement

	condition  intmod.IExpression
	statements intmod.IStatement
}

func (s *WhileStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *WhileStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *WhileStatement) Statements() intmod.IStatement {
	return s.statements
}

func (s *WhileStatement) IsWhileStatement() bool {
	return true
}

func (s *WhileStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitWhileStatement(s)
}
