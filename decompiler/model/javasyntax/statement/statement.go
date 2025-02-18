package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAbstractStatement() *AbstractStatement {
	s := &AbstractStatement{}
	s.SetValue(s)
	return s
}

type AbstractStatement struct {
	util.DefaultBase[IStatement]
}

func (s *AbstractStatement) AcceptStatement(visitor IStatementVisitor) {}

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

func (s *AbstractStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *AbstractStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *AbstractStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *AbstractStatement) GetElseStatements() IStatement { return NoStmt }
func (s *AbstractStatement) GetFinallyStatements() IStatement {
	return NoStmt
}
func (s *AbstractStatement) GetStatements() IStatement    { return NoStmt }
func (s *AbstractStatement) GetTryStatements() IStatement { return NoStmt }

func (s *AbstractStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *AbstractStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *AbstractStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *AbstractStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *AbstractStatement) String() string {
	return ""
}
