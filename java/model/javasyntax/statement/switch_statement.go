package statement

import (
	"bitbucket.org/coontec/javaClass/java/model/javasyntax/expression"
	"fmt"
)

var DefaultLabel = NewDefaultLabel()

func NewSwitchStatement(condition expression.Expression, blocks []Block) *SwitchStatement {
	return &SwitchStatement{
		condition: condition,
		blocks:    blocks,
	}
}

type SwitchStatement struct {
	AbstractStatement

	condition expression.Expression
	blocks    []Block
}

func (s *SwitchStatement) GetCondition() expression.Expression {
	return s.condition
}

func (s *SwitchStatement) SetCondition(condition expression.Expression) {
	s.condition = condition
}

func (s *SwitchStatement) GetBlocks() []Block {
	return s.blocks
}

func (s *SwitchStatement) IsSwitchStatement() bool {
	return true
}

func (s *SwitchStatement) Accept(visitor StatementVisitor) {
	visitor.VisitSwitchStatement(s)
}

// --- Label --- //

type ILabel interface {
	ignoreLabel()
}

func NewDefaultLabel() *DefaultLabe1 {
	return &DefaultLabe1{}
}

type DefaultLabe1 struct {
	AbstractStatement
}

func (l *DefaultLabe1) Accept(visitor StatementVisitor) {
	visitor.VisitSwitchStatementDefaultLabel(l)
}

func (l *DefaultLabe1) String() string {
	return "DefaultLabel"
}

func (l *DefaultLabe1) ignoreLabel() {}

func NewExpressionLabel(expression expression.Expression) *ExpressionLabel {
	return &ExpressionLabel{
		expression: expression,
	}
}

type ExpressionLabel struct {
	AbstractStatement

	expression expression.Expression
}

func (l *ExpressionLabel) GetExpression() expression.Expression {
	return l.expression
}

func (l *ExpressionLabel) SetExpression(expression expression.Expression) {
	l.expression = expression
}

func (l *ExpressionLabel) Accept(visitor StatementVisitor) {
	visitor.VisitSwitchStatementExpressionLabel(l)
}

func (l *ExpressionLabel) String() string {
	return fmt.Sprintf("ExpressionLabel{%s}", l.expression)
}

func (l *ExpressionLabel) ignoreLabel() {}

// --- Block --- //

func NewBlock(statements Statement) *Block {
	return &Block{
		statements: statements,
	}
}

type Block struct {
	AbstractStatement

	statements Statement
}

func (b *Block) GetStatements() Statement {
	return b.statements
}

func NewLabelBlock(label ILabel, statements Statement) *LabelBlock {
	return &LabelBlock{
		Block: Block{
			statements: statements,
		},
		label: label,
	}
}

type LabelBlock struct {
	Block

	label ILabel
}

func (b *LabelBlock) GetLabel() ILabel {
	return b.label
}

func (b *LabelBlock) IsSwitchStatementLabelBlock() bool {
	return true
}

func (b *LabelBlock) Accept(visitor StatementVisitor) {
	visitor.VisitSwitchStatementLabelBlock(b)
}

func (b *LabelBlock) String() string {
	return fmt.Sprintf("LabelBlock{label=%s}", b.label)
}

func NewMultiLabelsBlock(labels []ILabel, statements Statement) *MultiLabelsBlock {
	return &MultiLabelsBlock{
		Block: Block{
			statements: statements,
		},
		labels: labels,
	}
}

type MultiLabelsBlock struct {
	Block

	labels []ILabel
}

func (b *MultiLabelsBlock) GetLabels() []ILabel {
	return b.labels
}

func (b *MultiLabelsBlock) IsSwitchStatementMultiLabelsBlock() bool {
	return true
}

func (b *MultiLabelsBlock) Accept(visitor StatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{label=%s}", b.labels)
}
