package visitor

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/fragmenter/visitor/fragutil"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"strings"
)

func NewSearchImportsVisitor(loader api.Loader, mainInternalName string) *SearchImportsVisitor {
	v := &SearchImportsVisitor{
		loader:            loader,
		importsFragment:   fragutil.NewImportsFragment(),
		localTypeNames:    util.NewSet[string](),
		internalTypeNames: util.NewSet[string](),
		importTypeNames:   util.NewSet[string](),
	}

	index := strings.LastIndex(mainInternalName, "/")
	if index == -1 {
		v.internalPackagePrefix = ""
	} else {
		v.internalPackagePrefix = mainInternalName[:index+1]
	}

	return v
}

type SearchImportsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	loader                api.Loader
	internalPackagePrefix string
	importsFragment       intmod.IImportsFragment
	maxLineNumber         int
	localTypeNames        util.ISet[string]
	internalTypeNames     util.ISet[string]
	importTypeNames       util.ISet[string]
}

func (v *SearchImportsVisitor) ImportsFragment() intmod.IImportsFragment {
	v.importsFragment.InitLineCounts()
	return v.importsFragment
}

func (v *SearchImportsVisitor) MaxLineNumber() int {
	return v.maxLineNumber
}

func (v *SearchImportsVisitor) VisitCompilationUnit(compilationUnit *javasyntax.CompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(NewTypeVisitor2(v.localTypeNames))
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

func (v *SearchImportsVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	if !v.internalTypeNames.Contains(declaration.InternalTypeName()) {
		v.internalTypeNames.Add(declaration.InternalTypeName())
		v.SafeAcceptDeclaration(declaration.MemberDeclarations())
	}
}

func getTypeName(internalTypeName string) string {
	index := strings.LastIndex(internalTypeName, "$")
	if index != -1 {
		return internalTypeName[index+1:]
	}
	index = strings.LastIndex(internalTypeName, "/")
	if index != -1 {
		return internalTypeName[index+1:]
	}
	return internalTypeName
}

func (v *SearchImportsVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.AbstractJavaSyntaxVisitor.VisitAnnotationReference(ref)
	v.add(ref.Type())
}

func (v *SearchImportsVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.AbstractJavaSyntaxVisitor.VisitAnnotationElementValue(ref)
	v.add(ref.Type())
}

func (v *SearchImportsVisitor) VisitObjectType(typ intmod.IObjectType) {
	v.add(typ)
	v.SafeAcceptTypeArgumentVisitable(typ.TypeArguments())
}

func (v *SearchImportsVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *SearchImportsVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitBinaryOperatorExpression(expr)
}

func (v *SearchImportsVisitor) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitBooleanExpression(expr)
}

func (v *SearchImportsVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitCastExpression(expr)
}

func (v *SearchImportsVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitConstructorInvocationExpression(expr)
}

func (v *SearchImportsVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitConstructorReferenceExpression(expr)
}

func (v *SearchImportsVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	if v.maxLineNumber < expr.LineNumber() {
		v.maxLineNumber = expr.LineNumber()
	}
}

func (v *SearchImportsVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	if v.maxLineNumber < expr.LineNumber() {

	}
	v.SafeAcceptExpression(expr.Expression())
}

func (v *SearchImportsVisitor) VisitFloatConstantExpression(expression intmod.IFloatConstantExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
}

func (v *SearchImportsVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
}

func (v *SearchImportsVisitor) VisitInstanceOfExpression(expression intmod.IInstanceOfExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitInstanceOfExpression(expression)
}

func (v *SearchImportsVisitor) VisitLambdaFormalParametersExpression(expression intmod.ILambdaFormalParametersExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitLambdaFormalParametersExpression(expression)
}

func (v *SearchImportsVisitor) VisitLambdaIdentifiersExpression(expression intmod.ILambdaIdentifiersExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitLambdaIdentifiersExpression(expression)
}

func (v *SearchImportsVisitor) VisitLengthExpression(expression intmod.ILengthExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitLengthExpression(expression)
}

func (v *SearchImportsVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitLocalVariableReferenceExpression(expression)
}

