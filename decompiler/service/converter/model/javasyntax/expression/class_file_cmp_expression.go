package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model"
	expression2 "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
)

func NewClassFileCmpExpression(lineNumber int, leftExpress, rightExpression intmod.IExpression) intsrv.IClassFileCmpExpression {
	e := &ClassFileCmpExpression{
		BinaryOperatorExpression: *expression2.NewBinaryOperatorExpression(lineNumber,
			_type.PtTypeInt.(intmod.IType), leftExpress, "cmp", rightExpression,
			7).(*expression2.BinaryOperatorExpression),
	}
	e.SetValue(e)
	return e
}

type ClassFileCmpExpression struct {
	expression2.BinaryOperatorExpression
}
