package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	moddecl "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewInitInstanceFieldVisitor() intsrv.IInitInstanceFieldVisitor {
	return &InitInstanceFieldVisitor{
		searchFirstLineNumberVisitor:   NewSearchFirstLineNumberVisitor(),
		fieldDeclarators:               make(map[string]intmod.IFieldDeclarator),
		datas:                          util.NewDefaultList[intsrv.IData](),
		putFields:                      util.NewDefaultList[intmod.IExpression](),
		lineNumber:                     intmod.UnknownLineNumber,
		containsLocalVariableReference: false,
	}
}

type InitInstanceFieldVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor   intsrv.ISearchFirstLineNumberVisitor
	fieldDeclarators               map[string]intmod.IFieldDeclarator
	datas                          util.IList[intsrv.IData]
	putFields                      util.IList[intmod.IExpression]
	lineNumber                     int
	containsLocalVariableReference bool
}

func (v *InitInstanceFieldVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	// Init attributes
	v.fieldDeclarators = make(map[string]intmod.IFieldDeclarator)
	v.datas.Clear()
	v.putFields.Clear()
	// Visit fields
	tmp1 := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.FieldDeclarations() {
		tmp1 = append(tmp1, item)
	}
	v.SafeAcceptListDeclaration(tmp1)
	// Visit methods
	tmp2 := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.MethodDeclarations() {
		tmp2 = append(tmp2, item)
	}
	v.SafeAcceptListDeclaration(tmp2)
	// Init values
	v.updateFieldsAndConstructors()
}

func (v *InitInstanceFieldVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	if (declaration.Flags() & intmod.FlagStatic) == 0 {
		declaration.FieldDeclarators().AcceptDeclaration(v)
	}
}

func (v *InitInstanceFieldVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	cfcd := declaration.(intsrv.IClassFileConstructorDeclaration)

	if (cfcd.Statements() != nil) && cfcd.Statements().IsStatements() {
		statements := cfcd.Statements().(intmod.IStatements)
		iterator := statements.ListIterator()
		superConstructorCall := v.searchSuperConstructorCall(iterator)

		if superConstructorCall != nil {
			internalTypeName := cfcd.ClassFile().InternalTypeName()

			v.datas.Add(NewData(cfcd, statements, iterator.NextIndex()))

			if v.datas.Size() == 1 {
				var firstLineNumber int

				if (cfcd.Flags() & intmod.FlagSynthetic) != 0 {
					firstLineNumber = intmod.UnknownLineNumber
				} else {
					firstLineNumber = superConstructorCall.LineNumber()

					if (superConstructorCall.Descriptor() == "()V") &&
						(firstLineNumber != intmod.UnknownLineNumber) && iterator.HasNext() {
						if (v.lineNumber == intmod.UnknownLineNumber) || (v.lineNumber >= firstLineNumber) {
							v.searchFirstLineNumberVisitor.Init()
							iterator.Next().AcceptStatement(v.searchFirstLineNumberVisitor)
							iterator.Previous()

							ln := v.searchFirstLineNumberVisitor.LineNumber()
							if (ln != intmod.UnknownLineNumber) && (ln >= firstLineNumber) {
								firstLineNumber = intmod.UnknownLineNumber
							}
						}
					}
				}

				v.initPutFields(internalTypeName, firstLineNumber, iterator)
			} else {
				v.filterPutFields(internalTypeName, iterator)
			}
		}
	}
}

func (v *InitInstanceFieldVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.lineNumber = declaration.(intsrv.IClassFileMethodDeclaration).FirstLineNumber()
}

func (v *InitInstanceFieldVisitor) VisitNewExpression(expression intmod.INewExpression) {
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *InitInstanceFieldVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *InitInstanceFieldVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.fieldDeclarators[declaration.Name()] = declaration
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
	v.containsLocalVariableReference = true
}

func (v *InitInstanceFieldVisitor) searchSuperConstructorCall(iterator util.IListIterator[intmod.IStatement]) intmod.ISuperConstructorInvocationExpression {
	for iterator.HasNext() {
		expression := iterator.Next().Expression()

		if expression.IsSuperConstructorInvocationExpression() {
			return expression.(intmod.ISuperConstructorInvocationExpression)
		}

		if expression.IsConstructorInvocationExpression() {
			break
		}
	}

	return nil
}

func (v *InitInstanceFieldVisitor) initPutFields(internalTypeName string, firstLineNumber int, iterator util.IListIterator[intmod.IStatement]) {
	fieldNames := util.NewSet[string]()
	var expression intmod.IExpression

	for iterator.HasNext() {
		statement := iterator.Next()

		if !statement.IsExpressionStatement() {
			break
		}

		expression = statement.Expression()

		if !expression.IsBinaryOperatorExpression() {
			break
		}

		if (expression.Operator() != "=") || !expression.LeftExpression().IsFieldReferenceExpression() {
			break
		}

		fre := expression.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() != internalTypeName || !fre.Expression().IsThisExpression() {
			break
		}

		fieldName := fre.Name()

		if fieldNames.Contains(fieldName) {
			break
		}

		v.containsLocalVariableReference = false
		expression.RightExpression().Accept(v)

		if v.containsLocalVariableReference {
			break
		}

		v.putFields.Add(expression)
		fieldNames.Add(fieldName)
		expression = nil
	}

	var lastLineNumber int

	if expression == nil {
		if firstLineNumber == intmod.UnknownLineNumber {
			lastLineNumber = intmod.UnknownLineNumber
		} else {
			lastLineNumber = firstLineNumber + 1
		}
	} else {
		lastLineNumber = expression.LineNumber()
	}

	if firstLineNumber < lastLineNumber {
		ite := v.putFields.Iterator()

		for ite.HasNext() {
			lineNumber := ite.Next().LineNumber()

			if (firstLineNumber <= lineNumber) && (lastLineNumber <= lastLineNumber) {
				lastLineNumber++
				_ = ite.Remove()
			}
		}
	}
}

