package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"
	"fmt"
)

func NewCastExpression(typ intsyn.IType, expression intsyn.IExpression) intsyn.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithLineNumber(lineNumber int, typ intsyn.IType, expression intsyn.IExpression) intsyn.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithAll(lineNumber int, typ intsyn.IType, expression intsyn.IExpression, explicit bool) intsyn.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         explicit,
	}
}

type CastExpression struct {
	AbstractLineNumberTypeExpression

	expression intsyn.IExpression
	explicit   bool
}

func (e *CastExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *CastExpression) IsExplicit() bool {
	return e.explicit
}

func (e *CastExpression) Priority() int {
	return 3
}

func (e *CastExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *CastExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *CastExpression) IsCastExpression() bool {
	return true
}

func (e *CastExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitCastExpression(e)
}

func (e *CastExpression) String() string {
	return fmt.Sprintf("CastExpression{cast (%s) %s }", e.typ, e.expression)
}
