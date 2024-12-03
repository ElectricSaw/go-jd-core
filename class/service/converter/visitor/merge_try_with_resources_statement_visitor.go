package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
)

func NewMergeTryWithResourcesStatementVisitor() *MergeTryWithResourcesStatementVisitor {
	return &MergeTryWithResourcesStatementVisitor{}
}

type MergeTryWithResourcesStatementVisitor struct {
}

func (v *MergeTryWithResourcesStatementVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Statements().Accept(v)
	statement.ElseStatements().Accept(v)
}

func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	for _, block := range statement.Blocks() {
		block.Statements().Accept(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) VisitTryStatement(statement intmod.ITryStatement) {
	tryStatements := statement.TryStatements()

	v.SafeAcceptListStatement(statement.ResourceList())
	tryStatements.Accept(v)
	v.SafeAcceptListStatement(statement.CatchClauseList())
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
	v.AcceptListStatement(list.ToSlice())
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
	statement.Statements().Accept(v)
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementMultiLabelsBlock(statement intmod.IMultiLabelsBlock) {
	statement.Statements().Accept(v)
}

func (v *MergeTryWithResourcesStatementVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitBreakStatement(statement intmod.IBreakStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitByteCodeStatement(statement intmod.IByteCodeStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitCommentStatement(statement intmod.ICommentStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitContinueStatement(statement intmod.IContinueStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLabelStatement(statement intmod.ILabelStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitLocalVariableDeclarationStatement(statement intmod.ILocalVariableDeclarationStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitNoStatement(statement intmod.INoStatement) {}
func (v *MergeTryWithResourcesStatementVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitReturnStatement(statement intmod.IReturnStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementDefaultLabel(statement intmod.IDefaultLabel) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitTryStatementResource(statement intmod.IResource) {
}
func (v *MergeTryWithResourcesStatementVisitor) VisitTypeDeclarationStatement(statement intmod.ITypeDeclarationStatement) {
}

func (v *MergeTryWithResourcesStatementVisitor) SafeAccept(list intmod.IStatement) {
	if list != nil {
		list.Accept(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, state := range list {
		state.Accept(v)
	}
}

func (v *MergeTryWithResourcesStatementVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, state := range list {
			state.Accept(v)
		}
	}
}
