package statement

import (
	intsyn "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type AbstractStatement struct {
	util.DefaultBase[intsyn.IStatement]
}

func (s *AbstractStatement) Accept(visitor intsyn.IStatementVisitor) {}

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

func (s *AbstractStatement) Condition() intsyn.IExpression        { return expression.NeNoExpression }
func (s *AbstractStatement) Expression() intsyn.IExpression       { return expression.NeNoExpression }
func (s *AbstractStatement) Monitor() intsyn.IExpression          { return expression.NeNoExpression }
func (s *AbstractStatement) ElseStatements() intsyn.IStatement    { return NoStmt.(intsyn.IStatement) }
func (s *AbstractStatement) FinallyStatements() intsyn.IStatement { return NoStmt.(intsyn.IStatement) }
func (s *AbstractStatement) Statements() intsyn.IStatement        { return NoStmt.(intsyn.IStatement) }
func (s *AbstractStatement) TryStatements() intsyn.IStatement     { return NoStmt.(intsyn.IStatement) }

func (s *AbstractStatement) Init() intsyn.IExpression   { return expression.NeNoExpression }
func (s *AbstractStatement) Update() intsyn.IExpression { return expression.NeNoExpression }

func (s *AbstractStatement) CatchClauses() []intsyn.ICatchClause { return nil }

func (s *AbstractStatement) LineNumber() int { return intsyn.UnknownLineNumber }
