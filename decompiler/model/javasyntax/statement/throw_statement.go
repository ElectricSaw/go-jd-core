package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewThrowStatement(expression model.IExpression) ThrowStatement {
	return &ThrowStatement{
		Expression: expression,
	}
}

type ThrowStatement struct {
	util.DefaultBase[IStatement]

	Expression model.IExpression
}

func (s *ThrowStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitThrowStatement(s)
}

func (s *ThrowStatement) IsBreakStatement() bool                    { return false }
func (s *ThrowStatement) IsContinueStatement() bool                 { return false }
func (s *ThrowStatement) IsExpressionStatement() bool               { return false }
func (s *ThrowStatement) IsForStatement() bool                      { return false }
func (s *ThrowStatement) IsIfStatement() bool                       { return false }
func (s *ThrowStatement) IsIfElseStatement() bool                   { return false }
func (s *ThrowStatement) IsLabelStatement() bool                    { return false }
func (s *ThrowStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ThrowStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ThrowStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ThrowStatement) IsMonitorExitStatement() bool              { return false }
func (s *ThrowStatement) IsReturnStatement() bool                   { return false }
func (s *ThrowStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ThrowStatement) IsStatements() bool                        { return false }
func (s *ThrowStatement) IsSwitchStatement() bool                   { return false }
func (s *ThrowStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ThrowStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ThrowStatement) IsThrowStatement() bool                    { return true }
func (s *ThrowStatement) IsTryStatement() bool                      { return false }
func (s *ThrowStatement) IsWhileStatement() bool                    { return false }

func (s *ThrowStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ThrowStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *ThrowStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ThrowStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ThrowStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ThrowStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ThrowStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ThrowStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ThrowStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ThrowStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ThrowStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ThrowStatement) String() string {
	return fmt.Sprintf("ThrowStatement{throw %s}", s.Expression)
}
