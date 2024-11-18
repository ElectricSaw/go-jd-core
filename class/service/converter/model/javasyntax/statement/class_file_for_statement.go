package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileForStatement(fromOffset, toOffset int, init expression.IExpression,
	condition expression.IExpression, update expression.IExpression, state statement.IStatement) *ClassFileForStatement {
	return &ClassFileForStatement{
		ForStatement: *statement.NewForStatementWithInit(init, condition, update, state),
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
