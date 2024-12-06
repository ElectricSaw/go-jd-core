package statement

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewAbstractStatement() *AbstractStatement {
	s := &AbstractStatement{}
	s.SetValue(s)
	return s
}

type AbstractStatement struct {
	util.DefaultBase[intmod.IStatement]
}

func (s *AbstractStatement) AcceptStatement(visitor intmod.IStatementVisitor) {}

func (s *AbstractStatement) IsBreakStatement() bool                    { return false }
func (s *AbstractStatement) IsContinueStatement() bool                 { return false }
func (s *AbstractStatement) IsExpressionStatement() bool               { return false }
func (s *AbstractStatement) IsForStatement() bool                      { return false }
func (s *AbstractStatement) IsIfStatement() bool                       { return false }
func (s *AbstractStatement) IsIfElseStatement() bool                   { return false }
func (s *AbstractStatement) IsLabelStatement() bool                    { return false }
func (s *AbstractStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *AbstractStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *AbstractStatement) IsMonitorEnterStatement() bool             { return false }
func (s *AbstractStatement) IsMonitorExitStatement() bool              { return false }
func (s *AbstractStatement) IsReturnStatement() bool                   { return false }
func (s *AbstractStatement) IsReturnExpressionStatement() bool         { return false }
func (s *AbstractStatement) IsStatements() bool                        { return false }
func (s *AbstractStatement) IsSwitchStatement() bool                   { return false }
func (s *AbstractStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *AbstractStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *AbstractStatement) IsThrowStatement() bool                    { return false }
func (s *AbstractStatement) IsTryStatement() bool                      { return false }
func (s *AbstractStatement) IsWhileStatement() bool                    { return false }

func (s *AbstractStatement) Condition() intmod.IExpression        { return expression.NeNoExpression }
func (s *AbstractStatement) Expression() intmod.IExpression       { return expression.NeNoExpression }
func (s *AbstractStatement) Monitor() intmod.IExpression          { return expression.NeNoExpression }
func (s *AbstractStatement) ElseStatements() intmod.IStatement    { return NoStmt.(intmod.IStatement) }
func (s *AbstractStatement) FinallyStatements() intmod.IStatement { return NoStmt.(intmod.IStatement) }
func (s *AbstractStatement) Statements() intmod.IStatement        { return NoStmt.(intmod.IStatement) }
func (s *AbstractStatement) TryStatements() intmod.IStatement     { return NoStmt.(intmod.IStatement) }

func (s *AbstractStatement) Init() intmod.IExpression   { return expression.NeNoExpression }
func (s *AbstractStatement) Update() intmod.IExpression { return expression.NeNoExpression }

func (s *AbstractStatement) CatchClauses() []intmod.ICatchClause { return nil }

func (s *AbstractStatement) LineNumber() int { return intmod.UnknownLineNumber }
