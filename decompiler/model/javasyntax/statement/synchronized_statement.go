package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewSynchronizedStatement(monitor model.IExpression, statements IStatement) SynchronizedStatement {
	return &SynchronizedStatement{
		Monitor:    monitor,
		Statements: statements,
	}
}

type SynchronizedStatement struct {
	util.DefaultBase[IStatement]
	AbstractStatement

	Monitor    model.IExpression
	Statements IStatement
}

func (s *SynchronizedStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSynchronizedStatement(s)
}

func (s *SynchronizedStatement) IsBreakStatement() bool                    { return false }
func (s *SynchronizedStatement) IsContinueStatement() bool                 { return false }
func (s *SynchronizedStatement) IsExpressionStatement() bool               { return false }
func (s *SynchronizedStatement) IsForStatement() bool                      { return false }
func (s *SynchronizedStatement) IsIfStatement() bool                       { return false }
func (s *SynchronizedStatement) IsIfElseStatement() bool                   { return false }
func (s *SynchronizedStatement) IsLabelStatement() bool                    { return false }
func (s *SynchronizedStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *SynchronizedStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *SynchronizedStatement) IsMonitorEnterStatement() bool             { return false }
func (s *SynchronizedStatement) IsMonitorExitStatement() bool              { return false }
func (s *SynchronizedStatement) IsReturnStatement() bool                   { return false }
func (s *SynchronizedStatement) IsReturnExpressionStatement() bool         { return false }
func (s *SynchronizedStatement) IsStatements() bool                        { return false }
func (s *SynchronizedStatement) IsSwitchStatement() bool                   { return false }
func (s *SynchronizedStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *SynchronizedStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *SynchronizedStatement) IsThrowStatement() bool                    { return false }
func (s *SynchronizedStatement) IsTryStatement() bool                      { return false }
func (s *SynchronizedStatement) IsWhileStatement() bool                    { return false }

func (s *SynchronizedStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *SynchronizedStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *SynchronizedStatement) GetMonitor() model.IExpression    { return s.Monitor }

func (s *SynchronizedStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *SynchronizedStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *SynchronizedStatement) GetStatements() IStatement        { return s.Statements }
func (s *SynchronizedStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *SynchronizedStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *SynchronizedStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *SynchronizedStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *SynchronizedStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *SynchronizedStatement) String() string {
	return "SynchronizedStatement{}"
}
