package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model"
)

func NewClassFileCmpExpression(lineNumber int, leftExpress, rightExpression intmod.IExpression) intsrv.IClassFileCmpExpression {
	e := &ClassFileCmpExpression{
		BinaryOperatorExpression: *_type.NewBinaryOperatorExpression(lineNumber,
			_type.PtTypeInt.(intmod.IType), leftExpress, "cmp", rightExpression,
			7).(*_type.BinaryOperatorExpression),
	}
	e.SetValue(e)
	return e
}

type ClassFileCmpExpression struct {
	_type.BinaryOperatorExpression
}
