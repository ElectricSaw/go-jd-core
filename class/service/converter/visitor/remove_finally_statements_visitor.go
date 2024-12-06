package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewRemoveFinallyStatementsVisitor(localVariableMaker intsrv.ILocalVariableMaker) *RemoveFinallyStatementsVisitor {
	return &RemoveFinallyStatementsVisitor{
		localVariableMaker: localVariableMaker,
	}
}

type RemoveFinallyStatementsVisitor struct {
	declaredSyntheticLocalVariableVisitor *DeclaredSyntheticLocalVariableVisitor
	localVariableMaker                    intsrv.ILocalVariableMaker
	statementCountInFinally               int
	statementCountToRemove                int
	lastFinallyStatementIsATryStatement   bool
}

func (v *RemoveFinallyStatementsVisitor) Init() {
	v.statementCountInFinally = 0
	v.statementCountToRemove = 0
	v.lastFinallyStatementIsATryStatement = false
}

func (v *RemoveFinallyStatementsVisitor) VisitStatements(statements intmod.IStatements) {
	stmts := statements.ToList()
	size := statements.Size()

	if size > 0 {
		i := size
		oldStatementCountToRemove := v.statementCountToRemove

		// Check last statement
		lastStatement := stmts.Last()

		if lastStatement.IsReturnExpressionStatement() || lastStatement.IsReturnStatement() {
			v.statementCountToRemove = v.statementCountInFinally
			i--
		} else if lastStatement.IsThrowStatement() {
			v.statementCountToRemove = 0
			i--
		} else if lastStatement.IsContinueStatement() || lastStatement.IsBreakStatement() {
			i--
		} else {
			whileStatement := getInfiniteWhileStatement(lastStatement)

			if whileStatement != nil {
				// Infinite loop => Do not remove any statement
				v.statementCountToRemove = 0
				i--
				whileStatement.Statements().AcceptStatement(v)
			}
		}

		// Remove 'finally' statements
		if v.statementCountToRemove > 0 {
			if !v.lastFinallyStatementIsATryStatement && (i > 0) && stmts.Get(i-1).IsTryStatement() {
				stmts.Get(i - 1).AcceptStatement(v)
				v.statementCountToRemove = 0
				i--
			} else {
				v.declaredSyntheticLocalVariableVisitor.Init()

				// Remove 'finally' statements
				if i > v.statementCountToRemove {
					list := statements.SubList(i-v.statementCountToRemove, i)

					for _, statement := range list.ToSlice() {
						statement.AcceptStatement(v.declaredSyntheticLocalVariableVisitor)
					}

					lastStatement.AcceptStatement(v.declaredSyntheticLocalVariableVisitor)
					list.Clear()
					i -= v.statementCountToRemove
					v.statementCountToRemove = 0
				} else {
					list := statements.ToList()

					for _, statement := range list.ToSlice() {
						statement.AcceptStatement(v.declaredSyntheticLocalVariableVisitor)
					}

					list.Clear()
					if i < size {
						list.Add(lastStatement)
					}
					v.statementCountToRemove -= i
					i = 0
				}
			}
		}

		// Recursive visit
		for ; i > 0; i++ {
			stmts.Get(i).AcceptStatement(v)

			if v.statementCountToRemove > 0 {
				if (i + v.statementCountToRemove) < statements.Size() {
					statements.SubList(i+1, i+1+v.statementCountToRemove).Clear()
					v.statementCountToRemove = 0
				} else {
					//assert false : "Error. Eclipse try-finally ?";
				}
			}
		}

		v.statementCountToRemove = oldStatementCountToRemove
	}
}

func getInfiniteWhileStatement(statement intmod.IStatement) intmod.IWhileStatement {
	if statement.IsLabelStatement() {
		statement = statement.(intmod.ILabelStatement).Statement()
	}

	if (statement == nil) || !statement.IsWhileStatement() {
		return nil
	}

	if !statement.Condition().IsBooleanExpression() {
		return nil
	}

	booleanExpression := statement.Condition().(intmod.IBooleanExpression)

	if booleanExpression.IsFalse() {
		return nil
	}

	return statement.(intmod.IWhileStatement)
}

func (v *RemoveFinallyStatementsVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Statements().AcceptStatement(v)
	statement.ElseStatements().AcceptStatement(v)
}

func (v *RemoveFinallyStatementsVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	for _, block := range statement.Blocks() {
		block.Statements().AcceptStatement(v)
	}
}

