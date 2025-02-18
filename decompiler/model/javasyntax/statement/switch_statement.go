package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var DefaultLabe1 = NewDefaultLabel()

func NewSwitchStatement(condition model.IExpression, blocks util.DefaultList[*Block]) SwitchStatement {
	return &SwitchStatement{
		Condition: condition,
		Blocks:    blocks,
	}
}

type SwitchStatement struct {
	util.DefaultBase[IStatement]

	Condition model.IExpression
	Blocks    util.DefaultList[*Block]
}

func (s *SwitchStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatement(s)
}

func (s *SwitchStatement) IsBreakStatement() bool                    { return false }
func (s *SwitchStatement) IsContinueStatement() bool                 { return false }
func (s *SwitchStatement) IsExpressionStatement() bool               { return false }
func (s *SwitchStatement) IsForStatement() bool                      { return false }
func (s *SwitchStatement) IsIfStatement() bool                       { return false }
func (s *SwitchStatement) IsIfElseStatement() bool                   { return false }
func (s *SwitchStatement) IsLabelStatement() bool                    { return false }
func (s *SwitchStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *SwitchStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *SwitchStatement) IsMonitorEnterStatement() bool             { return false }
func (s *SwitchStatement) IsMonitorExitStatement() bool              { return false }
func (s *SwitchStatement) IsReturnStatement() bool                   { return false }
func (s *SwitchStatement) IsReturnExpressionStatement() bool         { return false }
func (s *SwitchStatement) IsStatements() bool                        { return false }
func (s *SwitchStatement) IsSwitchStatement() bool                   { return true }
func (s *SwitchStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *SwitchStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *SwitchStatement) IsThrowStatement() bool                    { return false }
func (s *SwitchStatement) IsTryStatement() bool                      { return false }
func (s *SwitchStatement) IsWhileStatement() bool                    { return false }

func (s *SwitchStatement) GetCondition() model.IExpression  { return s.Condition }
func (s *SwitchStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *SwitchStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *SwitchStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *SwitchStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *SwitchStatement) GetStatements() IStatement        { return &NoStmt }
func (s *SwitchStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *SwitchStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *SwitchStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *SwitchStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *SwitchStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *SwitchStatement) String() string {
	return "SwitchStatement{}"
}

// --- label --- //

func NewDefaultLabel() DefaultLabel {
	return DefaultLabel{}
}

type DefaultLabel struct {
	util.DefaultBase[IStatement]
}

func (l *DefaultLabel) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementDefaultLabel(l)
}

func (l *DefaultLabel) IsBreakStatement() bool                    { return false }
func (l *DefaultLabel) IsContinueStatement() bool                 { return false }
func (l *DefaultLabel) IsExpressionStatement() bool               { return false }
func (l *DefaultLabel) IsForStatement() bool                      { return false }
func (l *DefaultLabel) IsIfStatement() bool                       { return false }
func (l *DefaultLabel) IsIfElseStatement() bool                   { return false }
func (l *DefaultLabel) IsLabelStatement() bool                    { return false }
func (l *DefaultLabel) IsLambdaExpressionStatement() bool         { return false }
func (l *DefaultLabel) IsLocalVariableDeclarationStatement() bool { return false }
func (l *DefaultLabel) IsMonitorEnterStatement() bool             { return false }
func (l *DefaultLabel) IsMonitorExitStatement() bool              { return false }
func (l *DefaultLabel) IsReturnStatement() bool                   { return false }
func (l *DefaultLabel) IsReturnExpressionStatement() bool         { return false }
func (l *DefaultLabel) IsStatements() bool                        { return false }
func (l *DefaultLabel) IsSwitchStatement() bool                   { return false }
func (l *DefaultLabel) IsSwitchStatementLabelBlock() bool         { return false }
func (l *DefaultLabel) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (l *DefaultLabel) IsThrowStatement() bool                    { return false }
func (l *DefaultLabel) IsTryStatement() bool                      { return false }
func (l *DefaultLabel) IsWhileStatement() bool                    { return false }

