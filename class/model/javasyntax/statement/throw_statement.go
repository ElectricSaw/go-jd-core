package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewThrowStatement(expression intmod.IExpression) intmod.IThrowStatement {
	return &ThrowStatement{
		expression: expression,
	}
}

type ThrowStatement struct {
	AbstractStatement

	expression intmod.IExpression
}

func (s *ThrowStatement) Expression() intmod.IExpression {
	return s.expression
}

func (s *ThrowStatement) SetExpression(expression intmod.IExpression) {
	s.expression = expression
}

func (s *ThrowStatement) IsThrowStatement() bool {
	return true
}

func (s *ThrowStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitThrowStatement(s)
}

func (s *ThrowStatement) String() string {
	return fmt.Sprintf("ThrowStatement{throw %s}", s.expression)
}
