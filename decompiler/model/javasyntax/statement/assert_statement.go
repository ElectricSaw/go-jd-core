package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewAssertStatement(condition intmod.IExpression, message intmod.IExpression) intmod.IAssertStatement {
	return &AssertStatement{
		condition: condition,
		message:   message,
	}
}

type AssertStatement struct {
	AbstractStatement

	condition intmod.IExpression
	message   intmod.IExpression
}

func (s *AssertStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *AssertStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *AssertStatement) Message() intmod.IExpression {
	return s.message
}

func (s *AssertStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitAssertStatement(s)
}
