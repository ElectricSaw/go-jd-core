package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewLambdaExpressionStatement(expression intmod.IExpression) intmod.ILambdaExpressionStatement {
	return &LambdaExpressionStatement{
		expression: expression,
	}
}

type LambdaExpressionStatement struct {
	AbstractStatement

	expression intmod.IExpression
}

func (s *LambdaExpressionStatement) Expression() intmod.IExpression {
	return s.expression
}

func (s *LambdaExpressionStatement) SetExpression(expression intmod.IExpression) {
	s.expression = expression
}

func (s *LambdaExpressionStatement) IsLambdaExpressionStatement() bool {
	return true
}

func (s *LambdaExpressionStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitLambdaExpressionStatement(s)
}

func (s *LambdaExpressionStatement) String() string {
	return fmt.Sprintf("LambdaExpressionStatement{%s}", s.expression)
}
