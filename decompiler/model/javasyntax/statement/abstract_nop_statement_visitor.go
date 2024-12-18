package statement

import intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"

type AbstractNopStatementVisitor struct {
}

func (v AbstractNopStatementVisitor) VisitAssertStatement(statement intmod.IAssertStatement)     {}
func (v AbstractNopStatementVisitor) VisitBreakStatement(statement intmod.IBreakStatement)       {}
func (v AbstractNopStatementVisitor) VisitByteCodeStatement(statement intmod.IByteCodeStatement) {}
func (v AbstractNopStatementVisitor) VisitCommentStatement(statement intmod.ICommentStatement)   {}
func (v AbstractNopStatementVisitor) VisitContinueStatement(statement intmod.IContinueStatement) {}
func (v AbstractNopStatementVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement)   {}
func (v AbstractNopStatementVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {}
func (v AbstractNopStatementVisitor) VisitForStatement(statement intmod.IForStatement)         {}
func (v AbstractNopStatementVisitor) VisitIfStatement(statement intmod.IIfStatement)           {}
func (v AbstractNopStatementVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement)   {}
func (v AbstractNopStatementVisitor) VisitLabelStatement(statement intmod.ILabelStatement)     {}
func (v AbstractNopStatementVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitLocalVariableDeclarationStatement(statement intmod.ILocalVariableDeclarationStatement) {
}
func (v AbstractNopStatementVisitor) VisitNoStatement(statement intmod.INoStatement) {}
func (v AbstractNopStatementVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitReturnStatement(statement intmod.IReturnStatement) {}
func (v AbstractNopStatementVisitor) VisitStatements(statement intmod.IStatements)           {}
func (v AbstractNopStatementVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {}
func (v AbstractNopStatementVisitor) VisitSwitchStatementDefaultLabel(statement intmod.IDefaultLabel) {
}
func (v AbstractNopStatementVisitor) VisitSwitchStatementExpressionLabel(statement intmod.IExpressionLabel) {
}
func (v AbstractNopStatementVisitor) VisitSwitchStatementLabelBlock(statement intmod.ILabelBlock) {}
func (v AbstractNopStatementVisitor) VisitSwitchStatementMultiLabelsBlock(statement intmod.IMultiLabelsBlock) {
}
func (v AbstractNopStatementVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
}
func (v AbstractNopStatementVisitor) VisitThrowStatement(statement intmod.IThrowStatement)       {}
func (v AbstractNopStatementVisitor) VisitTryStatement(statement intmod.ITryStatement)           {}
func (v AbstractNopStatementVisitor) VisitTryStatementResource(statement intmod.IResource)       {}
func (v AbstractNopStatementVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {}
func (v AbstractNopStatementVisitor) VisitTypeDeclarationStatement(statement intmod.ITypeDeclarationStatement) {
}
func (v AbstractNopStatementVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {}
