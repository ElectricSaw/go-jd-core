package statement

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewExpressionStatement(expression intsyn.IExpression) intsyn.IExpressionStatement {
	return &ExpressionStatement{
		expression: expression,
	}
}

type ExpressionStatement struct {
	AbstractStatement

	expression intsyn.IExpression
}

func (s *ExpressionStatement) Expression() intsyn.IExpression {
	return s.expression
}

func (s *ExpressionStatement) SetExpression(expression intsyn.IExpression) {
	s.expression = expression
}

func (s *ExpressionStatement) IsExpressionStatement() bool {
	return true
}

func (s *ExpressionStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{%s}", s.expression)
}
