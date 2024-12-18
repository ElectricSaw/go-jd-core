package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
)

func NewClassFileCmpExpression(lineNumber int, leftExpress, rightExpression intmod.IExpression) intsrv.IClassFileCmpExpression {
	e := &ClassFileCmpExpression{
		BinaryOperatorExpression: *expression.NewBinaryOperatorExpression(lineNumber,
			_type.PtTypeInt.(intmod.IType), leftExpress, "cmp", rightExpression,
			7).(*expression.BinaryOperatorExpression),
	}
	e.SetValue(e)
	return e
}

type ClassFileCmpExpression struct {
	expression.BinaryOperatorExpression
}
