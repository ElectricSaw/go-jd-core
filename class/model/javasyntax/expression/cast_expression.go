package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

func NewCastExpression(typ intmod.IType, expression intmod.IExpression) intmod.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithLineNumber(lineNumber int, typ intmod.IType, expression intmod.IExpression) intmod.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         true,
	}
}

func NewCastExpressionWithAll(lineNumber int, typ intmod.IType, expression intmod.IExpression, explicit bool) intmod.ICastExpression {
	return &CastExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		expression:                       expression,
		explicit:                         explicit,
	}
}

type CastExpression struct {
	AbstractLineNumberTypeExpression

	expression intmod.IExpression
	explicit   bool
}

func (e *CastExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *CastExpression) IsExplicit() bool {
	return e.explicit
}

func (e *CastExpression) Priority() int {
	return 3
}

func (e *CastExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *CastExpression) SetExplicit(explicit bool) {
	e.explicit = explicit
}

func (e *CastExpression) IsCastExpression() bool {
	return true
}

func (e *CastExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitCastExpression(e)
}

func (e *CastExpression) String() string {
	return fmt.Sprintf("CastExpression{cast (%s) %s }", e.typ, e.expression)
}
