package model

import "bitbucket.org/coontec/go-jd-core/class/util"

type IAssertStatement interface {
	Condition() IExpression
	SetCondition(condition IExpression)
	Message() IExpression
	Accept(visitor IStatementVisitor)
}

type IBreakStatement interface {
	Label() string
	IsBreakStatement() bool
	Accept(visitor IStatementVisitor)
}

type IByteCodeStatement interface {
	Text() string
	Accept(visitor IStatementVisitor)
}

type ICommentStatement interface {
	Label() string
	IsContinueStatement() bool
	Accept(visitor IStatementVisitor)
}

type IContinueStatement interface {
	Label() string
	IsContinueStatement() bool
	Accept(visitor IStatementVisitor)
}

type IDoWhileStatement interface {
	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	Accept(visitor IStatementVisitor)
}

type IExpressionStatement interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	IsExpressionStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type IForEachStatement interface {
	Type() IType
	Name() string
	Expression() IExpression
	SetExpression(expression IExpression)
	Statement() IStatement
	Statements() IStatement
	Accept(visitor IStatementVisitor)
}

type IForStatement interface {
	Declaration() ILocalVariableDeclaration
	SetDeclaration(declaration ILocalVariableDeclaration)
	Init() IExpression
	SetInit(init IExpression)
	Condition() IExpression
	SetCondition(condition IExpression)
	Update() IExpression
	SetUpdate(update IExpression)
	Statements() IStatement
	Accept(visitor IStatementVisitor)
	String() string
}

type IIfElseStatement interface {
	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	ElseStatements() IStatement
	IsIfElseStatement() bool
	Accept(visitor IStatementVisitor)
}

type IIfStatement interface {
	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	IsIfStatement() bool
	Accept(visitor IStatementVisitor)
}

type ILabelStatement interface {
	Label() string
	Statement() IStatement
	Statements() IStatement
	IsLabelStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type ILambdaExpressionStatement interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	IsLambdaExpressionStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type ILocalVariableDeclarationStatement interface {
	IsFinal() bool
	SetFinal(final bool)
	Type() IType
	LocalVariableDeclarators() ILocalVariableDeclarator
	SetLocalVariableDeclarators(declarators ILocalVariableDeclarator)
	Accept(visitor IStatementVisitor)
}

type INoStatement interface {
	Accept(visitor IStatementVisitor)
	String() string
}

type IReturnExpressionStatement interface {
	LineNumber() int
	SetLineNumber(lineNumber int)
	Expression() IExpression
	SetExpression(expression IExpression)
	GenericExpression() IExpression
	IsReturnExpressionStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type IReturnStatement interface {
	IsReturnStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type IStatements interface {
	util.IList[IStatement]
	IsStatements() bool
	Accept(visitor IStatementVisitor)
}

type IStatement interface {
	util.Base[IStatement]

	Accept(visitor IStatementVisitor)

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
	Condition() IExpression
	SetCondition(condition IExpression)
	List() []IStatement
	Blocks() []IBlock
	IsSwitchStatement() bool
	Accept(visitor IStatementVisitor)
}

type ISynchronizedStatement interface {
	Monitor() IExpression
	SetMonitor(monitor IExpression)
	Statements() IStatement
	Accept(visitor IStatementVisitor)
}

type IThrowStatement interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	IsThrowStatement() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type ITryStatement interface {
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
	SetFinallyStatement(finallyStatement IStatement)
	IsTryStatement() bool
	Accept(visitor IStatementVisitor)
}

type ITypeDeclarationStatement interface {
	TypeDeclaration() ITypeDeclaration
	Accept(visitor IStatementVisitor)
}

type IWhileStatement interface {
	Condition() IExpression
	SetCondition(condition IExpression)
	Statements() IStatement
	IsWhileStatement() bool
	Accept(visitor IStatementVisitor)
}

type ILabel interface {
	IStatement
}

type IDefaultLabel interface {
	Accept(visitor IStatementVisitor)
	String() string
}

type IExpressionLabel interface {
	Expression() IExpression
	SetExpression(expression IExpression)
	Accept(visitor IStatementVisitor)
	String() string
}

type IBlock interface {
	Statements() IStatement
}

type ILabelBlock interface {
	Statements() IStatement
	Label() ILabel
	IsSwitchStatementLabelBlock() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type IMultiLabelsBlock interface {
	Statements() IStatement
	List() []IStatement
	Labels() []ILabel
	IsSwitchStatementMultiLabelsBlock() bool
	Accept(visitor IStatementVisitor)
	String() string
}

type IResource interface {
	Type() IObjectType
	Name() string
	Expression() IExpression
	SetExpression(expression IExpression)
	Accept(visitor IStatementVisitor)
}

type ICatchClause interface {
	LineNumber() int
	Type() IObjectType
	OtherType() []IObjectType
	Name() string
	Statements() IStatement
	AddType(typ IObjectType)
	Accept(visitor IStatementVisitor)
}
