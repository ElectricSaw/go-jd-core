package statement

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLambdaExpressionStatement(expression intsyn.IExpression) intsyn.ILambdaExpressionStatement {
	return &LambdaExpressionStatement{
		expression: expression,
	}
}

type LambdaExpressionStatement struct {
	AbstractStatement

	expression intsyn.IExpression
}

func (s *LambdaExpressionStatement) Expression() intsyn.IExpression {
	return s.expression
}

func (s *LambdaExpressionStatement) SetExpression(expression intsyn.IExpression) {
	s.expression = expression
}

func (s *LambdaExpressionStatement) IsLambdaExpressionStatement() bool {
	return true
}

func (s *LambdaExpressionStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitLambdaExpressionStatement(s)
}

func (s *LambdaExpressionStatement) String() string {
	return fmt.Sprintf("LambdaExpressionStatement{%s}", s.expression)
}
