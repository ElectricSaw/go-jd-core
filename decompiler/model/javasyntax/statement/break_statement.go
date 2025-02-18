package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var Break = NewBreakStatement("")

func NewBreakStatement(label string) BreakStatement {
	return BreakStatement{
		Label: label,
	}
}

type BreakStatement struct {
	util.DefaultBase[IStatement]

	Label string
}

func (s *BreakStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitBreakStatement(s)
}

func (s *BreakStatement) IsBreakStatement() bool                    { return true }
func (s *BreakStatement) IsContinueStatement() bool                 { return false }
func (s *BreakStatement) IsExpressionStatement() bool               { return false }
func (s *BreakStatement) IsForStatement() bool                      { return false }
func (s *BreakStatement) IsIfStatement() bool                       { return false }
func (s *BreakStatement) IsIfElseStatement() bool                   { return false }
func (s *BreakStatement) IsLabelStatement() bool                    { return false }
func (s *BreakStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *BreakStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *BreakStatement) IsMonitorEnterStatement() bool             { return false }
func (s *BreakStatement) IsMonitorExitStatement() bool              { return false }
func (s *BreakStatement) IsReturnStatement() bool                   { return false }
func (s *BreakStatement) IsReturnExpressionStatement() bool         { return false }
func (s *BreakStatement) IsStatements() bool                        { return false }
func (s *BreakStatement) IsSwitchStatement() bool                   { return false }
func (s *BreakStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *BreakStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *BreakStatement) IsThrowStatement() bool                    { return false }
func (s *BreakStatement) IsTryStatement() bool                      { return false }
func (s *BreakStatement) IsWhileStatement() bool                    { return false }

func (s *BreakStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *BreakStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *BreakStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *BreakStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *BreakStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *BreakStatement) GetStatements() IStatement        { return &NoStmt }
func (s *BreakStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *BreakStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *BreakStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *BreakStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *BreakStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *BreakStatement) String() string {
	return "BreakStatement{}"
}
