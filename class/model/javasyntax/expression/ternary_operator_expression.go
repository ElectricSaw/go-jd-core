package expression

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewTernaryOperatorExpression(typ _type.IType, condition Expression, trueExpression Expression, falseExpression Expression) *TernaryOperatorExpression {
	return &TernaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		condition:                        condition,
		trueExpression:                   trueExpression,
		falseExpression:                  falseExpression,
	}
}

func NewTernaryOperatorExpressionWithAll(lineNumber int, typ _type.IType, condition Expression, trueExpression Expression, falseExpression Expression) *TernaryOperatorExpression {
	return &TernaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		condition:                        condition,
		trueExpression:                   trueExpression,
		falseExpression:                  falseExpression,
	}
}

type TernaryOperatorExpression struct {
	AbstractLineNumberTypeExpression

	condition       Expression
	trueExpression  Expression
	falseExpression Expression
}

func (e *TernaryOperatorExpression) Condition() Expression {
	return e.condition
}

func (e *TernaryOperatorExpression) GetTrueCondition() Expression {
	return e.trueExpression
}

func (e *TernaryOperatorExpression) GetFalseCondition() Expression {
	return e.falseExpression
}

func (e *TernaryOperatorExpression) SetCondition(expression Expression) {
	e.condition = expression
}

func (e *TernaryOperatorExpression) SetTrueCondition(expression Expression) {
	e.trueExpression = expression
}

func (e *TernaryOperatorExpression) SetFalseCondition(expression Expression) {
	e.falseExpression = expression
}

func (e *TernaryOperatorExpression) Priority() int {
	return 15
}

func (e *TernaryOperatorExpression) IsTernaryOperatorExpression() bool {
	return true
}

func (e *TernaryOperatorExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitTernaryOperatorExpression(e)
}

func (e *TernaryOperatorExpression) String() string {
	return fmt.Sprintf("TernaryOperatorExpression{%s ? %s : %s}", e.condition, e.trueExpression, e.falseExpression)
}
