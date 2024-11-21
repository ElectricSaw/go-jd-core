package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"fmt"
)

func NewStringConstantExpression(str string) intmod.IStringConstantExpression {
	return &StringConstantExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		str:                          str,
	}
}

func NewStringConstantExpressionWithAll(lineNumber int, str string) intmod.IStringConstantExpression {
	return &StringConstantExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		str:                          str,
	}
}

type StringConstantExpression struct {
	AbstractLineNumberExpression

	str string
}

func (e *StringConstantExpression) StringValue() string {
	return e.str
}

func (e *StringConstantExpression) Type() intmod.IType {
	return _type.OtTypeString.(intmod.IType)
}

func (e *StringConstantExpression) IsStringConstantExpression() bool {
	return true
}

func (e *StringConstantExpression) Accept(visitor intmod.IExpressionVisitor) {
	visitor.VisitStringConstantExpression(e)
}

func (e *StringConstantExpression) String() string {
	return fmt.Sprintf("StringConstantExpression{\"%s\"}", e.str)
}
