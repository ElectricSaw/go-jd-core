package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTryStatement(tryStatements IStatement, catchClauses util.DefaultList[*CatchClause], finallyStatement IStatement) TryStatement {
	return &TryStatement{
		resources:        nil,
		tryStatements:    tryStatements,
		catchClauses:     catchClauses,
		finallyStatement: finallyStatement,
	}
}

func NewTryStatementWithAll(resource util.DefaultList[*Resource], tryStatements IStatement, catchClauses util.DefaultList[*CatchClause], finallyStatement IStatement) TryStatement {
	return &TryStatement{
		Resources:         resource,
		TryStatements:     tryStatements,
		CatchClause:       catchClauses,
		FinallyStatements: finallyStatement,
	}
}

type TryStatement struct {
	util.DefaultBase[IStatement]

	Resources         util.DefaultList[*Resource]
	TryStatements     IStatement
	CatchClause       util.DefaultList[*CatchClause]
	FinallyStatements IStatement
}

func (s *TryStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatement(s)
}

func (s *TryStatement) IsBreakStatement() bool                    { return false }
func (s *TryStatement) IsContinueStatement() bool                 { return false }
func (s *TryStatement) IsExpressionStatement() bool               { return false }
func (s *TryStatement) IsForStatement() bool                      { return false }
func (s *TryStatement) IsIfStatement() bool                       { return false }
func (s *TryStatement) IsIfElseStatement() bool                   { return false }
func (s *TryStatement) IsLabelStatement() bool                    { return false }
func (s *TryStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *TryStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *TryStatement) IsMonitorEnterStatement() bool             { return false }
func (s *TryStatement) IsMonitorExitStatement() bool              { return false }
func (s *TryStatement) IsReturnStatement() bool                   { return false }
func (s *TryStatement) IsReturnExpressionStatement() bool         { return false }
func (s *TryStatement) IsStatements() bool                        { return false }
func (s *TryStatement) IsSwitchStatement() bool                   { return false }
func (s *TryStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *TryStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *TryStatement) IsThrowStatement() bool                    { return false }
func (s *TryStatement) IsTryStatement() bool                      { return true }
func (s *TryStatement) IsWhileStatement() bool                    { return false }

func (s *TryStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *TryStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *TryStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *TryStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *TryStatement) GetFinallyStatements() IStatement { return s.FinallyStatements }
func (s *TryStatement) GetStatements() IStatement        { return &NoStmt }
func (s *TryStatement) GetTryStatements() IStatement     { return s.TryStatements }

func (s *TryStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *TryStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *TryStatement) GetCatchClauses() util.IList[*CatchClause] { return &s.CatchClause }
func (s *TryStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *TryStatement) String() string {
	return ""
}

func NewResource(typ *model.ObjectType, name string, expression model.IExpression) Resource {
	return &Resource{
		Type:       typ,
		Name:       name,
		Expression: expression,
	}
}

type Resource struct {
	Type       *model.ObjectType
	Name       string
	Expression model.IExpression
}

func (r *Resource) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatementResource(r)
}

func (r *Resource) IsBreakStatement() bool                    { return false }
func (r *Resource) IsContinueStatement() bool                 { return false }
func (r *Resource) IsExpressionStatement() bool               { return false }
func (r *Resource) IsForStatement() bool                      { return false }
func (r *Resource) IsIfStatement() bool                       { return false }
func (r *Resource) IsIfElseStatement() bool                   { return false }
func (r *Resource) IsLabelStatement() bool                    { return false }
func (r *Resource) IsLambdaExpressionStatement() bool         { return false }
func (r *Resource) IsLocalVariableDeclarationStatement() bool { return false }
func (r *Resource) IsMonitorEnterStatement() bool             { return false }
func (r *Resource) IsMonitorExitStatement() bool              { return false }
func (r *Resource) IsReturnStatement() bool                   { return false }
func (r *Resource) IsReturnExpressionStatement() bool         { return false }
func (r *Resource) IsStatements() bool                        { return false }
func (r *Resource) IsSwitchStatement() bool                   { return false }
func (r *Resource) IsSwitchStatementLabelBlock() bool         { return false }
func (r *Resource) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (r *Resource) IsThrowStatement() bool                    { return false }
func (r *Resource) IsTryStatement() bool                      { return false }
func (r *Resource) IsWhileStatement() bool                    { return false }

