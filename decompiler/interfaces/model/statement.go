package model

import "github.com/ElectricSaw/go-jd-core/decompiler/util"

type IAssertStatement interface {
	IStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	Message() IExpression
}

type IBreakStatement interface {
	IStatement

	Text() string
	IsBreakStatement() bool
}

type IByteCodeStatement interface {
	IStatement

	Text() string
}

type ICommentStatement interface {
	IStatement

	Text() string
	IsContinueStatement() bool
}

type IContinueStatement interface {
	IStatement

	Text() string
	IsContinueStatement() bool
}

type IDoWhileStatement interface {
	IStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
}

type IExpressionStatement interface {
	IStatement

	Expression() IExpression
	SetExpression(expression IExpression)
	IsExpressionStatement() bool
	String() string
}

type IForEachStatement interface {
	IStatement

	Type() IType
	Name() string
	Expression() IExpression
	SetExpression(expression IExpression)
	Statement() IStatement
	Statements() IStatement
}

type IForStatement interface {
	IStatement

	Declaration() ILocalVariableDeclaration
	SetDeclaration(declaration ILocalVariableDeclaration)
	Init() IExpression
	SetInit(init IExpression)
	Condition() IExpression
	SetCondition(condition IExpression)
	Update() IExpression
	SetUpdate(update IExpression)
	Statements() IStatement
	String() string
}

type IIfElseStatement interface {
	IIfStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	ElseStatements() IStatement
	IsIfElseStatement() bool
}

type IIfStatement interface {
	IStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	IsIfStatement() bool
}

type ILabelStatement interface {
	IStatement

	Text() string
	Statement() IStatement
	Statements() IStatement
	IsLabelStatement() bool
	String() string
}

type ILambdaExpressionStatement interface {
	IStatement

	Expression() IExpression
	SetExpression(expression IExpression)
	IsLambdaExpressionStatement() bool
	String() string
}

type ILocalVariableDeclarationStatement interface {
	IStatement
	ILocalVariableDeclaration

	IsFinal() bool
	SetFinal(final bool)
	Type() IType
	LocalVariableDeclarators() ILocalVariableDeclarator
	SetLocalVariableDeclarators(declarators ILocalVariableDeclarator)
}

type INoStatement interface {
	IStatement

	String() string
}

type IReturnExpressionStatement interface {
	IStatement

	LineNumber() int
	SetLineNumber(lineNumber int)
	Expression() IExpression
	SetExpression(expression IExpression)
	GenericExpression() IExpression
	IsReturnExpressionStatement() bool
	String() string
}

type IReturnStatement interface {
	IStatement

	IsReturnStatement() bool
	String() string
}

type IStatements interface {
	IStatement
	util.IList[IStatement]

	IsStatements() bool
}

type IStatement interface {
	util.IBase[IStatement]

	AcceptStatement(visitor IStatementVisitor)

	IsBreakStatement() bool
	IsContinueStatement() bool
	IsExpressionStatement() bool
	IsForStatement() bool
	IsIfStatement() bool
	IsIfElseStatement() bool
	IsLabelStatement() bool
	IsLambdaExpressionStatement() bool
	IsLocalVariableDeclarationStatement() bool
	IsMonitorEnterStatement() bool
	IsMonitorExitStatement() bool
	IsReturnStatement() bool
	IsReturnExpressionStatement() bool
	IsStatements() bool
	IsSwitchStatement() bool
	IsSwitchStatementLabelBlock() bool
	IsSwitchStatementMultiLabelsBlock() bool
	IsThrowStatement() bool
	IsTryStatement() bool
	IsWhileStatement() bool

	Condition() IExpression
	Expression() IExpression
	Monitor() IExpression
	ElseStatements() IStatement
	FinallyStatements() IStatement
	Statements() IStatement
	TryStatements() IStatement
	Init() IExpression
	Update() IExpression
	CatchClauses() []ICatchClause
	LineNumber() int
}

