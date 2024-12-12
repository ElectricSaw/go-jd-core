package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

var DefaultLabe1 = NewDefaultLabel()

func NewSwitchStatement(condition intmod.IExpression, blocks util.IList[intmod.IBlock]) intmod.ISwitchStatement {
	return &SwitchStatement{
		condition: condition,
		blocks:    blocks,
	}
}

type SwitchStatement struct {
	AbstractStatement

	condition intmod.IExpression
	blocks    util.IList[intmod.IBlock]
}

func (s *SwitchStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *SwitchStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *SwitchStatement) List() []intmod.IStatement {
	ret := make([]intmod.IStatement, 0, s.blocks.Size())
	for _, block := range s.blocks.ToSlice() {
		ret = append(ret, block.(intmod.IStatement))
	}
	return ret
}

func (s *SwitchStatement) Blocks() []intmod.IBlock {
	return s.blocks.ToSlice()
}

func (s *SwitchStatement) IsSwitchStatement() bool {
	return true
}

func (s *SwitchStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatement(s)
}

// --- label --- //

func NewDefaultLabel() intmod.IDefaultLabel {
	return &DefaultLabel{}
}

type DefaultLabel struct {
	AbstractStatement
}

func (l *DefaultLabel) AcceptStatement(visitor intmod.IStatementVisitor) {
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

func (l *ExpressionLabel) AcceptStatement(visitor intmod.IStatementVisitor) {
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

func (b *LabelBlock) AcceptStatement(visitor intmod.IStatementVisitor) {
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

func (b *MultiLabelsBlock) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{label=%s}", b.labels)
}