func (v *SearchImportsVisitor) VisitLongConstantExpression(expression intmod.ILongConstantExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
}

func (v *SearchImportsVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitMethodInvocationExpression(expression)
}

func (v *SearchImportsVisitor) VisitMethodReferenceExpression(expression intmod.IMethodReferenceExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	expression.Expression().Accept(v)
}

func (v *SearchImportsVisitor) VisitNewArray(expression intmod.INewArray) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.SafeAcceptExpression(expression.DimensionExpressionList())
}

func (v *SearchImportsVisitor) VisitNewExpression(expression intmod.INewExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}

	typ := expression.Type()
	typ.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expression.Parameters())
	v.SafeAcceptDeclaration(expression.BodyDeclaration())
}

func (v *SearchImportsVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitNewInitializedArray(expression)
}

func (v *SearchImportsVisitor) VisitNullExpression(expression intmod.INullExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
}

func (v *SearchImportsVisitor) VisitObjectTypeReferenceExpression(expression intmod.IObjectTypeReferenceExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitObjectTypeReferenceExpression(expression)
}

func (v *SearchImportsVisitor) VisitParenthesesExpression(expression intmod.IParenthesesExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitParenthesesExpression(expression)
}

func (v *SearchImportsVisitor) VisitPostOperatorExpression(expression intmod.IPostOperatorExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitPostOperatorExpression(expression)
}

func (v *SearchImportsVisitor) VisitPreOperatorExpression(expression intmod.IPreOperatorExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitPreOperatorExpression(expression)
}

func (v *SearchImportsVisitor) VisitStringConstantExpression(expression intmod.IStringConstantExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitStringConstantExpression(expression)
}

func (v *SearchImportsVisitor) VisitSuperExpression(expression intmod.ISuperExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitSuperExpression(expression)
}

func (v *SearchImportsVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitTernaryOperatorExpression(expression)
}

func (v *SearchImportsVisitor) VisitThisExpression(expression intmod.IThisExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitThisExpression(expression)
}
func (v *SearchImportsVisitor) VisitTypeReferenceDotClassExpression(expression intmod.ITypeReferenceDotClassExpression) {
	if v.maxLineNumber < expression.LineNumber() {
		v.maxLineNumber = expression.LineNumber()
	}
	v.AbstractJavaSyntaxVisitor.VisitTypeReferenceDotClassExpression(expression)
}

func (v *SearchImportsVisitor) add(typ intmod.IObjectType) {
	descriptor := typ.Descriptor()

	if descriptor[len(descriptor)-1] == ';' {
		internalTypeName := typ.InternalName()

		if !v.importsFragment.IncCounter(internalTypeName) {
			typName := getTypeName(internalTypeName)

			if !v.importTypeNames.Contains(typName) {
				if strings.HasPrefix(internalTypeName, "java/lang/") {
					if strings.Index(internalTypeName[10:], "/") != -1 { // 10 = "java/lang/".length()
						v.importsFragment.AddImport(internalTypeName, typ.QualifiedName())
						v.importTypeNames.Add(typName)
					}
				} else if strings.HasPrefix(internalTypeName, v.internalPackagePrefix) {
					if (strings.Index(internalTypeName[len(v.internalPackagePrefix):], "/") != -1) && !v.localTypeNames.Contains(typName) {
						v.importsFragment.AddImport(internalTypeName, typ.QualifiedName())
						v.importTypeNames.Add(typName)
					}
				} else if !v.localTypeNames.Contains(typName) && !v.loader.CanLoad(v.internalPackagePrefix+typName) {
					v.importsFragment.AddImport(internalTypeName, typ.QualifiedName())
					v.importTypeNames.Add(typName)
				}
			}
		}
	}
}

func NewTypeVisitor2(mainTypeNames util.ISet[string]) *TypeVisitor2 {
	return &TypeVisitor2{
		mainTypeNames: mainTypeNames,
	}
}

type TypeVisitor2 struct {
	javasyntax.AbstractJavaSyntaxVisitor

	mainTypeNames util.ISet[string]
}

func (v *TypeVisitor2) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

