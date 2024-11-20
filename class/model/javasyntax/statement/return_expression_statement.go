package statement

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewReturnExpressionStatement(expression intsyn.IExpression) intsyn.IReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: expression.LineNumber(),
		expression: expression,
	}
}

func NewReturnExpressionStatementWithAll(lineNumber int, expression intsyn.IExpression) intsyn.IReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: lineNumber,
		expression: expression,
	}
}

type ReturnExpressionStatement struct {
	AbstractStatement

	lineNumber int
	expression intsyn.IExpression
}

func (s *ReturnExpressionStatement) LineNumber() int {
	return s.lineNumber
}

func (s *ReturnExpressionStatement) SetLineNumber(lineNumber int) {
	s.lineNumber = lineNumber
}

func (s *ReturnExpressionStatement) Expression() intsyn.IExpression {
	return s.expression
}

func (s *ReturnExpressionStatement) SetExpression(expression intsyn.IExpression) {
	s.expression = expression
}

func (s *ReturnExpressionStatement) GenericExpression() intsyn.IExpression {
	return s.expression
}

func (s *ReturnExpressionStatement) IsReturnExpressionStatement() bool {
	return true
}

func (s *ReturnExpressionStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitReturnExpressionStatement(s)
}

func (s *ReturnExpressionStatement) String() string {
	return fmt.Sprintf("ReturnExpressionStatement{return %s}", s.expression)
}
