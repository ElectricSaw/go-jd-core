package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLambdaExpressionStatement(expression model.IExpression) LambdaExpressionStatement {
	return &LambdaExpressionStatement{
		Expression: expression,
	}
}

type LambdaExpressionStatement struct {
	util.DefaultBase[IStatement]

	Expression model.IExpression
}

func (s *LambdaExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitLambdaExpressionStatement(s)
}

func (s *LambdaExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *LambdaExpressionStatement) IsExpressionStatement() bool               { return false }
func (s *LambdaExpressionStatement) IsForStatement() bool                      { return false }
func (s *LambdaExpressionStatement) IsIfStatement() bool                       { return false }
func (s *LambdaExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsLambdaExpressionStatement() bool         { return true }
func (s *LambdaExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *LambdaExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *LambdaExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *LambdaExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsReturnExpressionStatement() bool         { return false }
func (s *LambdaExpressionStatement) IsStatements() bool                        { return false }
func (s *LambdaExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *LambdaExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *LambdaExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsTryStatement() bool                      { return false }
func (s *LambdaExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *LambdaExpressionStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *LambdaExpressionStatement) GetExpression() model.IExpression { return s.Expression }
func (s *LambdaExpressionStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *LambdaExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LambdaExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LambdaExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *LambdaExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LambdaExpressionStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *LambdaExpressionStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *LambdaExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LambdaExpressionStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *LambdaExpressionStatement) String() string {
	return fmt.Sprintf("LambdaExpressionStatement{%s}", s.Expression)
}
