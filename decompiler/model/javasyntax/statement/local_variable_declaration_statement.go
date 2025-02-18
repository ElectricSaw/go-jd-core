package statement

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewLocalVariableDeclarationStatement(typ model.IType,
	localVariableDeclarators model.ILocalVariableDeclarator) LocalVariableDeclarationStatement {
	return &LocalVariableDeclarationStatement{
		Type:                     typ,
		localVariableDeclarators: localVariableDeclarators,
	}
}

type LocalVariableDeclarationStatement struct {
	util.DefaultBase[IStatement]
	// FIXME: declaration.LocalVariableDeclaration 리펙토링 후 제작업 필요.
	declaration.LocalVariableDeclaration

	Final                    bool
	Type                     model.IType
	localVariableDeclarators model.ILocalVariableDeclarator
}

func (s *LocalVariableDeclarationStatement) LocalVariableDeclarators() model.ILocalVariableDeclarator {
	return s.localVariableDeclarators
}

func (s *LocalVariableDeclarationStatement) SetLocalVariableDeclarators(declarators model.ILocalVariableDeclarator) {
	s.localVariableDeclarators = declarators
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

func (s *LocalVariableDeclarationStatement) GetCondition() model.IExpression {
	return &model.NeNoExpression
}
func (s *LocalVariableDeclarationStatement) GetExpression() model.IExpression {
	return &model.NeNoExpression
}
func (s *LocalVariableDeclarationStatement) GetMonitor() model.IExpression {
	return &model.NeNoExpression
}

func (s *LocalVariableDeclarationStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetStatements() IStatement        { return &NoStmt }
func (s *LocalVariableDeclarationStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *LocalVariableDeclarationStatement) GetInit() model.IExpression { return &model.NeNoExpression }
func (s *LocalVariableDeclarationStatement) GetUpdate() model.IExpression {
	return &model.NeNoExpression
}

func (s *LocalVariableDeclarationStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *LocalVariableDeclarationStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *LocalVariableDeclarationStatement) String() string {
	return fmt.Sprintf("LocalVariableDeclarationStatement{ %s %s }", s.Type, s.localVariableDeclarators)
}
