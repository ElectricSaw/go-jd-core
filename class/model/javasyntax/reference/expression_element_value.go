package reference

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
	"fmt"
)

func NewExpressionElementValue(expression intmod.IExpression) intmod.IExpressionElementValue {
	v := &ExpressionElementValue{
		expression: expression,
	}
	v.SetValue(v)
	return v
}

type ExpressionElementValue struct {
	util.DefaultBase[intmod.IExpressionElementValue]
	
	expression intmod.IExpression
}

func (e *ExpressionElementValue) Expression() intmod.IExpression {
	return e.expression
}

func (e *ExpressionElementValue) SetExpression(expression intmod.IExpression) {
	e.expression = expression
}

func (e *ExpressionElementValue) Accept(visitor intmod.IReferenceVisitor) {
	visitor.VisitExpressionElementValue(e)
}

func (e *ExpressionElementValue) String() string {
	return fmt.Sprintf("ExpressionElementValue{%s}", *e)
}
