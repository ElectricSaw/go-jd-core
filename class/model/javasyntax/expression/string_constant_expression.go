package expression

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"fmt"
)

func NewStringConstantExpression(str string) intsyn.IStringConstantExpression {
	return &StringConstantExpression{
		AbstractLineNumberExpression: *NewAbstractLineNumberExpressionEmpty(),
		str:                          str,
	}
}

func NewStringConstantExpressionWithAll(lineNumber int, str string) intsyn.IStringConstantExpression {
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

func (e *StringConstantExpression) Type() intsyn.IType {
	return _type.OtTypeString.(intsyn.IType)
}

func (e *StringConstantExpression) IsStringConstantExpression() bool {
	return true
}

func (e *StringConstantExpression) Accept(visitor intsyn.IExpressionVisitor) {
	visitor.VisitStringConstantExpression(e)
}

func (e *StringConstantExpression) String() string {
	return fmt.Sprintf("StringConstantExpression{\"%s\"}", e.str)
}