func (l *DefaultLabel) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (l *DefaultLabel) GetExpression() model.IExpression { return &model.NeNoExpression }
func (l *DefaultLabel) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (l *DefaultLabel) GetElseStatements() IStatement    { return &NoStmt }
func (l *DefaultLabel) GetFinallyStatements() IStatement { return &NoStmt }
func (l *DefaultLabel) GetStatements() IStatement        { return &NoStmt }
func (l *DefaultLabel) GetTryStatements() IStatement     { return &NoStmt }

func (l *DefaultLabel) GetInit() model.IExpression   { return &model.NeNoExpression }
func (l *DefaultLabel) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (l *DefaultLabel) GetCatchClauses() []util.IList[*CatchClause] { return nil }
func (l *DefaultLabel) GetLineNumber() int                          { return model.UnknownLineNumber }

func (l *DefaultLabel) String() string {
	return "DefaultLabel{}"
}

func NewExpressionLabel(expression model.IExpression) ExpressionLabel {
	return ExpressionLabel{
		Expression: expression,
	}
}

type ExpressionLabel struct {
	Expression model.IExpression
}

func (l *ExpressionLabel) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementExpressionLabel(l)
}

func (l *ExpressionLabel) IsBreakStatement() bool                    { return false }
func (l *ExpressionLabel) IsContinueStatement() bool                 { return false }
func (l *ExpressionLabel) IsExpressionStatement() bool               { return false }
func (l *ExpressionLabel) IsForStatement() bool                      { return false }
func (l *ExpressionLabel) IsIfStatement() bool                       { return false }
func (l *ExpressionLabel) IsIfElseStatement() bool                   { return false }
func (l *ExpressionLabel) IsLabelStatement() bool                    { return false }
func (l *ExpressionLabel) IsLambdaExpressionStatement() bool         { return false }
func (l *ExpressionLabel) IsLocalVariableDeclarationStatement() bool { return false }
func (l *ExpressionLabel) IsMonitorEnterStatement() bool             { return false }
func (l *ExpressionLabel) IsMonitorExitStatement() bool              { return false }
func (l *ExpressionLabel) IsReturnStatement() bool                   { return false }
func (l *ExpressionLabel) IsReturnExpressionStatement() bool         { return false }
func (l *ExpressionLabel) IsStatements() bool                        { return false }
func (l *ExpressionLabel) IsSwitchStatement() bool                   { return false }
func (l *ExpressionLabel) IsSwitchStatementLabelBlock() bool         { return false }
func (l *ExpressionLabel) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (l *ExpressionLabel) IsThrowStatement() bool                    { return false }
func (l *ExpressionLabel) IsTryStatement() bool                      { return false }
func (l *ExpressionLabel) IsWhileStatement() bool                    { return false }

func (l *ExpressionLabel) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (l *ExpressionLabel) GetExpression() model.IExpression { return l.Expression }
func (l *ExpressionLabel) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (l *ExpressionLabel) GetElseStatements() IStatement    { return &NoStmt }
func (l *ExpressionLabel) GetFinallyStatements() IStatement { return &NoStmt }
func (l *ExpressionLabel) GetStatements() IStatement        { return &NoStmt }
func (l *ExpressionLabel) GetTryStatements() IStatement     { return &NoStmt }

func (l *ExpressionLabel) GetInit() model.IExpression   { return &model.NeNoExpression }
func (l *ExpressionLabel) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (l *ExpressionLabel) GetCatchClauses() []util.IList[*CatchClause] { return nil }
func (l *ExpressionLabel) GetLineNumber() int                          { return model.UnknownLineNumber }

func (l *ExpressionLabel) String() string {
	return fmt.Sprintf("ExpressionLabel{%s}", l.Expression)
}

// --- Block --- //

func NewBlock(statements IStatement) Block {
	return &Block{
		Statements: statements,
	}
}

