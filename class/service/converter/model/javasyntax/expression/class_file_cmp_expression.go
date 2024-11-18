package expression

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
)

func NewClassFileCmpExpression(lineNumber int, leftExpress, rightExpression expression.IExpression) *ClassFileCmpExpression {
	return &ClassFileCmpExpression{
		BinaryOperatorExpression: *expression.NewBinaryOperatorExpression(lineNumber, _type.PtTypeInt, leftExpress, "cmp", rightExpression, 7),
	}
}

type ClassFileCmpExpression struct {
	expression.BinaryOperatorExpression
}
