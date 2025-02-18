package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewByteCodeStatement(text string) ByteCodeStatement {
	return &ByteCodeStatement{
		Text: text,
	}
}

type ByteCodeStatement struct {
	util.DefaultBase[IStatement]

	Text string
}

func (s *ByteCodeStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitByteCodeStatement(s)
}

func (s *ByteCodeStatement) IsBreakStatement() bool                    { return false }
func (s *ByteCodeStatement) IsContinueStatement() bool                 { return false }
func (s *ByteCodeStatement) IsExpressionStatement() bool               { return false }
func (s *ByteCodeStatement) IsForStatement() bool                      { return false }
func (s *ByteCodeStatement) IsIfStatement() bool                       { return false }
func (s *ByteCodeStatement) IsIfElseStatement() bool                   { return false }
func (s *ByteCodeStatement) IsLabelStatement() bool                    { return false }
func (s *ByteCodeStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ByteCodeStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ByteCodeStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ByteCodeStatement) IsMonitorExitStatement() bool              { return false }
func (s *ByteCodeStatement) IsReturnStatement() bool                   { return false }
func (s *ByteCodeStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ByteCodeStatement) IsStatements() bool                        { return false }
func (s *ByteCodeStatement) IsSwitchStatement() bool                   { return false }
func (s *ByteCodeStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ByteCodeStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ByteCodeStatement) IsThrowStatement() bool                    { return false }
func (s *ByteCodeStatement) IsTryStatement() bool                      { return false }
func (s *ByteCodeStatement) IsWhileStatement() bool                    { return false }

func (s *ByteCodeStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *ByteCodeStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *ByteCodeStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *ByteCodeStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ByteCodeStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ByteCodeStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ByteCodeStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ByteCodeStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *ByteCodeStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *ByteCodeStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ByteCodeStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *ByteCodeStatement) String() string {
	return "ByteCodeStatement{}"
}
