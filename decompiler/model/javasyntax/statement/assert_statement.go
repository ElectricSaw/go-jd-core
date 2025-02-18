package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAssertStatement(condition model.IExpression, message model.IExpression) AssertStatement {
	s := AssertStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Message:     message,
	}
	s.SetValue(&s)
	return s
}

type AssertStatement struct {
	util.DefaultBase[IStatement]

	Condition model.IExpression
	Message   model.IExpression
}

func (s *AssertStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitAssertStatement(s)
}

func (s *AssertStatement) IsBreakStatement() bool                    { return false }
func (s *AssertStatement) IsContinueStatement() bool                 { return false }
func (s *AssertStatement) IsExpressionStatement() bool               { return false }
func (s *AssertStatement) IsForStatement() bool                      { return false }
func (s *AssertStatement) IsIfStatement() bool                       { return false }
func (s *AssertStatement) IsIfElseStatement() bool                   { return false }
func (s *AssertStatement) IsLabelStatement() bool                    { return false }
func (s *AssertStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *AssertStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *AssertStatement) IsMonitorEnterStatement() bool             { return false }
func (s *AssertStatement) IsMonitorExitStatement() bool              { return false }
func (s *AssertStatement) IsReturnStatement() bool                   { return false }
func (s *AssertStatement) IsReturnExpressionStatement() bool         { return false }
func (s *AssertStatement) IsStatements() bool                        { return false }
func (s *AssertStatement) IsSwitchStatement() bool                   { return false }
func (s *AssertStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *AssertStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *AssertStatement) IsThrowStatement() bool                    { return false }
func (s *AssertStatement) IsTryStatement() bool                      { return false }
func (s *AssertStatement) IsWhileStatement() bool                    { return false }

func (s *AssertStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *AssertStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *AssertStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *AssertStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *AssertStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *AssertStatement) GetStatements() IStatement        { return &NoStmt }
func (s *AssertStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *AssertStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *AssertStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *AssertStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *AssertStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *AssertStatement) String() string {
	return "AssertStatement{}"
}
