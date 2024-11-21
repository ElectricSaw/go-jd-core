package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/statement"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewAddCastExpressionVisitor(typeMaker *utils.TypeMaker) *AddCastExpressionVisitor {
	return &AddCastExpressionVisitor{
		searchFirstLineNumberVisitor: NewSearchFirstLineNumberVisitor(),
		typeMaker:                    typeMaker,
	}
}

type AddCastExpressionVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor *SearchFirstLineNumberVisitor
	typeMaker                    *utils.TypeMaker
	typeBounds                   map[string]intmod.IType
	returnedType                 intmod.IType
	exceptionType                intmod.IType
	typ                          intmod.IType
}

func (v *AddCastExpressionVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	memberDeclarations := decl.MemberDeclarations()

	if memberDeclarations != nil {
		//tb := v.typeBounds

		//if obj, ok := decl.(*servdecl.ClassFileBodyDeclaration); ok {
		//	v.typeBounds = obj.TypeBounds()
		//	memberDeclarations.Accept(v)
		//	v.typeBounds = tb
		//}
	}
}

func (v *AddCastExpressionVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	if (decl.Flags() & intmod.FlagSynthetic) == 0 {
		t := v.typ

		v.typ = decl.Type()
		decl.FieldDeclarators().Accept(v)
		v.typ = t
	}
}

func (v *AddCastExpressionVisitor) VisitFieldDeclarator(declarator intmod.IFieldDeclarator) {
	variableInitializer := declarator.VariableInitializer()

	if variableInitializer != nil {
		extraDimension := declarator.Dimension()

		if extraDimension == 0 {
			variableInitializer.Accept(v)
		} else {
			t := v.typ

			v.typ = v.typ.CreateType(v.typ.Dimension() + extraDimension)
			variableInitializer.Accept(v)
			v.typ = t
		}
	}
}