func (v *RemoveFinallyStatementsVisitor) VisitTryStatement(statement intmod.ITryStatement) {
	oldLastFinallyStatementIsTryStatement := v.lastFinallyStatementIsATryStatement
	ts := statement.(intsrv.IClassFileTryStatement)
	tryStatements := ts.TryStatements().(intmod.IStatements)
	finallyStatements := ts.FinallyStatements().(intmod.IStatements)

	if finallyStatements != nil {
		switch finallyStatements.Size() {
		case 0:
			break
		case 1:
			finallyStatements.First().AcceptStatement(v)
		default:
			for _, stmt := range finallyStatements.ToSlice() {
				stmt.AcceptStatement(v)
			}
		}

		if v.statementCountInFinally == 0 && finallyStatements.Size() > 0 {
			v.lastFinallyStatementIsATryStatement = finallyStatements.Last().IsTryStatement()
		}
	}

	if ts.IsJsr() || (finallyStatements == nil) || (finallyStatements.Size() == 0) {
		tryStatements.Accept(v)
		tmp := make([]intmod.IStatement, 0)
		for _, stmt := range statement.CatchClauses() {
			tmp = append(tmp, stmt)
		}
		v.safeAcceptListStatement(tmp)
	} else if ts.IsEclipse() {
		catchClauses := util.NewDefaultListWithSlice(statement.CatchClauses())
		oldStatementCountInFinally := v.statementCountInFinally
		finallyStatementsSize := finallyStatements.Size()

		v.statementCountInFinally += finallyStatementsSize

		tryStatements.Accept(v)

		v.statementCountToRemove = finallyStatementsSize

		if catchClauses != nil {
			for _, cc := range catchClauses.ToSlice() {
				cc.Statements().AcceptStatement(v)
			}
		}

		v.statementCountInFinally = oldStatementCountInFinally

		if statement.Resources() != nil {
			ts.SetFinallyStatements(nil)
		}
	} else {
		catchClauses := util.NewDefaultListWithSlice(statement.CatchClauses())
		oldStatementCountInFinally := v.statementCountInFinally
		oldStatementCountToRemove := v.statementCountToRemove
		finallyStatementsSize := finallyStatements.Size()

		v.statementCountInFinally += finallyStatementsSize
		v.statementCountToRemove += finallyStatementsSize

		tryStatements.Accept(v)

		if catchClauses != nil {
			for _, cc := range catchClauses.ToSlice() {
				cc.Statements().AcceptStatement(v)
			}
		}

		v.statementCountInFinally = oldStatementCountInFinally
		v.statementCountToRemove = oldStatementCountToRemove

		if statement.Resources() != nil {
			ts.SetFinallyStatements(nil)
		}
	}

	v.lastFinallyStatementIsATryStatement = oldLastFinallyStatementIsTryStatement
}

func (v *RemoveFinallyStatementsVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitForStatement(statement intmod.IForStatement) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitTryStatementCatchClause(statement intmod.ICatchClause) {
	v.safeAccept(statement.Statements())
}
func (v *RemoveFinallyStatementsVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	v.safeAccept(statement.Statements())
}

func (v *RemoveFinallyStatementsVisitor) VisitSwitchStatementLabelBlock(statement intmod.ILabelBlock) {
	statement.Statements().AcceptStatement(v)
}
func (v *RemoveFinallyStatementsVisitor) VisitSwitchStatementMultiLabelsBlock(statement intmod.IMultiLabelsBlock) {
	statement.Statements().AcceptStatement(v)
}

func (v *RemoveFinallyStatementsVisitor) VisitAssertStatement(_ intmod.IAssertStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitExpressionStatement(_ intmod.IExpressionStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitLabelStatement(_ intmod.ILabelStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitLambdaExpressionStatement(_ intmod.ILambdaExpressionStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitLocalVariableDeclarationStatement(_ intmod.ILocalVariableDeclarationStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitNoStatement(_ intmod.INoStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitReturnExpressionStatement(_ intmod.IReturnExpressionStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
}

func (v *RemoveFinallyStatementsVisitor) VisitSwitchStatementExpressionLabel(_ intmod.IExpressionLabel) {
}

func (v *RemoveFinallyStatementsVisitor) VisitThrowStatement(_ intmod.IThrowStatement) {
}

func (v *RemoveFinallyStatementsVisitor) VisitTryStatementResource(_ intmod.IResource) {
}

func (v *RemoveFinallyStatementsVisitor) VisitTypeDeclarationStatement(_ intmod.ITypeDeclarationStatement) {
}

func (v *RemoveFinallyStatementsVisitor) safeAccept(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *RemoveFinallyStatementsVisitor) safeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, statement := range list {
			statement.AcceptStatement(v)
		}
	}
}
