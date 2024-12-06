package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

func NewAddCastExpressionVisitor(typeMaker intsrv.ITypeMaker) *AddCastExpressionVisitor {
	return &AddCastExpressionVisitor{
		searchFirstLineNumberVisitor: NewSearchFirstLineNumberVisitor(),
		typeMaker:                    typeMaker,
	}
}

type AddCastExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor *SearchFirstLineNumberVisitor
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
	decl.LocalVariableDeclarators().Accept(v)
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

func ConvertArrayVariable(list []intmod.IVariableInitializer) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
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

func (v *AddCastExpressionVisitor) VisitFloatConstantExpression(expression intmod.IFloatConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitIntegerConstantExpression(expression intmod.IIntegerConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitConstructorReferenceExpression(expression intmod.IConstructorReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitDoubleConstantExpression(expression intmod.IDoubleConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitEnumConstantReferenceExpression(expression intmod.IEnumConstantReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitLocalVariableReferenceExpression(expression intmod.ILocalVariableReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitLongConstantExpression(expression intmod.ILongConstantExpression) {
}
func (v *AddCastExpressionVisitor) VisitBreakStatement(statement intmod.IBreakStatement)       {}
func (v *AddCastExpressionVisitor) VisitByteCodeStatement(statement intmod.IByteCodeStatement) {}
func (v *AddCastExpressionVisitor) VisitContinueStatement(statement intmod.IContinueStatement) {}
func (v *AddCastExpressionVisitor) VisitNullExpression(expression intmod.INullExpression)      {}
func (v *AddCastExpressionVisitor) VisitObjectTypeReferenceExpression(expression intmod.IObjectTypeReferenceExpression) {
}
func (v *AddCastExpressionVisitor) VisitSuperExpression(expression intmod.ISuperExpression) {}
func (v *AddCastExpressionVisitor) VisitThisExpression(expression intmod.IThisExpression)   {}
func (v *AddCastExpressionVisitor) VisitTypeReferenceDotClassExpression(expression intmod.ITypeReferenceDotClassExpression) {
}
func (v *AddCastExpressionVisitor) VisitObjectReference(reference intmod.IObjectReference) {}
func (v *AddCastExpressionVisitor) VisitInnerObjectReference(reference intmod.IInnerObjectReference) {
}
func (v *AddCastExpressionVisitor) VisitTypeArguments(typ intmod.ITypeArguments) {}
func (v *AddCastExpressionVisitor) VisitWildcardExtendsTypeArgument(typ intmod.IWildcardExtendsTypeArgument) {
}
func (v *AddCastExpressionVisitor) VisitObjectType(typ intmod.IObjectType)           {}
func (v *AddCastExpressionVisitor) VisitInnerObjectType(typ intmod.IInnerObjectType) {}
func (v *AddCastExpressionVisitor) VisitWildcardSuperTypeArgument(typ intmod.IWildcardSuperTypeArgument) {
}
func (v *AddCastExpressionVisitor) VisitTypes(list intmod.ITypes) {}
func (v *AddCastExpressionVisitor) VisitTypeParameterWithTypeBounds(typ intmod.ITypeParameterWithTypeBounds) {
}
