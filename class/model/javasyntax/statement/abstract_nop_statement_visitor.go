package statement

type AbstractNopStatementVisitor struct {
}

func (v *AbstractNopStatementVisitor) VisitAssertStatement(statement *AssertStatement)         {}
func (v *AbstractNopStatementVisitor) VisitBreakStatement(statement *BreakStatement)           {}
func (v *AbstractNopStatementVisitor) VisitByteCodeStatement(statement *ByteCodeStatement)     {}
func (v *AbstractNopStatementVisitor) VisitCommentStatement(statement *CommentStatement)       {}
func (v *AbstractNopStatementVisitor) VisitContinueStatement(statement *ContinueStatement)     {}
func (v *AbstractNopStatementVisitor) VisitDoWhileStatement(statement *DoWhileStatement)       {}
func (v *AbstractNopStatementVisitor) VisitExpressionStatement(statement *ExpressionStatement) {}
func (v *AbstractNopStatementVisitor) VisitForEachStatement(statement *ForEachStatement)       {}
func (v *AbstractNopStatementVisitor) VisitForStatement(statement *ForStatement)               {}
func (v *AbstractNopStatementVisitor) VisitIfStatement(statement *IfStatement)                 {}
func (v *AbstractNopStatementVisitor) VisitIfElseStatement(statement *IfElseStatement)         {}
func (v *AbstractNopStatementVisitor) VisitLabelStatement(statement *LabelStatement)           {}
func (v *AbstractNopStatementVisitor) VisitLambdaExpressionStatement(statement *LambdaExpressionStatement) {
}
func (v *AbstractNopStatementVisitor) VisitLocalVariableDeclarationStatement(statement *LocalVariableDeclarationStatement) {
}
func (v *AbstractNopStatementVisitor) VisitNoStatement(statement *NoStatement) {}
func (v *AbstractNopStatementVisitor) VisitReturnExpressionStatement(statement *ReturnExpressionStatement) {
}
func (v *AbstractNopStatementVisitor) VisitReturnStatement(statement *ReturnStatement)          {}
func (v *AbstractNopStatementVisitor) VisitStatements(statement *Statements)                    {}
func (v *AbstractNopStatementVisitor) VisitSwitchStatement(statement *SwitchStatement)          {}
func (v *AbstractNopStatementVisitor) VisitSwitchStatementDefaultLabel(statement *DefaultLabe1) {}
func (v *AbstractNopStatementVisitor) VisitSwitchStatementExpressionLabel(statement *ExpressionLabel) {
}
func (v *AbstractNopStatementVisitor) VisitSwitchStatementLabelBlock(statement *LabelBlock) {}
func (v *AbstractNopStatementVisitor) VisitSwitchStatementMultiLabelsBlock(statement *MultiLabelsBlock) {
}
func (v *AbstractNopStatementVisitor) VisitSynchronizedStatement(statement *SynchronizedStatement) {}
func (v *AbstractNopStatementVisitor) VisitThrowStatement(statement *ThrowStatement)               {}
func (v *AbstractNopStatementVisitor) VisitTryStatement(statement *TryStatement)                   {}
func (v *AbstractNopStatementVisitor) VisitTryStatementResource(statement *Resource)               {}
func (v *AbstractNopStatementVisitor) VisitTryStatementCatchClause(statement *CatchClause)         {}
func (v *AbstractNopStatementVisitor) VisitTypeDeclarationStatement(statement *TypeDeclarationStatement) {
}
func (v *AbstractNopStatementVisitor) VisitWhileStatement(statement *WhileStatement) {}