//func (v *AddCastExpressionVisitor) VisitStaticInitializerDeclaration( decl *declaration.StaticInitializerDeclaration) {
//	statements := decl.Statements();
//
//	if (statements != nil) {
//		tb := v.typeBounds;
//
//		typeBounds = decl.(*servdecl.ClassFileStaticInitializerDeclaration).TypeBounds();
//		statements.Accept(v);
//		typeBounds = tb;
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitConstructorDeclaration( declaration *declaration.ConstructorDeclaration) {
//	if ((declaration.Flags() & (FLAG_SYNTHETIC|FLAG_BRIDGE)) == 0) {
//		BaseStatement statements = declaration.Statements();
//
//		if (statements != nil) {
//			Map<String, BaseType> tb = typeBounds;
//			BaseType et = exceptionTypes;
//
//			typeBounds = ((ClassFileConstructorDeclaration) declaration).TypeBounds();
//			exceptionTypes = declaration.ExceptionTypes();
//			statements.Accept(v);
//			typeBounds = tb;
//			exceptionTypes = et;
//		}
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitMethodDeclaration( declaration *declaration.MethodDeclaration) {
//	if ((declaration.Flags() & (FLAG_SYNTHETIC|FLAG_BRIDGE)) == 0) {
//		BaseStatement statements = declaration.Statements();
//
//		if (statements != nil) {
//			Map<String, BaseType> tb = typeBounds;
//			Type rt = returnedType;
//			BaseType et = exceptionTypes;
//
//			typeBounds = ((ClassFileMethodDeclaration) declaration).TypeBounds();
//			returnedType = declaration.ReturnedType();
//			exceptionTypes = declaration.ExceptionTypes();
//			statements.Accept(v);
//			typeBounds = tb;
//			returnedType = rt;
//			exceptionTypes = et;
//		}
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitLambdaIdentifiersExpression( expression intmod.ILambdaIdentifiersExpression) {
//	BaseStatement statements = expression.Statements();
//
//	if (statements != nil) {
//		Type rt = returnedType;
//
//		returnedType = ObjectType.TYPE_OBJECT;
//		statements.Accept(v);
//		returnedType = rt;
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitReturnExpressionStatement( statement *statement.ReturnExpressionStatement) {
//	statement.setExpression(updateExpression(returnedType, statement.IExpression(), false, true));
//}
//
//func (v *AddCastExpressionVisitor) VisitThrowStatement( statement *ThrowStatement) {
//	if ((exceptionTypes != nil) && (exceptionTypes.size() == 1)) {
//		Type exceptionType = exceptionTypes.First();
//
//		if (exceptionType.isGenericType() && !statement.IExpression().Type().equals(exceptionType)) {
//			statement.setExpression(addCastExpression(exceptionType, statement.IExpression()));
//		}
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitLocalVariableDeclaration( declaration *declaration.LocalVariableDeclaration) {
//	Type t = type;
//
//	type = declaration.Type();
//	declaration.LocalVariableDeclarators().Accept(v);
//	type = t;
//}
//
//func (v *AddCastExpressionVisitor) VisitLocalVariableDeclarator( declarator *declaration.LocalVariableDeclarator) {
//	VariableInitializer variableInitializer = declarator.VariableInitializer();
//
//	if (variableInitializer != nil) {
//		int extraDimension = declarator.Dimension();
//
//		if (extraDimension == 0) {
//			variableInitializer.Accept(v);
//		} else {
//			Type t = type;
//
//			type = type.createType(type.Dimension() + extraDimension);
//			variableInitializer.Accept(v);
//			type = t;
//		}
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitArrayVariableInitializer( declaration *declaration.ArrayVariableInitializer) {
//	if (type.Dimension() == 0) {
//		AcceptListDeclaration(declaration);
//	} else {
//		Type t = type;
//
//		type = type.createType(type.Dimension() - 1);
//		AcceptListDeclaration(declaration);
//		type = t;
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitExpressionVariableInitializer( declaration *declaration.ExpressionVariableInitializer) {
//	IExpression expression = declaration.IExpression();
//
//	if (expression.isNewInitializedArray()) {
//		NewInitializedArray nia = (NewInitializedArray)expression;
//		Type t = type;
//
//		type = nia.Type();
//		nia.ArrayInitializer().Accept(v);
//		type = t;
//	} else {
//		declaration.setExpression(updateExpression(type, expression, false, true));
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitSuperConstructorInvocationExpression( expression intmod.ISuperConstructorInvocationExpression) {
//	BaseExpression parameters = expression.Parameters();
//
//	if ((parameters != nil) && (parameters.size() > 0)) {
//		boolean unique = typeMaker.matchCount(expression.ObjectType().InternalName(), "<init>", parameters.size(), true) <= 1;
//		boolean forceCast = !unique && (typeMaker.matchCount(typeBounds, expression.ObjectType().InternalName(), "<init>", parameters, true) > 1);
//		expression.setParameters(updateParameters(((ClassFileSuperConstructorInvocationExpression)expression).ParameterTypes(), parameters, forceCast, unique));
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitConstructorInvocationExpression( expression intmod.IConstructorInvocationExpression) {
//	BaseExpression parameters = expression.Parameters();
//
//	if ((parameters != nil) && (parameters.size() > 0)) {
//		boolean unique = typeMaker.matchCount(expression.ObjectType().InternalName(), "<init>", parameters.size(), true) <= 1;
//		boolean forceCast = !unique && (typeMaker.matchCount(typeBounds, expression.ObjectType().InternalName(), "<init>", parameters, true) > 1);
//		expression.setParameters(updateParameters(((ClassFileConstructorInvocationExpression)expression).ParameterTypes(), parameters, forceCast, unique));
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitMethodInvocationExpression( expression intmod.IMethodInvocationExpression) {
//	BaseExpression parameters = expression.Parameters();
//
//	if ((parameters != nil) && (parameters.size() > 0)) {
//		boolean unique = typeMaker.matchCount(expression.InternalTypeName(), expression.Name(), parameters.size(), false) <= 1;
//		boolean forceCast = !unique && (typeMaker.matchCount(typeBounds, expression.InternalTypeName(), expression.Name(), parameters, false) > 1);
//		expression.setParameters(updateParameters(((ClassFileMethodInvocationExpression)expression).ParameterTypes(), parameters, forceCast, unique));
//	}
//
//	expression.IExpression().Accept(v);
//}
//
//func (v *AddCastExpressionVisitor) VisitNewExpression( expression intmod.INewExpression) {
//	BaseExpression parameters = expression.Parameters();
//
//	if (parameters != nil) {
//		boolean unique = typeMaker.matchCount(expression.ObjectType().InternalName(), "<init>", parameters.size(), true) <= 1;
//		boolean forceCast = !unique && (typeMaker.matchCount(typeBounds, expression.ObjectType().InternalName(), "<init>", parameters, true) > 1);
//		expression.setParameters(updateParameters(((ClassFileNewExpression)expression).ParameterTypes(), parameters, forceCast, unique));
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitNewInitializedArray( expression intmod.INewInitializedArray) {
//	ArrayVariableInitializer arrayInitializer = expression.ArrayInitializer();
//
//	if (arrayInitializer != nil) {
//		Type t = type;
//
//		type = expression.Type();
//		arrayInitializer.Accept(v);
//		type = t;
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitFieldReferenceExpression( expression intmod.IFieldReferenceExpression) {
//	IExpression exp = expression.IExpression();
//
//	if ((exp != nil) && !exp.isObjectTypeReferenceExpression()) {
//		Type type = typeMaker.makeFromInternalTypeName(expression.InternalTypeName());
//
//		if (type.Name() != nil) {
//			expression.setExpression(updateExpression(type, exp, false, true));
//		}
//	}
//}
//
//func (v *AddCastExpressionVisitor) VisitBinaryOperatorExpression( expression intmod.IBinaryOperatorExpression) {
//	expression.LeftExpression().Accept(v);
//
//	IExpression rightExpression = expression.RightExpression();
//
//	if (expression.Operator().equals("=")) {
//		if (rightExpression.isMethodInvocationExpression()) {
//			ClassFileMethodInvocationExpression mie = (ClassFileMethodInvocationExpression)rightExpression;
//
//			if (mie.TypeParameters() != nil) {
//				// Do not add cast expression if method contains type parameters
//				rightExpression.Accept(v);
//				return;
//			}
//		}
//
//		expression.setRightExpression(updateExpression(expression.LeftExpression().Type(), rightExpression, false, true));
//		return;
//	}
//
//	rightExpression.Accept(v);
//}
//
//func (v *AddCastExpressionVisitor) VisitTernaryOperatorExpression( expression intmod.ITernaryOperatorExpression) {
//	Type expressionType = expression.Type();
//
//	expression.Condition().Accept(v);
//	expression.setTrueExpression(updateExpression(expressionType, expression.TrueExpression(), false, true));
//	expression.setFalseExpression(updateExpression(expressionType, expression.FalseExpression(), false, true));
//}
//
//func (v *AddCastExpressionVisitor)  updateParameters(types _type.IType, expression expression.IExpression,  forceCast bool,  unique bool) expression.IExpression
//if (expressions != nil) {
//if (expressions.isList()) {
//DefaultList<Type> typeList = types.List();
//DefaultList<IExpression> expressionList = expressions.List();
//
//for (int i = expressionList.size() - 1; i >= 0; i--) {
//expressionList.set(i, updateParameter(typeList.(i), expressionList.(i), forceCast, unique));
//}
//} else {
//expressions = updateParameter(types.First(), expressions.First(), forceCast, unique);
//}
//}
//
//return expressions;
//}
//
//func (v *AddCastExpressionVisitor)  updateParameter(typ _type.IType, expression expression.IExpression,  forceCast bool,  unique bool) expression.IExpression {
//	expression = updateExpression(type, expression, forceCast, unique);
//
//	if (type == TYPE_BYTE) {
//		if (expression.isIntegerConstantExpression()) {
//			expression = new CastExpression(TYPE_BYTE, expression);
//		} else if (expression.isTernaryOperatorExpression()) {
//			IExpression exp = expression.TrueExpression();
//
//			if (exp.isIntegerConstantExpression() || exp.isTernaryOperatorExpression()) {
//				expression = new CastExpression(TYPE_BYTE, expression);
//			} else {
//				exp = expression.FalseExpression();
//
//				if (exp.isIntegerConstantExpression() || exp.isTernaryOperatorExpression()) {
//					expression = new CastExpression(TYPE_BYTE, expression);
//				}
//			}
//		}
//	}
//
//	return expression;
//}
//
//func (v *AddCastExpressionVisitor)  updateExpression(typ _type.IType, expression expression.IExpression, forceCast bool,  unique bool)  expression.IExpression {
//	if (expression.isnilExpression()) {
//		if (forceCast) {
//			searchFirstLineNumberVisitor.init();
//			expression.Accept(searchFirstLineNumberVisitor);
//			expression = new CastExpression(searchFirstLineNumberVisitor.LineNumber(), type, expression);
//		}
//	} else {
//		Type expressionType = expression.Type();
//
//		if (!expressionType.equals(type)) {
//		if (type.isObjectType()) {
//		if (expressionType.isObjectType()) {
//		ObjectType objectType = (ObjectType) type;
//		ObjectType expressionObjectType = (ObjectType) expressionType;
//
//		if (forceCast && !objectType.rawEquals(expressionObjectType)) {
//		// Force disambiguation of method invocation => Add cast
//		if (expression.isNewExpression()) {
//		ClassFileNewExpression ne = (ClassFileNewExpression)expression;
//		ne.setObjectType(ne.ObjectType().createType(nil));
//		}
//		expression = addCastExpression(objectType, expression);
//		} else if (!ObjectType.TYPE_OBJECT.equals(type) && !typeMaker.isAssignable(typeBounds, objectType, expressionObjectType)) {
//		BaseTypeArgument ta1 = objectType.TypeArguments();
//		BaseTypeArgument ta2 = expressionObjectType.TypeArguments();
//		Type t = type;
//
//		if ((ta1 != nil) && (ta2 != nil) && !ta1.isTypeArgumentAssignableFrom(typeBounds, ta2)) {
//		// Incompatible typeArgument arguments => Add cast
//		t = objectType.createType(nil);
//		}
//		expression = addCastExpression(t, expression);
//		}
//		} else if (expressionType.isGenericType() && !ObjectType.TYPE_OBJECT.equals(type)) {
//		expression = addCastExpression(type, expression);
//		}
//		} else if (type.isGenericType()) {
//		if (expressionType.isObjectType() || expressionType.isGenericType()) {
//		expression = addCastExpression(type, expression);
//		}
//		}
//		}
//
//		if (expression.isCastExpression()) {
//			Type ceExpressionType = expression.IExpression().Type();
//
//			if (type.isObjectType() && ceExpressionType.isObjectType()) {
//				ObjectType ot1 = (ObjectType)type;
//				ObjectType ot2 = (ObjectType)ceExpressionType;
//
//				if (ot1.equals(ot2)) {
//					// Remove cast expression
//					expression = expression.IExpression();
//				} else if (unique && typeMaker.isAssignable(typeBounds, ot1, ot2)) {
//					// Remove cast expression
//					expression = expression.IExpression();
//				}
//			}
//		}
//
//		expression.Accept(v);
//	}
//
//	return expression;
//}
//
//func (v *AddCastExpressionVisitor)  addCastExpression( typ _type.IType,  expression expression.IExpression) expression.IExpression{
//	if (expression.isCastExpression()) {
//		if (type.equals(expression.IExpression().Type())) {
//			return expression.IExpression();
//		} else {
//			CastExpression ce = (CastExpression)expression;
//
//			ce.setType(type);
//			return ce;
//		}
//	} else {
//		searchFirstLineNumberVisitor.init();
//		expression.Accept(searchFirstLineNumberVisitor);
//		return new CastExpression(searchFirstLineNumberVisitor.LineNumber(), type, expression);
//	}
//}

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
func (v *AddCastExpressionVisitor) VisitBreakStatement(statement *statement.BreakStatement)       {}
func (v *AddCastExpressionVisitor) VisitByteCodeStatement(statement *statement.ByteCodeStatement) {}
func (v *AddCastExpressionVisitor) VisitContinueStatement(statement *statement.ContinueStatement) {}
func (v *AddCastExpressionVisitor) VisitNullExpression(expression intmod.INullExpression)         {}
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
