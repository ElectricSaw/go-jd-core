package statement

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewTypeDeclarationStatement(typeDeclaration TypeDeclaration) TypeDeclarationStatement {
	return &TypeDeclarationStatement{
		TypeDeclaration: typeDeclaration,
	}
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

func (s *TypeDeclarationStatement) GetCondition() model.IExpression  { return &model.NeNoExpression }
func (s *TypeDeclarationStatement) GetExpression() model.IExpression { return &model.NeNoExpression }
func (s *TypeDeclarationStatement) GetMonitor() model.IExpression    { return &model.NeNoExpression }

func (s *TypeDeclarationStatement) GetElseStatements() IStatement    { return &NoStmt }
func (s *TypeDeclarationStatement) GetFinallyStatements() IStatement { return &NoStmt }
func (s *TypeDeclarationStatement) GetStatements() IStatement        { return &NoStmt }
func (s *TypeDeclarationStatement) GetTryStatements() IStatement     { return &NoStmt }

func (s *TypeDeclarationStatement) GetInit() model.IExpression   { return &model.NeNoExpression }
func (s *TypeDeclarationStatement) GetUpdate() model.IExpression { return &model.NeNoExpression }

func (s *TypeDeclarationStatement) GetCatchClauses() util.IList[*CatchClause] { return nil }
func (s *TypeDeclarationStatement) GetLineNumber() int                        { return model.UnknownLineNumber }

func (s *TypeDeclarationStatement) String() string {
	return "TypeDeclarationStatement{}"
}
