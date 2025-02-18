package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewReturnExpressionStatement(expression model.IExpression) ReturnExpressionStatement {
	return &ReturnExpressionStatement{
		LineNumber: expression.GetLineNumber(),
		Expression: expression,
	}
}

func NewReturnExpressionStatementWithAll(lineNumber int, expression model.IExpression) ReturnExpressionStatement {
	return &ReturnExpressionStatement{
		LineNumber: lineNumber,
		Expression: expression,
	}
}

type ReturnExpressionStatement struct {
	util.DefaultBase[IStatement]

	LineNumber int
	Expression model.IExpression
}

//func (s *ReturnExpressionStatement) GenericExpression() model.IExpression {
//	return s.expression
//}

func (s *ReturnExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitReturnExpressionStatement(s)
}

func (s *ReturnExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *ReturnExpressionStatement) IsExpressionStatement() bool               { return false }
func (s *ReturnExpressionStatement) IsForStatement() bool                      { return false }
func (s *ReturnExpressionStatement) IsIfStatement() bool                       { return false }
func (s *ReturnExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ReturnExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ReturnExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ReturnExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *ReturnExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsReturnExpressionStatement() bool         { return true }
func (s *ReturnExpressionStatement) IsStatements() bool                        { return false }
func (s *ReturnExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ReturnExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ReturnExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsTryStatement() bool                      { return false }
func (s *ReturnExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *ReturnExpressionStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ReturnExpressionStatement) GetExpression() model.IExpression { return s.Expression }
func (s *ReturnExpressionStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ReturnExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ReturnExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ReturnExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ReturnExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ReturnExpressionStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ReturnExpressionStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ReturnExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ReturnExpressionStatement) GetLineNumber() int                        { return s.LineNumber }

func (s *ReturnExpressionStatement) String() string {
	return fmt.Sprintf("ReturnExpressionStatement{return %s}", s.Expression)
}