func (r *Resource) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (r *Resource) GetExpression() model.IExpression { return r.Expression }
func (r *Resource) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (r *Resource) GetElseStatements() IStatement    { return &NoStmt }
func (r *Resource) GetFinallyStatements() IStatement { return &NoStmt }
func (r *Resource) GetStatements() IStatement        { return &NoStmt }
func (r *Resource) GetTryStatements() IStatement     { return &NoStmt }

func (r *Resource) GetInit() model.IExpression   { return &model.NeNoExpression }
func (r *Resource) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (r *Resource) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (r *Resource) GetLineNumber() int                        { return model.UnknownLineNumber }

func (r *Resource) String() string {
	return "Resource{}"
}

func NewCatchClause(lineNumber int, typ *model.ObjectType, name string, statements IStatement) CatchClause {
	return &CatchClause{
		LineNumber: lineNumber,
		Type:       typ,
		Name:       name,
		Statements: statements,
	}
}

type CatchClause struct {
	LineNumber int
	Type       *model.ObjectType
	OtherType  []model.ObjectType
	Name       string
	Statements IStatement
}

func (c *CatchClause) AddType(typ model.ObjectType) {
	c.OtherType = append(c.OtherType, typ)
}

func (c *CatchClause) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatementCatchClause(c)
}

func (c *CatchClause) IsBreakStatement() bool                    { return false }
func (c *CatchClause) IsContinueStatement() bool                 { return false }
func (c *CatchClause) IsExpressionStatement() bool               { return false }
func (c *CatchClause) IsForStatement() bool                      { return false }
func (c *CatchClause) IsIfStatement() bool                       { return false }
func (c *CatchClause) IsIfElseStatement() bool                   { return false }
func (c *CatchClause) IsLabelStatement() bool                    { return false }
func (c *CatchClause) IsLambdaExpressionStatement() bool         { return false }
func (c *CatchClause) IsLocalVariableDeclarationStatement() bool { return false }
func (c *CatchClause) IsMonitorEnterStatement() bool             { return false }
func (c *CatchClause) IsMonitorExitStatement() bool              { return false }
func (c *CatchClause) IsReturnStatement() bool                   { return false }
func (c *CatchClause) IsReturnExpressionStatement() bool         { return false }
func (c *CatchClause) IsStatements() bool                        { return false }
func (c *CatchClause) IsSwitchStatement() bool                   { return false }
func (c *CatchClause) IsSwitchStatementLabelBlock() bool         { return false }
func (c *CatchClause) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (c *CatchClause) IsThrowStatement() bool                    { return false }
func (c *CatchClause) IsTryStatement() bool                      { return false }
func (c *CatchClause) IsWhileStatement() bool                    { return false }

func (c *CatchClause) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (c *CatchClause) GetExpression() model.IExpression { return &model.NeNoExpression }
func (c *CatchClause) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (c *CatchClause) GetElseStatements() IStatement    { return &NoStmt }
func (c *CatchClause) GetFinallyStatements() IStatement { return &NoStmt }
func (c *CatchClause) GetStatements() IStatement        { return c.Statements }
func (c *CatchClause) GetTryStatements() IStatement     { return &NoStmt }

func (c *CatchClause) GetInit() model.IExpression   { return &model.NeNoExpression }
func (c *CatchClause) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (c *CatchClause) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (c *CatchClause) GetLineNumber() int                        { return c.LineNumber }

func (c *CatchClause) String() string {
	return "CatchClause{}"
}
