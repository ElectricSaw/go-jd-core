package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewStatements() Statements {
	s := &Statements{
		DefaultList: *util.NewDefaultList[IStatement]().(*util.DefaultList[IStatement]),
	}
	return s
}

func NewStatementsWithList(list util.IList[IStatement]) Statements {
	return NewStatementsWithSlice(list.ToSlice())
}

func NewStatementsWithSlice(slice []IStatement) Statements {
	s := &Statements{
		DefaultList: *util.NewDefaultListWithSlice(slice).(*util.DefaultList[IStatement]),
	}
	return s
}

type Statements struct {
	util.DefaultList[IStatement]
}

func (s *Statements) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitStatements(s)
}

func (s *Statements) IsBreakStatement() bool                    { return false }
func (s *Statements) IsContinueStatement() bool                 { return false }
func (s *Statements) IsExpressionStatement() bool               { return false }
func (s *Statements) IsForStatement() bool                      { return false }
func (s *Statements) IsIfStatement() bool                       { return false }
func (s *Statements) IsIfElseStatement() bool                   { return false }
func (s *Statements) IsLabelStatement() bool                    { return false }
func (s *Statements) IsLambdaExpressionStatement() bool         { return false }
func (s *Statements) IsLocalVariableDeclarationStatement() bool { return false }
func (s *Statements) IsMonitorEnterStatement() bool             { return false }
func (s *Statements) IsMonitorExitStatement() bool              { return false }
func (s *Statements) IsReturnStatement() bool                   { return false }
func (s *Statements) IsReturnExpressionStatement() bool         { return false }
func (s *Statements) IsStatements() bool                        { return true }
func (s *Statements) IsSwitchStatement() bool                   { return false }
func (s *Statements) IsSwitchStatementLabelBlock() bool         { return false }
func (s *Statements) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *Statements) IsThrowStatement() bool                    { return false }
func (s *Statements) IsTryStatement() bool                      { return false }
func (s *Statements) IsWhileStatement() bool                    { return false }

func (s *Statements) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *Statements) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *Statements) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *Statements) GetElseStatements() IStatement    { return &NoStmt }
func (s *Statements) GetFinallyStatements() IStatement { return &NoStmt }
func (s *Statements) GetStatements() IStatement        { return &NoStmt }
func (s *Statements) GetTryStatements() IStatement     { return &NoStmt }

func (s *Statements) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *Statements) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *Statements) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *Statements) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *Statements) String() string {
	return "Statements{}"
}
