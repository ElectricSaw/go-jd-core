package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewForEachStatement(typ model.IType, name string, expression model.IExpression, statement IStatement) ForEachStatement {
	return &ForEachStatement{
		Type:       typ,
		Name:       name,
		Expression: expression,
		Statement:  statement,
	}
}

type ForEachStatement struct {
	util.DefaultBase[IStatement]

	Type       model.IType
	Name       string
	Expression model.IExpression
	Statement  IStatement
}

func (s *ForEachStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitForEachStatement(s)
}

func (s *ForEachStatement) IsBreakStatement() bool                    { return false }
func (s *ForEachStatement) IsContinueStatement() bool                 { return false }
func (s *ForEachStatement) IsExpressionStatement() bool               { return false }
func (s *ForEachStatement) IsForStatement() bool                      { return false }
func (s *ForEachStatement) IsIfStatement() bool                       { return false }
func (s *ForEachStatement) IsIfElseStatement() bool                   { return false }
func (s *ForEachStatement) IsLabelStatement() bool                    { return false }
func (s *ForEachStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ForEachStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ForEachStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ForEachStatement) IsMonitorExitStatement() bool              { return false }
func (s *ForEachStatement) IsReturnStatement() bool                   { return false }
func (s *ForEachStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ForEachStatement) IsStatements() bool                        { return false }
func (s *ForEachStatement) IsSwitchStatement() bool                   { return false }
func (s *ForEachStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ForEachStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ForEachStatement) IsThrowStatement() bool                    { return false }
func (s *ForEachStatement) IsTryStatement() bool                      { return false }
func (s *ForEachStatement) IsWhileStatement() bool                    { return false }

func (s *ForEachStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ForEachStatement) GetExpression() model.IExpression { return s.Expression }
func (s *ForEachStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ForEachStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ForEachStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ForEachStatement) GetStatements() IStatement        { return s.Statement }
func (s *ForEachStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ForEachStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ForEachStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ForEachStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ForEachStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ForEachStatement) String() string {
	return ""
}
