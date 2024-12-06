package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/api"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
)

func NewSearchImportsVisitor(loader api.Loader, mainInternalName string) *SearchImportsVisitor {
	v := &SearchImportsVisitor{
		loader: loader,
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
