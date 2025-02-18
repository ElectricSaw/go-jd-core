package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var NoStmt = NewNoStatement()

func NewNoStatement() NoStatement {
	return NoStatement{}
}

type NoStatement struct {
	util.DefaultBase[IStatement]
}

func (s *NoStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) IsBreakStatement() bool                    { return false }
func (s *NoStatement) IsContinueStatement() bool                 { return false }
func (s *NoStatement) IsExpressionStatement() bool               { return false }
func (s *NoStatement) IsForStatement() bool                      { return false }
func (s *NoStatement) IsIfStatement() bool                       { return false }
func (s *NoStatement) IsIfElseStatement() bool                   { return false }
func (s *NoStatement) IsLabelStatement() bool                    { return false }
func (s *NoStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *NoStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *NoStatement) IsMonitorEnterStatement() bool             { return false }
func (s *NoStatement) IsMonitorExitStatement() bool              { return false }
func (s *NoStatement) IsReturnStatement() bool                   { return false }
func (s *NoStatement) IsReturnExpressionStatement() bool         { return false }
func (s *NoStatement) IsStatements() bool                        { return false }
func (s *NoStatement) IsSwitchStatement() bool                   { return false }
func (s *NoStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *NoStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *NoStatement) IsThrowStatement() bool                    { return false }
func (s *NoStatement) IsTryStatement() bool                      { return false }
func (s *NoStatement) IsWhileStatement() bool                    { return false }

func (s *NoStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *NoStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *NoStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *NoStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *NoStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *NoStatement) GetStatements() IStatement        { return &NoStmt }
func (s *NoStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *NoStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *NoStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *NoStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *NoStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *NoStatement) String() string {
	return "NoStatement"
}
