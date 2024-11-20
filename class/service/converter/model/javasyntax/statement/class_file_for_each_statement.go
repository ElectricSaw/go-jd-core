package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"fmt"
)

func NewClassFileForEachStatement(localVariable intsrv.ILocalVariable, expr intmod.IExpression,
	state intmod.IStatement) intsrv.IClassFileForEachStatement {
	return &ClassFileForEachStatement{
		ForEachStatement: *statement.NewForEachStatement(localVariable.Type(), "", expr,
			state).(*statement.ForEachStatement),
		localVariable: localVariable,
	}
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
