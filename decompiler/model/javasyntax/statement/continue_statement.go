package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var Continue = NewContinueStatement("")

func NewContinueStatement(label string) ContinueStatement {
	return &ContinueStatement{
		Label: label,
	}
}

type ContinueStatement struct {
	util.DefaultBase[IStatement]

	Label string
}

func (s *ContinueStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitContinueStatement(s)
}

func (s *ContinueStatement) IsBreakStatement() bool                    { return false }
func (s *ContinueStatement) IsContinueStatement() bool                 { return true }
func (s *ContinueStatement) IsExpressionStatement() bool               { return false }
func (s *ContinueStatement) IsForStatement() bool                      { return false }
func (s *ContinueStatement) IsIfStatement() bool                       { return false }
func (s *ContinueStatement) IsIfElseStatement() bool                   { return false }
func (s *ContinueStatement) IsLabelStatement() bool                    { return false }
func (s *ContinueStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ContinueStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ContinueStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ContinueStatement) IsMonitorExitStatement() bool              { return false }
func (s *ContinueStatement) IsReturnStatement() bool                   { return false }
func (s *ContinueStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ContinueStatement) IsStatements() bool                        { return false }
func (s *ContinueStatement) IsSwitchStatement() bool                   { return false }
func (s *ContinueStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ContinueStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ContinueStatement) IsThrowStatement() bool                    { return false }
func (s *ContinueStatement) IsTryStatement() bool                      { return false }
func (s *ContinueStatement) IsWhileStatement() bool                    { return false }

func (s *ContinueStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ContinueStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *ContinueStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ContinueStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ContinueStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ContinueStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ContinueStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ContinueStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ContinueStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ContinueStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ContinueStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ContinueStatement) String() string {
	return "ContinueStatement{}"
}
