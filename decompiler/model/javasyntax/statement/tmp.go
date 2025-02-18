package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IStatement interface {
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

	GetCondition() model.IExpression
	GetExpression() model.IExpression
	GetMonitor() model.IExpression

	GetElseStatements() IStatement
	GetFinallyStatements() IStatement
	GetStatements() IStatement
	GetTryStatements() IStatement

	GetInit() model.IExpression
	GetUpdate() model.IExpression

	GetCatchClauses() util.IList[*CatchClause]
	GetLineNumber() int

	String() string
}

type IStatementVisitor interface {
	VisitAssertStatement(statement *AssertStatement)
	VisitBreakStatement(statement *BreakStatement)
	VisitByteCodeStatement(statement *ByteCodeStatement)
	VisitCommentStatement(statement *CommentStatement)
	VisitContinueStatement(statement *ContinueStatement)
	VisitDoWhileStatement(statement *DoWhileStatement)
	VisitExpressionStatement(statement *ExpressionStatement)
	VisitForEachStatement(statement *ForEachStatement)
	VisitForStatement(statement *ForStatement)
	VisitIfStatement(statement *IfStatement)
	VisitIfElseStatement(statement *IfElseStatement)
	VisitLabelStatement(statement *LabelStatement)
	VisitLambdaExpressionStatement(statement *LambdaExpressionStatement)
	VisitLocalVariableDeclarationStatement(statement *LocalVariableDeclarationStatement)
	VisitNoStatement(statement *NoStatement)
	VisitReturnExpressionStatement(statement *ReturnExpressionStatement)
	VisitReturnStatement(statement *ReturnStatement)
	VisitStatements(statement *Statements)
	VisitSwitchStatement(statement *SwitchStatement)
	VisitSwitchStatementDefaultLabel(statement *DefaultLabel)
	VisitSwitchStatementExpressionLabel(statement *ExpressionLabel)
	VisitSwitchStatementLabelBlock(statement *LabelBlock)
	VisitSwitchStatementMultiLabelsBlock(statement *MultiLabelsBlock)
	VisitSynchronizedStatement(statement *SynchronizedStatement)
	VisitThrowStatement(statement *ThrowStatement)
	VisitTryStatement(statement *TryStatement)
	VisitTryStatementResource(statement *Resource)
	VisitTryStatementCatchClause(statement *CatchClause)
	VisitTypeDeclarationStatement(statement *TypeDeclarationStatement)
	VisitWhileStatement(statement *WhileStatement)
}

type ILabel interface {
	IStatement
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////
