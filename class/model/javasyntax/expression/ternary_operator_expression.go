package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewTernaryOperatorExpression(typ intmod.IType, condition intmod.IExpression,
	trueExpression intmod.IExpression, falseExpression intmod.IExpression) intmod.ITernaryOperatorExpression {
	return NewTernaryOperatorExpressionWithAll(0, typ, condition, trueExpression, falseExpression)
}

func NewTernaryOperatorExpressionWithAll(lineNumber int, typ intmod.IType, condition intmod.IExpression,
	trueExpression intmod.IExpression, falseExpression intmod.IExpression) intmod.ITernaryOperatorExpression {
	e := &TernaryOperatorExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		condition:                        condition,
		trueExpression:                   trueExpression,
		falseExpression:                  falseExpression,
	}
	e.SetValue(e)
	return e
}

type TernaryOperatorExpression struct {
	AbstractLineNumberTypeExpression

	condition       intmod.IExpression
	trueExpression  intmod.IExpression
	falseExpression intmod.IExpression
}

func (e *TernaryOperatorExpression) Condition() intmod.IExpression {
	return e.condition
}

func (e *TernaryOperatorExpression) SetCondition(expression intmod.IExpression) {
	e.condition = expression
}

func (e *TernaryOperatorExpression) TrueExpression() intmod.IExpression {
	return e.trueExpression
}

func (e *TernaryOperatorExpression) SetTrueExpression(expression intmod.IExpression) {
	e.trueExpression = expression
}

func (e *TernaryOperatorExpression) FalseExpression() intmod.IExpression {
	return e.falseExpression
}

func (e *TernaryOperatorExpression) SetFalseExpression(expression intmod.IExpression) {
	e.falseExpression = expression
}

func (e *TernaryOperatorExpression) Priority() int {
	return 15
}

func (e *TernaryOperatorExpression) IsTernaryOperatorExpression() bool {
	return true
}

func (e *TernaryOperatorExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitTernaryOperatorExpression(e)
}

func (e *TernaryOperatorExpression) String() string {
	return fmt.Sprintf("TernaryOperatorExpression{%s ? %s : %s}", e.condition, e.trueExpression, e.falseExpression)
}
