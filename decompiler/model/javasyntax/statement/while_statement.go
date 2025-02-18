package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewWhileStatement(condition model.IExpression, statements IStatement) WhileStatement {
	return &WhileStatement{
		Condition:  condition,
		Statements: statements,
	}
}

type WhileStatement struct {
	util.DefaultBase[IStatement]

	Condition  model.IExpression
	Statements IStatement
}

func (s *WhileStatement) IsBreakStatement() bool                    { return false }
func (s *WhileStatement) IsContinueStatement() bool                 { return false }
func (s *WhileStatement) IsExpressionStatement() bool               { return false }
func (s *WhileStatement) IsForStatement() bool                      { return false }
func (s *WhileStatement) IsIfStatement() bool                       { return false }
func (s *WhileStatement) IsIfElseStatement() bool                   { return false }
func (s *WhileStatement) IsLabelStatement() bool                    { return false }
func (s *WhileStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *WhileStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *WhileStatement) IsMonitorEnterStatement() bool             { return false }
func (s *WhileStatement) IsMonitorExitStatement() bool              { return false }
func (s *WhileStatement) IsReturnStatement() bool                   { return false }
func (s *WhileStatement) IsReturnExpressionStatement() bool         { return false }
func (s *WhileStatement) IsStatements() bool                        { return false }
func (s *WhileStatement) IsSwitchStatement() bool                   { return false }
func (s *WhileStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *WhileStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *WhileStatement) IsThrowStatement() bool                    { return false }
func (s *WhileStatement) IsTryStatement() bool                      { return false }
func (s *WhileStatement) IsWhileStatement() bool                    { return true }

func (s *WhileStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *WhileStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *WhileStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *WhileStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *WhileStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *WhileStatement) GetStatements() IStatement        { return s.Statements }
func (s *WhileStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *WhileStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *WhileStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *WhileStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *WhileStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *WhileStatement) String() string {
	return ""
}
