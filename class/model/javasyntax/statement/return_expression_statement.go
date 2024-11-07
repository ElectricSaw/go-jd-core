package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"fmt"
)

func NewReturnExpressionStatement(expression expression.Expression) *ReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: expression.GetLineNumber(),
		expression: expression,
	}
}

func NewReturnExpressionStatementWithAll(lineNumber int, expression expression.Expression) *ReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: lineNumber,
		expression: expression,
	}
}

type ReturnExpressionStatement struct {
	AbstractStatement

	lineNumber int
	expression expression.Expression
}

func (s *ReturnExpressionStatement) GetLineNumber() int {
	return s.lineNumber
}

func (s *ReturnExpressionStatement) SetLineNumber(lineNumber int) {
	s.lineNumber = lineNumber
}

func (s *ReturnExpressionStatement) GetExpression() expression.Expression {
	return s.expression
}

func (s *ReturnExpressionStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *ReturnExpressionStatement) GetGenericExpression() expression.Expression {
	return s.expression
}

func (s *ReturnExpressionStatement) IsReturnExpressionStatement() bool {
	return true
}

func (s *ReturnExpressionStatement) Accept(visitor StatementVisitor) {
	visitor.VisitReturnExpressionStatement(s)
}

func (s *ReturnExpressionStatement) String() string {
	return fmt.Sprintf("ReturnExpressionStatement{return %s}", s.expression)
}
