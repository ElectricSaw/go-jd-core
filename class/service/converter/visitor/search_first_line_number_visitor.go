package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
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

func (v *SearchFirstLineNumberVisitor) VisitStatements(statements *statement.Statements) {
	if v.lineNumber == -1 {

		for _, value := range statements.Statements {
			value.Accept(v)
			if v.lineNumber != -1 {
				break
			}
		}
	}
}

func (v *SearchFirstLineNumberVisitor) VisitArrayExpression(expression *expression.ArrayExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitBinaryOperatorExpression(expression *expression.BinaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitBooleanExpression(expression *expression.BooleanExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitCastExpression(expression *expression.CastExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitCommentExpression(expression *expression.CommentExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitConstructorInvocationExpression(expression *expression.ConstructorInvocationExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitConstructorReferenceExpression(expression *expression.ConstructorReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitDoubleConstantExpression(expression *expression.DoubleConstantExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitEnumConstantReferenceExpression(expression *expression.EnumConstantReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitFieldReferenceExpression(expression *expression.FieldReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitFloatConstantExpression(expression *expression.FloatConstantExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitIntegerConstantExpression(expression *expression.IntegerConstantExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitInstanceOfExpression(expression *expression.InstanceOfExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitLambdaFormalParametersExpression(expression *expression.LambdaFormalParametersExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitLambdaIdentifiersExpression(expression *expression.LambdaIdentifiersExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitLengthExpression(expression *expression.LengthExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitLocalVariableReferenceExpression(expression *expression.LocalVariableReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitLongConstantExpression(expression *expression.LongConstantExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitMethodReferenceExpression(expression *expression.MethodReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitNewArray(expression *expression.NewArray) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitNewExpression(expression *expression.NewExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitNewInitializedArray(expression *expression.NewInitializedArray) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitNullExpression(expression *expression.NullExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitObjectTypeReferenceExpression(expression *expression.ObjectTypeReferenceExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitParenthesesExpression(expression *expression.ParenthesesExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitPostOperatorExpression(expression *expression.PostOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitPreOperatorExpression(expression *expression.PreOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitStringConstantExpression(expression *expression.StringConstantExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitSuperExpression(expression *expression.SuperExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitTernaryOperatorExpression(expression *expression.TernaryOperatorExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitThisExpression(expression *expression.ThisExpression) {
	v.lineNumber = expression.LineNumber()
}
func (v *SearchFirstLineNumberVisitor) VisitTypeReferenceDotClassExpression(expression *expression.TypeReferenceDotClassExpression) {
	v.lineNumber = expression.LineNumber()
}

func (v *SearchFirstLineNumberVisitor) VisitAssertStatement(statement *statement.AssertStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitExpressionStatement(statement *statement.ExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitForEachStatement(statement *statement.ForEachStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfStatement(statement *statement.IfStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitIfElseStatement(statement *statement.IfElseStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitLambdaExpressionStatement(statement *statement.LambdaExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitMethodInvocationExpression(expression *expression.MethodInvocationExpression) {
	expression.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSwitchStatement(statement *statement.SwitchStatement) {
	statement.Condition().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitSynchronizedStatement(statement *statement.SynchronizedStatement) {
	statement.Monitor().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitThrowStatement(statement *statement.ThrowStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitWhileStatement(statement *statement.WhileStatement) {
	if statement.Condition() != nil {
		statement.Condition().Accept(v)
	} else if statement.Statements() != nil {
		statement.Statements().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitDoWhileStatement(statement *statement.DoWhileStatement) {
	if statement.Statements() != nil {
		statement.Statements().Accept(v)
	} else if statement.Condition() != nil {
		statement.Condition().Accept(v)
	}
}

func (v *SearchFirstLineNumberVisitor) VisitForStatement(statement *statement.ForStatement) {
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

func (v *SearchFirstLineNumberVisitor) VisitReturnExpressionStatement(statement *statement.ReturnExpressionStatement) {
	statement.Expression().Accept(v)
}

func (v *SearchFirstLineNumberVisitor) VisitTryStatement(statement *statement.TryStatement) {
	if statement.Resources() != nil {
		v.AcceptListStatement(statement.ResourceList())
	} else {
		statement.TryStatements().Accept(v)
	}
}
