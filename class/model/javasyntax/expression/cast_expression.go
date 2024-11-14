package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewCastExpression(typ _type.IType, expression Expression) *CastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithLineNumber(lineNumber int, typ _type.IType, expression Expression) *CastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithAll(lineNumber int, typ _type.IType, expression Expression, explicit bool) *CastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         explicit,
	}
}

type CastExpression struct {
	AbstractLineNumberTypeExpression

	expression Expression
	explicit   bool
}

func (e *CastExpression) Expression() Expression {
	return e.expression
}

func (e *CastExpression) IsExplicit() bool {
	return e.explicit
}

func (e *CastExpression) Priority() int {
	return 3
}

func (e *CastExpression) SetExpression(expression Expression) {
	e.expression = expression
}

func (e *CastExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *CastExpression) IsCastExpression() bool {
	return true
}

func (e *CastExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitCastExpression(e)
}

func (e *CastExpression) String() string {
	return fmt.Sprintf("CastExpression{cast (%s) %s }", e.typ, e.expression)
}
