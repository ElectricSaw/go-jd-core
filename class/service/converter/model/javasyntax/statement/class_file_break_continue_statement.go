package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
)

func NewClassFileBreakContinueStatement(offset int, targetOffset int) intsrv.IClassFileBreakContinueStatement {
	s := &ClassFileBreakContinueStatement{
		offset:       offset,
		targetOffset: targetOffset,
	}
	s.SetStatement(s)
	return s
}

type ClassFileBreakContinueStatement struct {
	intmod.IStatement

	offset        int
	targetOffset  int
	continueLabel bool
	statement     intmod.IStatement
}

func (s *ClassFileBreakContinueStatement) Offset() int {
	return s.offset
}

func (s *ClassFileBreakContinueStatement) TargetOffset() int {
	return s.targetOffset
}

func (s *ClassFileBreakContinueStatement) Statement() intmod.IStatement {
	return s.statement
}

func (s *ClassFileBreakContinueStatement) SetStatement(statement intmod.IStatement) {
	s.statement = statement
}

func (s *ClassFileBreakContinueStatement) IsContinueLabel() bool {
	return s.continueLabel
}

func (s *ClassFileBreakContinueStatement) SetContinueLabel(continueLabel bool) {
	s.continueLabel = continueLabel
}

func (s *ClassFileBreakContinueStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	if s.statement == nil {
		s.statement.AcceptStatement(visitor)
	}
}

func (s *ClassFileBreakContinueStatement) String() string {
	if s.statement == nil {
		return "ClassFileBreakContinueStatement{}"
	}
	return fmt.Sprintf("ClassFileBreakContinueStatement{%s}", s.statement)
}
