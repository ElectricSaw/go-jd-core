package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileMonitorEnterStatement(monitor intmod.IExpression) intsrv.IClassFileMonitorEnterStatement {
	s := &ClassFileMonitorEnterStatement{
		CommentStatement: *model.NewCommentStatement(fmt.Sprintf("/* monitor enter %s */",
			monitor)).(*model.CommentStatement),
		monitor: monitor,
	}
	s.SetValue(s)
	return s
}

type ClassFileMonitorEnterStatement struct {
	model.CommentStatement

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
