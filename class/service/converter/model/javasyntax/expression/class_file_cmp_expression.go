package expression

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewClassFileCmpExpression(lineNumber int, leftExpress, rightExpression intmod.IExpression) intsrv.IClassFileCmpExpression {
	return &ClassFileCmpExpression{
		BinaryOperatorExpression: *expression.NewBinaryOperatorExpression(lineNumber,
			_type.PtTypeInt.(intmod.IType), leftExpress, "cmp", rightExpression,
			7).(*expression.BinaryOperatorExpression),
	}
}

type ClassFileCmpExpression struct {
	expression.BinaryOperatorExpression
}
