package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileMonitorExitStatement(monitor intmod.IExpression) intsrv.IClassFileMonitorExitStatement {
	s := &ClassFileMonitorExitStatement{
		CommentStatement: *model.NewCommentStatement(fmt.Sprintf("/* monitor exit %s */", monitor)).(*model.CommentStatement),
		monitor:          monitor,
	}
	s.SetValue(s)
	return s
}

type ClassFileMonitorExitStatement struct {
	model.CommentStatement

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
