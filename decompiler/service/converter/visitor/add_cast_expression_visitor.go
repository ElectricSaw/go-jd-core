package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewAddCastExpressionVisitor(typeMaker intsrv.ITypeMaker) intsrv.IAddCastExpressionVisitor {
	return &AddCastExpressionVisitor{
		searchFirstLineNumberVisitor: NewSearchFirstLineNumberVisitor(),
		typeMaker:                    typeMaker,
	}
}

type AddCastExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor intsrv.ISearchFirstLineNumberVisitor
	typeMaker                    intsrv.ITypeMaker
	typeBounds                   map[string]intmod.IType
	returnedType                 intmod.IType
	exceptionType                intmod.IType
	typ                          intmod.IType
}

func (v *AddCastExpressionVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	memberDeclarations := decl.MemberDeclarations()

	if memberDeclarations != nil {
		tb := v.typeBounds

		if obj, ok := decl.(intsrv.IClassFileBodyDeclaration); ok {
			v.typeBounds = obj.TypeBounds()
			memberDeclarations.AcceptDeclaration(v)
			v.typeBounds = tb
		}
	}
}

func (v *AddCastExpressionVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	if (decl.Flags() & intmod.FlagSynthetic) == 0 {
		t := v.typ

		v.typ = decl.Type()
		decl.FieldDeclarators().AcceptDeclaration(v)
		v.typ = t
	}
}

func (v *AddCastExpressionVisitor) VisitFieldDeclarator(declarator intmod.IFieldDeclarator) {
	variableInitializer := declarator.VariableInitializer()

	if variableInitializer != nil {
		extraDimension := declarator.Dimension()

		if extraDimension == 0 {
			variableInitializer.AcceptDeclaration(v)
		} else {
			t := v.typ

			v.typ = v.typ.CreateType(v.typ.Dimension() + extraDimension)
			variableInitializer.AcceptDeclaration(v)
			v.typ = t
		}
	}
}

func (v *AddCastExpressionVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	state := decl.Statements()

	if state != nil {
		tb := v.typeBounds

		v.typeBounds = decl.(intsrv.IClassFileStaticInitializerDeclaration).TypeBounds()
		state.AcceptStatement(v)
		v.typeBounds = tb
	}
}

func (v *AddCastExpressionVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	if (decl.Flags() & (intmod.FlagSynthetic | intmod.FlagBridge)) == 0 {
		statements := decl.Statements()

		if statements != nil {
			tb := v.typeBounds
			et := v.exceptionType

			v.typeBounds = decl.(intsrv.IClassFileConstructorDeclaration).TypeBounds()
			v.exceptionType = decl.ExceptionTypes()
			statements.AcceptStatement(v)
			v.typeBounds = tb
			v.exceptionType = et
		}
	}
}

func (v *AddCastExpressionVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	if (decl.Flags() & (intmod.FlagSynthetic | intmod.FlagBridge)) == 0 {
		statements := decl.Statements()

		if statements != nil {
			tb := v.typeBounds
			rt := v.returnedType
			et := v.exceptionType

			v.typeBounds = decl.(intsrv.IClassFileMethodDeclaration).TypeBounds()
			v.returnedType = decl.ReturnedType()
			v.exceptionType = decl.ExceptionTypes()
			statements.AcceptStatement(v)
			v.typeBounds = tb
			v.returnedType = rt
			v.exceptionType = et
		}
	}
}

func (v *AddCastExpressionVisitor) VisitLambdaIdentifiersExpression(expression intmod.ILambdaIdentifiersExpression) {
	statements := expression.Statements()

	if statements != nil {
		rt := v.returnedType

		v.returnedType = _type.OtTypeObject.(intmod.IType)
		statements.AcceptStatement(v)
		v.returnedType = rt
	}
}

func (v *AddCastExpressionVisitor) VisitReturnExpressionStatement(statement intmod.IReturnExpressionStatement) {
	statement.SetExpression(v.updateExpression(v.returnedType, statement.Expression(), false, true))
}

func (v *AddCastExpressionVisitor) VisitThrowStatement(statement intmod.IThrowStatement) {
	if v.exceptionType != nil && v.exceptionType.Size() == 1 {
		exceptionType := v.exceptionType.First()

		if exceptionType.IsGenericType() && statement.Expression().Type() != exceptionType {
			statement.SetExpression(v.addCastExpression(exceptionType, statement.Expression()))
		}
	}
}

func (v *AddCastExpressionVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	t := v.typ

	v.typ = decl.Type()
	decl.LocalVariableDeclarators().AcceptDeclaration(v)
	v.typ = t
}

