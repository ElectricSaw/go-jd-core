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

func (s *AbstractStatement) Condition() expression.Expression  { return expression.NeNoExpression }
func (s *AbstractStatement) Expression() expression.Expression { return expression.NeNoExpression }
func (s *AbstractStatement) Monitor() expression.Expression    { return expression.NeNoExpression }
func (s *AbstractStatement) ElseStatements() Statement         { return NoStmt }
func (s *AbstractStatement) FinallyStatements() Statement      { return NoStmt }
func (s *AbstractStatement) Statements() Statement             { return NoStmt }
func (s *AbstractStatement) TryStatements() Statement          { return NoStmt }

func (s *AbstractStatement) Init() expression.Expression   { return expression.NeNoExpression }
func (s *AbstractStatement) Update() expression.Expression { return expression.NeNoExpression }

func (s *AbstractStatement) CatchClauses() []CatchClause { return nil }

func (s *AbstractStatement) LineNumber() int { return expression.UnknownLineNumber }

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

	Condition() expression.Expression
	Expression() expression.Expression
	Monitor() expression.Expression
	ElseStatements() Statement
	FinallyStatements() Statement
	Statements() Statement
	TryStatements() Statement
	Init() expression.Expression
	Update() expression.Expression
	CatchClauses() []CatchClause
	LineNumber() int
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
