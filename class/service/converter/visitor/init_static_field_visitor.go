package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	moddec "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	modsts "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvdecl "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewInitStaticFieldVisitor() intsrv.IInitStaticFieldVisitor {
	return &InitStaticFieldVisitor{
		searchFirstLineNumberVisitor:        NewSearchFirstLineNumberVisitor(),
		searchLocalVariableReferenceVisitor: NewSearchLocalVariableReferenceVisitor(),
		fields:                              make(map[string]intmod.IFieldDeclarator),
	}
}

type InitStaticFieldVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor        intsrv.ISearchFirstLineNumberVisitor
	searchLocalVariableReferenceVisitor intsrv.ISearchLocalVariableReferenceVisitor
	internalTypeName                    string
	fields                              map[string]intmod.IFieldDeclarator
	methods                             util.IList[intsrv.IClassFileConstructorOrMethodDeclaration]
	deleteStaticDeclaration             bool
}

func (v *InitStaticFieldVisitor) SetInternalTypeName(internalTypeName string) {
	v.internalTypeName = internalTypeName
}

func (v *InitStaticFieldVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)

	// Store field declarations
	v.fields = make(map[string]intmod.IFieldDeclarator)
	tmp := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.FieldDeclarations() {
		tmp = append(tmp, item)
	}
	v.SafeAcceptListDeclaration(tmp)

	if !(len(v.fields) == 0) {
		methods := util.NewDefaultListWithSlice(bodyDeclaration.MethodDeclarations())

		if methods != nil {
			v.deleteStaticDeclaration = false

			length := methods.Size()
			for i := 0; i < length; i++ {
				methods.Get(i).AcceptDeclaration(v)

				if !v.deleteStaticDeclaration {
					if v.deleteStaticDeclaration {
						methods.RemoveAt(i)
					}
					break
				}
			}
		}
	}

	tmp = make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.InnerTypeDeclarations() {
		tmp = append(tmp, item)
	}
	v.SafeAcceptListDeclaration(tmp)
}

func (v *InitStaticFieldVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.fields[decl.Name()] = decl
}

func (v *InitStaticFieldVisitor) VisitConstructorDeclaration(_ intmod.IConstructorDeclaration) {}

func (v *InitStaticFieldVisitor) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {}

func (v *InitStaticFieldVisitor) VisitInstanceInitializerDeclaration(_ intmod.IInstanceInitializerDeclaration) {
}

func (v *InitStaticFieldVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	sid := decl.(intsrv.IClassFileStaticInitializerDeclaration)
	statements := sid.Statements()

	if statements != nil {
		if statements.IsList() {
			list := statements.ToList()

			// Multiple statements
			if (list.Size() > 0) && v.isAssertionsDisabledStatement(list.First()) {
				// Remove assert initialization statement
				list.RemoveFirst()
			}

			length := list.Size()
			for i := 0; i < length; i++ {
				if v.setStaticFieldInitializer(list.Get(i)) {
					if i > 0 {
						// Split 'static' block
						var newStatements intmod.IStatement

						if i == 1 {
							newStatements = list.RemoveFirst()
						} else {
							subList := list.SubList(0, i)
							newStatements = modsts.NewStatementsWithList(subList)
							subList.Clear()
						}

						// Removes statements from original list
						length -= newStatements.Size()
						i = 0

						v.addStaticInitializerDeclaration(sid, v.getFirstLineNumber(newStatements), newStatements)
					}
					// Remove field initialization statement
					list.RemoveAt(i)
					i--
					length--
				}
			}
		} else {
			// Single statement
			if v.isAssertionsDisabledStatement(statements.First()) {
				// Remove assert initialization statement
				statements = nil
			}
			if (statements != nil) && v.setStaticFieldInitializer(statements.First()) {
				// Remove field initialization statement
				statements = nil
			}
		}

		if (statements == nil) || (statements.Size() == 0) {
			v.deleteStaticDeclaration = true
		} else {
			firstLineNumber := v.getFirstLineNumber(statements)
			if firstLineNumber == -1 {
				sid.SetFirstLineNumber(0)
			} else {
				sid.SetFirstLineNumber(firstLineNumber)
			}
			v.deleteStaticDeclaration = false
		}
	}
}

