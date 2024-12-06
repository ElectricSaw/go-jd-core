package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewTryStatement(tryStatements intmod.IStatement, catchClauses []intmod.ICatchClause, finallyStatement intmod.IStatement) intmod.ITryStatement {
	return &TryStatement{
		resources:        nil,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

func NewTryStatementWithAll(resource []intmod.IResource, tryStatements intmod.IStatement, catchClauses []intmod.ICatchClause, finallyStatement intmod.IStatement) intmod.ITryStatement {
	return &TryStatement{
		resources:        resource,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

type TryStatement struct {
	AbstractStatement

	resources        []intmod.IResource
	tryStatements    intmod.IStatement
	catchClauses     []intmod.ICatchClause
	finallyStatement intmod.IStatement
}

func (s *TryStatement) ResourceList() []intmod.IStatement {
	ret := make([]intmod.IStatement, 0, len(s.resources))
	for _, resource := range s.resources {
		ret = append(ret, resource.(intmod.IStatement))
	}
	return ret
}

func (s *TryStatement) Resources() []intmod.IResource {
	return s.resources
}

func (s *TryStatement) SetResources(resources []intmod.IResource) {
	s.resources = resources
}

func (s *TryStatement) AddResources(resources []intmod.IResource) {
	s.resources = append(s.resources, resources...)
}

func (s *TryStatement) AddResource(resource intmod.IResource) {
	s.resources = append(s.resources, resource)
}

func (s *TryStatement) TryStatements() intmod.IStatement {
	return s.tryStatements
}

func (s *TryStatement) SetTryStatements(tryStatement intmod.IStatement) {
	s.tryStatements = tryStatement
}

func (s *TryStatement) CatchClauseList() []intmod.IStatement {
	ret := make([]intmod.IStatement, 0, len(s.catchClauses))
	for _, resource := range s.catchClauses {
		ret = append(ret, resource.(intmod.IStatement))
	}
	return ret
}

func (s *TryStatement) CatchClauses() []intmod.ICatchClause {
	return s.catchClauses
}

func (s *TryStatement) FinallyStatements() intmod.IStatement {
	return s.finallyStatement
}

func (s *TryStatement) SetFinallyStatements(finallyStatement intmod.IStatement) {
	s.finallyStatement = finallyStatement
}

func (s *TryStatement) IsTryStatement() bool {
	return true
}

func (s *TryStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitTryStatement(s)
}

func NewResource(typ intmod.IObjectType, name string, expression intmod.IExpression) intmod.IResource {
	return &Resource{
		typ:        typ,
		name:       name,
		expression: expression,
	}
}

type Resource struct {
	AbstractStatement

	typ        intmod.IObjectType
	name       string
	expression intmod.IExpression
}

func (r *Resource) Type() intmod.IObjectType {
	return r.typ
}

func (r *Resource) Name() string {
	return r.name
}

func (r *Resource) Expression() intmod.IExpression {
	return r.expression
}

func (r *Resource) SetExpression(expression intmod.IExpression) {
	r.expression = expression
}

func (r *Resource) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitTryStatementResource(r)
}

func NewCatchClause(lineNumber int, typ intmod.IObjectType, name string, statements intmod.IStatement) intmod.ICatchClause {
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
	typ        intmod.IObjectType
	otherType  []intmod.IObjectType
	name       string
	statements intmod.IStatement
}

func (c *CatchClause) LineNumber() int {
	return c.lineNumber
}

func (c *CatchClause) Type() intmod.IObjectType {
	return c.typ
}

func (c *CatchClause) OtherType() []intmod.IObjectType {
	return c.otherType
}

func (c *CatchClause) Name() string {
	return c.name
}

func (c *CatchClause) Statements() intmod.IStatement {
	return c.statements
}

func (c *CatchClause) AddType(typ intmod.IObjectType) {
	c.otherType = append(c.otherType, typ)
}

func (c *CatchClause) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitTryStatementCatchClause(c)
}
