package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var Return = NewReturnStatement()

func NewReturnStatement() ReturnStatement {
	return &ReturnStatement{}
}

type ReturnStatement struct {
	util.DefaultBase[IStatement]
}

func (s *ReturnStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) IsBreakStatement() bool                    { return false }
func (s *ReturnStatement) IsContinueStatement() bool                 { return false }
func (s *ReturnStatement) IsExpressionStatement() bool               { return false }
func (s *ReturnStatement) IsForStatement() bool                      { return false }
func (s *ReturnStatement) IsIfStatement() bool                       { return false }
func (s *ReturnStatement) IsIfElseStatement() bool                   { return false }
func (s *ReturnStatement) IsLabelStatement() bool                    { return false }
func (s *ReturnStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ReturnStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ReturnStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ReturnStatement) IsMonitorExitStatement() bool              { return false }
func (s *ReturnStatement) IsReturnStatement() bool                   { return true }
func (s *ReturnStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ReturnStatement) IsStatements() bool                        { return false }
func (s *ReturnStatement) IsSwitchStatement() bool                   { return false }
func (s *ReturnStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ReturnStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ReturnStatement) IsThrowStatement() bool                    { return false }
func (s *ReturnStatement) IsTryStatement() bool                      { return false }
func (s *ReturnStatement) IsWhileStatement() bool                    { return false }

func (s *ReturnStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ReturnStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *ReturnStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ReturnStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ReturnStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ReturnStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ReturnStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ReturnStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ReturnStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ReturnStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ReturnStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ReturnStatement) String() string {
	return "ReturnStatement{}"
}
