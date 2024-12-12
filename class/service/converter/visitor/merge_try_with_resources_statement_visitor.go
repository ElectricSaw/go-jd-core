package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
)

func NewMergeTryWithResourcesStatementVisitor() intsrv.IMergeTryWithResourcesStatementVisitor {
	return &MergeTryWithResourcesStatementVisitor{}
}

type MergeTryWithResourcesStatementVisitor struct {
}

func (v *MergeTryWithResourcesStatementVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Statements().AcceptStatement(v)
	statement.ElseStatements().AcceptStatement(v)
}

func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	for _, block := range statement.Blocks() {
		block.Statements().AcceptStatement(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) VisitTryStatement(statement intmod.ITryStatement) {
	tryStatements := statement.TryStatements()

	v.SafeAcceptSlice(statement.ResourceList())
	tryStatements.AcceptStatement(v)
	v.SafeAcceptSlice(statement.CatchClauseList())
	v.SafeAccept(statement.FinallyStatements())

	if tryStatements.Size() == 1 {
		first := tryStatements.First()

		if first.IsTryStatement() {
			cfswrs1 := statement.(intsrv.IClassFileTryStatement)
			cfswrs2 := first.(intsrv.IClassFileTryStatement)

			if cfswrs2.Resources() != nil && cfswrs2.CatchClauses() == nil && cfswrs2.FinallyStatements() == nil {
				// Merge 'try' and 'try-with-resources" statements
				cfswrs1.SetTryStatements(cfswrs2.TryStatements())
				cfswrs1.AddResources(cfswrs2.Resources())
			}
		}
	}
}

func (v *MergeTryWithResourcesStatementVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitStatements(list intmod.IStatements) {
	v.AcceptSlice(list.ToSlice())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.SafeAccept(statement.Statements())
}
func (v *MergeTryWithResourcesStatementVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	v.SafeAccept(statement.Statements())
}

func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementLabelBlock(statement intmod.ILabelBlock) {
	statement.Statements().AcceptStatement(v)
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementMultiLabelsBlock(statement intmod.IMultiLabelsBlock) {
	statement.Statements().AcceptStatement(v)
}

func (v *MergeTryWithResourcesStatementVisitor) VisitAssertStatement(_ intmod.IAssertStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitExpressionStatement(_ intmod.IExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLabelStatement(_ intmod.ILabelStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLambdaExpressionStatement(_ intmod.ILambdaExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLocalVariableDeclarationStatement(_ intmod.ILocalVariableDeclarationStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitNoStatement(_ intmod.INoStatement) {}
func (v *MergeTryWithResourcesStatementVisitor) VisitReturnExpressionStatement(_ intmod.IReturnExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementExpressionLabel(_ intmod.IExpressionLabel) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitThrowStatement(_ intmod.IThrowStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitTryStatementResource(_ intmod.IResource) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitTypeDeclarationStatement(_ intmod.ITypeDeclarationStatement) {
}

func (v *MergeTryWithResourcesStatementVisitor) SafeAccept(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) AcceptSlice(list []intmod.IStatement) {
	for _, state := range list {
		state.AcceptStatement(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) SafeAcceptSlice(list []intmod.IStatement) {
	if list != nil {
		for _, state := range list {
			state.AcceptStatement(v)
		}
	}
}
