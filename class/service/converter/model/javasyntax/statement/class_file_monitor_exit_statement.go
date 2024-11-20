package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileMonitorExitStatement() intsrv.IClassFileMonitorExitStatement {
	return &ClassFileMonitorExitStatement{}
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
