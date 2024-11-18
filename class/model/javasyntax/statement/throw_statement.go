package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewThrowStatement(expression intsyn.IExpression) intsyn.IThrowStatement {
	return &ThrowStatement{
		expression: expression,
	}
}

type ThrowStatement struct {
	AbstractStatement

	expression intsyn.IExpression
}

func (s *ThrowStatement) Expression() intsyn.IExpression {
	return s.expression
}

func (s *ThrowStatement) SetExpression(expression intsyn.IExpression) {
	s.expression = expression
}

func (s *ThrowStatement) IsThrowStatement() bool {
	return true
}

func (s *ThrowStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitThrowStatement(s)
}

func (s *ThrowStatement) String() string {
	return fmt.Sprintf("ThrowStatement{throw %s}", s.expression)
}