type IStatementVisitor interface {
	VisitAssertStatement(statement IAssertStatement)
	VisitBreakStatement(statement IBreakStatement)
	VisitByteCodeStatement(statement IByteCodeStatement)
	VisitCommentStatement(statement ICommentStatement)
	VisitContinueStatement(statement IContinueStatement)
	VisitDoWhileStatement(statement IDoWhileStatement)
	VisitExpressionStatement(statement IExpressionStatement)
	VisitForEachStatement(statement IForEachStatement)
	VisitForStatement(statement IForStatement)
	VisitIfStatement(statement IIfStatement)
	VisitIfElseStatement(statement IIfElseStatement)
	VisitLabelStatement(statement ILabelStatement)
	VisitLambdaExpressionStatement(statement ILambdaExpressionStatement)
	VisitLocalVariableDeclarationStatement(statement ILocalVariableDeclarationStatement)
	VisitNoStatement(statement INoStatement)
	VisitReturnExpressionStatement(statement IReturnExpressionStatement)
	VisitReturnStatement(statement IReturnStatement)
	VisitStatements(statement IStatements)
	VisitSwitchStatement(statement ISwitchStatement)
	VisitSwitchStatementDefaultLabel(statement IDefaultLabel)
	VisitSwitchStatementExpressionLabel(statement IExpressionLabel)
	VisitSwitchStatementLabelBlock(statement ILabelBlock)
	VisitSwitchStatementMultiLabelsBlock(statement IMultiLabelsBlock)
	VisitSynchronizedStatement(statement ISynchronizedStatement)
	VisitThrowStatement(statement IThrowStatement)
	VisitTryStatement(statement ITryStatement)
	VisitTryStatementResource(statement IResource)
	VisitTryStatementCatchClause(statement ICatchClause)
	VisitTypeDeclarationStatement(statement ITypeDeclarationStatement)
	VisitWhileStatement(statement IWhileStatement)
}

type ISwitchStatement interface {
	IStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	List() []IStatement
	Blocks() []IBlock
	IsSwitchStatement() bool
}

type ISynchronizedStatement interface {
	IStatement

	Monitor() IExpression
	SetMonitor(monitor IExpression)
	Statements() IStatement
}

type IThrowStatement interface {
	IStatement

	Expression() IExpression
	SetExpression(expression IExpression)
	IsThrowStatement() bool
	String() string
}

type ITryStatement interface {
	IStatement

	ResourceList() []IStatement
	Resources() []IResource
	SetResources(resources []IResource)
	AddResources(resources []IResource)
	AddResource(resource IResource)
	TryStatements() IStatement
	SetTryStatements(tryStatement IStatement)
	CatchClauseList() []IStatement
	CatchClauses() []ICatchClause
	FinallyStatements() IStatement
	SetFinallyStatements(finallyStatement IStatement)
	IsTryStatement() bool
}

type ITypeDeclarationStatement interface {
	IStatement

	TypeDeclaration() ITypeDeclaration
}

type IWhileStatement interface {
	IStatement

	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	IsWhileStatement() bool
}

type ILabel interface {
	IStatement
}

type IDefaultLabel interface {
	IStatement

	AcceptStatement(visitor IStatementVisitor)
	String() string
}

type IExpressionLabel interface {
	IStatement

	Expression() IExpression
	SetExpression(expression IExpression)
	AcceptStatement(visitor IStatementVisitor)
	String() string
}

type IBlock interface {
	IStatement

	Statements() IStatement
}

type ILabelBlock interface {
	IBlock

	Label() ILabel
	IsSwitchStatementLabelBlock() bool
	AcceptStatement(visitor IStatementVisitor)
	String() string
}

type IMultiLabelsBlock interface {
	IBlock

	Labels() []ILabel
	IsSwitchStatementMultiLabelsBlock() bool
	AcceptStatement(visitor IStatementVisitor)
	String() string
}

type IResource interface {
	IStatement

	Type() IObjectType
	Name() string
	Expression() IExpression
	SetExpression(expression IExpression)
	AcceptStatement(visitor IStatementVisitor)
}

type ICatchClause interface {
	util.IBase[IStatement]

	AcceptStatement(visitor IStatementVisitor)

	IsBreakStatement() bool
	IsContinueStatement() bool
	IsExpressionStatement() bool
	IsForStatement() bool
	IsIfStatement() bool
	IsIfElseStatement() bool
	IsLabelStatement() bool
	IsLambdaExpressionStatement() bool
	IsLocalVariableDeclarationStatement() bool
	IsMonitorEnterStatement() bool
	IsMonitorExitStatement() bool
	IsReturnStatement() bool
	IsReturnExpressionStatement() bool
	IsStatements() bool
	IsSwitchStatement() bool
	IsSwitchStatementLabelBlock() bool
	IsSwitchStatementMultiLabelsBlock() bool
	IsThrowStatement() bool
	IsTryStatement() bool
	IsWhileStatement() bool

	Condition() IExpression
	Expression() IExpression
	Monitor() IExpression
	ElseStatements() IStatement
	FinallyStatements() IStatement
	Statements() IStatement
	TryStatements() IStatement
	Init() IExpression
	Update() IExpression
	CatchClauses() []ICatchClause
	LineNumber() int

	Type() IObjectType
	OtherType() []IObjectType
	Name() string
	AddType(typ IObjectType)
}
