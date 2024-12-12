package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileMonitorEnterStatement(monitor intmod.IExpression) intsrv.IClassFileMonitorEnterStatement {
	s := &ClassFileMonitorEnterStatement{
		CommentStatement: *statement.NewCommentStatement(fmt.Sprintf("/* monitor enter %s */",
			monitor)).(*statement.CommentStatement),
		monitor: monitor,
	}
	s.SetValue(s)
	return s
}

type ClassFileMonitorEnterStatement struct {
	statement.CommentStatement

	monitor intmod.IExpression
}

func (s *ClassFileMonitorEnterStatement) Monitor() intmod.IExpression {
	return s.monitor
}

func (s *ClassFileMonitorEnterStatement) IsMonitorEnterStatement() bool {
	return true
}

func (s *ClassFileMonitorEnterStatement) String() string {
	return fmt.Sprintf("ClassFileMonitorEnterStatement{%s}", s.monitor)
}
