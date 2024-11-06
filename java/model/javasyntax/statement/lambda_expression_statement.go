package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	"fmt"
)

func NewLambdaExpressionStatement(expression expression.Expression) *LambdaExpressionStatement {
	return &LambdaExpressionStatement{
		expression: expression,
	}
}

type LambdaExpressionStatement struct {
	AbstractStatement

	expression expression.Expression
}

func (s *LambdaExpressionStatement) GetExpression() expression.Expression {
	return s.expression
}

func (s *LambdaExpressionStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *LambdaExpressionStatement) IsLambdaExpressionStatement() bool {
	return true
}

func (s *LambdaExpressionStatement) Accept(visitor StatementVisitor) {
	visitor.VisitLambdaExpressionStatement(s)
}

func (s *LambdaExpressionStatement) String() string {
	return fmt.Sprintf("LambdaExpressionStatement{%s}", s.expression)
}
