package expression

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewPostOperatorExpression(operator string, expression intmod.IExpression) intmod.IPostOperatorExpression {
	return NewPostOperatorExpressionWithAll(0, operator, expression)
}

func NewPostOperatorExpressionWithAll(lineNumber int, operator string, expression intmod.IExpression) intmod.IPostOperatorExpression {
	e := &PostOperatorExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		operator:                     operator,
		expression:                   expression,
	}
	e.SetValue(e)
	return e
}

type PostOperatorExpression struct {
	AbstractLineNumberExpression

	operator   string
	expression intmod.IExpression
}

func (e *PostOperatorExpression) Operator() string {
	return e.operator
}

func (e *PostOperatorExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *PostOperatorExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *PostOperatorExpression) Type() intmod.IType {
	return e.expression.Type()
}

func (e *PostOperatorExpression) Priority() int {
	return 1
}

func (e *PostOperatorExpression) IsPostOperatorExpression() bool {
	return true
}

func (e *PostOperatorExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitPostOperatorExpression(e)
}

func (e *PostOperatorExpression) String() string {
	return fmt.Sprintf("PostOperatorExpression{%s %s}", e.expression, e.operator)
}
