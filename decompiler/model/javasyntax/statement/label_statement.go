package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLabelStatement(label string, statement IStatement) LabelStatement {
	return &LabelStatement{
		Label:     label,
		Statement: statement,
	}
}

type LabelStatement struct {
	util.DefaultBase[IStatement]

	Label     string
	Statement IStatement
}

func (s *LabelStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) IsBreakStatement() bool                    { return false }
func (s *LabelStatement) IsContinueStatement() bool                 { return false }
func (s *LabelStatement) IsExpressionStatement() bool               { return false }
func (s *LabelStatement) IsForStatement() bool                      { return false }
func (s *LabelStatement) IsIfStatement() bool                       { return false }
func (s *LabelStatement) IsIfElseStatement() bool                   { return false }
func (s *LabelStatement) IsLabelStatement() bool                    { return true }
func (s *LabelStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *LabelStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *LabelStatement) IsMonitorEnterStatement() bool             { return false }
func (s *LabelStatement) IsMonitorExitStatement() bool              { return false }
func (s *LabelStatement) IsReturnStatement() bool                   { return false }
func (s *LabelStatement) IsReturnExpressionStatement() bool         { return false }
func (s *LabelStatement) IsStatements() bool                        { return false }
func (s *LabelStatement) IsSwitchStatement() bool                   { return false }
func (s *LabelStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *LabelStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *LabelStatement) IsThrowStatement() bool                    { return false }
func (s *LabelStatement) IsTryStatement() bool                      { return false }
func (s *LabelStatement) IsWhileStatement() bool                    { return false }

func (s *LabelStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *LabelStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *LabelStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *LabelStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LabelStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LabelStatement) GetStatements() IStatement        { return s.Statement }
func (s *LabelStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LabelStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *LabelStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *LabelStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LabelStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.Label, s.Statement)
}
