package expression

import (
	_type "bitbucket.org/coontec/javaClass/java/model/javasyntax/type"
	"fmt"
)

func NewStringConstantExpression(str string) *StringConstantExpression {
	return &StringConstantExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		str:                          str,
	}
}

func NewStringConstantExpressionWithAll(lineNumber int, str string) *StringConstantExpression {
	return &StringConstantExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpression(lineNumber),
		str:                          str,
	}
}

type StringConstantExpression struct {
	AbstractLineNumberExpression

	str string
}

func (e *StringConstantExpression) GetStringValue() string {
	return e.str
}

func (e *StringConstantExpression) GetType() _type.IType {
	return _type.OtTypeString
}

func (e *StringConstantExpression) IsStringConstantExpression() bool {
	return true
}

func (e *StringConstantExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitStringConstantExpression(e)
}

func (e *StringConstantExpression) String() string {
	return fmt.Sprintf("StringConstantExpression{\"%s\"}", e.str)
}
