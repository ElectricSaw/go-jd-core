package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
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

func (s *ForEachStatement) GetType() _type.IType {
	return s.typ
}

func (s *ForEachStatement) GetName() string {
	return s.name
}

func (s *ForEachStatement) GetExpression() expression.Expression {
	return s.expression
}

func (s *ForEachStatement) SetExpression(expression expression.Expression) {
	s.expression = expression
}

func (s *ForEachStatement) GetStatement() Statement {
	return s.statement
}

func (s *ForEachStatement) Accept(visitor StatementVisitor) {
	visitor.VisitForEachStatement(s)
}
