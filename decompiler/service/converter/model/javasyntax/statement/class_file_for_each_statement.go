package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileForEachStatement(localVariable intsrv.ILocalVariable, expr intmod.IExpression,
	state intmod.IStatement) intsrv.IClassFileForEachStatement {
	s := &ClassFileForEachStatement{
		ForEachStatement: *model.NewForEachStatement(localVariable.Type(), "", expr,
			state).(*model.ForEachStatement),
		localVariable: localVariable,
	}
	s.SetValue(s)
	return s
}

type ClassFileForEachStatement struct {
	model.ForEachStatement

	localVariable intsrv.ILocalVariable
}

func (s *ClassFileForEachStatement) Name() string {
	return s.localVariable.Name()
}

func (s *ClassFileForEachStatement) String() string {
	return fmt.Sprintf("ClassFileForEachStatement{%s %s : %s", s.Type(), s.localVariable.Name(), s.Expression())
}
