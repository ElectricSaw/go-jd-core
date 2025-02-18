package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewForStatementWithDeclaration(declaration model.ILocalVariableDeclaration,
	condition model.IExpression, update model.IExpression,
	statements IStatement) ForStatement {
	return &ForStatement{
		declaration: declaration,
		condition:   condition,
		update:      update,
		statements:  statements,
	}
}

func NewForStatementWithInit(init, condition, update model.IExpression, statements IStatement) ForStatement {
	return &ForStatement{
		Init:       init,
		Condition:  condition,
		Update:     update,
		Statements: statements,
	}
}

type ForStatement struct {
	util.DefaultBase[IStatement]

	Declaration model.ILocalVariableDeclaration
	Init        model.IExpression
	Condition   model.IExpression
	Update      model.IExpression
	Statements  IStatement
}

func (s *ForStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitForStatement(s)
}

func (s *ForStatement) IsBreakStatement() bool                    { return false }
func (s *ForStatement) IsContinueStatement() bool                 { return false }
func (s *ForStatement) IsExpressionStatement() bool               { return false }
func (s *ForStatement) IsForStatement() bool                      { return false }
func (s *ForStatement) IsIfStatement() bool                       { return false }
func (s *ForStatement) IsIfElseStatement() bool                   { return false }
func (s *ForStatement) IsLabelStatement() bool                    { return false }
func (s *ForStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ForStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ForStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ForStatement) IsMonitorExitStatement() bool              { return false }
func (s *ForStatement) IsReturnStatement() bool                   { return false }
func (s *ForStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ForStatement) IsStatements() bool                        { return false }
func (s *ForStatement) IsSwitchStatement() bool                   { return false }
func (s *ForStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ForStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ForStatement) IsThrowStatement() bool                    { return false }
func (s *ForStatement) IsTryStatement() bool                      { return false }
func (s *ForStatement) IsWhileStatement() bool                    { return false }

func (s *ForStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *ForStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *ForStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ForStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ForStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ForStatement) GetStatements() IStatement        { return s.Statements }
func (s *ForStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ForStatement) GetInit() model.IExpression   { return s.Init }
func (s *ForStatement) GetUpdate() model.IExpression { return s.Update }

func (s *ForStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ForStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ForStatement) String() string {
	return fmt.Sprintf("ForStatement{%s or %s; %s; %s", s.Declaration, s.Init, s.Condition, s.Update)
}
