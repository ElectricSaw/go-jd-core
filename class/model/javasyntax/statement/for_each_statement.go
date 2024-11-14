package statement

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewForEachStatement(typ _type.IType, name string, expression expression.Expression, statement Statement) *ForEachStatement {
	return &ForEachStatement{
		typ:        typ,
		name:       name,
		expression: expression,
		statement:  statement,
	}
}

type ForEachStatement struct {
	AbstractStatement

	typ        _type.IType
	name       string
	expression expression.Expression
	statement  Statement
}

func (s *ForEachStatement) Type() _type.IType {
	return s.typ
}

func (s *ForEachStatement) Name() string {
	return s.name
}

func (s *ForEachStatement) Expression() expression.Expression {
	return s.expression
}

func (s *ForEachStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *ForEachStatement) Statement() Statement {
	return s.statement
}

func (s *ForEachStatement) Accept(visitor StatementVisitor) {
	visitor.VisitForEachStatement(s)
}
