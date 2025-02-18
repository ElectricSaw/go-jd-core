package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewIfStatement(condition model.IExpression, statements IStatement) IfStatement {
	return &IfStatement{
		Condition:  condition,
		Statements: statements,
	}
}

type IfStatement struct {
	util.DefaultBase[IStatement]

	Condition  model.IExpression
	Statements IStatement
}

func (s *IfStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitIfStatement(s)
}

func (s *IfStatement) IsBreakStatement() bool                    { return false }
func (s *IfStatement) IsContinueStatement() bool                 { return false }
func (s *IfStatement) IsExpressionStatement() bool               { return false }
func (s *IfStatement) IsForStatement() bool                      { return false }
func (s *IfStatement) IsIfStatement() bool                       { return s.Condition != nil }
func (s *IfStatement) IsIfElseStatement() bool                   { return false }
func (s *IfStatement) IsLabelStatement() bool                    { return false }
func (s *IfStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *IfStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *IfStatement) IsMonitorEnterStatement() bool             { return false }
func (s *IfStatement) IsMonitorExitStatement() bool              { return false }
func (s *IfStatement) IsReturnStatement() bool                   { return false }
func (s *IfStatement) IsReturnExpressionStatement() bool         { return false }
func (s *IfStatement) IsStatements() bool                        { return false }
func (s *IfStatement) IsSwitchStatement() bool                   { return false }
func (s *IfStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *IfStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *IfStatement) IsThrowStatement() bool                    { return false }
func (s *IfStatement) IsTryStatement() bool                      { return false }
func (s *IfStatement) IsWhileStatement() bool                    { return false }

func (s *IfStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *IfStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *IfStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *IfStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *IfStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *IfStatement) GetStatements() IStatement        { return s.Statements }
func (s *IfStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *IfStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *IfStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *IfStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *IfStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *IfStatement) String() string {
	return ""
}
