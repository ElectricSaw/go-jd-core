package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

func NewTernaryOperatorExpression(typ intsyn.IType, condition intsyn.IExpression,
	trueExpression intsyn.IExpression, falseExpression intsyn.IExpression) intsyn.ITernaryOperatorExpression {
	return &TernaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		condition:                        condition,
		trueExpression:                   trueExpression,
		falseExpression:                  falseExpression,
	}
}

func NewTernaryOperatorExpressionWithAll(lineNumber int, typ intsyn.IType, condition intsyn.IExpression,
	trueExpression intsyn.IExpression, falseExpression intsyn.IExpression) intsyn.ITernaryOperatorExpression {
	return &TernaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		condition:                        condition,
		trueExpression:                   trueExpression,
		falseExpression:                  falseExpression,
	}
}

type TernaryOperatorExpression struct {
	AbstractLineNumberTypeExpression

	condition       intsyn.IExpression
	trueExpression  intsyn.IExpression
	falseExpression intsyn.IExpression
}

func (e *TernaryOperatorExpression) Condition() intsyn.IExpression {
	return e.condition
}

func (e *TernaryOperatorExpression) SetCondition(expression intsyn.IExpression) {
	e.condition = expression
}

func (e *TernaryOperatorExpression) TrueExpression() intsyn.IExpression {
	return e.trueExpression
}

func (e *TernaryOperatorExpression) SetTrueExpression(expression intsyn.IExpression) {
	e.trueExpression = expression
}

func (e *TernaryOperatorExpression) FalseExpression() intsyn.IExpression {
	return e.falseExpression
}

func (e *TernaryOperatorExpression) SetFalseExpression(expression intsyn.IExpression) {
	e.falseExpression = expression
}

func (e *TernaryOperatorExpression) Priority() int {
	return 15
}

func (e *TernaryOperatorExpression) IsTernaryOperatorExpression() bool {
	return true
}

func (e *TernaryOperatorExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitTernaryOperatorExpression(e)
}

func (e *TernaryOperatorExpression) String() string {
	return fmt.Sprintf("TernaryOperatorExpression{%s ? %s : %s}", e.condition, e.trueExpression, e.falseExpression)
}
