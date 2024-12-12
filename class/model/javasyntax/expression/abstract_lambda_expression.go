package expression

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
)

func NewAbstractLambdaExpression(typ intmod.IType, statements intmod.IStatement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpression(typ),
		statements:                       statements,
	}
}

func NewAbstractLambdaExpressionWithAll(lineNumber int, typ intmod.IType, statements intmod.IStatement) *AbstractLambdaExpression {
	return &AbstractLambdaExpression{
		AbstractLineNumberTypeExpression: *NewAbstractLineNumberTypeExpressionWithAll(lineNumber, typ),
		statements:                       statements,
	}
}

type AbstractLambdaExpression struct {
	AbstractLineNumberTypeExpression

	statements intmod.IStatement
}

func (e *AbstractLambdaExpression) Priority() int {
	return 17
}

func (e *AbstractLambdaExpression) Statements() intmod.IStatement {
	return e.statements
}