func (v *AddCastExpressionVisitor) VisitLocalVariableDeclarator(declarator intmod.ILocalVariableDeclarator) {
	variableInitializer := declarator.VariableInitializer()

	if variableInitializer != nil {
		extraDimension := declarator.Dimension()

		if extraDimension == 0 {
			variableInitializer.AcceptDeclaration(v)
		} else {
			t := v.typ

			v.typ = v.typ.CreateType(v.typ.Dimension() + extraDimension)
			variableInitializer.AcceptDeclaration(v)
			v.typ = t
		}
	}
}

func (v *AddCastExpressionVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	if v.typ.Dimension() == 0 {
		v.AcceptListDeclaration(ConvertArrayVariable(decl.ToSlice()))
	} else {
		t := v.typ

		v.typ = v.typ.CreateType(v.typ.Dimension() - 1)
		v.AcceptListDeclaration(ConvertArrayVariable(decl.ToSlice()))
		v.typ = t
	}
}

func (v *AddCastExpressionVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	expr := decl.Expression()

	if expr.IsNewInitializedArray() {
		nia := expr.(intmod.INewInitializedArray)
		t := v.typ

		v.typ = nia.Type()
		nia.ArrayInitializer().AcceptDeclaration(v)
		v.typ = t
	} else {
		decl.SetExpression(v.updateExpression(v.typ, expr, false, true))
	}
}

func (v *AddCastExpressionVisitor) VisitSuperConstructorInvocationExpression(expression intmod.ISuperConstructorInvocationExpression) {
	parameters := expression.Parameters()

	if (parameters != nil) && (parameters.Size() > 0) {
		unique := v.typeMaker.MatchCount(expression.ObjectType().InternalName(),
			"<init>", parameters.Size(), true) <= 1
		forceCast := !unique && (v.typeMaker.MatchCount2(v.typeBounds,
			expression.ObjectType().InternalName(), "<init>", parameters, true) > 1)
		expression.SetParameters(v.updateParameters(
			expression.(intsrv.IClassFileSuperConstructorInvocationExpression).ParameterTypes(), parameters, forceCast, unique))
	}
}

func (v *AddCastExpressionVisitor) VisitConstructorInvocationExpression(expression intmod.IConstructorInvocationExpression) {
	parameters := expression.Parameters()

	if (parameters != nil) && (parameters.Size() > 0) {
		unique := v.typeMaker.MatchCount(expression.ObjectType().InternalName(), "<init>", parameters.Size(), true) <= 1
		forceCast := !unique && (v.typeMaker.MatchCount2(v.typeBounds, expression.ObjectType().InternalName(), "<init>", parameters, true) > 1)
		expression.SetParameters(v.updateParameters(expression.(intsrv.IClassFileConstructorInvocationExpression).ParameterTypes(), parameters, forceCast, unique))
	}
}