func (v *InitInstanceFieldVisitor) filterPutFields(internalTypeName string, iterator util.IListIterator[intmod.IStatement]) {
	putFieldIterator := v.putFields.Iterator()
	index := 0

	for iterator.HasNext() && putFieldIterator.HasNext() {
		expression := iterator.Next().Expression()

		if !expression.IsBinaryOperatorExpression() {
			break
		}

		if expression.Operator() != "=" || !expression.LeftExpression().IsFieldReferenceExpression() {
			break
		}

		fre := expression.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() != internalTypeName {
			break
		}

		putField := putFieldIterator.Next()

		if expression.LineNumber() != putField.LineNumber() {
			break
		}

		if fre.Name() != putField.LeftExpression().Name() {
			break
		}

		index++
	}

	if index < v.putFields.Size() {
		// Cut extra putFields
		v.putFields.SubList(index, v.putFields.Size()).Clear()
	}
}

func (v *InitInstanceFieldVisitor) updateFieldsAndConstructors() {
	count := v.putFields.Size()

	if count > 0 {
		// Init values
		for _, putField := range v.putFields.ToSlice() {
			decl := v.fieldDeclarators[putField.LeftExpression().Name()]

			if decl != nil {
				expression := putField.RightExpression()
				decl.SetVariableInitializer(moddecl.NewExpressionVariableInitializer(expression))
				decl.FieldDeclaration().(intsrv.IClassFileFieldDeclaration).SetFirstLineNumber(expression.LineNumber())
			}
		}

		// Update data : remove init field statements
		for _, data := range v.datas.ToSlice() {
			data.Statements().SubList(data.Index(), data.Index()+count).Clear()

			if data.Statements().IsEmpty() {
				data.Declaration().SetStatements(nil)
				data.Declaration().SetFirstLineNumber(0)
			} else {
				v.searchFirstLineNumberVisitor.Init()
				v.searchFirstLineNumberVisitor.VisitStatements(data.Statements())

				firstLineNumber := v.searchFirstLineNumberVisitor.LineNumber()

				if firstLineNumber == -1 {
					data.Declaration().SetFirstLineNumber(0)
				} else {
					data.Declaration().SetFirstLineNumber(firstLineNumber)
				}
			}
		}
	}
}

func NewData(declaration intsrv.IClassFileConstructorDeclaration, statements intmod.IStatements, index int) intsrv.IData {
	return &Data{
		declaration: declaration,
		statements:  statements,
		index:       index,
	}
}

type Data struct {
	declaration intsrv.IClassFileConstructorDeclaration
	statements  intmod.IStatements
	index       int
}

func (d *Data) Declaration() intsrv.IClassFileConstructorDeclaration {
	return d.declaration
}

func (d *Data) SetDeclaration(declaration intsrv.IClassFileConstructorDeclaration) {
	d.declaration = declaration
}

func (d *Data) Statements() intmod.IStatements {
	return d.statements
}

func (d *Data) SetStatements(statements intmod.IStatements) {
	d.statements = statements
}

func (d *Data) Index() int {
	return d.index
}

func (d *Data) SetIndex(index int) {
	d.index = index
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *InitInstanceFieldVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *InitInstanceFieldVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitInstanceFieldVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitInstanceFieldVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *InitInstanceFieldVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitInstanceFieldVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitInstanceFieldVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *InitInstanceFieldVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *InitInstanceFieldVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *InitInstanceFieldVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *InitInstanceFieldVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *InitInstanceFieldVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *InitInstanceFieldVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *InitInstanceFieldVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *InitInstanceFieldVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *InitInstanceFieldVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *InitInstanceFieldVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *InitInstanceFieldVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *InitInstanceFieldVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *InitInstanceFieldVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *InitInstanceFieldVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitInstanceFieldVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *InitInstanceFieldVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitInstanceFieldVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *InitInstanceFieldVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *InitInstanceFieldVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *InitInstanceFieldVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *InitInstanceFieldVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *InitInstanceFieldVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *InitInstanceFieldVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *InitInstanceFieldVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *InitInstanceFieldVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *InitInstanceFieldVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *InitInstanceFieldVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *InitInstanceFieldVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *InitInstanceFieldVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *InitInstanceFieldVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *InitInstanceFieldVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *InitInstanceFieldVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *InitInstanceFieldVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *InitInstanceFieldVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *InitInstanceFieldVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *InitInstanceFieldVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *InitInstanceFieldVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *InitInstanceFieldVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *InitInstanceFieldVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *InitInstanceFieldVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *InitInstanceFieldVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
