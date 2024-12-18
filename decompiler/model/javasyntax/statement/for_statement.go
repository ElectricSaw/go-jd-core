package statement

import (
	"fmt"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
)

func NewForStatementWithDeclaration(declaration intmod.ILocalVariableDeclaration,
	condition intmod.IExpression, update intmod.IExpression,
	statements intmod.IStatement) intmod.IForStatement {
	return &ForStatement{
		declaration: declaration,
		condition:   condition,
		update:      update,
		statements:  statements,
	}
}

func NewForStatementWithInit(init intmod.IExpression, condition intmod.IExpression,
	update intmod.IExpression, statements intmod.IStatement) intmod.IForStatement {
	return &ForStatement{
		init:       init,
		condition:  condition,
		update:     update,
		statements: statements,
	}
}

type ForStatement struct {
	AbstractStatement

	declaration intmod.ILocalVariableDeclaration
	init        intmod.IExpression
	condition   intmod.IExpression
	update      intmod.IExpression
	statements  intmod.IStatement
}

func (s *ForStatement) Declaration() intmod.ILocalVariableDeclaration {
	return s.declaration
}

func (s *ForStatement) SetDeclaration(declaration intmod.ILocalVariableDeclaration) {
	s.declaration = declaration
}

func (s *ForStatement) Init() intmod.IExpression {
	return s.init
}

func (s *ForStatement) SetInit(init intmod.IExpression) {
	s.init = init
}

func (s *ForStatement) Condition() intmod.IExpression {
	return s.condition
}

func (s *ForStatement) SetCondition(condition intmod.IExpression) {
	s.condition = condition
}

func (s *ForStatement) Update() intmod.IExpression {
	return s.update
}

func (s *ForStatement) SetUpdate(update intmod.IExpression) {
	s.update = update
}

func (s *ForStatement) Statements() intmod.IStatement {
	return s.statements
}

func (s *ForStatement) AcceptStatement(visitor intmod.IStatementVisitor) {
	visitor.VisitForStatement(s)
}

func (s *ForStatement) String() string {
	return fmt.Sprintf("ForStatement{%s or %s; %s; %s", s.declaration, s.init, s.condition, s.update)
}
