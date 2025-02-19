package statement

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileTryStatement(tryStatements intmod.IStatement, catchClauses []intmod.ICatchClause,
	finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	return NewClassFileTryStatement2(nil, tryStatements, catchClauses, finallyStatements, jsr, eclipse)
}

func NewClassFileTryStatement2(resources []intmod.IResource, tryStatements intmod.IStatement,
	catchClauses []intmod.ICatchClause, finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	tryStatement := *model.NewTryStatementWithAll(resources, tryStatements,
		catchClauses, finallyStatements).(*model.TryStatement)
	s := &ClassFileTryStatement{
		TryStatement: tryStatement,
		jsr:          jsr,
		eclipse:      eclipse,
	}
	s.SetValue(s)
	return s
}

type ClassFileTryStatement struct {
	model.TryStatement

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
		CatchClause:   *model.NewCatchClause(lineNumber, typ, "", statements).(*model.CatchClause),
		localVariable: localVariable,
	}
}

type CatchClause struct {
	model.CatchClause

	localVariable intsrv.ILocalVariable
}

func (c *CatchClause) Name() string {
	return c.localVariable.Name()
}

func (c *CatchClause) LocalVariable() intsrv.ILocalVariable {
	return c.localVariable
}
