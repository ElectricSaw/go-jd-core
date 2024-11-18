package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileBreakContinueStatement(offset int, targetOffset int) *ClassFileBreakContinueStatement {
	return &ClassFileBreakContinueStatement{
		offset:       offset,
		targetOffset: targetOffset,
	}
}

type ClassFileBreakContinueStatement struct {
	statement.AbstractStatement

	offset        int
	targetOffset  int
	continueLabel bool
	statement     statement.IStatement
}

func (s *ClassFileBreakContinueStatement) Offset() int {
	return s.offset
}

func (s *ClassFileBreakContinueStatement) TargetOffset() int {
	return s.targetOffset
}

func (s *ClassFileBreakContinueStatement) Statement() statement.IStatement {
	return s.statement
}

func (s *ClassFileBreakContinueStatement) SetStatement(statement statement.IStatement) {
	s.statement = statement
}

func (s *ClassFileBreakContinueStatement) IsContinueLabel() bool {
	return s.continueLabel
}

func (s *ClassFileBreakContinueStatement) SetContinueLabel(continueLabel bool) {
	s.continueLabel = continueLabel
}

func (s *ClassFileBreakContinueStatement) Accept(visitor statement.IStatementVisitor) {
	if s.statement == nil {
		s.statement.Accept(visitor)
	}
}

func (s *ClassFileBreakContinueStatement) String() string {
	if s.statement == nil {
		return "ClassFileBreakContinueStatement{}"
	}
	return fmt.Sprintf("ClassFileBreakContinueStatement{%s}", s.statement)
}
