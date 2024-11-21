package statement

import intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"

func NewIfStatement(condition intmod.IExpression, statements intmod.IStatement) intmod.IIfStatement {
	return &IfStatement{
		condition:  condition,
		statements: statements,
	}
}

type IfStatement struct {
	AbstractStatement

	condition  intmod.IExpression
	statements intmod.IStatement
}

func (s *IfStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *IfStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *IfStatement) Statements() intmod.IStatement {
	return s.statements
}

func (s *IfStatement) IsIfStatement() bool {
	return s.condition != nil
}

func (s *IfStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitIfStatement(s)
}