func (v *AddCastExpressionVisitor) VisitMethodInvocationExpression(expression intmod.IMethodInvocationExpression) {
	parameters := expression.Parameters()

	if (parameters != nil) && (parameters.Size() > 0) {
		unique := v.typeMaker.MatchCount(expression.InternalTypeName(), expression.Name(), parameters.Size(), false) <= 1
		forceCast := !unique && (v.typeMaker.MatchCount2(v.typeBounds, expression.InternalTypeName(), expression.Name(), parameters, false) > 1)
		expression.SetParameters(v.updateParameters(expression.(intsrv.IClassFileMethodInvocationExpression).ParameterTypes(), parameters, forceCast, unique))
	}

	expression.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitNewExpression(expression intmod.INewExpression) {
	parameters := expression.Parameters()

	if parameters != nil {
		unique := v.typeMaker.MatchCount(expression.ObjectType().InternalName(), "<init>", parameters.Size(), true) <= 1
		forceCast := !unique && (v.typeMaker.MatchCount2(v.typeBounds, expression.ObjectType().InternalName(), "<init>", parameters, true) > 1)
		expression.SetParameters(v.updateParameters(expression.(intsrv.IClassFileNewExpression).ParameterTypes(), parameters, forceCast, unique))
	}
}

func (v *AddCastExpressionVisitor) VisitNewInitializedArray(expression intmod.INewInitializedArray) {
	arrayInitializer := expression.ArrayInitializer()

	if arrayInitializer != nil {
		t := v.typ

		v.typ = expression.Type()
		arrayInitializer.AcceptDeclaration(v)
		v.typ = t
	}
}

func (v *AddCastExpressionVisitor) VisitFieldReferenceExpression(expression intmod.IFieldReferenceExpression) {
	exp := expression.Expression()

	if (exp != nil) && !exp.IsObjectTypeReferenceExpression() {
		typ := v.typeMaker.MakeFromInternalTypeName(expression.InternalTypeName())

		if typ.Name() != "" {
			expression.SetExpression(v.updateExpression(typ.(intmod.IType), exp, false, true))
		}
	}
}

func (v *AddCastExpressionVisitor) VisitBinaryOperatorExpression(expression intmod.IBinaryOperatorExpression) {
	expression.LeftExpression().Accept(v)
	rightExpression := expression.RightExpression()

	if expression.Operator() == "=" {
		if rightExpression.IsMethodInvocationExpression() {
			mie := rightExpression.(intsrv.IClassFileMethodInvocationExpression)

			if mie.TypeParameters() != nil {
				// Do not add cast expression if method contains type parameters
				rightExpression.Accept(v)
				return
			}
		}

		expression.SetRightExpression(v.updateExpression(expression.LeftExpression().Type(), rightExpression, false, true))
		return
	}

	rightExpression.Accept(v)
}

func (v *AddCastExpressionVisitor) VisitTernaryOperatorExpression(expression intmod.ITernaryOperatorExpression) {
	expressionType := expression.Type()

	expression.Condition().Accept(v)
	expression.SetTrueExpression(v.updateExpression(expressionType, expression.TrueExpression(), false, true))
	expression.SetFalseExpression(v.updateExpression(expressionType, expression.FalseExpression(), false, true))
}

func (v *AddCastExpressionVisitor) updateParameters(types intmod.IType, expression intmod.IExpression, forceCast bool, unique bool) intmod.IExpression {
	if expression != nil {
		if expression.IsList() {
			typeList := util.NewDefaultListWithSlice[intmod.IType](types.ToSlice())
			expressionList := util.NewDefaultListWithSlice[intmod.IExpression](expression.ToSlice())

			for i := expressionList.Size() - 1; i >= 0; i-- {
				expressionList.Set(i, v.updateParameter(typeList.Get(i), expressionList.Get(i), forceCast, unique))
			}
		} else {
			expression = v.updateParameter(types.First(), expression.First(), forceCast, unique)
		}
	}

	return expression
}

func (v *AddCastExpressionVisitor) updateParameter(typ intmod.IType, expr intmod.IExpression, forceCast bool, unique bool) intmod.IExpression {
	expr = v.updateExpression(typ, expr, forceCast, unique)

	if typ == _type.PtTypeByte {
		if expr.IsIntegerConstantExpression() {
			expr = expression.NewCastExpression(_type.PtTypeByte, expr)
		} else if expr.IsTernaryOperatorExpression() {
			exp := expr.TrueExpression()
			if exp.IsIntegerConstantExpression() || exp.IsTernaryOperatorExpression() {
				expr = expression.NewCastExpression(_type.PtTypeByte, expr)
			} else {
				exp = expr.FalseExpression()
				if exp.IsIntegerConstantExpression() || exp.IsTernaryOperatorExpression() {
					expr = expression.NewCastExpression(_type.PtTypeByte, expr)
				}
			}
		}
	}

	return expr
}

func (v *AddCastExpressionVisitor) updateExpression(typ intmod.IType, expr intmod.IExpression, forceCast bool, unique bool) intmod.IExpression {
	if expr.IsNullExpression() {
		if forceCast {
			v.searchFirstLineNumberVisitor.Init()
			expr.Accept(v.searchFirstLineNumberVisitor)
			expr = expression.NewCastExpressionWithLineNumber(v.searchFirstLineNumberVisitor.LineNumber(), typ, expr)
		}
	} else {
		expressionType := expr.Type()

		if expressionType != typ {
			if typ.IsObjectType() {
				if expressionType.IsObjectType() {
					objectType := typ.(intmod.IObjectType)
					expressionObjectType := expressionType.(intmod.IObjectType)

					if forceCast && !objectType.RawEquals(expressionObjectType) {
						// Force disambiguation of method invocation => Add cast
						if expr.IsNewExpression() {
							ne := expr.(intsrv.IClassFileNewExpression)
							ne.SetObjectType(ne.ObjectType().CreateTypeWithArgs(nil))
						}
						expr = v.addCastExpression(objectType.(intmod.IType), expr)
					} else if _type.OtTypeObject.(intmod.IType) != typ && !v.typeMaker.IsAssignable(v.typeBounds, objectType, expressionObjectType) {
						ta1 := objectType.TypeArguments()
						ta2 := expressionObjectType.TypeArguments()
						t := typ

						if (ta1 != nil) && (ta2 != nil) && !ta1.IsTypeArgumentAssignableFrom(v.typeBounds, ta2) {
							// Incompatible typeArgument arguments => Add cast
							t = objectType.CreateTypeWithArgs(nil).(intmod.IType)
						}
						expr = v.addCastExpression(t, expr)
					}
				} else if expressionType.IsGenericType() && _type.OtTypeObject.(intmod.IType) != typ {
					expr = v.addCastExpression(typ, expr)
				}
			} else if typ.IsGenericType() {
				if expressionType.IsObjectType() || expressionType.IsGenericType() {
					expr = v.addCastExpression(typ, expr)
				}
			}
		}

		if expr.IsCastExpression() {
			ceExpressionType := expr.Expression().Type()

			if typ.IsObjectType() && ceExpressionType.IsObjectType() {
				ot1 := typ.(intmod.IObjectType)
				ot2 := ceExpressionType.(intmod.IObjectType)

				if ot1 == ot2 {
					// Remove cast expr
					expr = expr.Expression()
				} else if unique && v.typeMaker.IsAssignable(v.typeBounds, ot1, ot2) {
					// Remove cast expr
					expr = expr.Expression()
				}
			}
		}

		expr.Accept(v)
	}

	return expr
}

func (v *AddCastExpressionVisitor) addCastExpression(typ intmod.IType, expr intmod.IExpression) intmod.IExpression {
	if expr.IsCastExpression() {
		if typ == expr.Expression().Type() {
			return expr.Expression()
		} else {
			ce := expr.(intmod.ICastExpression)
			ce.SetType(typ)
			return ce
		}
	} else {
		v.searchFirstLineNumberVisitor.Init()
		expr.Accept(v.searchFirstLineNumberVisitor)
		return expression.NewCastExpressionWithLineNumber(v.searchFirstLineNumberVisitor.LineNumber(), typ, expr)
	}
}

func (v *AddCastExpressionVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitBreakStatement(_ intmod.IBreakStatement)       {}
func (v *AddCastExpressionVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {}
func (v *AddCastExpressionVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {}
func (v *AddCastExpressionVisitor) VisitNullExpression(_ intmod.INullExpression)       {}
func (v *AddCastExpressionVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {}
func (v *AddCastExpressionVisitor) VisitThisExpression(_ intmod.IThisExpression)   {}
func (v *AddCastExpressionVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *AddCastExpressionVisitor) VisitObjectReference(_ intmod.IObjectReference) {}
func (v *AddCastExpressionVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *AddCastExpressionVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {}
func (v *AddCastExpressionVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *AddCastExpressionVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *AddCastExpressionVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *AddCastExpressionVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *AddCastExpressionVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *AddCastExpressionVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}

func ConvertArrayVariable(list []intmod.IVariableInitializer) []intmod.IDeclaration {
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

func (v *AddCastExpressionVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *AddCastExpressionVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AddCastExpressionVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AddCastExpressionVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AddCastExpressionVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AddCastExpressionVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AddCastExpressionVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AddCastExpressionVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AddCastExpressionVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *AddCastExpressionVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *AddCastExpressionVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AddCastExpressionVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *AddCastExpressionVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---
func (v *AddCastExpressionVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	expr.Expression().Accept(v)
	expr.Index().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *AddCastExpressionVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *AddCastExpressionVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitNewArray(expr intmod.INewArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptExpression(expr.DimensionExpressionList())
}

func (v *AddCastExpressionVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

// --- IReferenceVisitor ---

func (v *AddCastExpressionVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AddCastExpressionVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *AddCastExpressionVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AddCastExpressionVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *AddCastExpressionVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AddCastExpressionVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *AddCastExpressionVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

// --- IStatementVisitor ---

func (v *AddCastExpressionVisitor) VisitAssertStatement(stat intmod.IAssertStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptExpression(stat.Message())
}

func (v *AddCastExpressionVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitDoWhileStatement(stat intmod.IDoWhileStatement) {
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitForStatement(stat intmod.IForStatement) {
	v.SafeAcceptDeclaration(stat.Declaration())
	v.SafeAcceptExpression(stat.Init())
	v.SafeAcceptExpression(stat.Condition())
	v.SafeAcceptExpression(stat.Update())
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitIfStatement(stat intmod.IIfStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitIfElseStatement(stat intmod.IIfElseStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
	stat.ElseStatements().AcceptStatement(v)
}

func (v *AddCastExpressionVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *AddCastExpressionVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *AddCastExpressionVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *AddCastExpressionVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *AddCastExpressionVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *AddCastExpressionVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *AddCastExpressionVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *AddCastExpressionVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *AddCastExpressionVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

func (v *AddCastExpressionVisitor) VisitWhileStatement(stat intmod.IWhileStatement) {
	stat.Condition().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

// --- ITypeVisitor ---

// --- ITypeParameterVisitor --- //

func (v *AddCastExpressionVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *AddCastExpressionVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *AddCastExpressionVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *AddCastExpressionVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *AddCastExpressionVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AddCastExpressionVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *AddCastExpressionVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *AddCastExpressionVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
