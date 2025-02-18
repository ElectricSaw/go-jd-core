package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewExpressionStatement(expression model.IExpression) ExpressionStatement {
	return &ExpressionStatement{
		Expression: expression,
	}
}

type ExpressionStatement struct {
	util.DefaultBase[IStatement]

	Expression model.IExpression
}

func (s *ExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *ExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *ExpressionStatement) IsExpressionStatement() bool               { return true }
func (s *ExpressionStatement) IsForStatement() bool                      { return false }
func (s *ExpressionStatement) IsIfStatement() bool                       { return false }
func (s *ExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *ExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *ExpressionStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *ExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *ExpressionStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ExpressionStatement) IsStatements() bool                        { return false }
func (s *ExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *ExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *ExpressionStatement) IsTryStatement() bool                      { return false }
func (s *ExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *ExpressionStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ExpressionStatement) GetExpression() model.IExpression { return s.Expression }
func (s *ExpressionStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ExpressionStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ExpressionStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ExpressionStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{%s}", s.expression)
}
