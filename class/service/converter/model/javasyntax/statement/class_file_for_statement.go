package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileForStatement(fromOffset, toOffset int, init intmod.IExpression,
	condition intmod.IExpression, update intmod.IExpression, state intmod.IStatement) intsrv.IClassFileForStatement {
	return &ClassFileForStatement{
		ForStatement: *statement.NewForStatementWithInit(init, condition, update, state).(*statement.ForStatement),
		fromOffset:   fromOffset,
		toOffset:     toOffset,
	}
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
