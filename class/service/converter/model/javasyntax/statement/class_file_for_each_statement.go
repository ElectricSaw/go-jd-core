package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
	"fmt"
)

func NewClassFileForEachStatement(localVariable localvariable.ILocalVariable, expr expression.IExpression,
	state statement.IStatement) *ClassFileForEachStatement {
	return &ClassFileForEachStatement{
		ForEachStatement: *statement.NewForEachStatement(localVariable.Type(), "", expr, state),
		localVariable:    localVariable,
	}
}

type ClassFileForEachStatement struct {
	statement.ForEachStatement

	localVariable localvariable.ILocalVariable
}

func (s *ClassFileForEachStatement) Name() string {
	return s.localVariable.Name()
}

func (s *ClassFileForEachStatement) String() string {
	return fmt.Sprintf("ClassFileForEachStatement{%s %s : %s", s.Type(), s.localVariable.Name(), s.Expression())
}
