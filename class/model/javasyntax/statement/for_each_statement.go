package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

func NewForEachStatement(typ intmod.IType, name string, expression intmod.IExpression, statement intmod.IStatement) intmod.IForEachStatement {
	return &ForEachStatement{
		typ:        typ,
		name:       name,
		expression: expression,
		statement:  statement,
	}
}

type ForEachStatement struct {
	AbstractStatement

	typ        intmod.IType
	name       string
	expression intmod.IExpression
	statement  intmod.IStatement
}

func (s *ForEachStatement) Type() intmod.IType {
	return s.typ
}

func (s *ForEachStatement) Name() string {
	return s.name
}

func (s *ForEachStatement) Expression() intmod.IExpression {
	return s.expression
}

func (s *ForEachStatement) SetExpression(expression intmod.IExpression) {
	s.expression = expression
}

func (s *ForEachStatement) Statement() intmod.IStatement {
	return s.statement
}

func (s *ForEachStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitForEachStatement(s)
}
