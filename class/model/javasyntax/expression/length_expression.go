package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewLengthExpression(expression intsyn.IExpression) intsyn.ILengthExpression {
	return &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		expression:                   expression,
	}
}

func NewLengthExpressionWithAll(lineNumber int, expression intsyn.IExpression) intsyn.ILengthExpression {
	return &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		expression:                   expression,
	}
}

type LengthExpression struct {
	AbstractLineNumberExpression

	expression intsyn.IExpression
}

func (e *LengthExpression) Type() intsyn.IType {
	return _type.PtTypeInt.(intsyn.IType)
}

func (e *LengthExpression) Expression() intsyn.IExpression {
	return e.expression
}

func (e *LengthExpression) SetExpression(expression intsyn.IExpression) {
	e.expression = expression
}

func (e *LengthExpression) IsLengthExpression() bool {
	return true
}

func (e *LengthExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitLengthExpression(e)
}

func (e *LengthExpression) String() string {
	return fmt.Sprintf("LengthExpression{%s}", e.expression)
}