func (v *InitStaticFieldVisitor) isAssertionsDisabledStatement(state intmod.IStatement) bool {
	expr := state.Expression()

	if expr.LeftExpression().IsFieldReferenceExpression() {
		fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.Type() == _type.PtTypeBoolean &&
			fre.InternalTypeName() == v.internalTypeName &&
			fre.Name() == "$assertionsDisabled" {
			return true
		}
	}

	return false
}

func (v *InitStaticFieldVisitor) setStaticFieldInitializer(state intmod.IStatement) bool {
	expr := state.Expression()

	if expr.LeftExpression().IsFieldReferenceExpression() {
		fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() == v.internalTypeName {
			fdr := v.fields[fre.Name()].(intmod.IFieldDeclarator)

			if fdr != nil && fdr.VariableInitializer() == nil {
				fdn := fdr.FieldDeclaration()

				if fdn.Flags()&intmod.FlagStatic != 0 && fdn.Type().Descriptor() == fre.Descriptor() {
					expr = expr.RightExpression()

					v.searchLocalVariableReferenceVisitor.Init(-1)
					expr.Accept(v.searchLocalVariableReferenceVisitor)

					if !v.searchLocalVariableReferenceVisitor.ContainsReference() {
						fdr.SetVariableInitializer(moddec.NewExpressionVariableInitializer(expr))
						fdr.FieldDeclaration().(intsrv.IClassFileFieldDeclaration).SetFirstLineNumber(expr.LineNumber())
						return true
					}
				}
			}
		}
	}

	return false
}

func (v *InitStaticFieldVisitor) getFirstLineNumber(state intmod.IStatement) int {
	v.searchFirstLineNumberVisitor.Init()
	state.AcceptStatement(v.searchFirstLineNumberVisitor)
	return v.searchFirstLineNumberVisitor.LineNumber()
}

func (v *InitStaticFieldVisitor) addStaticInitializerDeclaration(
	sid intsrv.IClassFileStaticInitializerDeclaration,
	lineNumber int, statements intmod.IStatement) {
	v.methods.Add(srvdecl.NewClassFileStaticInitializerDeclaration2(
		sid.BodyDeclaration(), sid.ClassFile(), sid.Method(), sid.Bindings(),
		sid.TypeBounds(), lineNumber, statements))
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *InitStaticFieldVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *InitStaticFieldVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitStaticFieldVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *InitStaticFieldVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitStaticFieldVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *InitStaticFieldVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitStaticFieldVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *InitStaticFieldVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *InitStaticFieldVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitStaticFieldVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitStaticFieldVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *InitStaticFieldVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *InitStaticFieldVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *InitStaticFieldVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *InitStaticFieldVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *InitStaticFieldVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *InitStaticFieldVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *InitStaticFieldVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *InitStaticFieldVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *InitStaticFieldVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *InitStaticFieldVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *InitStaticFieldVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *InitStaticFieldVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *InitStaticFieldVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *InitStaticFieldVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitStaticFieldVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *InitStaticFieldVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitStaticFieldVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitStaticFieldVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *InitStaticFieldVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *InitStaticFieldVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *InitStaticFieldVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *InitStaticFieldVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *InitStaticFieldVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *InitStaticFieldVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *InitStaticFieldVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *InitStaticFieldVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *InitStaticFieldVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *InitStaticFieldVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *InitStaticFieldVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitStaticFieldVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *InitStaticFieldVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *InitStaticFieldVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *InitStaticFieldVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *InitStaticFieldVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *InitStaticFieldVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *InitStaticFieldVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *InitStaticFieldVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *InitStaticFieldVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *InitStaticFieldVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *InitStaticFieldVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *InitStaticFieldVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
