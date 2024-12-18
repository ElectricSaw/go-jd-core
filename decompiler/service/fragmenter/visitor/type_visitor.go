package visitor

import (
	"strings"

	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

var UnknownLineNumber = api.UnknownLineNumber

func NewTypeVisitor(loader api.Loader, mainInternalTypeName string, majorVersion int,
	importsFragment intmod.IImportsFragment) ITypeVisitor {
	v := &TypeVisitor{
		loader:                loader,
		genericTypesSupported: majorVersion > 49,
		importsFragment:       importsFragment,
	}

	index := strings.Index(mainInternalTypeName, "/")
	if index == -1 {
		v.internalPackageName = ""
	} else {
		v.internalPackageName = mainInternalTypeName[:index+1]
	}

	return v
}

type ITypeVisitor interface {
	intmod.ITypeVisitor
	intmod.ITypeArgumentVisitor
	intmod.ITypeParameterVisitor
}

type TypeVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	loader                  api.Loader
	internalPackageName     string
	genericTypesSupported   bool
	importsFragment         intmod.IImportsFragment
	tokens                  ITokens
	maxLineNumber           int
	currentInternalTypeName string
	textTokenCache          map[string]intmod.ITextToken
}

func (v *TypeVisitor) VisitTypeArguments(arguments intmod.ITypeArguments) {
	tmp := make([]intmod.IType, 0)
	for _, item := range arguments.ToSlice() {
		tmp = append(tmp, item.(intmod.IType))
	}
	v.buildTokensForList(util.NewDefaultListWithSlice(tmp), token.CommaSpace)
}

func (v *TypeVisitor) VisitDiamondTypeArgument(_ intmod.IDiamondTypeArgument) {}

