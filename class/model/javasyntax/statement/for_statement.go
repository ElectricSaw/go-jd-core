package statement

import (
	intsyn "bitbucket.org/coontec/javaClass/class/interfaces/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	"fmt"
)

func NewForStatementWithDeclaration(declaration intsyn.ILocalVariableDeclaration,
	condition intsyn.IExpression, update intsyn.IExpression,
	statements intsyn.IStatement) intsyn.IForStatement {
	return &ForStatement{
		declaration: declaration,
		condition:   condition,
		update:      update,
		statements:  statements,
	}
}

func NewForStatementWithInit(init intsyn.IExpression, condition intsyn.IExpression,
	update intsyn.IExpression, statements intsyn.IStatement) intsyn.IForStatement {
	return &ForStatement{
		init:       init,
		condition:  condition,
		update:     update,
		statements: statements,
	}
}

type ForStatement struct {
	AbstractStatement

	declaration intsyn.ILocalVariableDeclaration
	init        intsyn.IExpression
	condition   intsyn.IExpression
	update      intsyn.IExpression
	statements  intsyn.IStatement
}

func (s *ForStatement) Declaration() intsyn.ILocalVariableDeclaration {
	return s.declaration
}

func (s *ForStatement) SetDeclaration(declaration intsyn.ILocalVariableDeclaration) {
	s.declaration = declaration
}

func (s *ForStatement) Init() intsyn.IExpression {
	return s.init
}

func (s *ForStatement) SetInit(init intsyn.IExpression) {
	s.init = init
}

func (s *ForStatement) Condition() intsyn.IExpression {
	return s.condition
}

func (s *ForStatement) SetCondition(condition intsyn.IExpression) {
	s.condition = condition
}

func (s *ForStatement) Update() intsyn.IExpression {
	return s.update
}

func (s *ForStatement) SetUpdate(update intsyn.IExpression) {
	s.update = update
}

func (s *ForStatement) Statements() intsyn.IStatement {
	return s.statements
}

func (s *ForStatement) Accept(visitor intsyn.IStatementVisitor) {
	visitor.VisitForStatement(s)
}

func (s *ForStatement) String() string {
	return fmt.Sprintf("ForStatement{%s or %s; %s; %s", s.declaration, s.init, s.condition, s.update)
}
