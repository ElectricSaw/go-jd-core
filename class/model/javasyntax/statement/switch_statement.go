package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"fmt"
)

var DefaultLabe1 = NewDefaultLabel()

func NewSwitchStatement(condition intmod.IExpression, blocks []intmod.IBlock) intmod.ISwitchStatement {
	return &SwitchStatement{
		condition: condition,
		blocks:    blocks,
	}
}

type SwitchStatement struct {
	AbstractStatement

	condition intmod.IExpression
	blocks    []intmod.IBlock
}

func (s *SwitchStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *SwitchStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *SwitchStatement) List() []intmod.IStatement {
	ret := make([]intmod.IStatement, 0, len(s.blocks))
	for _, block := range s.blocks {
		ret = append(ret, block.(intmod.IStatement))
	}
	return ret
}

func (s *SwitchStatement) Blocks() []intmod.IBlock {
	return s.blocks
}

func (s *SwitchStatement) IsSwitchStatement() bool {
	return true
}

func (s *SwitchStatement) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatement(s)
}

// --- Label --- //

func NewDefaultLabel() intmod.IDefaultLabel {
	return &DefaultLabel{}
}

type DefaultLabel struct {
	AbstractStatement
}

func (l *DefaultLabel) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatementDefaultLabel(l)
}

func (l *DefaultLabel) String() string {
	return "DefaultLabel"
}

func NewExpressionLabel(expression intmod.IExpression) intmod.IExpressionLabel {
	return &ExpressionLabel{
		expression: expression,
	}
}

type ExpressionLabel struct {
	AbstractStatement

	expression intmod.IExpression
}

func (l *ExpressionLabel) Expression() intmod.IExpression {
	return l.expression
}

func (l *ExpressionLabel) SetExpression(expression intmod.IExpression) {
	l.expression = expression
}

func (l *ExpressionLabel) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatementExpressionLabel(l)
}

func (l *ExpressionLabel) String() string {
	return fmt.Sprintf("ExpressionLabel{%s}", l.expression)
}

// --- Block --- //

func NewBlock(statements intmod.IStatement) intmod.IBlock {
	return &Block{
		statements: statements,
	}
}

type Block struct {
	AbstractStatement

	statements intmod.IStatement
}

func (b *Block) Statements() intmod.IStatement {
	return b.statements
}

func NewLabelBlock(label intmod.ILabel, statements intmod.IStatement) intmod.ILabelBlock {
	return &LabelBlock{
		Block: Block{
			statements: statements,
		},
		label: label,
	}
}

type LabelBlock struct {
	Block

	label intmod.ILabel
}

func (b *LabelBlock) Label() intmod.ILabel {
	return b.label
}

func (b *LabelBlock) IsSwitchStatementLabelBlock() bool {
	return true
}

func (b *LabelBlock) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatementLabelBlock(b)
}

func (b *LabelBlock) String() string {
	return fmt.Sprintf("LabelBlock{label=%s}", b.label)
}

func NewMultiLabelsBlock(labels []intmod.ILabel, statements intmod.IStatement) intmod.IMultiLabelsBlock {
	return &MultiLabelsBlock{
		Block: Block{
			statements: statements,
		},
		labels: labels,
	}
}

type MultiLabelsBlock struct {
	Block

	labels []intmod.ILabel
}

func (b *MultiLabelsBlock) List() []intmod.IStatement {
	ret := make([]intmod.IStatement, 0, len(b.labels))
	for _, label := range b.labels {
		ret = append(ret, label)
	}
	return ret
}

func (b *MultiLabelsBlock) Labels() []intmod.ILabel {
	return b.labels
}

func (b *MultiLabelsBlock) IsSwitchStatementMultiLabelsBlock() bool {
	return true
}

func (b *MultiLabelsBlock) Accept(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{label=%s}", b.labels)
}
