package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileMonitorExitStatement(monitor intmod.IExpression) intsrv.IClassFileMonitorExitStatement {
	s := &ClassFileMonitorExitStatement{
		CommentStatement: *statement.NewCommentStatement(fmt.Sprintf("/* monitor exit %s */", monitor)).(*statement.CommentStatement),
		monitor:          monitor,
	}
	s.SetValue(s)
	return s
}

type ClassFileMonitorExitStatement struct {
	statement.CommentStatement

	monitor intmod.IExpression
}

func (s *ClassFileMonitorExitStatement) Monitor() intmod.IExpression {
	return s.monitor
}

func (s *ClassFileMonitorExitStatement) IsMonitorExitStatement() bool {
	return true
}

func (s *ClassFileMonitorExitStatement) String() string {
	return fmt.Sprintf("ClassFileMonitorExitStatement{%s}", s.monitor)
}
