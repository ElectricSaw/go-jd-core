package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewSingleLineStatementVisitor() *SingleLineStatementVisitor {
	return &SingleLineStatementVisitor{
		minLineNumber:  -2,
		maxLineNumber:  -1,
		statementCount: 0,
	}
}

type SingleLineStatementVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	minLineNumber  int
	maxLineNumber  int
	statementCount int
}

func (v *SingleLineStatementVisitor) Init() {
	v.minLineNumber = -2
	v.maxLineNumber = -1
	v.statementCount = 0
}

func (v *SingleLineStatementVisitor) IsSingleLineStatement() bool {
	if v.minLineNumber <= 0 {
		// Line numbers are unknown
		return v.statementCount <= 1
	} else {
		return v.minLineNumber == v.maxLineNumber
	}
}

// -- Statement -- //
func (v *SingleLineStatementVisitor) VisitAssertStatement(state intmod.IAssertStatement) {
	state.Condition().Accept(v)
	v.SafeAcceptExpression(state.Message())
	v.minLineNumber = state.Condition().LineNumber()
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitDoWhileStatement(state intmod.IDoWhileStatement) {
	v.minLineNumber = state.Condition().LineNumber()
	v.SafeAcceptStatement(state.Statements())
	state.Condition().Accept(v)
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitExpressionStatement(state intmod.IExpressionStatement) {
	state.Expression().Accept(v)
	v.minLineNumber = state.Expression().LineNumber()
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitForEachStatement(state intmod.IForEachStatement) {
	state.Expression().Accept(v)
	v.SafeAcceptStatement(state.Statements())
	v.minLineNumber = state.Expression().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitForStatement(state intmod.IForStatement) {
	if state.Statements() != nil {
		state.Statements().AcceptStatement(v)
	} else if state.Update() != nil {
		state.Update().Accept(v)
	} else if state.Condition() != nil {
		state.Condition().Accept(v)
	} else if state.Init() != nil {
		state.Init().Accept(v)
	} else {
		v.maxLineNumber = 0
	}

	if state.Condition() != nil {
		v.minLineNumber = state.Condition().LineNumber()
	} else if state.Condition() != nil {
		v.minLineNumber = state.Condition().LineNumber()
	} else {
		v.minLineNumber = v.maxLineNumber
	}

	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitIfStatement(state intmod.IIfStatement) {
	state.Condition().Accept(v)
	v.SafeAcceptStatement(state.Statements())
	v.minLineNumber = state.Condition().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	state.Condition().Accept(v)
	state.ElseStatements().AcceptStatement(v)
	v.minLineNumber = state.Condition().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitLabelStatement(state intmod.ILabelStatement) {
	v.minLineNumber = 0
	v.maxLineNumber = 0
	v.SafeAcceptStatement(state.Statement())
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitLambdaExpressionStatement(state intmod.ILambdaExpressionStatement) {
	state.Expression().Accept(v)
	v.minLineNumber = state.Expression().LineNumber()
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitLocalVariableDeclarationStatement(state intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(state.(intmod.ILocalVariableDeclaration))
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitReturnExpressionStatement(state intmod.IReturnExpressionStatement) {
	state.Expression().Accept(v)
	v.minLineNumber = state.Expression().LineNumber()
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitSwitchStatement(state intmod.ISwitchStatement) {
	state.Condition().Accept(v)
	list := util.NewDefaultList[intmod.IStatement]()
	for _, block := range state.Blocks() {
		list.Add(block.(intmod.IStatement))
	}
	v.AcceptListStatement(list)
	v.minLineNumber = state.Condition().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitSynchronizedStatement(state intmod.ISynchronizedStatement) {
	state.Monitor().Accept(v)
	v.SafeAcceptStatement(state.Statements())
	v.minLineNumber = state.Monitor().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitThrowStatement(state intmod.IThrowStatement) {
	state.Expression().Accept(v)
	v.minLineNumber = state.Expression().LineNumber()
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitTryStatement(state intmod.ITryStatement) {
	state.TryStatements().AcceptStatement(v)

	minimum := v.minLineNumber

	list := util.NewDefaultList[intmod.IStatement]()
	for _, block := range state.CatchClauses() {
		list.Add(block.(intmod.IStatement))
	}
	v.SafeAcceptListStatement(list)
	v.SafeAcceptStatement(state.FinallyStatements())
	v.minLineNumber = minimum
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) VisitTryStatementCatchClause(state intmod.ICatchClause) {
	v.SafeAcceptStatement(state.Statements())
}

func (v *SingleLineStatementVisitor) VisitTypeDeclarationStatement(_ intmod.ITypeDeclarationStatement) {
	v.minLineNumber = 0
	v.maxLineNumber = 0
	v.statementCount = 1
}

func (v *SingleLineStatementVisitor) VisitWhileStatement(state intmod.IWhileStatement) {
	state.Condition().Accept(v)
	v.SafeAcceptStatement(state.Statements())
	v.minLineNumber = state.Condition().LineNumber()
	v.statementCount = 2
}

func (v *SingleLineStatementVisitor) AcceptListStatement(list util.IList[intmod.IStatement]) {
	size := list.Size()

	switch size {
	case 0:
		v.minLineNumber = 0
		v.maxLineNumber = 0
		break
	case 1:
		list.Get(0).AcceptStatement(v)
		break
	default:
		list.Get(0).AcceptStatement(v)
		minimum := v.minLineNumber
		list.Get(size - 1).AcceptStatement(v)
		v.minLineNumber = minimum
		v.statementCount = size
		break
	}
}

func (v *SingleLineStatementVisitor) SafeAcceptListStatement(list util.IList[intmod.IStatement]) {
	if list == nil {
		v.minLineNumber = 0
		v.maxLineNumber = 0
	} else {
		v.AcceptListStatement(list)
	}
}

// -- Expression -- //
func (v *SingleLineStatementVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	parameters := expr.Parameters()

	if parameters == nil {
		v.maxLineNumber = expr.LineNumber()
	} else {
		parameters.Accept(v)
	}
}

func (v *SingleLineStatementVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	parameters := expr.Parameters()

	if parameters == nil {
		v.maxLineNumber = expr.LineNumber()
	} else {
		parameters.Accept(v)
	}
}

func (v *SingleLineStatementVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	parameters := expr.Parameters()

	if parameters == nil {
		v.maxLineNumber = expr.LineNumber()
	} else {
		parameters.Accept(v)
	}
}

func (v *SingleLineStatementVisitor) VisitNewArray(expr intmod.INewArray) {
	dimensionExpressionList := expr.DimensionExpressionList()

	if dimensionExpressionList == nil {
		v.maxLineNumber = expr.LineNumber()
	} else {
		dimensionExpressionList.Accept(v)
	}
}

func (v *SingleLineStatementVisitor) VisitNewExpression(expr intmod.INewExpression) {
	bodyDeclaration := expr.BodyDeclaration()

	if bodyDeclaration == nil {
		parameters := expr.Parameters()

		if parameters == nil {
			v.maxLineNumber = expr.LineNumber()
		} else {
			parameters.Accept(v)
		}
	} else {
		v.maxLineNumber = expr.LineNumber() + 1
	}
}

func (v *SingleLineStatementVisitor) AcceptListExpression(list util.IList[intmod.IExpression]) {
	size := list.Size()

	if size == 0 {
		v.maxLineNumber = 0
	} else {
		list.Get(size - 1).Accept(v)
	}
}

func (v *SingleLineStatementVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	expr.Index().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.RightExpression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	expr.Expression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	expr.Statements().AcceptStatement(v)
}

func (v *SingleLineStatementVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *SingleLineStatementVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	expr.ArrayInitializer().AcceptDeclaration(v)
}

func (v *SingleLineStatementVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.FalseExpression().Accept(v)
}

func (v *SingleLineStatementVisitor) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitNullExpression(expr intmod.INullExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	v.maxLineNumber = expr.LineNumber()
}

func (v *SingleLineStatementVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	v.maxLineNumber = expr.LineNumber()
}
