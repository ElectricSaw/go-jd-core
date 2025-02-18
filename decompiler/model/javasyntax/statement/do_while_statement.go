package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewDoWhileStatement(condition model.IExpression, statements IStatement) DoWhileStatement {
	return &DoWhileStatement{
		Condition:  condition,
		Statements: statements,
	}
}

type DoWhileStatement struct {
	util.DefaultBase[IStatement]

	Condition  model.IExpression
	Statements IStatement
}

func (s *DoWhileStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitDoWhileStatement(s)
}

func (s *DoWhileStatement) IsBreakStatement() bool                    { return false }
func (s *DoWhileStatement) IsContinueStatement() bool                 { return false }
func (s *DoWhileStatement) IsExpressionStatement() bool               { return false }
func (s *DoWhileStatement) IsForStatement() bool                      { return false }
func (s *DoWhileStatement) IsIfStatement() bool                       { return false }
func (s *DoWhileStatement) IsIfElseStatement() bool                   { return false }
func (s *DoWhileStatement) IsLabelStatement() bool                    { return false }
func (s *DoWhileStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *DoWhileStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *DoWhileStatement) IsMonitorEnterStatement() bool             { return false }
func (s *DoWhileStatement) IsMonitorExitStatement() bool              { return false }
func (s *DoWhileStatement) IsReturnStatement() bool                   { return false }
func (s *DoWhileStatement) IsReturnExpressionStatement() bool         { return false }
func (s *DoWhileStatement) IsStatements() bool                        { return false }
func (s *DoWhileStatement) IsSwitchStatement() bool                   { return false }
func (s *DoWhileStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *DoWhileStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *DoWhileStatement) IsThrowStatement() bool                    { return false }
func (s *DoWhileStatement) IsTryStatement() bool                      { return false }
func (s *DoWhileStatement) IsWhileStatement() bool                    { return false }

func (s *DoWhileStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *DoWhileStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *DoWhileStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *DoWhileStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *DoWhileStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *DoWhileStatement) GetStatements() IStatement        { return s.Statements }
func (s *DoWhileStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *DoWhileStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *DoWhileStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *DoWhileStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *DoWhileStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *DoWhileStatement) String() string {
	return ""
}
