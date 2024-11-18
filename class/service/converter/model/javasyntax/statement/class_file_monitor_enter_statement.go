package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileMonitorEnterStatement(monitor expression.IExpression) *ClassFileMonitorEnterStatement {
	return &ClassFileMonitorEnterStatement{
		CommentStatement: *statement.NewCommentStatement(fmt.Sprintf("/* monitor enter %s */", monitor)),
		monitor:          monitor,
	}
}

type ClassFileMonitorEnterStatement struct {
	statement.CommentStatement

	monitor expression.IExpression
}

func (s *ClassFileMonitorEnterStatement) Monitor() expression.IExpression {
	return s.monitor
}

func (s *ClassFileMonitorEnterStatement) IsMonitorEnterStatement() bool {
	return true
}

func (s *ClassFileMonitorEnterStatement) String() string {
	return fmt.Sprintf("ClassFileMonitorEnterStatement{%s}", s.monitor)
}
