package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileForEachStatement(localVariable intsrv.ILocalVariable, expr intmod.IExpression,
	state intmod.IStatement) intsrv.IClassFileForEachStatement {
	s := &ClassFileForEachStatement{
		ForEachStatement: *statement.NewForEachStatement(localVariable.Type(), "", expr,
			state).(*statement.ForEachStatement),
		localVariable: localVariable,
	}
	s.SetValue(s)
	return s
}

type ClassFileForEachStatement struct {
	statement.ForEachStatement

	localVariable intsrv.ILocalVariable
}

func (s *ClassFileForEachStatement) Name() string {
	return s.localVariable.Name()
}

func (s *ClassFileForEachStatement) String() string {
	return fmt.Sprintf("ClassFileForEachStatement{%s %s : %s", s.Type(), s.localVariable.Name(), s.Expression())
}
