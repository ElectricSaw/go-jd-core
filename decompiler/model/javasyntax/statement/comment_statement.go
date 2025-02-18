package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewCommentStatement(text string) CommentStatement {
	return &CommentStatement{
		Text: text,
	}
}

type CommentStatement struct {
	util.DefaultBase[IStatement]

	Text string
}

func (s *CommentStatement) IsContinueStatement() bool {
	return true
}

func (s *CommentStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitCommentStatement(s)
}

func (s *CommentStatement) IsBreakStatement() bool                    { return false }
func (s *CommentStatement) IsContinueStatement() bool                 { return false }
func (s *CommentStatement) IsExpressionStatement() bool               { return false }
func (s *CommentStatement) IsForStatement() bool                      { return false }
func (s *CommentStatement) IsIfStatement() bool                       { return false }
func (s *CommentStatement) IsIfElseStatement() bool                   { return false }
func (s *CommentStatement) IsLabelStatement() bool                    { return false }
func (s *CommentStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *CommentStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *CommentStatement) IsMonitorEnterStatement() bool             { return false }
func (s *CommentStatement) IsMonitorExitStatement() bool              { return false }
func (s *CommentStatement) IsReturnStatement() bool                   { return false }
func (s *CommentStatement) IsReturnExpressionStatement() bool         { return false }
func (s *CommentStatement) IsStatements() bool                        { return false }
func (s *CommentStatement) IsSwitchStatement() bool                   { return false }
func (s *CommentStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *CommentStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *CommentStatement) IsThrowStatement() bool                    { return false }
func (s *CommentStatement) IsTryStatement() bool                      { return false }
func (s *CommentStatement) IsWhileStatement() bool                    { return false }

func (s *CommentStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *CommentStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *CommentStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *CommentStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *CommentStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *CommentStatement) GetStatements() IStatement        { return &NoStmt }
func (s *CommentStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *CommentStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *CommentStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *CommentStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *CommentStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *CommentStatement) String() string {
	return "CommentStatement{}"
}