type Block struct {
	Statements IStatement
}

func (b *Block) AcceptStatement(visitor IStatementVisitor) {}

func (b *Block) IsBreakStatement() bool                    { return false }
func (b *Block) IsContinueStatement() bool                 { return false }
func (b *Block) IsExpressionStatement() bool               { return false }
func (b *Block) IsForStatement() bool                      { return false }
func (b *Block) IsIfStatement() bool                       { return false }
func (b *Block) IsIfElseStatement() bool                   { return false }
func (b *Block) IsLabelStatement() bool                    { return false }
func (b *Block) IsLambdaExpressionStatement() bool         { return false }
func (b *Block) IsLocalVariableDeclarationStatement() bool { return false }
func (b *Block) IsMonitorEnterStatement() bool             { return false }
func (b *Block) IsMonitorExitStatement() bool              { return false }
func (b *Block) IsReturnStatement() bool                   { return false }
func (b *Block) IsReturnExpressionStatement() bool         { return false }
func (b *Block) IsStatements() bool                        { return false }
func (b *Block) IsSwitchStatement() bool                   { return false }
func (b *Block) IsSwitchStatementLabelBlock() bool         { return false }
func (b *Block) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (b *Block) IsThrowStatement() bool                    { return false }
func (b *Block) IsTryStatement() bool                      { return false }
func (b *Block) IsWhileStatement() bool                    { return false }

func (b *Block) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (b *Block) GetExpression() model.IExpression { return &model.NeNoExpression }
func (b *Block) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (b *Block) GetElseStatements() IStatement    { return &NoStmt }
func (b *Block) GetFinallyStatements() IStatement { return &NoStmt }
func (b *Block) GetStatements() IStatement        { return b.Statements }
func (b *Block) GetTryStatements() IStatement     { return &NoStmt }

func (b *Block) GetInit() model.IExpression   { return &model.NeNoExpression }
func (b *Block) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (b *Block) GetCatchClauses() []util.IList[*CatchClause] { return nil }
func (b *Block) GetLineNumber() int                          { return model.UnknownLineNumber }

func (b *Block) String() string {
	return ""
}

func NewLabelBlock(label ILabel, statements IStatement) LabelBlock {
	return &LabelBlock{
		Statements: statements,
		Label:      label,
	}
}

type LabelBlock struct {
	Statements IStatement
	Label      ILabel
}

func (b *LabelBlock) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementLabelBlock(b)
}

func (b *LabelBlock) IsBreakStatement() bool                    { return false }
func (b *LabelBlock) IsContinueStatement() bool                 { return false }
func (b *LabelBlock) IsExpressionStatement() bool               { return false }
func (b *LabelBlock) IsForStatement() bool                      { return false }
func (b *LabelBlock) IsIfStatement() bool                       { return false }
func (b *LabelBlock) IsIfElseStatement() bool                   { return false }
func (b *LabelBlock) IsLabelStatement() bool                    { return false }
func (b *LabelBlock) IsLambdaExpressionStatement() bool         { return false }
func (b *LabelBlock) IsLocalVariableDeclarationStatement() bool { return false }
func (b *LabelBlock) IsMonitorEnterStatement() bool             { return false }
func (b *LabelBlock) IsMonitorExitStatement() bool              { return false }
func (b *LabelBlock) IsReturnStatement() bool                   { return false }
func (b *LabelBlock) IsReturnExpressionStatement() bool         { return false }
func (b *LabelBlock) IsStatements() bool                        { return false }
func (b *LabelBlock) IsSwitchStatement() bool                   { return false }
func (b *LabelBlock) IsSwitchStatementLabelBlock() bool         { return true }
func (b *LabelBlock) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (b *LabelBlock) IsThrowStatement() bool                    { return false }
func (b *LabelBlock) IsTryStatement() bool                      { return false }
func (b *LabelBlock) IsWhileStatement() bool                    { return false }

