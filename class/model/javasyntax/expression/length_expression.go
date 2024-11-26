package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
)

func NewLengthExpression(expression intmod.IExpression) intmod.ILengthExpression {
	return NewLengthExpressionWithAll(0, expression)
}

func NewLengthExpressionWithAll(lineNumber int, expression intmod.IExpression) intmod.ILengthExpression {
	e := &LengthExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		expression:                   expression,
	}
	e.SetValue(e)
	return e
}

type LengthExpression struct {
	AbstractLineNumberExpression
	util.DefaultBase[intmod.ILengthExpression]

	expression intmod.IExpression
}

func (e *LengthExpression) Type() intmod.IType {
	return _type.PtTypeInt.(intmod.IType)
}

func (e *LengthExpression) Expression() intmod.IExpression {
	return e.expression
}

func (e *LengthExpression) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *LengthExpression) IsLengthExpression() bool {
	return true
}

func (e *LengthExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitLengthExpression(e)
}

func (e *LengthExpression) String() string {
	return fmt.Sprintf("LengthExpression{%s}", e.expression)
}
