package statement

import intsyn "bitbucket.org/coontec/javaClass/class/interfaces/model"

type AbstractNopStatementVisitor struct {
}

func (v AbstractNopStatementVisitor) VisitAssertStatement(statement intsyn.IAssertStatement)     {}
func (v AbstractNopStatementVisitor) VisitBreakStatement(statement intsyn.IBreakStatement)       {}
func (v AbstractNopStatementVisitor) VisitByteCodeStatement(statement intsyn.IByteCodeStatement) {}
func (v AbstractNopStatementVisitor) VisitCommentStatement(statement intsyn.ICommentStatement)   {}
func (v AbstractNopStatementVisitor) VisitContinueStatement(statement intsyn.IContinueStatement) {}
func (v AbstractNopStatementVisitor) VisitDoWhileStatement(statement intsyn.IDoWhileStatement)   {}
func (v AbstractNopStatementVisitor) VisitExpressionStatement(statement intsyn.IExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitForEachStatement(statement intsyn.IForEachStatement) {}
func (v AbstractNopStatementVisitor) VisitForStatement(statement intsyn.IForStatement)         {}
func (v AbstractNopStatementVisitor) VisitIfStatement(statement intsyn.IIfStatement)           {}
func (v AbstractNopStatementVisitor) VisitIfElseStatement(statement intsyn.IIfElseStatement)   {}
func (v AbstractNopStatementVisitor) VisitLabelStatement(statement intsyn.ILabelStatement)     {}
func (v AbstractNopStatementVisitor) VisitLambdaExpressionStatement(statement intsyn.ILambdaExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitLocalVariableDeclarationStatement(statement intsyn.ILocalVariableDeclarationStatement) {
}
func (v AbstractNopStatementVisitor) VisitNoStatement(statement intsyn.INoStatement) {}
func (v AbstractNopStatementVisitor) VisitReturnExpressionStatement(statement intsyn.IReturnExpressionStatement) {
}
func (v AbstractNopStatementVisitor) VisitReturnStatement(statement intsyn.IReturnStatement) {}
func (v AbstractNopStatementVisitor) VisitStatements(statement intsyn.IStatements)           {}
func (v AbstractNopStatementVisitor) VisitSwitchStatement(statement intsyn.ISwitchStatement) {}
func (v AbstractNopStatementVisitor) VisitSwitchStatementDefaultLabel(statement intsyn.IDefaultLabel) {
}
func (v AbstractNopStatementVisitor) VisitSwitchStatementExpressionLabel(statement intsyn.IExpressionLabel) {
}
func (v AbstractNopStatementVisitor) VisitSwitchStatementLabelBlock(statement intsyn.ILabelBlock) {}
func (v AbstractNopStatementVisitor) VisitSwitchStatementMultiLabelsBlock(statement intsyn.IMultiLabelsBlock) {
}
func (v AbstractNopStatementVisitor) VisitSynchronizedStatement(statement intsyn.ISynchronizedStatement) {
}
func (v AbstractNopStatementVisitor) VisitThrowStatement(statement intsyn.IThrowStatement)       {}
func (v AbstractNopStatementVisitor) VisitTryStatement(statement intsyn.ITryStatement)           {}
func (v AbstractNopStatementVisitor) VisitTryStatementResource(statement intsyn.IResource)       {}
func (v AbstractNopStatementVisitor) VisitTryStatementCatchClause(statement intsyn.ICatchClause) {}
func (v AbstractNopStatementVisitor) VisitTypeDeclarationStatement(statement intsyn.ITypeDeclarationStatement) {
}
func (v AbstractNopStatementVisitor) VisitWhileStatement(statement intsyn.IWhileStatement) {}
