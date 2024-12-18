package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewReturnExpressionStatement(expression intmod.IExpression) intmod.IReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: expression.LineNumber(),
		expression: expression,
	}
}

func NewReturnExpressionStatementWithAll(lineNumber int, expression intmod.IExpression) intmod.IReturnExpressionStatement {
	return &ReturnExpressionStatement{
		lineNumber: lineNumber,
		expression: expression,
	}
}

type ReturnExpressionStatement struct {
	AbstractStatement

	lineNumber int
	expression intmod.IExpression
}

func (s *ReturnExpressionStatement) LineNumber() int {
	return s.lineNumber
}

func (s *ReturnExpressionStatement) SetLineNumber(lineNumber int) {
	s.lineNumber = lineNumber
}

func (s *ReturnExpressionStatement) Expression() intmod.IExpression {
	return s.expression
}

func (s *ReturnExpressionStatement) SetExpression(expression intmod.IExpression) {
	s.expression = expression
}

func (s *ReturnExpressionStatement) GenericExpression() intmod.IExpression {
	return s.expression
}

func (s *ReturnExpressionStatement) IsReturnExpressionStatement() bool {
	return true
}

func (s *ReturnExpressionStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitReturnExpressionStatement(s)
}

func (s *ReturnExpressionStatement) String() string {
	return fmt.Sprintf("ReturnExpressionStatement{return %s}", s.expression)
}
