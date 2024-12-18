package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
)

func NewClassFileTryStatement(tryStatements intmod.IStatement, catchClauses []intmod.ICatchClause,
	finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	return NewClassFileTryStatement2(nil, tryStatements, catchClauses, finallyStatements, jsr, eclipse)
}

func NewClassFileTryStatement2(resources []intmod.IResource, tryStatements intmod.IStatement,
	catchClauses []intmod.ICatchClause, finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	tryStatement := *statement.NewTryStatementWithAll(resources, tryStatements,
		catchClauses, finallyStatements).(*statement.TryStatement)
	s := &ClassFileTryStatement{
		TryStatement: tryStatement,
		jsr:          jsr,
		eclipse:      eclipse,
	}
	s.SetValue(s)
	return s
}

type ClassFileTryStatement struct {
	statement.TryStatement

	jsr     bool
	eclipse bool
}

func (s *ClassFileTryStatement) IsJsr() bool {
	return s.jsr
}

func (s *ClassFileTryStatement) IsEclipse() bool {
	return s.eclipse
}

func NewCatchClause(lineNumber int, typ intmod.IObjectType, localVariable intsrv.ILocalVariable, statements intmod.IStatement) intsrv.ICatchClause {
	return &CatchClause{
		CatchClause:   *statement.NewCatchClause(lineNumber, typ, "", statements).(*statement.CatchClause),
		localVariable: localVariable,
	}
}

type CatchClause struct {
	statement.CatchClause

	localVariable intsrv.ILocalVariable
}

func (c *CatchClause) Name() string {
	return c.localVariable.Name()
}

func (c *CatchClause) LocalVariable() intsrv.ILocalVariable {
	return c.localVariable
}