func (v *TypeVisitor2) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.mainTypeNames.Add(getTypeName(decl.InternalTypeName()))
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor2) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.mainTypeNames.Add(getTypeName(decl.InternalTypeName()))
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor2) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.mainTypeNames.Add(getTypeName(decl.InternalTypeName()))
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor2) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.mainTypeNames.Add(getTypeName(decl.InternalTypeName()))
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor2) VisitFieldDeclaration(_ intmod.IFieldDeclaration) {
}

func (v *TypeVisitor2) VisitConstructorDeclaration(_ intmod.IConstructorDeclaration) {
}

func (v *TypeVisitor2) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {
}

func (v *TypeVisitor2) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor2) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor2) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *TypeVisitor2) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor2) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *TypeVisitor2) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor2) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *TypeVisitor2) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *TypeVisitor2) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *TypeVisitor2) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor2) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor2) VisitModuleDeclaration(decl intmod.IModuleDeclaration) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *TypeVisitor2) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *TypeVisitor2) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *TypeVisitor2) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitCommentExpression(expr intmod.ICommentExpression) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *TypeVisitor2) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *TypeVisitor2) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *TypeVisitor2) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *TypeVisitor2) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *TypeVisitor2) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor2) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *TypeVisitor2) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor2) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *TypeVisitor2) VisitNoExpression(expr intmod.INoExpression) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor2) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *TypeVisitor2) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *TypeVisitor2) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *TypeVisitor2) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *TypeVisitor2) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor2) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *TypeVisitor2) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor2) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *TypeVisitor2) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor2) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *TypeVisitor2) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *TypeVisitor2) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *TypeVisitor2) VisitBreakStatement(stat intmod.IBreakStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitByteCodeStatement(stat intmod.IByteCodeStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitCommentStatement(stat intmod.ICommentStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitContinueStatement(stat intmod.IContinueStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *TypeVisitor2) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *TypeVisitor2) VisitNoStatement(stat intmod.INoStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitReturnStatement(stat intmod.IReturnStatement) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *TypeVisitor2) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *TypeVisitor2) VisitSwitchStatementDefaultLabel(stat intmod.IDefaultLabel) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *TypeVisitor2) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *TypeVisitor2) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *TypeVisitor2) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *TypeVisitor2) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor2) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *TypeVisitor2) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---
// func (v *TypeVisitor2) VisitPrimitiveType(y *PrimitiveType)     {}
// func (v *TypeVisitor2) VisitObjectType(y *ObjectType)           {}
// func (v *TypeVisitor2) VisitInnerObjectType(y *InnerObjectType) {}

func (v *TypeVisitor2) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

//func (v *TypeVisitor2) VisitGenericType(y *GenericType)         {}

// --- ITypeParameterVisitor --- //

func (v *TypeVisitor2) VisitTypeParameter(parameter intmod.ITypeParameter) {
	// TODO: Empty
}

func (v *TypeVisitor2) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *TypeVisitor2) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

//func (v *TypeVisitor2) VisitTypeArguments(arguments intmod.ITypeArguments) {}
//
//func (v *TypeVisitor2) VisitDiamondTypeArgument(argument intmod.IDiamondTypeArgument) {}
//
//func (v *TypeVisitor2) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
//}
//
//func (v *TypeVisitor2) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
//}
//
//func (v *TypeVisitor2) VisitWildcardTypeArgument(argument intmod.IWildcardTypeArgument) {}
//
//func (v *TypeVisitor2) VisitPrimitiveType(t intmod.IPrimitiveType) {}
//
//func (v *TypeVisitor2) VisitObjectType(t intmod.IObjectType) {}
//
//func (v *TypeVisitor2) VisitInnerObjectType(t intmod.IInnerObjectType) {}
//
//func (v *TypeVisitor2) VisitGenericType(t intmod.IGenericType) {}

func (v *TypeVisitor2) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *TypeVisitor2) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *TypeVisitor2) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *TypeVisitor2) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *TypeVisitor2) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *TypeVisitor2) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *TypeVisitor2) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *TypeVisitor2) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *TypeVisitor2) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *TypeVisitor2) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *TypeVisitor2) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *TypeVisitor2) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *TypeVisitor2) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *TypeVisitor2) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
