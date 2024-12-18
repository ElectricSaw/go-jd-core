package visitor

import (
	"strings"

	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/classfile/attribute"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/classfile/constant"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
)

func NewUpdateBridgeMethodTypeVisitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateBridgeMethodTypeVisitor {
	return &UpdateBridgeMethodTypeVisitor{
		typeMaker: typeMaker,
	}
}

type UpdateBridgeMethodTypeVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker intsrv.ITypeMaker
}

func (v *UpdateBridgeMethodTypeVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
	v.SafeAcceptListDeclaration(ConvertTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
}

func (v *UpdateBridgeMethodTypeVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	if declaration.IsStatic() && declaration.ReturnedType().IsObjectType() && strings.HasPrefix(declaration.Name(), "access$") {
		typeTypes := v.typeMaker.MakeTypeTypes(declaration.ReturnedType().InternalName())

		if (typeTypes != nil) && (typeTypes.TypeParameters != nil) {
			cfmd := declaration.(intsrv.IClassFileMethodDeclaration)
			method := cfmd.Method()
			code := method.Attributes()["Code"].(attribute.AttributeCode).Code()
			offset := 0
			opcode := code[offset] & 255

			for (21 <= opcode && opcode <= 45) || // ILOAD, LLOAD, FLOAD, DLOAD, ..., ILOAD_0 ... ILOAD_3, ..., ALOAD_1, ..., ALOAD_3
				(89 <= opcode && opcode <= 95) { // DUP, ..., DUP2_X2, SWAP
				offset++
				opcode = code[offset] & 255
			}

			switch opcode {
			case 178, 179, 180, 181: // GETSTATIC, PUTSTATIC, GETFIELD, PUTFIELD
				offset++
				r1 := int(code[offset]&255) << 8
				offset++
				r2 := int(code[offset]&255) << 8
				index := r1 | r2

				constants := method.Constants()
				constantMemberRef := constants.Constant(index).(constant.ConstantMemberRef)
				typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(constant.ConstantNameAndType)
				name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				typ := v.typeMaker.MakeFieldType(typeName, name, descriptor)
				// Update returned generic type of bridge method
				v.typeMaker.SetMethodReturnedType(typeName, cfmd.Name(), cfmd.Descriptor(), typ)
			case 182, 183, 184, 185: // INVOKEVIRTUAL, INVOKESPECIAL, INVOKESTATIC, INVOKEINTERFACE
				offset++
				r1 := int(code[offset]&255) << 8
				offset++
				r2 := int(code[offset]&255) << 8
				index := r1 | r2

				constants := method.Constants()
				constantMemberRef := constants.Constant(index).(constant.ConstantMemberRef)
				typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(constant.ConstantNameAndType)
				name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				methodTypes := v.typeMaker.MakeMethodTypes2(typeName, name, descriptor)

				// Update returned generic type of bridge method
				v.typeMaker.SetMethodReturnedType(typeName, cfmd.Name(), cfmd.Descriptor(), methodTypes.ReturnedType())
			}
		}
	}
}

func (v *UpdateBridgeMethodTypeVisitor) VisitConstructorDeclaration(_ intmod.IConstructorDeclaration) {
}
func (v *UpdateBridgeMethodTypeVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *UpdateBridgeMethodTypeVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitAnnotationDeclaration(_ intmod.IAnnotationDeclaration) {
}
func (v *UpdateBridgeMethodTypeVisitor) VisitEnumDeclaration(_ intmod.IEnumDeclaration) {}

func ConvertMethodDeclarations(list []intsrv.IClassFileConstructorOrMethodDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertFieldDeclarations(list []intsrv.IClassFileFieldDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertTypeDeclarations(list []intsrv.IClassFileTypeDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertMemberDeclaration(list []intsrv.IClassFileMemberDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateBridgeMethodTypeVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *UpdateBridgeMethodTypeVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	decl.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
	decl.FieldDeclarators().AcceptDeclaration(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	v.SafeAcceptDeclaration(decl.LocalVariableDeclarators())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *UpdateBridgeMethodTypeVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	expr.LeftExpression().Accept(v)
	expr.RightExpression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	t := expression.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.SafeAcceptStatement(expr.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.Expression().Accept(v)
	v.SafeAcceptTypeArgumentVisitable(expr.NonWildcardTypeArguments().(intmod.IWildcardSuperTypeArgument))
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNewExpression(expr intmod.INewExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNullExpression(expr intmod.INullExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.Parameters())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	expr.Condition().Accept(v)
	expr.TrueExpression().Accept(v)
	expr.FalseExpression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
}

// --- IReferenceVisitor ---

func (v *UpdateBridgeMethodTypeVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitInnerObjectReference(ref intmod.IInnerObjectReference) {
	v.VisitInnerObjectType(ref.(intmod.IInnerObjectType))
}

func (v *UpdateBridgeMethodTypeVisitor) VisitObjectReference(ref intmod.IObjectReference) {
	v.VisitObjectType(ref.(intmod.IObjectType))
}

// --- IStatementVisitor ---

func (v *UpdateBridgeMethodTypeVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitBreakStatement(_ intmod.IBreakStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitReturnExpressionStatement(stat intmod.IReturnExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

func (v *UpdateBridgeMethodTypeVisitor) VisitTypes(types intmod.ITypes) {
	for _, value := range types.ToSlice() {
		value.AcceptTypeVisitor(v)
	}
}

// --- ITypeParameterVisitor --- //

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	parameter.TypeBounds().AcceptTypeVisitor(v)
}

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *UpdateBridgeMethodTypeVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateBridgeMethodTypeVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateBridgeMethodTypeVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
