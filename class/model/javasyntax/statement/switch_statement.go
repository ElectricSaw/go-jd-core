package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"fmt"
)

var DefaultLabe1 = NewDefaultLabel()

func NewSwitchStatement(condition intsyn.IExpression, blocks []intsyn.IBlock) intsyn.ISwitchStatement {
	return &SwitchStatement{
		condition: condition,
		blocks:    blocks,
	}
}

type SwitchStatement struct {
	AbstractStatement

	condition intsyn.IExpression
	blocks    []intsyn.IBlock
}

func (s *SwitchStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *SwitchStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *SwitchStatement) List() []intsyn.IStatement {
	ret := make([]intsyn.IStatement, 0, len(s.blocks))
	for _, block := range s.blocks {
		ret = append(ret, block.(intsyn.IStatement))
	}
	return ret
}

func (s *SwitchStatement) Blocks() []intsyn.IBlock {
	return s.blocks
}

func (s *SwitchStatement) IsSwitchStatement() bool {
	return true
}

func (s *SwitchStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSwitchStatement(s)
}

// --- Label --- //

func NewDefaultLabel() intsyn.IDefaultLabel {
	return &DefaultLabel{}
}

type DefaultLabel struct {
	AbstractStatement
}

func (l *DefaultLabel) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSwitchStatementDefaultLabel(l)
}

func (l *DefaultLabel) String() string {
	return "DefaultLabel"
}

func NewExpressionLabel(expression intsyn.IExpression) intsyn.IExpressionLabel {
	return &ExpressionLabel{
		expression: expression,
	}
}

type ExpressionLabel struct {
	AbstractStatement

	expression intsyn.IExpression
}

func (l *ExpressionLabel) Expression() intsyn.IExpression {
	return l.expression
}

func (l *ExpressionLabel) SetExpression(expression intsyn.IExpression) {
	l.expression = expression
}

func (l *ExpressionLabel) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSwitchStatementExpressionLabel(l)
}

func (l *ExpressionLabel) String() string {
	return fmt.Sprintf("ExpressionLabel{%s}", l.expression)
}

// --- Block --- //

func NewBlock(statements intsyn.IStatement) intsyn.IBlock {
	return &Block{
		statements: statements,
	}
}

type Block struct {
	AbstractStatement

	statements intsyn.IStatement
}

func (b *Block) Statements() intsyn.IStatement {
	return b.statements
}

func NewLabelBlock(label intsyn.ILabel, statements intsyn.IStatement) intsyn.ILabelBlock {
	return &LabelBlock{
		Block: Block{
			statements: statements,
		},
		label: label,
	}
}

type LabelBlock struct {
	Block

	label intsyn.ILabel
}

func (b *LabelBlock) Label() intsyn.ILabel {
	return b.label
}

func (b *LabelBlock) IsSwitchStatementLabelBlock() bool {
	return true
}

func (b *LabelBlock) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSwitchStatementLabelBlock(b)
}

func (b *LabelBlock) String() string {
	return fmt.Sprintf("LabelBlock{label=%s}", b.label)
}

func NewMultiLabelsBlock(labels []intsyn.ILabel, statements intsyn.IStatement) intsyn.IMultiLabelsBlock {
	return &MultiLabelsBlock{
		Block: Block{
			statements: statements,
		},
		labels: labels,
	}
}

type MultiLabelsBlock struct {
	Block

	labels []intsyn.ILabel
}

func (b *MultiLabelsBlock) List() []intsyn.IStatement {
	ret := make([]intsyn.IStatement, 0, len(b.labels))
	for _, label := range b.labels {
		ret = append(ret, label)
	}
	return ret
}

func (b *MultiLabelsBlock) Labels() []intsyn.ILabel {
	return b.labels
}

func (b *MultiLabelsBlock) IsSwitchStatementMultiLabelsBlock() bool {
	return true
}

func (b *MultiLabelsBlock) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{label=%s}", b.labels)
}
