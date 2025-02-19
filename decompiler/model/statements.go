package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

/////////////////////////////////////////////////////////////////////////
//  Global Variable
/////////////////////////////////////////////////////////////////////////

var (
	Break        = NewBreakStatement("")
	Continue     = NewContinueStatement("")
	NoStmt       = NewNoStatement()
	Return       = NewReturnStatement()
	DefaultLabe1 = NewDefaultLabel()
)

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

func NewAssertStatement(condition IExpression, message IExpression) AssertStatement {
	s := AssertStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Message:     message,
	}
	s.SetValue(&s)
	return s
}

func NewBreakStatement(label string) BreakStatement {
	s := BreakStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Label:       label,
	}
	s.SetValue(&s)
	return s
}

func NewByteCodeStatement(text string) ByteCodeStatement {
	s := ByteCodeStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Text:        text,
	}
	s.SetValue(&s)
	return s
}

func NewCommentStatement(text string) CommentStatement {
	s := CommentStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Text:        text,
	}
	s.SetValue(&s)
	return s
}

func NewContinueStatement(label string) ContinueStatement {
	s := ContinueStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Label:       label,
	}
	s.SetValue(&s)
	return s
}

func NewDoWhileStatement(condition IExpression, statements IStatement) DoWhileStatement {
	s := DoWhileStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewExpressionStatement(expression IExpression) ExpressionStatement {
	s := ExpressionStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewForEachStatement(typ IType, name string, expression IExpression, statement IStatement) ForEachStatement {
	s := ForEachStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Type:        typ,
		Name:        name,
		Expression:  expression,
		Statement:   statement,
	}
	s.SetValue(&s)
	return s
}

func NewForStatementWithDeclaration(declaration ILocalVariableDeclaration,
	condition, update IExpression, statements IStatement) ForStatement {
	s := ForStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Declaration: declaration,
		Condition:   condition,
		Update:      update,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewForStatementWithInit(init, condition, update IExpression, statements IStatement) ForStatement {
	s := ForStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Init:        init,
		Condition:   condition,
		Update:      update,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewIfElseStatement(condition IExpression, statements, elseStatements IStatement) IfElseStatement {
	s := IfElseStatement{
		DefaultBase:    *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:      condition,
		IfStatements:   statements,
		ElseStatements: elseStatements,
	}
	s.SetValue(&s)
	return s
}

func NewIfStatement(condition IExpression, statements IStatement) IfStatement {
	s := IfStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewLabelStatement(label string, statement IStatement) LabelStatement {
	s := LabelStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Label:       label,
		Statement:   statement,
	}
	s.SetValue(&s)
	return s
}

func NewLambdaExpressionStatement(expression IExpression) LambdaExpressionStatement {
	s := LambdaExpressionStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewLocalVariableDeclarationStatement(typ IType,
	localVariableDeclarators ILocalVariableDeclarator) LocalVariableDeclarationStatement {
	s := LocalVariableDeclarationStatement{
		DefaultBase:              *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Type:                     typ,
		LocalVariableDeclarators: localVariableDeclarators,
	}
	s.SetValue(&s)
	return s
}

func NewNoStatement() NoStatement {
	s := NoStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
	}
	s.SetValue(&s)
	return s
}

func NewReturnExpressionStatement(expression IExpression) ReturnExpressionStatement {
	return NewReturnExpressionStatementWithAll(expression.GetLineNumber(), expression)
}

func NewReturnExpressionStatementWithAll(lineNumber int, expression IExpression) ReturnExpressionStatement {
	s := ReturnExpressionStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		LineNumber:  lineNumber,
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewReturnStatement() ReturnStatement {
	s := ReturnStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
	}
	s.SetValue(&s)
	return s
}

func NewStatements() Statements {
	s := Statements{
		DefaultList: *util.NewDefaultList[IStatement]().(*util.DefaultList[IStatement]),
	}
	return s
}

func NewStatementsWithList(list util.IList[IStatement]) Statements {
	return NewStatementsWithElements(list.ToSlice()...)
}

func NewStatementsWithElements(slice ...IStatement) Statements {
	s := Statements{
		DefaultList: *util.NewDefaultListWithElements[IStatement](slice...).(*util.DefaultList[IStatement]),
	}
	return s
}

func NewSwitchStatement(condition IExpression, blocks util.DefaultList[*Block]) SwitchStatement {
	s := SwitchStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Blocks:      blocks,
	}
	s.SetValue(&s)
	return s
}

func NewDefaultLabel() DefaultLabel {
	s := DefaultLabel{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
	}
	s.SetValue(&s)
	return s
}

func NewExpressionLabel(expression IExpression) ExpressionLabel {
	s := ExpressionLabel{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewBlock(statements IStatement) Block {
	s := Block{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewLabelBlock(label ILabel, statements IStatement) LabelBlock {
	s := LabelBlock{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Statements:  statements,
		Label:       label,
	}
	s.SetValue(&s)
	return s
}

func NewMultiLabelsBlock(labels util.DefaultList[ILabel], statements IStatement) MultiLabelsBlock {
	s := MultiLabelsBlock{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Statements:  statements,
		Labels:      labels,
	}
	s.SetValue(&s)
	return s
}

func NewSynchronizedStatement(monitor IExpression, statements IStatement) SynchronizedStatement {
	s := SynchronizedStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Monitor:     monitor,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewThrowStatement(expression IExpression) ThrowStatement {
	s := ThrowStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewTryStatement(tryStatements IStatement, catchClauses util.DefaultList[*CatchClause], finallyStatement IStatement) TryStatement {
	return NewTryStatementWithAll(*util.NewArrayList[*Resource]().(*util.DefaultList[*Resource]),
		tryStatements, catchClauses, finallyStatement)
}

func NewTryStatementWithAll(resource util.DefaultList[*Resource], tryStatements IStatement,
	catchClauses util.DefaultList[*CatchClause], finallyStatement IStatement) TryStatement {
	s := TryStatement{
		DefaultBase:       *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Resources:         resource,
		TryStatements:     tryStatements,
		CatchClause:       catchClauses,
		FinallyStatements: finallyStatement,
	}
	s.SetValue(&s)
	return s
}

func NewResource(typ *ObjectType, name string, expression IExpression) Resource {
	s := Resource{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Type:        typ,
		Name:        name,
		Expression:  expression,
	}
	s.SetValue(&s)
	return s
}

func NewCatchClause(lineNumber int, typ *ObjectType, name string, statements IStatement) CatchClause {
	s := CatchClause{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		LineNumber:  lineNumber,
		Type:        typ,
		OtherType:   *util.NewArrayList[*ObjectType]().(*util.DefaultList[*ObjectType]),
		Name:        name,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

func NewTypeDeclarationStatement(typeDeclaration TypeDeclaration) TypeDeclarationStatement {
	s := TypeDeclarationStatement{
		DefaultBase:     *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		TypeDeclaration: typeDeclaration,
	}
	s.SetValue(&s)
	return s
}

func NewWhileStatement(condition IExpression, statements IStatement) WhileStatement {
	s := WhileStatement{
		DefaultBase: *util.NewDefaultBase[IStatement]().(*util.DefaultBase[IStatement]),
		Condition:   condition,
		Statements:  statements,
	}
	s.SetValue(&s)
	return s
}

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type IStatement interface {
	AcceptStatement(visitor IStatementVisitor)

	IsBreakStatement() bool
	IsContinueStatement() bool
	IsExpressionStatement() bool
	IsForStatement() bool
	IsIfStatement() bool
	IsIfElseStatement() bool
	IsLabelStatement() bool
	IsLambdaExpressionStatement() bool
	IsLocalVariableDeclarationStatement() bool
	IsMonitorEnterStatement() bool
	IsMonitorExitStatement() bool
	IsReturnStatement() bool
	IsReturnExpressionStatement() bool
	IsStatements() bool
	IsSwitchStatement() bool
	IsSwitchStatementLabelBlock() bool
	IsSwitchStatementMultiLabelsBlock() bool
	IsThrowStatement() bool
	IsTryStatement() bool
	IsWhileStatement() bool

	GetCondition() IExpression
	GetExpression() IExpression
	GetMonitor() IExpression

	GetElseStatements() IStatement
	GetFinallyStatements() IStatement
	GetStatements() IStatement
	GetTryStatements() IStatement

	GetInit() IExpression
	GetUpdate() IExpression

	GetCatchClauses() util.IList[*CatchClause]
	GetLineNumber() int

	String() string
}

type IStatementVisitor interface {
	VisitAssertStatement(statement *AssertStatement)
	VisitBreakStatement(statement *BreakStatement)
	VisitByteCodeStatement(statement *ByteCodeStatement)
	VisitCommentStatement(statement *CommentStatement)
	VisitContinueStatement(statement *ContinueStatement)
	VisitDoWhileStatement(statement *DoWhileStatement)
	VisitExpressionStatement(statement *ExpressionStatement)
	VisitForEachStatement(statement *ForEachStatement)
	VisitForStatement(statement *ForStatement)
	VisitIfStatement(statement *IfStatement)
	VisitIfElseStatement(statement *IfElseStatement)
	VisitLabelStatement(statement *LabelStatement)
	VisitLambdaExpressionStatement(statement *LambdaExpressionStatement)
	VisitLocalVariableDeclarationStatement(statement *LocalVariableDeclarationStatement)
	VisitNoStatement(statement *NoStatement)
	VisitReturnExpressionStatement(statement *ReturnExpressionStatement)
	VisitReturnStatement(statement *ReturnStatement)
	VisitStatements(statement *Statements)
	VisitSwitchStatement(statement *SwitchStatement)
	VisitSwitchStatementDefaultLabel(statement *DefaultLabel)
	VisitSwitchStatementExpressionLabel(statement *ExpressionLabel)
	VisitSwitchStatementLabelBlock(statement *LabelBlock)
	VisitSwitchStatementMultiLabelsBlock(statement *MultiLabelsBlock)
	VisitSynchronizedStatement(statement *SynchronizedStatement)
	VisitThrowStatement(statement *ThrowStatement)
	VisitTryStatement(statement *TryStatement)
	VisitTryStatementResource(statement *Resource)
	VisitTryStatementCatchClause(statement *CatchClause)
	VisitTypeDeclarationStatement(statement *TypeDeclarationStatement)
	VisitWhileStatement(statement *WhileStatement)
}

type ILabel interface {
	IStatement

	IsLabel() bool
}

type IBlock interface {
	IStatement

	IsBlock() bool
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

type AssertStatement struct {
	util.DefaultBase[IStatement]

	Condition IExpression
	Message   IExpression
}

func (s *AssertStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitAssertStatement(s)
}

func (s *AssertStatement) IsBreakStatement() bool                    { return false }
func (s *AssertStatement) IsContinueStatement() bool                 { return false }
func (s *AssertStatement) IsExpressionStatement() bool               { return false }
func (s *AssertStatement) IsForStatement() bool                      { return false }
func (s *AssertStatement) IsIfStatement() bool                       { return false }
func (s *AssertStatement) IsIfElseStatement() bool                   { return false }
func (s *AssertStatement) IsLabelStatement() bool                    { return false }
func (s *AssertStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *AssertStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *AssertStatement) IsMonitorEnterStatement() bool             { return false }
func (s *AssertStatement) IsMonitorExitStatement() bool              { return false }
func (s *AssertStatement) IsReturnStatement() bool                   { return false }
func (s *AssertStatement) IsReturnExpressionStatement() bool         { return false }
func (s *AssertStatement) IsStatements() bool                        { return false }
func (s *AssertStatement) IsSwitchStatement() bool                   { return false }
func (s *AssertStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *AssertStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *AssertStatement) IsThrowStatement() bool                    { return false }
func (s *AssertStatement) IsTryStatement() bool                      { return false }
func (s *AssertStatement) IsWhileStatement() bool                    { return false }

func (s *AssertStatement) GetCondition() IExpression  { return s.Condition }
func (s *AssertStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *AssertStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *AssertStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *AssertStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *AssertStatement) GetStatements() IStatement        { return &NoStmt }
func (s *AssertStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *AssertStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *AssertStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *AssertStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *AssertStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *AssertStatement) String() string {
	return fmt.Sprintf("AssertStatement{ condition=%s, message=%s }", s.Condition, s.Message)
}

type BreakStatement struct {
	util.DefaultBase[IStatement]

	Label string
}

func (s *BreakStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitBreakStatement(s)
}

func (s *BreakStatement) IsBreakStatement() bool                    { return true }
func (s *BreakStatement) IsContinueStatement() bool                 { return false }
func (s *BreakStatement) IsExpressionStatement() bool               { return false }
func (s *BreakStatement) IsForStatement() bool                      { return false }
func (s *BreakStatement) IsIfStatement() bool                       { return false }
func (s *BreakStatement) IsIfElseStatement() bool                   { return false }
func (s *BreakStatement) IsLabelStatement() bool                    { return false }
func (s *BreakStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *BreakStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *BreakStatement) IsMonitorEnterStatement() bool             { return false }
func (s *BreakStatement) IsMonitorExitStatement() bool              { return false }
func (s *BreakStatement) IsReturnStatement() bool                   { return false }
func (s *BreakStatement) IsReturnExpressionStatement() bool         { return false }
func (s *BreakStatement) IsStatements() bool                        { return false }
func (s *BreakStatement) IsSwitchStatement() bool                   { return false }
func (s *BreakStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *BreakStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *BreakStatement) IsThrowStatement() bool                    { return false }
func (s *BreakStatement) IsTryStatement() bool                      { return false }
func (s *BreakStatement) IsWhileStatement() bool                    { return false }

func (s *BreakStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *BreakStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *BreakStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *BreakStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *BreakStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *BreakStatement) GetStatements() IStatement        { return &NoStmt }
func (s *BreakStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *BreakStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *BreakStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *BreakStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *BreakStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *BreakStatement) String() string {
	return fmt.Sprintf("BreakStatement{ label=%s }", s.Label)
}

type ByteCodeStatement struct {
	util.DefaultBase[IStatement]

	Text string
}

func (s *ByteCodeStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitByteCodeStatement(s)
}

func (s *ByteCodeStatement) IsBreakStatement() bool                    { return false }
func (s *ByteCodeStatement) IsContinueStatement() bool                 { return false }
func (s *ByteCodeStatement) IsExpressionStatement() bool               { return false }
func (s *ByteCodeStatement) IsForStatement() bool                      { return false }
func (s *ByteCodeStatement) IsIfStatement() bool                       { return false }
func (s *ByteCodeStatement) IsIfElseStatement() bool                   { return false }
func (s *ByteCodeStatement) IsLabelStatement() bool                    { return false }
func (s *ByteCodeStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ByteCodeStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ByteCodeStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ByteCodeStatement) IsMonitorExitStatement() bool              { return false }
func (s *ByteCodeStatement) IsReturnStatement() bool                   { return false }
func (s *ByteCodeStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ByteCodeStatement) IsStatements() bool                        { return false }
func (s *ByteCodeStatement) IsSwitchStatement() bool                   { return false }
func (s *ByteCodeStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ByteCodeStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ByteCodeStatement) IsThrowStatement() bool                    { return false }
func (s *ByteCodeStatement) IsTryStatement() bool                      { return false }
func (s *ByteCodeStatement) IsWhileStatement() bool                    { return false }

func (s *ByteCodeStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ByteCodeStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *ByteCodeStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ByteCodeStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ByteCodeStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ByteCodeStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ByteCodeStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ByteCodeStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ByteCodeStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ByteCodeStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ByteCodeStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ByteCodeStatement) String() string {
	return fmt.Sprintf("ByteCodeStatement{ text=%s }", s.Text)
}

type CommentStatement struct {
	util.DefaultBase[IStatement]

	Text string
}

func (s *CommentStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitCommentStatement(s)
}

func (s *CommentStatement) IsBreakStatement() bool                    { return false }
func (s *CommentStatement) IsContinueStatement() bool                 { return true }
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

func (s *CommentStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *CommentStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *CommentStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *CommentStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *CommentStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *CommentStatement) GetStatements() IStatement        { return &NoStmt }
func (s *CommentStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *CommentStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *CommentStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *CommentStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *CommentStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *CommentStatement) String() string {
	return fmt.Sprintf("CommentStatement{ text=%s }", s.Text)
}

type ContinueStatement struct {
	util.DefaultBase[IStatement]

	Label string
}

func (s *ContinueStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitContinueStatement(s)
}

func (s *ContinueStatement) IsBreakStatement() bool                    { return false }
func (s *ContinueStatement) IsContinueStatement() bool                 { return true }
func (s *ContinueStatement) IsExpressionStatement() bool               { return false }
func (s *ContinueStatement) IsForStatement() bool                      { return false }
func (s *ContinueStatement) IsIfStatement() bool                       { return false }
func (s *ContinueStatement) IsIfElseStatement() bool                   { return false }
func (s *ContinueStatement) IsLabelStatement() bool                    { return false }
func (s *ContinueStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ContinueStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ContinueStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ContinueStatement) IsMonitorExitStatement() bool              { return false }
func (s *ContinueStatement) IsReturnStatement() bool                   { return false }
func (s *ContinueStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ContinueStatement) IsStatements() bool                        { return false }
func (s *ContinueStatement) IsSwitchStatement() bool                   { return false }
func (s *ContinueStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ContinueStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ContinueStatement) IsThrowStatement() bool                    { return false }
func (s *ContinueStatement) IsTryStatement() bool                      { return false }
func (s *ContinueStatement) IsWhileStatement() bool                    { return false }

func (s *ContinueStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ContinueStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *ContinueStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ContinueStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ContinueStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ContinueStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ContinueStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ContinueStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ContinueStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ContinueStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ContinueStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ContinueStatement) String() string {
	return fmt.Sprintf("ContinueStatement{ label=%s }", s.Label)
}

type DoWhileStatement struct {
	util.DefaultBase[IStatement]

	Condition  IExpression
	Statements IStatement
}

func (s *DoWhileStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitDoWhileStatement(s)
}

func (s *DoWhileStatement) IsBreakStatement() bool                    { return false }
func (s *DoWhileStatement) IsContinueStatement() bool                 { return false }
func (s *DoWhileStatement) IsExpressionStatement() bool               { return false }
func (s *DoWhileStatement) IsForStatement() bool                      { return false }
func (s *DoWhileStatement) IsIfStatement() bool                       { return false }
func (s *DoWhileStatement) IsIfElseStatement() bool                   { return false }
func (s *DoWhileStatement) IsLabelStatement() bool                    { return false }
func (s *DoWhileStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *DoWhileStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *DoWhileStatement) IsMonitorEnterStatement() bool             { return false }
func (s *DoWhileStatement) IsMonitorExitStatement() bool              { return false }
func (s *DoWhileStatement) IsReturnStatement() bool                   { return false }
func (s *DoWhileStatement) IsReturnExpressionStatement() bool         { return false }
func (s *DoWhileStatement) IsStatements() bool                        { return false }
func (s *DoWhileStatement) IsSwitchStatement() bool                   { return false }
func (s *DoWhileStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *DoWhileStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *DoWhileStatement) IsThrowStatement() bool                    { return false }
func (s *DoWhileStatement) IsTryStatement() bool                      { return false }
func (s *DoWhileStatement) IsWhileStatement() bool                    { return false }

func (s *DoWhileStatement) GetCondition() IExpression  { return s.Condition }
func (s *DoWhileStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *DoWhileStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *DoWhileStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *DoWhileStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *DoWhileStatement) GetStatements() IStatement        { return s.Statements }
func (s *DoWhileStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *DoWhileStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *DoWhileStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *DoWhileStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *DoWhileStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *DoWhileStatement) String() string {
	return fmt.Sprintf("DoWhileStatement{ condition=%s, statements=%s }", s.Condition, s.Statements)
}

type ExpressionStatement struct {
	util.DefaultBase[IStatement]

	Expression IExpression
}

func (s *ExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitExpressionStatement(s)
}

func (s *ExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *ExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *ExpressionStatement) IsExpressionStatement() bool               { return true }
func (s *ExpressionStatement) IsForStatement() bool                      { return false }
func (s *ExpressionStatement) IsIfStatement() bool                       { return false }
func (s *ExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *ExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *ExpressionStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *ExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *ExpressionStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ExpressionStatement) IsStatements() bool                        { return false }
func (s *ExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *ExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *ExpressionStatement) IsTryStatement() bool                      { return false }
func (s *ExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *ExpressionStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ExpressionStatement) GetExpression() IExpression { return s.Expression }
func (s *ExpressionStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ExpressionStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ExpressionStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ExpressionStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ExpressionStatement) String() string {
	return fmt.Sprintf("ExpressionStatement{ expression=%s }", s.Expression)
}

type ForEachStatement struct {
	util.DefaultBase[IStatement]

	Type       IType
	Name       string
	Expression IExpression
	Statement  IStatement
}

func (s *ForEachStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitForEachStatement(s)
}

func (s *ForEachStatement) IsBreakStatement() bool                    { return false }
func (s *ForEachStatement) IsContinueStatement() bool                 { return false }
func (s *ForEachStatement) IsExpressionStatement() bool               { return false }
func (s *ForEachStatement) IsForStatement() bool                      { return false }
func (s *ForEachStatement) IsIfStatement() bool                       { return false }
func (s *ForEachStatement) IsIfElseStatement() bool                   { return false }
func (s *ForEachStatement) IsLabelStatement() bool                    { return false }
func (s *ForEachStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ForEachStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ForEachStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ForEachStatement) IsMonitorExitStatement() bool              { return false }
func (s *ForEachStatement) IsReturnStatement() bool                   { return false }
func (s *ForEachStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ForEachStatement) IsStatements() bool                        { return false }
func (s *ForEachStatement) IsSwitchStatement() bool                   { return false }
func (s *ForEachStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ForEachStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ForEachStatement) IsThrowStatement() bool                    { return false }
func (s *ForEachStatement) IsTryStatement() bool                      { return false }
func (s *ForEachStatement) IsWhileStatement() bool                    { return false }

func (s *ForEachStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ForEachStatement) GetExpression() IExpression { return s.Expression }
func (s *ForEachStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ForEachStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ForEachStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ForEachStatement) GetStatements() IStatement        { return s.Statement }
func (s *ForEachStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ForEachStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ForEachStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ForEachStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ForEachStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ForEachStatement) String() string {
	return fmt.Sprintf("ForEachStatement{ name=%s }", s.Name)
}

type ForStatement struct {
	util.DefaultBase[IStatement]

	Declaration ILocalVariableDeclaration
	Init        IExpression
	Condition   IExpression
	Update      IExpression
	Statements  IStatement
}

func (s *ForStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitForStatement(s)
}

func (s *ForStatement) IsBreakStatement() bool                    { return false }
func (s *ForStatement) IsContinueStatement() bool                 { return false }
func (s *ForStatement) IsExpressionStatement() bool               { return false }
func (s *ForStatement) IsForStatement() bool                      { return false }
func (s *ForStatement) IsIfStatement() bool                       { return false }
func (s *ForStatement) IsIfElseStatement() bool                   { return false }
func (s *ForStatement) IsLabelStatement() bool                    { return false }
func (s *ForStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ForStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ForStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ForStatement) IsMonitorExitStatement() bool              { return false }
func (s *ForStatement) IsReturnStatement() bool                   { return false }
func (s *ForStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ForStatement) IsStatements() bool                        { return false }
func (s *ForStatement) IsSwitchStatement() bool                   { return false }
func (s *ForStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ForStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ForStatement) IsThrowStatement() bool                    { return false }
func (s *ForStatement) IsTryStatement() bool                      { return false }
func (s *ForStatement) IsWhileStatement() bool                    { return false }

func (s *ForStatement) GetCondition() IExpression  { return s.Condition }
func (s *ForStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *ForStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ForStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ForStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ForStatement) GetStatements() IStatement        { return s.Statements }
func (s *ForStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ForStatement) GetInit() IExpression   { return s.Init }
func (s *ForStatement) GetUpdate() IExpression { return s.Update }

func (s *ForStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ForStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ForStatement) String() string {
	return fmt.Sprintf("ForStatement{ %s or %s; %s; %s }", s.Declaration, s.Init, s.Condition, s.Update)
}

type IfElseStatement struct {
	util.DefaultBase[IStatement]

	Condition      IExpression
	IfStatements   IStatement
	ElseStatements IStatement
}

func (s *IfElseStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitIfElseStatement(s)
}

func (s *IfElseStatement) IsBreakStatement() bool                    { return false }
func (s *IfElseStatement) IsContinueStatement() bool                 { return false }
func (s *IfElseStatement) IsExpressionStatement() bool               { return false }
func (s *IfElseStatement) IsForStatement() bool                      { return false }
func (s *IfElseStatement) IsIfStatement() bool                       { return s.Condition != nil }
func (s *IfElseStatement) IsIfElseStatement() bool                   { return true }
func (s *IfElseStatement) IsLabelStatement() bool                    { return false }
func (s *IfElseStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *IfElseStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *IfElseStatement) IsMonitorEnterStatement() bool             { return false }
func (s *IfElseStatement) IsMonitorExitStatement() bool              { return false }
func (s *IfElseStatement) IsReturnStatement() bool                   { return false }
func (s *IfElseStatement) IsReturnExpressionStatement() bool         { return false }
func (s *IfElseStatement) IsStatements() bool                        { return false }
func (s *IfElseStatement) IsSwitchStatement() bool                   { return false }
func (s *IfElseStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *IfElseStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *IfElseStatement) IsThrowStatement() bool                    { return false }
func (s *IfElseStatement) IsTryStatement() bool                      { return false }
func (s *IfElseStatement) IsWhileStatement() bool                    { return false }

func (s *IfElseStatement) GetCondition() IExpression  { return s.Condition }
func (s *IfElseStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *IfElseStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *IfElseStatement) GetElseStatements() IStatement    { return s.ElseStatements }
func (s *IfElseStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *IfElseStatement) GetStatements() IStatement        { return s.IfStatements }
func (s *IfElseStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *IfElseStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *IfElseStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *IfElseStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *IfElseStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *IfElseStatement) String() string {
	return fmt.Sprintf("IfElseStatement{ condition=%s, if=%s, else=%s }", s.Condition, s.IfStatements, s.ElseStatements)
}

type IfStatement struct {
	util.DefaultBase[IStatement]

	Condition  IExpression
	Statements IStatement
}

func (s *IfStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitIfStatement(s)
}

func (s *IfStatement) IsBreakStatement() bool                    { return false }
func (s *IfStatement) IsContinueStatement() bool                 { return false }
func (s *IfStatement) IsExpressionStatement() bool               { return false }
func (s *IfStatement) IsForStatement() bool                      { return false }
func (s *IfStatement) IsIfStatement() bool                       { return s.Condition != nil }
func (s *IfStatement) IsIfElseStatement() bool                   { return false }
func (s *IfStatement) IsLabelStatement() bool                    { return false }
func (s *IfStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *IfStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *IfStatement) IsMonitorEnterStatement() bool             { return false }
func (s *IfStatement) IsMonitorExitStatement() bool              { return false }
func (s *IfStatement) IsReturnStatement() bool                   { return false }
func (s *IfStatement) IsReturnExpressionStatement() bool         { return false }
func (s *IfStatement) IsStatements() bool                        { return false }
func (s *IfStatement) IsSwitchStatement() bool                   { return false }
func (s *IfStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *IfStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *IfStatement) IsThrowStatement() bool                    { return false }
func (s *IfStatement) IsTryStatement() bool                      { return false }
func (s *IfStatement) IsWhileStatement() bool                    { return false }

func (s *IfStatement) GetCondition() IExpression  { return s.Condition }
func (s *IfStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *IfStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *IfStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *IfStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *IfStatement) GetStatements() IStatement        { return s.Statements }
func (s *IfStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *IfStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *IfStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *IfStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *IfStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *IfStatement) String() string {
	return fmt.Sprintf("IfStatement{ condition=%s, if=%s }", s.Condition, s.Statements)
}

type LabelStatement struct {
	util.DefaultBase[IStatement]

	Label     string
	Statement IStatement
}

func (s *LabelStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitLabelStatement(s)
}

func (s *LabelStatement) IsBreakStatement() bool                    { return false }
func (s *LabelStatement) IsContinueStatement() bool                 { return false }
func (s *LabelStatement) IsExpressionStatement() bool               { return false }
func (s *LabelStatement) IsForStatement() bool                      { return false }
func (s *LabelStatement) IsIfStatement() bool                       { return false }
func (s *LabelStatement) IsIfElseStatement() bool                   { return false }
func (s *LabelStatement) IsLabelStatement() bool                    { return true }
func (s *LabelStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *LabelStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *LabelStatement) IsMonitorEnterStatement() bool             { return false }
func (s *LabelStatement) IsMonitorExitStatement() bool              { return false }
func (s *LabelStatement) IsReturnStatement() bool                   { return false }
func (s *LabelStatement) IsReturnExpressionStatement() bool         { return false }
func (s *LabelStatement) IsStatements() bool                        { return false }
func (s *LabelStatement) IsSwitchStatement() bool                   { return false }
func (s *LabelStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *LabelStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *LabelStatement) IsThrowStatement() bool                    { return false }
func (s *LabelStatement) IsTryStatement() bool                      { return false }
func (s *LabelStatement) IsWhileStatement() bool                    { return false }

func (s *LabelStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *LabelStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *LabelStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *LabelStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LabelStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LabelStatement) GetStatements() IStatement        { return s.Statement }
func (s *LabelStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LabelStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *LabelStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *LabelStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LabelStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *LabelStatement) String() string {
	return fmt.Sprintf("LabelStatement{%s: %s}", s.Label, s.Statement)
}

type LambdaExpressionStatement struct {
	util.DefaultBase[IStatement]

	Expression IExpression
}

func (s *LambdaExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitLambdaExpressionStatement(s)
}

func (s *LambdaExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *LambdaExpressionStatement) IsExpressionStatement() bool               { return false }
func (s *LambdaExpressionStatement) IsForStatement() bool                      { return false }
func (s *LambdaExpressionStatement) IsIfStatement() bool                       { return false }
func (s *LambdaExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsLambdaExpressionStatement() bool         { return true }
func (s *LambdaExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *LambdaExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *LambdaExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *LambdaExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsReturnExpressionStatement() bool         { return false }
func (s *LambdaExpressionStatement) IsStatements() bool                        { return false }
func (s *LambdaExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *LambdaExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *LambdaExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *LambdaExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *LambdaExpressionStatement) IsTryStatement() bool                      { return false }
func (s *LambdaExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *LambdaExpressionStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *LambdaExpressionStatement) GetExpression() IExpression { return s.Expression }
func (s *LambdaExpressionStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *LambdaExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LambdaExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LambdaExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *LambdaExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LambdaExpressionStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *LambdaExpressionStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *LambdaExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LambdaExpressionStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *LambdaExpressionStatement) String() string {
	return fmt.Sprintf("LambdaExpressionStatement{ %s }", s.Expression)
}

type LocalVariableDeclarationStatement struct {
	util.DefaultBase[IStatement]
	// FIXME: declaration.LocalVariableDeclaration 리펙토링 후 제작업 필요.
	LocalVariableDeclaration

	Final                    bool
	Type                     IType
	LocalVariableDeclarators ILocalVariableDeclarator
}

func (s *LocalVariableDeclarationStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitLocalVariableDeclarationStatement(s)
}

func (s *LocalVariableDeclarationStatement) IsBreakStatement() bool                    { return false }
func (s *LocalVariableDeclarationStatement) IsContinueStatement() bool                 { return false }
func (s *LocalVariableDeclarationStatement) IsExpressionStatement() bool               { return false }
func (s *LocalVariableDeclarationStatement) IsForStatement() bool                      { return false }
func (s *LocalVariableDeclarationStatement) IsIfStatement() bool                       { return false }
func (s *LocalVariableDeclarationStatement) IsIfElseStatement() bool                   { return false }
func (s *LocalVariableDeclarationStatement) IsLabelStatement() bool                    { return false }
func (s *LocalVariableDeclarationStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *LocalVariableDeclarationStatement) IsLocalVariableDeclarationStatement() bool { return true }
func (s *LocalVariableDeclarationStatement) IsMonitorEnterStatement() bool             { return false }
func (s *LocalVariableDeclarationStatement) IsMonitorExitStatement() bool              { return false }
func (s *LocalVariableDeclarationStatement) IsReturnStatement() bool                   { return false }
func (s *LocalVariableDeclarationStatement) IsReturnExpressionStatement() bool         { return false }
func (s *LocalVariableDeclarationStatement) IsStatements() bool                        { return false }
func (s *LocalVariableDeclarationStatement) IsSwitchStatement() bool                   { return false }
func (s *LocalVariableDeclarationStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *LocalVariableDeclarationStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *LocalVariableDeclarationStatement) IsThrowStatement() bool                    { return false }
func (s *LocalVariableDeclarationStatement) IsTryStatement() bool                      { return false }
func (s *LocalVariableDeclarationStatement) IsWhileStatement() bool                    { return false }

func (s *LocalVariableDeclarationStatement) GetCondition() IExpression {
	return &NeNoExpression
}
func (s *LocalVariableDeclarationStatement) GetExpression() IExpression {
	return &NeNoExpression
}
func (s *LocalVariableDeclarationStatement) GetMonitor() IExpression {
	return &NeNoExpression
}

func (s *LocalVariableDeclarationStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetStatements() IStatement        { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LocalVariableDeclarationStatement) GetInit() IExpression { return &NeNoExpression }
func (s *LocalVariableDeclarationStatement) GetUpdate() IExpression {
	return &NeNoExpression
}

func (s *LocalVariableDeclarationStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LocalVariableDeclarationStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *LocalVariableDeclarationStatement) String() string {
	return fmt.Sprintf("LocalVariableDeclarationStatement{ %s %s }", s.Type, s.LocalVariableDeclarators)
}

type NoStatement struct {
	util.DefaultBase[IStatement]
}

func (s *NoStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitNoStatement(s)
}

func (s *NoStatement) IsBreakStatement() bool                    { return false }
func (s *NoStatement) IsContinueStatement() bool                 { return false }
func (s *NoStatement) IsExpressionStatement() bool               { return false }
func (s *NoStatement) IsForStatement() bool                      { return false }
func (s *NoStatement) IsIfStatement() bool                       { return false }
func (s *NoStatement) IsIfElseStatement() bool                   { return false }
func (s *NoStatement) IsLabelStatement() bool                    { return false }
func (s *NoStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *NoStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *NoStatement) IsMonitorEnterStatement() bool             { return false }
func (s *NoStatement) IsMonitorExitStatement() bool              { return false }
func (s *NoStatement) IsReturnStatement() bool                   { return false }
func (s *NoStatement) IsReturnExpressionStatement() bool         { return false }
func (s *NoStatement) IsStatements() bool                        { return false }
func (s *NoStatement) IsSwitchStatement() bool                   { return false }
func (s *NoStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *NoStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *NoStatement) IsThrowStatement() bool                    { return false }
func (s *NoStatement) IsTryStatement() bool                      { return false }
func (s *NoStatement) IsWhileStatement() bool                    { return false }

func (s *NoStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *NoStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *NoStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *NoStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *NoStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *NoStatement) GetStatements() IStatement        { return &NoStmt }
func (s *NoStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *NoStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *NoStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *NoStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *NoStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *NoStatement) String() string {
	return "NoStatement{}"
}

type ReturnExpressionStatement struct {
	util.DefaultBase[IStatement]

	LineNumber int
	Expression IExpression
}

//func (s *ReturnExpressionStatement) GenericExpression() model.IExpression {
//	return s.expression
//}

func (s *ReturnExpressionStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitReturnExpressionStatement(s)
}

func (s *ReturnExpressionStatement) IsBreakStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsContinueStatement() bool                 { return false }
func (s *ReturnExpressionStatement) IsExpressionStatement() bool               { return false }
func (s *ReturnExpressionStatement) IsForStatement() bool                      { return false }
func (s *ReturnExpressionStatement) IsIfStatement() bool                       { return false }
func (s *ReturnExpressionStatement) IsIfElseStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsLabelStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ReturnExpressionStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ReturnExpressionStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ReturnExpressionStatement) IsMonitorExitStatement() bool              { return false }
func (s *ReturnExpressionStatement) IsReturnStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsReturnExpressionStatement() bool         { return true }
func (s *ReturnExpressionStatement) IsStatements() bool                        { return false }
func (s *ReturnExpressionStatement) IsSwitchStatement() bool                   { return false }
func (s *ReturnExpressionStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ReturnExpressionStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ReturnExpressionStatement) IsThrowStatement() bool                    { return false }
func (s *ReturnExpressionStatement) IsTryStatement() bool                      { return false }
func (s *ReturnExpressionStatement) IsWhileStatement() bool                    { return false }

func (s *ReturnExpressionStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ReturnExpressionStatement) GetExpression() IExpression { return s.Expression }
func (s *ReturnExpressionStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ReturnExpressionStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ReturnExpressionStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ReturnExpressionStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ReturnExpressionStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ReturnExpressionStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ReturnExpressionStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ReturnExpressionStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ReturnExpressionStatement) GetLineNumber() int                        { return s.LineNumber }

func (s *ReturnExpressionStatement) String() string {
	return fmt.Sprintf("ReturnExpressionStatement{return %s}", s.Expression)
}

type ReturnStatement struct {
	util.DefaultBase[IStatement]
}

func (s *ReturnStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitReturnStatement(s)
}

func (s *ReturnStatement) IsBreakStatement() bool                    { return false }
func (s *ReturnStatement) IsContinueStatement() bool                 { return false }
func (s *ReturnStatement) IsExpressionStatement() bool               { return false }
func (s *ReturnStatement) IsForStatement() bool                      { return false }
func (s *ReturnStatement) IsIfStatement() bool                       { return false }
func (s *ReturnStatement) IsIfElseStatement() bool                   { return false }
func (s *ReturnStatement) IsLabelStatement() bool                    { return false }
func (s *ReturnStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ReturnStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ReturnStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ReturnStatement) IsMonitorExitStatement() bool              { return false }
func (s *ReturnStatement) IsReturnStatement() bool                   { return true }
func (s *ReturnStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ReturnStatement) IsStatements() bool                        { return false }
func (s *ReturnStatement) IsSwitchStatement() bool                   { return false }
func (s *ReturnStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ReturnStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ReturnStatement) IsThrowStatement() bool                    { return false }
func (s *ReturnStatement) IsTryStatement() bool                      { return false }
func (s *ReturnStatement) IsWhileStatement() bool                    { return false }

func (s *ReturnStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ReturnStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *ReturnStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ReturnStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ReturnStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ReturnStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ReturnStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ReturnStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ReturnStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ReturnStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ReturnStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ReturnStatement) String() string {
	return "ReturnStatement{}"
}

type Statements struct {
	util.DefaultList[IStatement]
}

func (s *Statements) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitStatements(s)
}

func (s *Statements) IsBreakStatement() bool                    { return false }
func (s *Statements) IsContinueStatement() bool                 { return false }
func (s *Statements) IsExpressionStatement() bool               { return false }
func (s *Statements) IsForStatement() bool                      { return false }
func (s *Statements) IsIfStatement() bool                       { return false }
func (s *Statements) IsIfElseStatement() bool                   { return false }
func (s *Statements) IsLabelStatement() bool                    { return false }
func (s *Statements) IsLambdaExpressionStatement() bool         { return false }
func (s *Statements) IsLocalVariableDeclarationStatement() bool { return false }
func (s *Statements) IsMonitorEnterStatement() bool             { return false }
func (s *Statements) IsMonitorExitStatement() bool              { return false }
func (s *Statements) IsReturnStatement() bool                   { return false }
func (s *Statements) IsReturnExpressionStatement() bool         { return false }
func (s *Statements) IsStatements() bool                        { return true }
func (s *Statements) IsSwitchStatement() bool                   { return false }
func (s *Statements) IsSwitchStatementLabelBlock() bool         { return false }
func (s *Statements) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *Statements) IsThrowStatement() bool                    { return false }
func (s *Statements) IsTryStatement() bool                      { return false }
func (s *Statements) IsWhileStatement() bool                    { return false }

func (s *Statements) GetCondition() IExpression  { return &NeNoExpression }
func (s *Statements) GetExpression() IExpression { return &NeNoExpression }
func (s *Statements) GetMonitor() IExpression    { return &NeNoExpression }

func (s *Statements) GetElseStatements() IStatement    { return &NoStmt }
func (s *Statements) GetFinallyStatements() IStatement { return &NoStmt }
func (s *Statements) GetStatements() IStatement        { return &NoStmt }
func (s *Statements) GetTryStatements() IStatement     { return &NoStmt }

func (s *Statements) GetInit() IExpression   { return &NeNoExpression }
func (s *Statements) GetUpdate() IExpression { return &NeNoExpression }

func (s *Statements) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *Statements) GetLineNumber() int                        { return UnknownLineNumber }

func (s *Statements) String() string {
	return fmt.Sprintf("Statements{ %d }", s.Size())
}

type SwitchStatement struct {
	util.DefaultBase[IStatement]

	Condition IExpression
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

func (s *SwitchStatement) GetCondition() IExpression  { return s.Condition }
func (s *SwitchStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *SwitchStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *SwitchStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *SwitchStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *SwitchStatement) GetStatements() IStatement        { return &NoStmt }
func (s *SwitchStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *SwitchStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *SwitchStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *SwitchStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *SwitchStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *SwitchStatement) String() string {
	return "SwitchStatement{}"
}

type SynchronizedStatement struct {
	util.DefaultBase[IStatement]

	Monitor    IExpression
	Statements IStatement
}

func (s *SynchronizedStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSynchronizedStatement(s)
}

func (s *SynchronizedStatement) IsBreakStatement() bool                    { return false }
func (s *SynchronizedStatement) IsContinueStatement() bool                 { return false }
func (s *SynchronizedStatement) IsExpressionStatement() bool               { return false }
func (s *SynchronizedStatement) IsForStatement() bool                      { return false }
func (s *SynchronizedStatement) IsIfStatement() bool                       { return false }
func (s *SynchronizedStatement) IsIfElseStatement() bool                   { return false }
func (s *SynchronizedStatement) IsLabelStatement() bool                    { return false }
func (s *SynchronizedStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *SynchronizedStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *SynchronizedStatement) IsMonitorEnterStatement() bool             { return false }
func (s *SynchronizedStatement) IsMonitorExitStatement() bool              { return false }
func (s *SynchronizedStatement) IsReturnStatement() bool                   { return false }
func (s *SynchronizedStatement) IsReturnExpressionStatement() bool         { return false }
func (s *SynchronizedStatement) IsStatements() bool                        { return false }
func (s *SynchronizedStatement) IsSwitchStatement() bool                   { return false }
func (s *SynchronizedStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *SynchronizedStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *SynchronizedStatement) IsThrowStatement() bool                    { return false }
func (s *SynchronizedStatement) IsTryStatement() bool                      { return false }
func (s *SynchronizedStatement) IsWhileStatement() bool                    { return false }

func (s *SynchronizedStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *SynchronizedStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *SynchronizedStatement) GetMonitor() IExpression    { return s.Monitor }

func (s *SynchronizedStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *SynchronizedStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *SynchronizedStatement) GetStatements() IStatement        { return s.Statements }
func (s *SynchronizedStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *SynchronizedStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *SynchronizedStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *SynchronizedStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *SynchronizedStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *SynchronizedStatement) String() string {
	return fmt.Sprintf("SynchronizedStatement{ monitor=%s, statements=%s }", s.Monitor, s.Statements)
}

type ThrowStatement struct {
	util.DefaultBase[IStatement]

	Expression IExpression
}

func (s *ThrowStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitThrowStatement(s)
}

func (s *ThrowStatement) IsBreakStatement() bool                    { return false }
func (s *ThrowStatement) IsContinueStatement() bool                 { return false }
func (s *ThrowStatement) IsExpressionStatement() bool               { return false }
func (s *ThrowStatement) IsForStatement() bool                      { return false }
func (s *ThrowStatement) IsIfStatement() bool                       { return false }
func (s *ThrowStatement) IsIfElseStatement() bool                   { return false }
func (s *ThrowStatement) IsLabelStatement() bool                    { return false }
func (s *ThrowStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *ThrowStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *ThrowStatement) IsMonitorEnterStatement() bool             { return false }
func (s *ThrowStatement) IsMonitorExitStatement() bool              { return false }
func (s *ThrowStatement) IsReturnStatement() bool                   { return false }
func (s *ThrowStatement) IsReturnExpressionStatement() bool         { return false }
func (s *ThrowStatement) IsStatements() bool                        { return false }
func (s *ThrowStatement) IsSwitchStatement() bool                   { return false }
func (s *ThrowStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *ThrowStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *ThrowStatement) IsThrowStatement() bool                    { return true }
func (s *ThrowStatement) IsTryStatement() bool                      { return false }
func (s *ThrowStatement) IsWhileStatement() bool                    { return false }

func (s *ThrowStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *ThrowStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *ThrowStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *ThrowStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *ThrowStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *ThrowStatement) GetStatements() IStatement        { return &NoStmt }
func (s *ThrowStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *ThrowStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *ThrowStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *ThrowStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *ThrowStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *ThrowStatement) String() string {
	return fmt.Sprintf("ThrowStatement{throw %s}", s.Expression)
}

type TryStatement struct {
	util.DefaultBase[IStatement]

	Resources         util.DefaultList[*Resource]
	TryStatements     IStatement
	CatchClause       util.DefaultList[*CatchClause]
	FinallyStatements IStatement
}

func (s *TryStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatement(s)
}

func (s *TryStatement) IsBreakStatement() bool                    { return false }
func (s *TryStatement) IsContinueStatement() bool                 { return false }
func (s *TryStatement) IsExpressionStatement() bool               { return false }
func (s *TryStatement) IsForStatement() bool                      { return false }
func (s *TryStatement) IsIfStatement() bool                       { return false }
func (s *TryStatement) IsIfElseStatement() bool                   { return false }
func (s *TryStatement) IsLabelStatement() bool                    { return false }
func (s *TryStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *TryStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *TryStatement) IsMonitorEnterStatement() bool             { return false }
func (s *TryStatement) IsMonitorExitStatement() bool              { return false }
func (s *TryStatement) IsReturnStatement() bool                   { return false }
func (s *TryStatement) IsReturnExpressionStatement() bool         { return false }
func (s *TryStatement) IsStatements() bool                        { return false }
func (s *TryStatement) IsSwitchStatement() bool                   { return false }
func (s *TryStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *TryStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *TryStatement) IsThrowStatement() bool                    { return false }
func (s *TryStatement) IsTryStatement() bool                      { return true }
func (s *TryStatement) IsWhileStatement() bool                    { return false }

func (s *TryStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *TryStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *TryStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *TryStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *TryStatement) GetFinallyStatements() IStatement { return s.FinallyStatements }
func (s *TryStatement) GetStatements() IStatement        { return &NoStmt }
func (s *TryStatement) GetTryStatements() IStatement     { return s.TryStatements }

func (s *TryStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *TryStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *TryStatement) GetCatchClauses() util.IList[*CatchClause] { return &s.CatchClause }
func (s *TryStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *TryStatement) String() string {
	return ""
}

type TypeDeclarationStatement struct {
	util.DefaultBase[IStatement]

	TypeDeclaration TypeDeclaration
}

func (s *TypeDeclarationStatement) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTypeDeclarationStatement(s)
}

func (s *TypeDeclarationStatement) IsBreakStatement() bool                    { return false }
func (s *TypeDeclarationStatement) IsContinueStatement() bool                 { return false }
func (s *TypeDeclarationStatement) IsExpressionStatement() bool               { return false }
func (s *TypeDeclarationStatement) IsForStatement() bool                      { return false }
func (s *TypeDeclarationStatement) IsIfStatement() bool                       { return false }
func (s *TypeDeclarationStatement) IsIfElseStatement() bool                   { return false }
func (s *TypeDeclarationStatement) IsLabelStatement() bool                    { return false }
func (s *TypeDeclarationStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *TypeDeclarationStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *TypeDeclarationStatement) IsMonitorEnterStatement() bool             { return false }
func (s *TypeDeclarationStatement) IsMonitorExitStatement() bool              { return false }
func (s *TypeDeclarationStatement) IsReturnStatement() bool                   { return false }
func (s *TypeDeclarationStatement) IsReturnExpressionStatement() bool         { return false }
func (s *TypeDeclarationStatement) IsStatements() bool                        { return false }
func (s *TypeDeclarationStatement) IsSwitchStatement() bool                   { return false }
func (s *TypeDeclarationStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *TypeDeclarationStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *TypeDeclarationStatement) IsThrowStatement() bool                    { return false }
func (s *TypeDeclarationStatement) IsTryStatement() bool                      { return false }
func (s *TypeDeclarationStatement) IsWhileStatement() bool                    { return false }

func (s *TypeDeclarationStatement) GetCondition() IExpression  { return &NeNoExpression }
func (s *TypeDeclarationStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *TypeDeclarationStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *TypeDeclarationStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *TypeDeclarationStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *TypeDeclarationStatement) GetStatements() IStatement        { return &NoStmt }
func (s *TypeDeclarationStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *TypeDeclarationStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *TypeDeclarationStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *TypeDeclarationStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *TypeDeclarationStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *TypeDeclarationStatement) String() string {
	return fmt.Sprintf("TypeDeclarationStatement{ %s }", s.TypeDeclaration)
}

type WhileStatement struct {
	util.DefaultBase[IStatement]

	Condition  IExpression
	Statements IStatement
}

func (s *WhileStatement) IsBreakStatement() bool                    { return false }
func (s *WhileStatement) IsContinueStatement() bool                 { return false }
func (s *WhileStatement) IsExpressionStatement() bool               { return false }
func (s *WhileStatement) IsForStatement() bool                      { return false }
func (s *WhileStatement) IsIfStatement() bool                       { return false }
func (s *WhileStatement) IsIfElseStatement() bool                   { return false }
func (s *WhileStatement) IsLabelStatement() bool                    { return false }
func (s *WhileStatement) IsLambdaExpressionStatement() bool         { return false }
func (s *WhileStatement) IsLocalVariableDeclarationStatement() bool { return false }
func (s *WhileStatement) IsMonitorEnterStatement() bool             { return false }
func (s *WhileStatement) IsMonitorExitStatement() bool              { return false }
func (s *WhileStatement) IsReturnStatement() bool                   { return false }
func (s *WhileStatement) IsReturnExpressionStatement() bool         { return false }
func (s *WhileStatement) IsStatements() bool                        { return false }
func (s *WhileStatement) IsSwitchStatement() bool                   { return false }
func (s *WhileStatement) IsSwitchStatementLabelBlock() bool         { return false }
func (s *WhileStatement) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (s *WhileStatement) IsThrowStatement() bool                    { return false }
func (s *WhileStatement) IsTryStatement() bool                      { return false }
func (s *WhileStatement) IsWhileStatement() bool                    { return true }

func (s *WhileStatement) GetCondition() IExpression  { return s.Condition }
func (s *WhileStatement) GetExpression() IExpression { return &NeNoExpression }
func (s *WhileStatement) GetMonitor() IExpression    { return &NeNoExpression }

func (s *WhileStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *WhileStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *WhileStatement) GetStatements() IStatement        { return s.Statements }
func (s *WhileStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *WhileStatement) GetInit() IExpression   { return &NeNoExpression }
func (s *WhileStatement) GetUpdate() IExpression { return &NeNoExpression }

func (s *WhileStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *WhileStatement) GetLineNumber() int                        { return UnknownLineNumber }

func (s *WhileStatement) String() string {
	return fmt.Sprintf("WhileStatement{ condition=%s, expression=%s }", s.Condition, s.Statements)
}

/////////////////////////////////////////////////////////////////////////
//  Additional Structures
/////////////////////////////////////////////////////////////////////////

type DefaultLabel struct {
	util.DefaultBase[IStatement]
}

func (l *DefaultLabel) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementDefaultLabel(l)
}

func (l *DefaultLabel) IsLabel() bool                             { return true }
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

func (l *DefaultLabel) GetCondition() IExpression  { return &NeNoExpression }
func (l *DefaultLabel) GetExpression() IExpression { return &NeNoExpression }
func (l *DefaultLabel) GetMonitor() IExpression    { return &NeNoExpression }

func (l *DefaultLabel) GetElseStatements() IStatement    { return &NoStmt }
func (l *DefaultLabel) GetFinallyStatements() IStatement { return &NoStmt }
func (l *DefaultLabel) GetStatements() IStatement        { return &NoStmt }
func (l *DefaultLabel) GetTryStatements() IStatement     { return &NoStmt }

func (l *DefaultLabel) GetInit() IExpression   { return &NeNoExpression }
func (l *DefaultLabel) GetUpdate() IExpression { return &NeNoExpression }

func (l *DefaultLabel) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (l *DefaultLabel) GetLineNumber() int                        { return UnknownLineNumber }

func (l *DefaultLabel) String() string {
	return "DefaultLabel{}"
}

type ExpressionLabel struct {
	util.DefaultBase[IStatement]

	Expression IExpression
}

func (l *ExpressionLabel) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementExpressionLabel(l)
}

func (l *ExpressionLabel) IsLabel() bool                             { return true }
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

func (l *ExpressionLabel) GetCondition() IExpression  { return &NeNoExpression }
func (l *ExpressionLabel) GetExpression() IExpression { return l.Expression }
func (l *ExpressionLabel) GetMonitor() IExpression    { return &NeNoExpression }

func (l *ExpressionLabel) GetElseStatements() IStatement    { return &NoStmt }
func (l *ExpressionLabel) GetFinallyStatements() IStatement { return &NoStmt }
func (l *ExpressionLabel) GetStatements() IStatement        { return &NoStmt }
func (l *ExpressionLabel) GetTryStatements() IStatement     { return &NoStmt }

func (l *ExpressionLabel) GetInit() IExpression   { return &NeNoExpression }
func (l *ExpressionLabel) GetUpdate() IExpression { return &NeNoExpression }

func (l *ExpressionLabel) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (l *ExpressionLabel) GetLineNumber() int                        { return UnknownLineNumber }

func (l *ExpressionLabel) String() string {
	return fmt.Sprintf("ExpressionLabel{ expression=%s }", l.Expression)
}

type Block struct {
	util.DefaultBase[IStatement]

	Statements IStatement
}

func (b *Block) AcceptStatement(visitor IStatementVisitor) {}

func (b *Block) IsBlock() bool                             { return true }
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

func (b *Block) GetCondition() IExpression  { return &NeNoExpression }
func (b *Block) GetExpression() IExpression { return &NeNoExpression }
func (b *Block) GetMonitor() IExpression    { return &NeNoExpression }

func (b *Block) GetElseStatements() IStatement    { return &NoStmt }
func (b *Block) GetFinallyStatements() IStatement { return &NoStmt }
func (b *Block) GetStatements() IStatement        { return b.Statements }
func (b *Block) GetTryStatements() IStatement     { return &NoStmt }

func (b *Block) GetInit() IExpression   { return &NeNoExpression }
func (b *Block) GetUpdate() IExpression { return &NeNoExpression }

func (b *Block) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (b *Block) GetLineNumber() int                        { return UnknownLineNumber }

func (b *Block) String() string {
	return "Block{}"
}

type LabelBlock struct {
	util.DefaultBase[IStatement]

	Statements IStatement
	Label      ILabel
}

func (b *LabelBlock) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementLabelBlock(b)
}

func (b *LabelBlock) IsBlock() bool                             { return true }
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

func (b *LabelBlock) GetCondition() IExpression  { return &NeNoExpression }
func (b *LabelBlock) GetExpression() IExpression { return &NeNoExpression }
func (b *LabelBlock) GetMonitor() IExpression    { return &NeNoExpression }

func (b *LabelBlock) GetElseStatements() IStatement    { return &NoStmt }
func (b *LabelBlock) GetFinallyStatements() IStatement { return &NoStmt }
func (b *LabelBlock) GetStatements() IStatement        { return b.Statements }
func (b *LabelBlock) GetTryStatements() IStatement     { return &NoStmt }

func (b *LabelBlock) GetInit() IExpression   { return &NeNoExpression }
func (b *LabelBlock) GetUpdate() IExpression { return &NeNoExpression }

func (b *LabelBlock) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (b *LabelBlock) GetLineNumber() int                        { return UnknownLineNumber }

func (b *LabelBlock) String() string {
	return fmt.Sprintf("LabelBlock{ label=%s }", b.Label)
}

type MultiLabelsBlock struct {
	util.DefaultBase[IStatement]

	Statements IStatement
	Labels     util.DefaultList[ILabel]
}

func (b *MultiLabelsBlock) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitSwitchStatementMultiLabelsBlock(b)
}

func (b *MultiLabelsBlock) IsBlock() bool                             { return true }
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

func (b *MultiLabelsBlock) GetCondition() IExpression  { return &NeNoExpression }
func (b *MultiLabelsBlock) GetExpression() IExpression { return &NeNoExpression }
func (b *MultiLabelsBlock) GetMonitor() IExpression    { return &NeNoExpression }

func (b *MultiLabelsBlock) GetElseStatements() IStatement    { return &NoStmt }
func (b *MultiLabelsBlock) GetFinallyStatements() IStatement { return &NoStmt }
func (b *MultiLabelsBlock) GetStatements() IStatement        { return b.Statements }
func (b *MultiLabelsBlock) GetTryStatements() IStatement     { return &NoStmt }

func (b *MultiLabelsBlock) GetInit() IExpression   { return &NeNoExpression }
func (b *MultiLabelsBlock) GetUpdate() IExpression { return &NeNoExpression }

func (b *MultiLabelsBlock) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (b *MultiLabelsBlock) GetLineNumber() int                        { return UnknownLineNumber }

func (b *MultiLabelsBlock) String() string {
	return fmt.Sprintf("MultiLabelsBlock{ label=%s }", b.Labels.ToSlice())
}

type Resource struct {
	util.DefaultBase[IStatement]

	Type       *ObjectType
	Name       string
	Expression IExpression
}

func (r *Resource) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatementResource(r)
}

func (r *Resource) IsBreakStatement() bool                    { return false }
func (r *Resource) IsContinueStatement() bool                 { return false }
func (r *Resource) IsExpressionStatement() bool               { return false }
func (r *Resource) IsForStatement() bool                      { return false }
func (r *Resource) IsIfStatement() bool                       { return false }
func (r *Resource) IsIfElseStatement() bool                   { return false }
func (r *Resource) IsLabelStatement() bool                    { return false }
func (r *Resource) IsLambdaExpressionStatement() bool         { return false }
func (r *Resource) IsLocalVariableDeclarationStatement() bool { return false }
func (r *Resource) IsMonitorEnterStatement() bool             { return false }
func (r *Resource) IsMonitorExitStatement() bool              { return false }
func (r *Resource) IsReturnStatement() bool                   { return false }
func (r *Resource) IsReturnExpressionStatement() bool         { return false }
func (r *Resource) IsStatements() bool                        { return false }
func (r *Resource) IsSwitchStatement() bool                   { return false }
func (r *Resource) IsSwitchStatementLabelBlock() bool         { return false }
func (r *Resource) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (r *Resource) IsThrowStatement() bool                    { return false }
func (r *Resource) IsTryStatement() bool                      { return false }
func (r *Resource) IsWhileStatement() bool                    { return false }

func (r *Resource) GetCondition() IExpression  { return &NeNoExpression }
func (r *Resource) GetExpression() IExpression { return r.Expression }
func (r *Resource) GetMonitor() IExpression    { return &NeNoExpression }

func (r *Resource) GetElseStatements() IStatement    { return &NoStmt }
func (r *Resource) GetFinallyStatements() IStatement { return &NoStmt }
func (r *Resource) GetStatements() IStatement        { return &NoStmt }
func (r *Resource) GetTryStatements() IStatement     { return &NoStmt }

func (r *Resource) GetInit() IExpression   { return &NeNoExpression }
func (r *Resource) GetUpdate() IExpression { return &NeNoExpression }

func (r *Resource) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (r *Resource) GetLineNumber() int                        { return UnknownLineNumber }

func (r *Resource) String() string {
	return "Resource{}"
}

type CatchClause struct {
	util.DefaultBase[IStatement]

	LineNumber int
	Type       *ObjectType
	OtherType  util.DefaultList[*ObjectType]
	Name       string
	Statements IStatement
}

func (c *CatchClause) AddType(typ ObjectType) {
	c.OtherType.Add(&typ)
}

func (c *CatchClause) AcceptStatement(visitor IStatementVisitor) {
	visitor.VisitTryStatementCatchClause(c)
}

func (c *CatchClause) IsBreakStatement() bool                    { return false }
func (c *CatchClause) IsContinueStatement() bool                 { return false }
func (c *CatchClause) IsExpressionStatement() bool               { return false }
func (c *CatchClause) IsForStatement() bool                      { return false }
func (c *CatchClause) IsIfStatement() bool                       { return false }
func (c *CatchClause) IsIfElseStatement() bool                   { return false }
func (c *CatchClause) IsLabelStatement() bool                    { return false }
func (c *CatchClause) IsLambdaExpressionStatement() bool         { return false }
func (c *CatchClause) IsLocalVariableDeclarationStatement() bool { return false }
func (c *CatchClause) IsMonitorEnterStatement() bool             { return false }
func (c *CatchClause) IsMonitorExitStatement() bool              { return false }
func (c *CatchClause) IsReturnStatement() bool                   { return false }
func (c *CatchClause) IsReturnExpressionStatement() bool         { return false }
func (c *CatchClause) IsStatements() bool                        { return false }
func (c *CatchClause) IsSwitchStatement() bool                   { return false }
func (c *CatchClause) IsSwitchStatementLabelBlock() bool         { return false }
func (c *CatchClause) IsSwitchStatementMultiLabelsBlock() bool   { return false }
func (c *CatchClause) IsThrowStatement() bool                    { return false }
func (c *CatchClause) IsTryStatement() bool                      { return false }
func (c *CatchClause) IsWhileStatement() bool                    { return false }

func (c *CatchClause) GetCondition() IExpression  { return &NeNoExpression }
func (c *CatchClause) GetExpression() IExpression { return &NeNoExpression }
func (c *CatchClause) GetMonitor() IExpression    { return &NeNoExpression }

func (c *CatchClause) GetElseStatements() IStatement    { return &NoStmt }
func (c *CatchClause) GetFinallyStatements() IStatement { return &NoStmt }
func (c *CatchClause) GetStatements() IStatement        { return c.Statements }
func (c *CatchClause) GetTryStatements() IStatement     { return &NoStmt }

func (c *CatchClause) GetInit() IExpression   { return &NeNoExpression }
func (c *CatchClause) GetUpdate() IExpression { return &NeNoExpression }

func (c *CatchClause) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (c *CatchClause) GetLineNumber() int                        { return c.LineNumber }

func (c *CatchClause) String() string {
	return "CatchClause{}"
}

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////
