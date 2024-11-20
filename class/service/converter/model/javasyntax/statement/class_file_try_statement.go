package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
)

func NewClassFileTryStatement(tryStatements intmod.IStatement, catchClauses []intmod.ICatchClause,
	finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	return &ClassFileTryStatement{
		TryStatement: *statement.NewTryStatement(tryStatements, catchClauses,
			finallyStatements).(*statement.TryStatement),
		jsr:     jsr,
		eclipse: eclipse,
	}
}

func NewClassFileTryStatement2(resources []intmod.IResource, tryStatements intmod.IStatement,
	catchClauses []intmod.ICatchClause, finallyStatements intmod.IStatement, jsr, eclipse bool) intsrv.IClassFileTryStatement {
	return &ClassFileTryStatement{
		TryStatement: *statement.NewTryStatementWithAll(resources, tryStatements,
			catchClauses, finallyStatements).(*statement.TryStatement),
		jsr:     jsr,
		eclipse: eclipse,
	}
}

type ClassFileTryStatement struct {
	statement.TryStatement

	jsr     bool
	eclipse bool
}

func (s *ClassFileTryStatement) AddResources(resources []intmod.IResource) {
	if resources != nil {
		if s.Resources() == nil {
			s.SetResources(s.Resources())
		} else {
			s.TryStatement.AddResources(resources)
		}
	}
}

func (s *ClassFileTryStatement) isJsr() bool {
	return s.jsr
}

func (s *ClassFileTryStatement) isEclipse() bool {
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
