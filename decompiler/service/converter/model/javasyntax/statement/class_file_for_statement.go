package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
)

func NewClassFileForStatement(fromOffset, toOffset int, init intmod.IExpression,
	condition intmod.IExpression, update intmod.IExpression, state intmod.IStatement) intsrv.IClassFileForStatement {
	s := &ClassFileForStatement{
		ForStatement: *statement.NewForStatementWithInit(init, condition, update, state).(*statement.ForStatement),
		fromOffset:   fromOffset,
		toOffset:     toOffset,
	}
	s.SetValue(s)
	return s
}

type ClassFileForStatement struct {
	statement.ForStatement

	fromOffset int
	toOffset   int
}

func (s *ClassFileForStatement) FromOffset() int {
	return s.fromOffset
}

func (s *ClassFileForStatement) ToOffset() int {
	return s.toOffset
}

func (s *ClassFileForStatement) IsForStatement() bool {
	return true
}

func (s *ClassFileForStatement) String() string {
	sb := "ClassFileForStatement{"

	if s.Declaration() != nil {
		sb += fmt.Sprintf("%v", s.Declaration())
	} else {
		sb += fmt.Sprintf("%s", s.Init())
	}

	sb += fmt.Sprintf("; %s; %s}", s.Condition(), s.Update())

	return sb
}
