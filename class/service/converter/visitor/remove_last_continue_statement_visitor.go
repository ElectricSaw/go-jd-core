package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	modsts "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"reflect"
)

func NewRemoveLastContinueStatementVisitor() *RemoveLastContinueStatementVisitor {
	return &RemoveLastContinueStatementVisitor{}
}

type RemoveLastContinueStatementVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *RemoveLastContinueStatementVisitor) VisitStatements(list intmod.IStatements) {
	if !list.IsEmpty() {
		last := list.Last()

		if reflect.TypeOf(last) == reflect.TypeOf(modsts.ContinueStatement{}) {
			list.RemoveLast()
			v.VisitStatements(list)
		} else {
			last.Accept(v)
		}
	}
}

func (v *RemoveLastContinueStatementVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	v.SafeAcceptStatement(state.Statements())
	state.ElseStatements().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatement(state intmod.ITryStatement) {
	state.TryStatements().Accept(v)
	tmp := make([]intmod.IStatement, 0)
	for _, block := range state.CatchClauses() {
		tmp = append(tmp, block)
	}
	v.SafeAcceptListStatement(tmp)
	v.SafeAcceptStatement(state.FinallyStatements())
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement) {
	tmp := make([]intmod.IStatement, 0)
	for _, block := range state.Blocks() {
		tmp = append(tmp, block)
	}
	v.AcceptListStatement(tmp)
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementLabelBlock(state intmod.ILabelBlock) {
	state.Statements().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementMultiLabelsBlock(state intmod.IMultiLabelsBlock) {
	state.Statements().Accept(v)
}

func (v *RemoveLastContinueStatementVisitor) VisitIfStatement(state intmod.IIfStatement) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitSynchronizedStatement(state intmod.ISynchronizedStatement) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatementCatchClause(state intmod.ICatchClause) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *RemoveLastContinueStatementVisitor) VisitDoWhileStatement(_ intmod.IDoWhileStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitForEachStatement(_ intmod.IForEachStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitForStatement(_ intmod.IForStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitWhileStatement(_ intmod.IWhileStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitAssertStatement(_ intmod.IAssertStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitExpressionStatement(_ intmod.IExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLabelStatement(_ intmod.ILabelStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLambdaExpressionStatement(_ intmod.ILambdaExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitLocalVariableDeclarationStatement(_ intmod.ILocalVariableDeclarationStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitReturnExpressionStatement(_ intmod.IReturnExpressionStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitSwitchStatementExpressionLabel(_ intmod.IExpressionLabel) {
}

func (v *RemoveLastContinueStatementVisitor) VisitThrowStatement(_ intmod.IThrowStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitTypeDeclarationStatement(_ intmod.ITypeDeclarationStatement) {
}

func (v *RemoveLastContinueStatementVisitor) VisitTryStatementResource(_ intmod.IResource) {
}
