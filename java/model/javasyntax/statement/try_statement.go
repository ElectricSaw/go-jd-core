package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
)

func NewTryStatement(tryStatements Statement, catchClauses []CatchClause, finallyStatement Statement) *TryStatement {
	return &TryStatement{
		resources:        nil,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

func NewTryStatementWithAll(resource []Resource, tryStatements Statement, catchClauses []CatchClause, finallyStatement Statement) *TryStatement {
	return &TryStatement{
		resources:        resource,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

type TryStatement struct {
	AbstractStatement

	resources        []Resource
	tryStatements    Statement
	catchClauses     []CatchClause
	finallyStatement Statement
}

func (s *TryStatement) GetResources() []Resource {
	return s.resources
}

func (s *TryStatement) GetTryStatement() Statement {
	return s.tryStatements
}

func (s *TryStatement) SetTryStatement(tryStatement Statement) {
	s.tryStatements = tryStatement
}

func (s *TryStatement) GetCatchClauses() []CatchClause {
	return s.catchClauses
}

func (s *TryStatement) GetFinallyStatements() Statement {
	return s.finallyStatement
}

func (s *TryStatement) SetFinallyStatement(finallyStatement Statement) {
	s.finallyStatement = finallyStatement
}

func (s *TryStatement) IsTryStatement() bool {
	return true
}

func (s *TryStatement) Accept(visitor StatementVisitor) {
	visitor.VisitTryStatement(s)
}

func NewResource(typ _type.ObjectType, name string, expression expression.Expression) *Resource {
	return &Resource{
		typ:        typ,
		name:       name,
		expression: expression,
	}
}

type Resource struct {
	AbstractStatement

	typ        _type.ObjectType
	name       string
	expression expression.Expression
}

func (r *Resource) GetType() *_type.ObjectType {
	return &r.typ
}

func (r *Resource) GetName() string {
	return r.name
}

func (r *Resource) GetExpression() expression.Expression {
	return r.expression
}

func (r *Resource) SetExpression(expression expression.Expression) {
	r.expression = expression
}

func (r *Resource) Accept(visitor StatementVisitor) {
	visitor.VisitTryStatementResource(r)
}

func NewCatchClause(lineNumber int, typ _type.ObjectType, name string, statements Statement) *CatchClause {
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
	typ        _type.ObjectType
	otherType  []_type.ObjectType
	name       string
	statements Statement
}

func (c *CatchClause) GetLineNumber() int {
	return c.lineNumber
}

func (c *CatchClause) GetType() *_type.ObjectType {
	return &c.typ
}

func (c *CatchClause) GetOtherType() []_type.ObjectType {
	return c.otherType
}

func (c *CatchClause) GetName() string {
	return c.name
}

func (c *CatchClause) GetStatements() Statement {
	return c.statements
}

func (c *CatchClause) AddType(typ _type.ObjectType) {
	c.otherType = append(c.otherType, typ)
}

func (c *CatchClause) Accept(visitor StatementVisitor) {
	visitor.VisitTryStatementCatchClause(c)
}
