package service

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
)

type IClassFileBreakContinueStatement interface {
	intmod.IStatement

	Offset() int
	TargetOffset() int
	Statement() intmod.IStatement
	SetStatement(statement intmod.IStatement)
	IsContinueLabel() bool
	SetContinueLabel(continueLabel bool)
	Accept(visitor intmod.IStatementVisitor)
	String() string
}

type IClassFileForEachStatement interface {
	intmod.IForEachStatement

	Name() string
	String() string
}

type IClassFileForStatement interface {
	intmod.IForStatement

	FromOffset() int
	ToOffset() int
	IsForStatement() bool
	String() string
}

type IClassFileMonitorEnterStatement interface {
	intmod.ICommentStatement

	Monitor() intmod.IExpression
	IsMonitorEnterStatement() bool
	String() string
}

type IClassFileMonitorExitStatement interface {
	intmod.ICommentStatement

	Monitor() intmod.IExpression
	IsMonitorExitStatement() bool
	String() string
}

type IClassFileTryStatement interface {
	intmod.ITryStatement

	isJsr() bool
	isEclipse() bool
}

type ICatchClause interface {
	intmod.ICatchClause

	Name() string
	LocalVariable() ILocalVariable
}