func (v *TypeVisitor) VisitWildcardExtendsTypeArgument(argument intmod.IWildcardExtendsTypeArgument) {
	v.tokens.Add(token.QuestionMarkSpace)
	v.tokens.Add(token.Extends)
	v.tokens.Add(token.Space)

	typ := argument.Type()
	typ.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {
	switch typ.JavaPrimitiveFlags() {
	case intmod.FlagBoolean:
		v.tokens.Add(token.Boolean)
		break
	case intmod.FlagChar:
		v.tokens.Add(token.Char)
		break
	case intmod.FlagFloat:
		v.tokens.Add(token.Float)
		break
	case intmod.FlagDouble:
		v.tokens.Add(token.Double)
		break
	case intmod.FlagByte:
		v.tokens.Add(token.Byte)
		break
	case intmod.FlagShort:
		v.tokens.Add(token.Short)
		break
	case intmod.FlagInt:
		v.tokens.Add(token.Int)
		break
	case intmod.FlagLong:
		v.tokens.Add(token.Long)
		break
	case intmod.FlagVoid:
		v.tokens.Add(token.Void)
		break
	}

	// Build token for dimension
	v.visitDimension(typ.Dimension())
}

func (v *TypeVisitor) VisitObjectType(typ intmod.IObjectType) {
	// Build token for type reference
	v.tokens.Add(v.newTypeReferenceToken(typ, v.currentInternalTypeName))

	if v.genericTypesSupported {
		// Build token for type arguments
		typeArguments := typ.TypeArguments()

		if typeArguments != nil {
			v.visitTypeArgumentList(typeArguments)
		}
	}

	// Build token for dimension
	v.visitDimension(typ.Dimension())
}

func (v *TypeVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {
	if v.currentInternalTypeName == "" || v.currentInternalTypeName != typ.InternalName() &&
		v.currentInternalTypeName != typ.OuterType().InternalName() {
		outerType := typ.OuterType()

		outerType.AcceptTypeVisitor(v)
		v.tokens.Add(token.Dot)
	}

	// Build token for type reference
	v.tokens.Add(token.NewReferenceToken(intmod.TypeToken, typ.InternalName(),
		typ.Name(), "", v.currentInternalTypeName))

	if v.genericTypesSupported {
		// Build token for type arguments
		typeArguments := typ.TypeArguments()

		if typeArguments != nil {
			v.visitTypeArgumentList(typeArguments)
		}
	}

	// Build token for dimension
	v.visitDimension(typ.Dimension())
}

func (v *TypeVisitor) visitTypeArgumentList(arguments intmod.ITypeArgument) {
	if arguments != nil {
		v.tokens.Add(token.LeftAngleBracket)
		arguments.AcceptTypeArgumentVisitor(v)
		v.tokens.Add(token.RightAngleBracket)
	}
}

func (v *TypeVisitor) visitDimension(dimension int) {
	switch dimension {
	case 0:
		break
	case 1:
		v.tokens.Add(token.Dimension1)
		break
	case 2:
		v.tokens.Add(token.Dimension2)
		break
	default:
		str := ""
		for i := 0; i < dimension; i++ {
			str += "[]"
		}
		v.tokens.Add(v.newTextToken(str))
		break
	}
}

func (v *TypeVisitor) VisitWildcardSuperTypeArgument(argument intmod.IWildcardSuperTypeArgument) {
	v.tokens.Add(token.QuestionMarkSpace)
	v.tokens.Add(token.Super)
	v.tokens.Add(token.Space)

	typ := argument.Type()
	typ.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitTypes(types intmod.ITypes) {
	tmp := make([]intmod.IType, 0)
	for _, item := range types.ToSlice() {
		tmp = append(tmp, item.(intmod.IType))
	}
	v.buildTokensForList(util.NewDefaultListWithSlice(tmp), token.CommaSpace)
}

func (v *TypeVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	v.tokens.Add(v.newTextToken(parameter.Identifier()))
}

func (v *TypeVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	v.tokens.Add(v.newTextToken(parameter.Identifier()))
	v.tokens.Add(token.Space)
	v.tokens.Add(token.Extends)
	v.tokens.Add(token.Space)

	types := parameter.TypeBounds()
	if types.IsList() {
		tmp := make([]intmod.IType, 0)
		for _, item := range types.ToSlice() {
			tmp = append(tmp, item.(intmod.IType))
		}
		v.buildTokensForList(util.NewDefaultListWithSlice(tmp), token.SpaceAndSpace)
	} else {
		typ := types.First()
		typ.AcceptTypeVisitor(v)
	}
}

func (v *TypeVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	size := parameters.Size()

	if size > 0 {
		parameters.Get(0).AcceptTypeParameterVisitor(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(token.CommaSpace)
			parameters.Get(i).AcceptTypeParameterVisitor(v)
		}
	}
}

func (v *TypeVisitor) VisitGenericType(typ intmod.IGenericType) {
	v.tokens.Add(v.newTextToken(typ.Name()))
	v.visitDimension(typ.Dimension())
}

func (v *TypeVisitor) VisitWildcardTypeArgument(_ intmod.IWildcardTypeArgument) {
	v.tokens.Add(token.QuestionMark)
}

func (v *TypeVisitor) buildTokensForList(list util.IList[intmod.IType], separator intmod.ITextToken) {
	size := list.Size()

	if size > 0 {
		list.Get(0).AcceptTypeArgumentVisitor(v)

		for i := 1; i < size; i++ {
			v.tokens.Add(separator)
			list.Get(i).AcceptTypeArgumentVisitor(v)
		}
	}
}

func (v *TypeVisitor) newTypeReferenceToken(ot intmod.IObjectType, ownerInternalName string) intmod.IReferenceToken {
	internalName := ot.InternalName()
	qualifiedName := ot.QualifiedName()
	name := ot.Name()

	if packageContainsType(v.internalPackageName, internalName) {
		// In the current package
		return token.NewReferenceToken(intmod.TypeToken, internalName, name, "", ownerInternalName)
	} else {
		if packageContainsType("java/lang/", internalName) {
			// A 'java.lang' class
			internalLocalTypeName := v.internalPackageName + name

			if v.loader.CanLoad(internalLocalTypeName) {
				return token.NewReferenceToken(intmod.TypeToken, internalName, qualifiedName,
					"", ownerInternalName)
			} else {
				return token.NewReferenceToken(intmod.TypeToken, internalName, name,
					"", ownerInternalName)
			}
		} else {
			return NewTypeReferenceToken(v, v.importsFragment, internalName, qualifiedName,
				name, ownerInternalName).(intmod.IReferenceToken)
		}
	}
}

func packageContainsType(internalPackageName, internalClassName string) bool {
	if strings.HasPrefix(internalClassName, internalPackageName) {
		return strings.Index(internalClassName[len(internalPackageName):], "/") == -1
	} else {
		return false
	}
}

func (v *TypeVisitor) newTextToken(text string) intmod.ITextToken {
	textToken := v.textTokenCache[text]

	if textToken == nil {
		textToken = token.NewTextToken(text)
		v.textTokenCache[text] = textToken
	}

	return textToken
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *TypeVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *TypeVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *TypeVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	v.SafeAcceptDeclaration(decl.MemberDeclarations())
}

func (v *TypeVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *TypeVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *TypeVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *TypeVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *TypeVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *TypeVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *TypeVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *TypeVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *TypeVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *TypeVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *TypeVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *TypeVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *TypeVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *TypeVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *TypeVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *TypeVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *TypeVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *TypeVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *TypeVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *TypeVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *TypeVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *TypeVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *TypeVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *TypeVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *TypeVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *TypeVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *TypeVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *TypeVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *TypeVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *TypeVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *TypeVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *TypeVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *TypeVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *TypeVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *TypeVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *TypeVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *TypeVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *TypeVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *TypeVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *TypeVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *TypeVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *TypeVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *TypeVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *TypeVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *TypeVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *TypeVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *TypeVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *TypeVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *TypeVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *TypeVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *TypeVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *TypeVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *TypeVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *TypeVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *TypeVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *TypeVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *TypeVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *TypeVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

// --- ITypeParameterVisitor --- //

// --- ITypeArgumentVisitor ---

func (v *TypeVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *TypeVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *TypeVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *TypeVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *TypeVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *TypeVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *TypeVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *TypeVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *TypeVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *TypeVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *TypeVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *TypeVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *TypeVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *TypeVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}

func NewTypeReferenceToken(parent *TypeVisitor, importsFragment intmod.IImportsFragment, internalTypeName,
	qualifiedName, name, ownerInternalName string) ITypeReferenceToken {
	t := &TypeReferenceToken{
		ReferenceToken: *token.NewReferenceToken(intmod.TypeToken, internalTypeName, name, "",
			ownerInternalName).(*token.ReferenceToken),
		parent:          parent,
		importsFragment: importsFragment,
		qualifiedName:   qualifiedName,
	}
	return t
}

type ITypeReferenceToken interface {
	intmod.IReferenceToken

	Name() string
}

type TypeReferenceToken struct {
	token.ReferenceToken

	parent          *TypeVisitor
	importsFragment intmod.IImportsFragment
	qualifiedName   string
}

func (t *TypeReferenceToken) Name() string {
	if t.importsFragment.Contains(t.InternalTypeName()) {
		return t.ReferenceToken.Name()
	} else {
		return t.qualifiedName
	}
}

func NewTokens(parent ITypeVisitor) ITokens {
	return &Tokens{
		DefaultList:       *util.NewDefaultList[intmod.IToken]().(*util.DefaultList[intmod.IToken]),
		parent:            parent,
		currentLineNumber: UnknownLineNumber,
	}
}

type ITokens interface {
	util.IList[intmod.IToken]

	CurrentLineNumber() int
	SetCurrentLineNumber(currentLineNumber int)
	AddLineNumberToken(expression intmod.IExpression)
	AddLineNumberTokenAt(lineNumber int)
}

type Tokens struct {
	util.DefaultList[intmod.IToken]

	parent            ITypeVisitor
	currentLineNumber int
}

func (t *Tokens) CurrentLineNumber() int {
	return t.currentLineNumber
}

func (t *Tokens) SetCurrentLineNumber(currentLineNumber int) {
	t.currentLineNumber = currentLineNumber
}

func (t *Tokens) AddLineNumberToken(expression intmod.IExpression) {
	t.AddLineNumberTokenAt(expression.LineNumber())
}

func (t *Tokens) AddLineNumberTokenAt(lineNumber int) {
	if lineNumber != UnknownLineNumber {
		if lineNumber >= t.parent.(*TypeVisitor).maxLineNumber {
			t.Add(token.NewLineNumberToken(lineNumber))
			t.parent.(*TypeVisitor).maxLineNumber = lineNumber
			t.currentLineNumber = lineNumber
		}
	}
}
