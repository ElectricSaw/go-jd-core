package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	modsts "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	srvsts "github.com/ElectricSaw/go-jd-core/decompiler/service/converter/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func MakeTryWithResourcesStatementMaker(localVariableMaker intsrv.ILocalVariableMaker,
	statements intmod.IStatements, tryStatements intmod.IStatements,
	catchClauses util.IList[intmod.ICatchClause], finallyStatements intmod.IStatements) intmod.IStatement {
	size := statements.Size()

	if size < 2 || finallyStatements == nil ||
		finallyStatements.Size() != 1 ||
		!checkThrowable(catchClauses.ToSlice()) {
		return nil
	}

	state := statements.Get(size - 2)
	if !state.IsExpressionStatement() {
		return nil
	}

	expr := state.Expression()
	if !expr.IsBinaryOperatorExpression() {
		return nil
	}

	boe := expr
	expr = boe.LeftExpression()

	if !expr.IsLocalVariableReferenceExpression() {
		return nil
	}

	lv1 := ConvertTo(expr)
	state = statements.Get(size - 1)
	if !state.IsExpressionStatement() {
		return nil
	}

	expr = state.Expression()
	if !expr.IsBinaryOperatorExpression() {
		return nil
	}

	expr = expr.LeftExpression()
	if !expr.IsLocalVariableReferenceExpression() {
		return nil
	}

	lv2 := ConvertTo(expr)
	state = finallyStatements.First()
	if state.IsIfStatement() && (lv1 == getLocalVariable(state.Condition())) {
		state = state.Statements().First()

		if state.IsIfElseStatement() {
			return parsePatternAddSuppressed(localVariableMaker,
				statements, tryStatements, finallyStatements, boe, lv1, lv2, state)
		}
		if state.IsExpressionStatement() {
			return parsePatternCloseResource(localVariableMaker,
				statements, tryStatements, finallyStatements, boe, lv1, lv2, state)
		}
	}

	if state.IsExpressionStatement() {
		return parsePatternCloseResource(localVariableMaker,
			statements, tryStatements, finallyStatements, boe, lv1, lv2, state)
	}

	return nil
}

func parsePatternAddSuppressed(localVariableMaker intsrv.ILocalVariableMaker,
	statements intmod.IStatements, tryStatements intmod.IStatements,
	finallyStatements intmod.IStatements, boe intmod.IExpression,
	lv1, lv2 intsrv.ILocalVariable, statement intmod.IStatement) intmod.IStatement {
	if !statement.IsIfElseStatement() {
		return nil
	}

	ies := statement
	statement = ies.Statements().First()

	if !statement.IsTryStatement() {
		return nil
	}

	ts := statement
	statement = ies.ElseStatements().First()
	if !statement.IsExpressionStatement() {
		return nil
	}

	expression := statement.Expression()
	if !expression.IsMethodInvocationExpression() {
		return nil
	}

	mie := expression.(intmod.IMethodInvocationExpression)
	if ts.FinallyStatements() != nil || lv2 != getLocalVariable(ies.Condition()) ||
		!checkThrowable(ts.CatchClauses()) || !checkCloseInvocation(mie, lv1) {
		return nil
	}

	statement = ts.TryStatements().First()
	if !statement.IsExpressionStatement() {
		return nil
	}

	expression = statement.Expression()
	if !expression.IsMethodInvocationExpression() {
		return nil
	}

	mie = expression.(intmod.IMethodInvocationExpression)

	if !checkCloseInvocation(mie, lv1) {
		return nil
	}

	statement = ts.CatchClauses()[0].Statements().First()
	if !statement.IsExpressionStatement() {
		return nil
	}

	expression = statement.Expression()

	if !expression.IsMethodInvocationExpression() {
		return nil
	}

	mie = expression.(intmod.IMethodInvocationExpression)

	if mie.Name() != "addSuppressed" || mie.Descriptor() != "(Ljava/lang/Throwable;)V" {
		return nil
	}

	expression = mie.Expression()

	if !expression.IsLocalVariableReferenceExpression() {
		return nil
	}
	if ConvertTo(expression) != lv2 {
		return nil
	}

	return newTryStatement(localVariableMaker, statements, tryStatements, finallyStatements, boe, lv1, lv2)
}

func checkThrowable(catchClauses []intmod.ICatchClause) bool {
	return len(catchClauses) == 1 && catchClauses[0].Type() == _type.OtTypeThrowable
}

func getLocalVariable(condition intmod.IExpression) intsrv.ILocalVariable {
	if !condition.IsBinaryOperatorExpression() {
		return nil
	}

	if condition.Operator() != "!=" ||
		!condition.RightExpression().IsNullExpression() ||
		!condition.LeftExpression().IsLocalVariableReferenceExpression() {
		return nil
	}

	return ConvertTo(condition.LeftExpression())
}

func checkCloseInvocation(mie intmod.IMethodInvocationExpression, lv intsrv.ILocalVariable) bool {
	if mie.Name() == "close" && mie.Descriptor() == "()V" {
		expression := mie.Expression()
		if expression.IsLocalVariableReferenceExpression() {
			return ConvertTo(expression) == lv
		}
	}

	return false
}

func parsePatternCloseResource(localVariableMaker intsrv.ILocalVariableMaker,
	statements intmod.IStatements, tryStatements intmod.IStatements,
	finallyStatements intmod.IStatements, boe intmod.IExpression,
	lv1, lv2 intsrv.ILocalVariable, statement intmod.IStatement) intmod.IStatement {
	expression := statement.Expression()
	if !expression.IsMethodInvocationExpression() {
		return nil
	}

	mie := expression.(intmod.IMethodInvocationExpression)

	if mie.Name() != "$closeResource" || mie.Descriptor() != "(Ljava/lang/Throwable;Ljava/lang/AutoCloseable;)V" {
		return nil
	}

	parameters := mie.Parameters().ToList()
	parameter0 := parameters.First()

	if !parameter0.IsLocalVariableReferenceExpression() {
		return nil
	}
	if ConvertTo(parameter0) != lv2 {
		return nil
	}

	parameter1 := parameters.Get(1)
	if !parameter1.IsLocalVariableReferenceExpression() {
		return nil
	}

	if ConvertTo(parameter1) != lv1 {
		return nil
	}

	return newTryStatement(localVariableMaker, statements, tryStatements, finallyStatements, boe, lv1, lv2)
}

func newTryStatement(localVariableMaker intsrv.ILocalVariableMaker,
	statements intmod.IStatements, tryStatements intmod.IStatements,
	finallyStatements intmod.IStatements, boe intmod.IExpression, lv1, lv2 intsrv.ILocalVariable) intsrv.IClassFileTryStatement {

	// Remove resource & synthetic local variables
	statements.RemoveLast()
	statements.RemoveLast()
	lv1.SetDeclared(true)
	localVariableMaker.RemoveLocalVariable(lv2)

	// Create try-with-resources statement
	resources := util.NewDefaultList[intmod.IResource]()
	resources.Add(modsts.NewResource(lv1.Type().(intmod.IObjectType), lv1.Name(), boe.RightExpression()))

	return srvsts.NewClassFileTryStatement2(resources.ToSlice(),
		tryStatements, nil, finallyStatements, false, false)
}
