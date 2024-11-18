package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/statement"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/model/localvariable"
)

func NewClassFileTryStatement(tryStatements statement.IStatement, catchClauses []statement.CatchClause,
	finallyStatements statement.IStatement, jsr, eclipse bool) *ClassFileTryStatement {
	return &ClassFileTryStatement{
		TryStatement: *statement.NewTryStatement(tryStatements, catchClauses, finallyStatements),
		jsr:          jsr,
		eclipse:      eclipse,
	}
}

func NewClassFileTryStatement2(resources []statement.Resource, tryStatements statement.IStatement,
	catchClauses []statement.CatchClause, finallyStatements statement.IStatement, jsr, eclipse bool) *ClassFileTryStatement {
	return &ClassFileTryStatement{
		TryStatement: *statement.NewTryStatementWithAll(resources, tryStatements, catchClauses, finallyStatements),
		jsr:          jsr,
		eclipse:      eclipse,
	}
}

type ClassFileTryStatement struct {
	statement.TryStatement

	jsr     bool
	eclipse bool
}

func (s *ClassFileTryStatement) addResources(resources []statement.Resource) {
	if resources != nil {
		if s.Resources() == nil {
			s.SetResources(s.Resources())
		} else {
			s.AddResources(resources)
		}
	}
}

func (s *ClassFileTryStatement) isJsr() bool {
	return s.jsr
}

func (s *ClassFileTryStatement) isEclipse() bool {
	return s.eclipse
}

func NewCatchClause(lineNumber int, typ _type.ObjectType, localVariable localvariable.ILocalVariable, statements statement.IStatement) *CatchClause {
	return &CatchClause{
		CatchClause:   *statement.NewCatchClause(lineNumber, typ, "", statements),
		localVariable: localVariable,
	}
}

type CatchClause struct {
	statement.CatchClause

	localVariable localvariable.ILocalVariable
}

func (c *CatchClause) Name() string {
	return c.localVariable.Name()
}

func (c *CatchClause) LocalVariable() localvariable.ILocalVariable {
	return c.localVariable
}