func (b *LabelBlock) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (b *LabelBlock) GetExpression() model.IExpression { return &model.NeNoExpression }
func (b *LabelBlock) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (b *LabelBlock) GetElseStatements() IStatement    { return &NoStmt }
func (b *LabelBlock) GetFinallyStatements() IStatement { return &NoStmt }
func (b *LabelBlock) GetStatements() IStatement        { return b.Statements }
func (b *LabelBlock) GetTryStatements() IStatement     { return &NoStmt }

func (b *LabelBlock) GetInit() model.IExpression   { return &model.NeNoExpression }
func (b *LabelBlock) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (b *LabelBlock) GetCatchClauses() []util.IList[*CatchClause] { return nil }
func (b *LabelBlock) GetLineNumber() int                          { return model.UnknownLineNumber }

func (b *LabelBlock) String() string {
	return fmt.Sprintf("LabelBlock{label=%s}", b.Label)
}

func NewMultiLabelsBlock(labels util.DefaultList[ILabel], statements IStatement) MultiLabelsBlock {
	return &MultiLabelsBlock{
		Statements: statements,
		Labels:     labels,
	}
}

type MultiLabelsBlock struct {
	Statements IStatement
	Labels     util.DefaultList[ILabel]
}

func (b *MultiLabelsBlock) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) IsBreakStatement() bool                    { return false }
func (b *MultiLabelsBlock) IsContinueStatement() bool                 { return false }
func (b *MultiLabelsBlock) IsExpressionStatement() bool               { return false }
func (b *MultiLabelsBlock) IsForStatement() bool                      { return false }
func (b *MultiLabelsBlock) IsIfStatement() bool                       { return false }
func (b *MultiLabelsBlock) IsIfElseStatement() bool                   { return false }
func (b *MultiLabelsBlock) IsLabelStatement() bool                    { return false }
func (b *MultiLabelsBlock) IsLambdaExpressionStatement() bool         { return false }
func (b *MultiLabelsBlock) IsLocalVariableDeclarationStatement() bool { return false }
func (b *MultiLabelsBlock) IsMonitorEnterStatement() bool             { return false }
func (b *MultiLabelsBlock) IsMonitorExitStatement() bool              { return false }
func (b *MultiLabelsBlock) IsReturnStatement() bool                   { return false }
func (b *MultiLabelsBlock) IsReturnExpressionStatement() bool         { return false }
func (b *MultiLabelsBlock) IsStatements() bool                        { return false }
func (b *MultiLabelsBlock) IsSwitchStatement() bool                   { return false }
func (b *MultiLabelsBlock) IsSwitchStatementLabelBlock() bool         { return false }
func (b *MultiLabelsBlock) IsSwitchStatementMultiLabelsBlock() bool   { return true }
func (b *MultiLabelsBlock) IsThrowStatement() bool                    { return false }
func (b *MultiLabelsBlock) IsTryStatement() bool                      { return false }
func (b *MultiLabelsBlock) IsWhileStatement() bool                    { return false }

func (b *MultiLabelsBlock) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (b *MultiLabelsBlock) GetExpression() model.IExpression { return &model.NeNoExpression }
func (b *MultiLabelsBlock) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (b *MultiLabelsBlock) GetElseStatements() IStatement    { return &NoStmt }
func (b *MultiLabelsBlock) GetFinallyStatements() IStatement { return &NoStmt }
func (b *MultiLabelsBlock) GetStatements() IStatement        { return b.Statements }
func (b *MultiLabelsBlock) GetTryStatements() IStatement     { return &NoStmt }

func (b *MultiLabelsBlock) GetInit() model.IExpression   { return &model.NeNoExpression }
func (b *MultiLabelsBlock) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (b *MultiLabelsBlock) GetCatchClauses() []util.IList[*CatchClause] { return nil }
func (b *MultiLabelsBlock) GetLineNumber() int                          { return model.UnknownLineNumber }

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{label=%s}", b.Labels.ToSlice())
}
