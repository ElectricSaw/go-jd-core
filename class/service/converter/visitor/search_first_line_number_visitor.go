package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
)

func NewSearchFirstLineNumberVisitor() *SearchFirstLineNumberVisitor {
	return &SearchFirstLineNumberVisitor{
		lineNumber: -1,
	}
}

type SearchFirstLineNumberVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	lineNumber int
}

func (v *SearchFirstLineNumberVisitor) Init() {
	v.lineNumber = -1
}

func (v *SearchFirstLineNumberVisitor) LineNumber() int {
	return v.lineNumber
}

func (v *SearchFirstLineNumberVisitor) VisitStatements(statements intmod.IStatements) {
	if v.lineNumber == -1 {

		for _, value := range statements.Elements() {
			value.Accept(v)
			if v.lineNumber != -1 {
				break
			}
		}
	}
}

func (v *SearchFirstLineNumberVisitor) VisitArrayExpression(expression intmod.IArrayExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitBooleanExpression(expression intmod.IBooleanExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitCastExpression(expression intmod.ICastExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitCommentExpression(expression intmod.ICommentExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitConstructorReferenceExpression(expression intmod.IConstructorReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitDoubleConstantExpression(expression intmod.IDoubleConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitEnumConstantReferenceExpression(expression intmod.IEnumConstantReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitFloatConstantExpression(expression intmod.IFloatConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaIdentifiersExpression(expression intmod.ILambdaIdentifiersExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitLongConstantExpression(expression intmod.ILongConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewArray(expression intmod.INewArray) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewExpression(expression intmod.INewExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitNullExpression(expression intmod.INullExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitObjectTypeReferenceExpression(expression intmod.IObjectTypeReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitStringConstantExpression(expression intmod.IStringConstantExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitSuperExpression(expression intmod.ISuperExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitThisExpression(expression intmod.IThisExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitTypeReferenceDotClassExpression(expression intmod.ITypeReferenceDotClassExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitAssertStatement(statement intmod.IAssertStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitExpressionStatement(statement intmod.IExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitForEachStatement(statement intmod.IForEachStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfStatement(statement intmod.IIfStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfElseStatement(statement intmod.IIfElseStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaExpressionStatement(statement intmod.ILambdaExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	expression.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatement(statement intmod.ISwitchStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSynchronizedStatement(statement intmod.ISynchronizedStatement) {
	statement.Monitor().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitWhileStatement(statement intmod.IWhileStatement) {
	if statement.Condition() != nil {
		statement.Condition().Accept(v)
	} else if statement.Statements() != nil {
		statement.Statements().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitDoWhileStatement(statement intmod.IDoWhileStatement) {
	if statement.Statements() != nil {
		statement.Statements().Accept(v)
	} else if statement.Condition() != nil {
		statement.Condition().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitForStatement(statement intmod.IForStatement) {
	if statement.Init() != nil {
		statement.Init().Accept(v)
	} else if statement.Condition() != nil {
		statement.Condition().Accept(v)
	} else if statement.Update() != nil {
		statement.Update().Accept(v)
	} else if statement.Statements() != nil {
		statement.Statements().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTryStatement(statement intmod.ITryStatement) {
	if statement.Resources() != nil {
		v.AcceptListStatement(statement.ResourceList())
	} else {
		statement.TryStatements().Accept(v)
	}
}
