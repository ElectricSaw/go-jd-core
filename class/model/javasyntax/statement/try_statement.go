package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
)

func NewTryStatement(tryStatements intsyn.IStatement, catchClauses []intsyn.ICatchClause, finallyStatement intsyn.IStatement) intsyn.ITryStatement {
	return &TryStatement{
		resources:        nil,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

func NewTryStatementWithAll(resource []intsyn.IResource, tryStatements intsyn.IStatement, catchClauses []intsyn.ICatchClause, finallyStatement intsyn.IStatement) intsyn.ITryStatement {
	return &TryStatement{
		resources:        resource,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

type TryStatement struct {
	AbstractStatement

	resources        []intsyn.IResource
	tryStatements    intsyn.IStatement
	catchClauses     []intsyn.ICatchClause
	finallyStatement intsyn.IStatement
}

func (s *TryStatement) ResourceList() []intsyn.IStatement {
	ret := make([]intsyn.IStatement, 0, len(s.resources))
	for _, resource := range s.resources {
		ret = append(ret, resource.(intsyn.IStatement))
	}
	return ret
}

func (s *TryStatement) Resources() []intsyn.IResource {
	return s.resources
}

func (s *TryStatement) SetResources(resources []intsyn.IResource) {
	s.resources = resources
}

func (s *TryStatement) AddResources(resources []intsyn.IResource) {
	s.resources = append(s.resources, resources...)
}

func (s *TryStatement) AddResource(resource intsyn.IResource) {
	s.resources = append(s.resources, resource)
}

func (s *TryStatement) TryStatements() intsyn.IStatement {
	return s.tryStatements
}

func (s *TryStatement) SetTryStatements(tryStatement intsyn.IStatement) {
	s.tryStatements = tryStatement
}

func (s *TryStatement) CatchClauseList() []intsyn.IStatement {
	ret := make([]intsyn.IStatement, 0, len(s.catchClauses))
	for _, resource := range s.catchClauses {
		ret = append(ret, resource.(intsyn.IStatement))
	}
	return ret
}

func (s *TryStatement) CatchClauses() []intsyn.ICatchClause {
	return s.catchClauses
}

func (s *TryStatement) FinallyStatements() intsyn.IStatement {
	return s.finallyStatement
}

func (s *TryStatement) SetFinallyStatement(finallyStatement intsyn.IStatement) {
	s.finallyStatement = finallyStatement
}

func (s *TryStatement) IsTryStatement() bool {
	return true
}

func (s *TryStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitTryStatement(s)
}

func NewResource(typ intsyn.IObjectType, name string, expression intsyn.IExpression) intsyn.IResource {
	return &Resource{
		typ:        typ,
		name:       name,
		expression: expression,
	}
}

type Resource struct {
	AbstractStatement

	typ        intsyn.IObjectType
	name       string
	expression intsyn.IExpression
}

func (r *Resource) Type() intsyn.IObjectType {
	return r.typ
}

func (r *Resource) Name() string {
	return r.name
}

func (r *Resource) Expression() intsyn.IExpression {
	return r.expression
}

func (r *Resource) SetExpression(expression intsyn.IExpression) {
	r.expression = expression
}

func (r *Resource) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitTryStatementResource(r)
}

func NewCatchClause(lineNumber int, typ intsyn.IObjectType, name string, statements intsyn.IStatement) intsyn.ICatchClause {
	return &CatchClause{
		lineNumber: lineNumber,
		typ:        typ,
		name:       name,
		statements: statements,
	}
}

type CatchClause struct {
	AbstractStatement

	lineNumber int
	typ        intsyn.IObjectType
	otherType  []intsyn.IObjectType
	name       string
	statements intsyn.IStatement
}

func (c *CatchClause) LineNumber() int {
	return c.lineNumber
}

func (c *CatchClause) Type() intsyn.IObjectType {
	return c.typ
}

func (c *CatchClause) OtherType() []intsyn.IObjectType {
	return c.otherType
}

func (c *CatchClause) Name() string {
	return c.name
}

func (c *CatchClause) Statements() intsyn.IStatement {
	return c.statements
}

func (c *CatchClause) AddType(typ intsyn.IObjectType) {
	c.otherType = append(c.otherType, typ)
}

func (c *CatchClause) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitTryStatementCatchClause(c)
}
