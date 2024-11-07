package statement

import "bitbucket.org/coontec/javaClass/class/model/javasyntax/expression"

type AbstractStatement struct {
}

func (s *AbstractStatement) Accept(visitor StatementVisitor) {}

func (s *AbstractStatement) IsBreakStatement() bool                    { return false }
func (s *AbstractStatement) IsContinueStatement() bool                 { return false }
func (s *AbstractStatement) IsExpressionStatement() bool               { return false }
func (s *AbstractStatement) IsForStatement() bool                      { return false }
func (s *AbstractStatement) IsIfStatement() bool                       { return false }
func (s *AbstractStatement) IsIfElseStatement() bool                   { return false }
func (s *AbstractStatement) IsLabelStatement() bool                    { return false }
func (s *AbstractStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *AbstractStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *AbstractStatement) IsMonitorEnterStatement() bool             { return false }
func (s *AbstractStatement) IsMonitorExitStatement() bool              { return false }
func (s *AbstractStatement) IsReturnStatement() bool                   { return false }
func (s *AbstractStatement) IsReturnExpressionStatement() bool         { return false }
func (s *AbstractStatement) IsStatements() bool                        { return false }
func (s *AbstractStatement) IsSwitchStatement() bool                   { return false }
func (s *AbstractStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *AbstractStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *AbstractStatement) IsThrowStatement() bool                    { return false }
func (s *AbstractStatement) IsTryStatement() bool                      { return false }
func (s *AbstractStatement) IsWhileStatement() bool                    { return false }

func (s *AbstractStatement) GetCondition() expression.Expression  { return expression.NeNoExpression }
func (s *AbstractStatement) GetExpression() expression.Expression { return expression.NeNoExpression }
func (s *AbstractStatement) GetMonitor() expression.Expression    { return expression.NeNoExpression }
func (s *AbstractStatement) GetElseStatements() Statement         { return NoStmt }
func (s *AbstractStatement) GetFinallyStatements() Statement      { return NoStmt }
func (s *AbstractStatement) GetStatements() Statement             { return NoStmt }
func (s *AbstractStatement) GetTryStatements() Statement          { return NoStmt }

func (s *AbstractStatement) GetInit() expression.Expression   { return expression.NeNoExpression }
func (s *AbstractStatement) GetUpdate() expression.Expression { return expression.NeNoExpression }

func (s *AbstractStatement) GetCatchClauses() []CatchClause { return nil }

func (s *AbstractStatement) GetLineNumber() int { return expression.UnknownLineNumber }

type Statement interface {
	Accept(visitor StatementVisitor)

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

	GetCondition() expression.Expression
	GetExpression() expression.Expression
	GetMonitor() expression.Expression
	GetElseStatements() Statement
	GetFinallyStatements() Statement
	GetStatements() Statement
	GetTryStatements() Statement
	GetInit() expression.Expression
	GetUpdate() expression.Expression
	GetCatchClauses() []CatchClause
	GetLineNumber() int
}

type StatementVisitor interface {
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
	VisitSwitchStatementDefaultLabel(statement *DefaultLabe1)
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
