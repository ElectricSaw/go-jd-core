package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
)

func NewForEachStatement(typ intsyn.IType, name string, expression intsyn.IExpression, statement intsyn.IStatement) intsyn.IForEachStatement {
	return &ForEachStatement{
		typ:        typ,
		name:       name,
		expression: expression,
		statement:  statement,
	}
}

type ForEachStatement struct {
	AbstractStatement

	typ        intsyn.IType
	name       string
	expression intsyn.IExpression
	statement  intsyn.IStatement
}

func (s *ForEachStatement) Type() intsyn.IType {
	return s.typ
}

func (s *ForEachStatement) Name() string {
	return s.name
}

func (s *ForEachStatement) Expression() intsyn.IExpression {
	return s.expression
}

func (s *ForEachStatement) SetExpression(expression intsyn.IExpression) {
	s.expression = expression
}

func (s *ForEachStatement) Statement() intsyn.IStatement {
	return s.statement
}

func (s *ForEachStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitForEachStatement(s)
}
