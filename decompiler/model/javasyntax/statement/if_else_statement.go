package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewIfElseStatement(condition model.IExpression, statements, elseStatements IStatement) IfElseStatement {
	return &IfElseStatement{
		Condition:      condition,
		IfStatements:   statements,
		ElseStatements: elseStatements,
	}
}

type IfElseStatement struct {
	util.DefaultBase[IStatement]

	Condition      model.IExpression
	IfStatements   IStatement
	ElseStatements IStatement
}

func (s *IfElseStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitIfElseStatement(s)
}

func (s *IfElseStatement) IsBreakStatement() bool                    { return false }
func (s *IfElseStatement) IsContinueStatement() bool                 { return false }
func (s *IfElseStatement) IsExpressionStatement() bool               { return false }
func (s *IfElseStatement) IsForStatement() bool                      { return false }
func (s *IfElseStatement) IsIfStatement() bool                       { return s.Condition != nil }
func (s *IfElseStatement) IsIfElseStatement() bool                   { return true }
func (s *IfElseStatement) IsLabelStatement() bool                    { return false }
func (s *IfElseStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *IfElseStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *IfElseStatement) IsMonitorEnterStatement() bool             { return false }
func (s *IfElseStatement) IsMonitorExitStatement() bool              { return false }
func (s *IfElseStatement) IsReturnStatement() bool                   { return false }
func (s *IfElseStatement) IsReturnExpressionStatement() bool         { return false }
func (s *IfElseStatement) IsStatements() bool                        { return false }
func (s *IfElseStatement) IsSwitchStatement() bool                   { return false }
func (s *IfElseStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *IfElseStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *IfElseStatement) IsThrowStatement() bool                    { return false }
func (s *IfElseStatement) IsTryStatement() bool                      { return false }
func (s *IfElseStatement) IsWhileStatement() bool                    { return false }

func (s *IfElseStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *IfElseStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *IfElseStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *IfElseStatement) GetElseStatements() IStatement    { return s.ElseStatements }
func (s *IfElseStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *IfElseStatement) GetStatements() IStatement        { return s.IfStatements }
func (s *IfElseStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *IfElseStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *IfElseStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *IfElseStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *IfElseStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *IfElseStatement) String() string {
	return ""
}
