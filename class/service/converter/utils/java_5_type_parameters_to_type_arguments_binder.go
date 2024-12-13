package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	modexp "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/expression"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvexp "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/class/service/converter/visitor"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

var GlobalRemoveNonWildcardTypeArgumentsVisitor = NewRemoveNonWildcardTypeArgumentsVisitor()

func NewJava5TypeParametersToTypeArgumentsBinder(typeMaker intsrv.ITypeMaker,
	internalTypeName string,
	comd intsrv.IClassFileConstructorOrMethodDeclaration) intsrv.ITypeParametersToTypeArgumentsBinder {
	b := &Java5TypeParametersToTypeArgumentsBinder{
		populateBindingsWithTypeParameterVisitor:            visitor.NewPopulateBindingsWithTypeParameterVisitor(),
		bindTypesToTypesVisitor:                             visitor.NewBindTypesToTypesVisitor(),
		searchInTypeArgumentVisitor:                         visitor.NewSearchInTypeArgumentVisitor(),
		typeArgumentToTypeVisitor:                           visitor.NewTypeArgumentToTypeVisitor(),
		baseTypeToTypeArgumentVisitor:                       visitor.NewBaseTypeToTypeArgumentVisitor(),
		getTypeArgumentVisitor:                              visitor.NewGetTypeArgumentVisitor(),
		bindTypeParametersToNonWildcardTypeArgumentsVisitor: visitor.NewBindTypeParametersToNonWildcardTypeArgumentsVisitor(),

		typeMaker:                               typeMaker,
		internalTypeName:                        internalTypeName,
		staticMethod:                            comd.Flags()&intmod.FlagStatic != 0,
		populateBindingsWithTypeArgumentVisitor: visitor.NewPopulateBindingsWithTypeArgumentVisitor(typeMaker),
		contextualBindings:                      comd.Bindings(),
		contextualTypeBounds:                    comd.TypeBounds(),
	}
	return b
}

type Java5TypeParametersToTypeArgumentsBinder struct {
	AbstractTypeParametersToTypeArgumentsBinder

	populateBindingsWithTypeParameterVisitor            intsrv.IPopulateBindingsWithTypeParameterVisitor
	bindTypesToTypesVisitor                             intsrv.IBindTypesToTypesVisitor
	searchInTypeArgumentVisitor                         intsrv.ISearchInTypeArgumentVisitor
	typeArgumentToTypeVisitor                           intsrv.ITypeArgumentToTypeVisitor
	baseTypeToTypeArgumentVisitor                       intsrv.IBaseTypeToTypeArgumentVisitor
	getTypeArgumentVisitor                              intsrv.IGetTypeArgumentVisitor
	bindTypeParametersToNonWildcardTypeArgumentsVisitor intsrv.IBindTypeParametersToNonWildcardTypeArgumentsVisitor

	typeMaker                               intsrv.ITypeMaker
	internalTypeName                        string
	staticMethod                            bool
	populateBindingsWithTypeArgumentVisitor intsrv.IPopulateBindingsWithTypeArgumentVisitor
	contextualBindings                      map[string]intmod.ITypeArgument
	contextualTypeBounds                    map[string]intmod.IType
	typ                                     intmod.IType
}

func (v *Java5TypeParametersToTypeArgumentsBinder) NewConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileConstructorInvocationExpression {

	bindings := make(map[string]intmod.ITypeArgument)
	parameterTypes := Clone(methodTypes.ParameterTypes())
	methodTypeParameters := methodTypes.TypeParameters()

	v.populateBindings(bindings, nil, nil, nil,
		methodTypeParameters, _type.OtTypeObject, nil, nil, nil)

	parameterTypes = v.bind(bindings, parameterTypes)
	v.bindParameters(parameterTypes, parameters)

	return srvexp.NewClassFileConstructorInvocationExpression(lineNumber, objectType, descriptor, parameterTypes, parameters)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) NewSuperConstructorInvocationExpression(
	lineNumber int, objectType intmod.IObjectType, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileSuperConstructorInvocationExpression {

	parameterTypes := Clone(methodTypes.ParameterTypes())
	bindings := v.contextualBindings
	typeTypes := v.typeMaker.MakeTypeTypes(v.internalTypeName)

	if (typeTypes != nil) && (typeTypes.SuperType() != nil) && (typeTypes.SuperType().TypeArguments() != nil) {
		superTypeTypes := v.typeMaker.MakeTypeTypes(objectType.InternalName())

		if superTypeTypes != nil {
			bindings = make(map[string]intmod.ITypeArgument)
			typeParameters := superTypeTypes.TypeParameters()
			typeArguments := typeTypes.SuperType().TypeArguments()
			methodTypeParameters := methodTypes.TypeParameters()

			v.populateBindings(bindings, nil, typeParameters, typeArguments,
				methodTypeParameters, _type.OtTypeObject, nil, nil, nil)
		}
	}

	parameterTypes = v.bind(bindings, parameterTypes)
	v.bindParameters(parameterTypes, parameters)

	return srvexp.NewClassFileSuperConstructorInvocationExpression(lineNumber, objectType, descriptor, parameterTypes, parameters)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) NewMethodInvocationExpression(
	lineNumber int, expression intmod.IExpression, objectType intmod.IObjectType, name, descriptor string,
	methodTypes intsrv.IMethodTypes, parameters intmod.IExpression) intsrv.IClassFileMethodInvocationExpression {
	return srvexp.NewClassFileMethodInvocationExpression(
		lineNumber, methodTypes.TypeParameters(), methodTypes.ReturnedType(), expression,
		objectType.InternalName(), name, descriptor, Clone(methodTypes.ParameterTypes()), parameters)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) NewFieldReferenceExpression(
	lineNumber int, typ intmod.IType, expr intmod.IExpression,
	objectType intmod.IObjectType, name, descriptor string) intmod.IFieldReferenceExpression {

	expressionType := expr.Type()

	if expressionType.IsObjectType() {
		expressionObjectType := expressionType.(intmod.IObjectType)

		if v.staticMethod || !(expressionObjectType.InternalName() == v.internalTypeName) {
			if typ.IsObjectType() {
				ot := typ.(intmod.IObjectType)

				if ot.TypeArguments() != nil {
					typeTypes := v.typeMaker.MakeTypeTypes(expressionObjectType.InternalName())

					if typeTypes == nil {
						typ = v.bind(v.contextualBindings, typ)
					} else {
						bindings := make(map[string]intmod.ITypeArgument)
						typeParameters := typeTypes.TypeParameters()
						typeArguments := expressionObjectType.TypeArguments()
						partialBinding := v.populateBindings(bindings, expr, typeParameters,
							typeArguments, nil, _type.OtTypeObject,
							nil, nil, nil)

						if !partialBinding {
							typ = v.bind(bindings, typ)
						}
					}
				}
			}
		}
	}

	return modexp.NewFieldReferenceExpressionWithAll(lineNumber, typ, expr,
		objectType.InternalName(), name, descriptor)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) BindParameterTypesWithArgumentTypes(
	typ intmod.IType, expr intmod.IExpression) {
	v.typ = typ
	expr.Accept(v)
	expr.Accept(GlobalRemoveNonWildcardTypeArgumentsVisitor)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) checkTypeArguments(typ intmod.IType, localVariable intsrv.ILocalVariable) intmod.IType {
	if typ.IsObjectType() {
		objectType := typ.(intmod.IObjectType)

		if objectType.TypeArguments() != nil {
			localVariableType := localVariable.Type()

			if localVariableType.IsObjectType() {
				localVariableObjectType := localVariableType.(intmod.IObjectType)
				typeTypes := v.typeMaker.MakeTypeTypes(localVariableObjectType.InternalName())

				if (typeTypes != nil) && (typeTypes.TypeParameters() == nil) {
					typ = typ.(intmod.IObjectType).CreateTypeWithArgs(nil)
				}
			}
		}
	}

	return typ
}

func (v *Java5TypeParametersToTypeArgumentsBinder) bindParameters(parameterTypes intmod.IType, parameters intmod.IExpression) {
	if parameterTypes != nil {
		if parameterTypes.IsList() && parameters.IsList() {
			parameterTypesIterator := parameterTypes.Iterator()
			parametersIterator := parameters.Iterator()

			for parametersIterator.HasNext() {
				parameter := parametersIterator.Next()
				v.typ = parameterTypesIterator.Next()
				parameter.Accept(v)
				parameter.Accept(GlobalRemoveNonWildcardTypeArgumentsVisitor)
			}
		} else {
			parameter := parameters.First()
			v.typ = parameterTypes.First()
			parameter.Accept(v)
			parameter.Accept(GlobalRemoveNonWildcardTypeArgumentsVisitor)
		}
	}
}

func (v *Java5TypeParametersToTypeArgumentsBinder) populateBindings(
	bindings map[string]intmod.ITypeArgument, expr intmod.IExpression,
	typeParameters intmod.ITypeParameter, typeArguments intmod.ITypeArgument, methodTypeParameters intmod.ITypeParameter,
	returnType intmod.IType, returnExpressionType intmod.IType, parameterTypes intmod.IType, parameters intmod.IExpression) bool {
	typeBounds := make(map[string]intmod.IType)
	statik := (expr != nil) && expr.IsObjectTypeReferenceExpression()

	if !statik {
		for key, value := range v.contextualBindings {
			bindings[key] = value
		}

		if typeParameters != nil {
			v.populateBindingsWithTypeParameterVisitor.Init(bindings, typeBounds)
			typeParameters.AcceptTypeParameterVisitor(v.populateBindingsWithTypeParameterVisitor)

			if typeArguments != nil {
				if typeParameters.IsList() && typeArguments.IsTypeArgumentList() {
					iteratorTypeParameter := typeParameters.Iterator()
					iteratorTypeArgument := util.NewIteratorWithSlice[intmod.ITypeArgument](typeArguments.TypeArgumentList())

					for iteratorTypeParameter.HasNext() {
						bindings[iteratorTypeParameter.Next().Identifier()] = iteratorTypeArgument.Next()
					}
				} else {
					bindings[typeParameters.First().Identifier()] = typeArguments.TypeArgumentFirst()
				}
			}
		}
	}

	if methodTypeParameters != nil {
		v.populateBindingsWithTypeParameterVisitor.Init(bindings, typeBounds)
		methodTypeParameters.AcceptTypeParameterVisitor(v.populateBindingsWithTypeParameterVisitor)
	}

	if !(_type.OtTypeObject == returnType) && returnExpressionType != nil {
		v.populateBindingsWithTypeArgumentVisitor.Init(v.contextualTypeBounds, bindings, typeBounds, returnType)
		returnExpressionType.AcceptTypeArgumentVisitor(v.populateBindingsWithTypeArgumentVisitor)
	}

	if parameterTypes != nil {
		if parameterTypes.IsList() && parameters.IsList() {
			parameterTypesIterator := parameterTypes.Iterator()
			parametersIterator := parameters.Iterator()

			for parametersIterator.HasNext() {
				v.populateBindingsWithTypeArgument(bindings, typeBounds, parameterTypesIterator.Next(), parametersIterator.Next())
			}
		} else {
			v.populateBindingsWithTypeArgument(bindings, typeBounds, parameterTypes.First(), parameters.First())
		}
	}

	bindingsContainsNull := ContainsValue(bindings, nil)

	if bindingsContainsNull {
		if v.eraseTypeArguments(expr, typeParameters, typeArguments) {
			for key, _ := range bindings {
				bindings[key] = nil
			}
		} else {
			for key, value := range bindings {
				if value == nil {
					baseType := typeBounds[key]

					if baseType == nil {
						bindings[key] = _type.WildcardTypeArgumentEmpty
					} else {
						v.bindTypesToTypesVisitor.SetBindings(bindings)
						v.bindTypesToTypesVisitor.Init()
						baseType.AcceptTypeVisitor(v.bindTypesToTypesVisitor)
						baseType = v.bindTypesToTypesVisitor.Type()

						v.baseTypeToTypeArgumentVisitor.Init()
						baseType.AcceptTypeVisitor(v.baseTypeToTypeArgumentVisitor)
						bindings[key] = v.baseTypeToTypeArgumentVisitor.TypeArgument()
					}
				}
			}
		}
	}

	return bindingsContainsNull
}

func (v *Java5TypeParametersToTypeArgumentsBinder) eraseTypeArguments(expr intmod.IExpression,
	typeParameters intmod.ITypeParameter, typeArguments intmod.ITypeArgument) bool {
	if (typeParameters != nil) && (typeArguments == nil) && (expr != nil) {
		if expr.IsCastExpression() {
			expr = expr.Expression()
		}
		if expr.IsFieldReferenceExpression() || expr.IsMethodInvocationExpression() || expr.IsLocalVariableReferenceExpression() {
			return true
		}
	}

	return false
}

func (v *Java5TypeParametersToTypeArgumentsBinder) populateBindingsWithTypeArgument(
	bindings map[string]intmod.ITypeArgument, typeBounds map[string]intmod.IType, typ intmod.IType, expr intmod.IExpression) {
	t := v.getExpressionType(expr)

	if (t != nil) && (t != _type.OtTypeUndefinedObject) {
		v.populateBindingsWithTypeArgumentVisitor.Init(v.contextualTypeBounds, bindings, typeBounds, t)
		v.typ.AcceptTypeArgumentVisitor(v.populateBindingsWithTypeArgumentVisitor)
	}
}

func (v *Java5TypeParametersToTypeArgumentsBinder) bind(bindings map[string]intmod.ITypeArgument, parameterTypes intmod.IType) intmod.IType {
	if (parameterTypes != nil) && len(bindings) != 0 {
		v.bindTypesToTypesVisitor.SetBindings(bindings)
		v.bindTypesToTypesVisitor.Init()
		parameterTypes.AcceptTypeVisitor(v.bindTypesToTypesVisitor)
		parameterTypes = v.bindTypesToTypesVisitor.Type()
	}

	return parameterTypes
}

func (v *Java5TypeParametersToTypeArgumentsBinder) getExpressionType(expr intmod.IExpression) intmod.IType {
	if expr.IsMethodInvocationExpression() {
		return v.getExpressionType2(expr.(intsrv.IClassFileMethodInvocationExpression))
	} else if expr.IsNewExpression() {
		return v.getExpressionType3(expr.(intsrv.IClassFileNewExpression))
	}

	return expr.Type()
}

func (v *Java5TypeParametersToTypeArgumentsBinder) getExpressionType2(mie intsrv.IClassFileMethodInvocationExpression) intmod.IType {
	t := mie.Type()

	v.searchInTypeArgumentVisitor.Init()
	t.AcceptTypeArgumentVisitor(v.searchInTypeArgumentVisitor)

	if !v.searchInTypeArgumentVisitor.ContainsGeneric() {
		return t
	}

	if mie.TypeParameters() != nil {
		return nil
	}

	if v.staticMethod || mie.InternalTypeName() != v.internalTypeName {
		typeTypes := v.typeMaker.MakeTypeTypes(mie.InternalTypeName())

		if (typeTypes != nil) && (typeTypes.TypeParameters() != nil) {
			return nil
		}
	}

	return t
}

func (v *Java5TypeParametersToTypeArgumentsBinder) getExpressionType3(ne intsrv.IClassFileNewExpression) intmod.IType {
	ot := ne.ObjectType()

	if v.staticMethod || ot.InternalName() != v.internalTypeName {
		typeTypes := v.typeMaker.MakeTypeTypes(ot.InternalName())

		if (typeTypes != nil) && (typeTypes.TypeParameters() != nil) {
			return nil
		}
	}

	return ot
}

// --- ExpressionVisitor --- //

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	mie := expr.(intsrv.IClassFileMethodInvocationExpression)

	if !mie.IsBound() {
		parameterTypes := mie.ParameterTypes()
		parameters := mie.Parameters()
		exp := mie.Expression()
		expressionType := exp.Type()

		if v.staticMethod || (mie.TypeParameters() != nil) || mie.InternalTypeName() != v.internalTypeName {
			typeTypes := v.typeMaker.MakeTypeTypes(mie.InternalTypeName())

			if typeTypes != nil {
				typeParameters := typeTypes.TypeParameters()
				methodTypeParameters := mie.TypeParameters()
				var typeArguments intmod.ITypeArgument

				if exp.IsSuperExpression() {
					typeTypes = v.typeMaker.MakeTypeTypes(v.internalTypeName)
					if typeTypes == nil || typeTypes.SuperType() == nil {
						typeArguments = nil
					} else {
						typeArguments = typeTypes.SuperType().TypeArguments()
					}
				} else if exp.IsMethodInvocationExpression() {
					t := v.getExpressionType2(exp.(intsrv.IClassFileMethodInvocationExpression))
					if (t != nil) && t.IsObjectType() {
						typeArguments = t.(intmod.IObjectType).TypeArguments()
					} else {
						typeArguments = nil
					}
				} else if expressionType.IsGenericType() {
					typeBound := v.contextualTypeBounds[expressionType.Name()]

					if typeBound != nil {
						v.getTypeArgumentVisitor.Init()
						typeBound.AcceptTypeVisitor(v.getTypeArgumentVisitor)
						typeArguments = v.getTypeArgumentVisitor.TypeArguments()
					} else {
						typeArguments = nil
					}
				} else {
					typeArguments = expressionType.(intmod.IObjectType).TypeArguments()
				}

				t := mie.Type()

				if v.typ.IsObjectType() && t.IsObjectType() {
					objectType := v.typ.(intmod.IObjectType)
					mieTypeObjectType := t.(intmod.IObjectType)
					t = v.typeMaker.SearchSuperParameterizedType(objectType, mieTypeObjectType)
					if t == nil {
						t = mie.Type()
					}
				}

				bindings := make(map[string]intmod.ITypeArgument)
				partialBinding := v.populateBindings(bindings, exp, typeParameters,
					typeArguments, methodTypeParameters, v.typ, t, parameterTypes, parameters)

				parameterTypes = v.bind(bindings, parameterTypes)
				mie.SetParameterTypes(parameterTypes)
				mie.SetType(v.bind(bindings, mie.Type()))

				if (methodTypeParameters != nil) && !partialBinding {
					v.bindTypeParametersToNonWildcardTypeArgumentsVisitor.Init(bindings)
					methodTypeParameters.AcceptTypeParameterVisitor(v.bindTypeParametersToNonWildcardTypeArgumentsVisitor)
					mie.SetNonWildcardTypeArguments(v.bindTypeParametersToNonWildcardTypeArgumentsVisitor.TypeArgument())
				}

				if expressionType.IsObjectType() {
					expressionObjectType := expressionType.(intmod.IObjectType)

					if len(bindings) == 0 || partialBinding {
						expressionType = expressionObjectType.CreateTypeWithArgs(nil)
					} else {
						if exp.IsObjectTypeReferenceExpression() || (typeParameters == nil) {
							expressionType = expressionObjectType.CreateTypeWithArgs(nil)
						} else if typeParameters.IsList() {
							tas := _type.NewTypeArgumentsWithCapacity(typeParameters.Size())
							for _, typeParameter := range typeParameters.ToSlice() {
								tas.Add(bindings[typeParameter.Identifier()])
							}
							expressionType = expressionObjectType.CreateTypeWithArgs(tas)
						} else {
							expressionType = expressionObjectType.CreateTypeWithArgs(bindings[typeParameters.First().Identifier()])
						}
					}
				} else if expressionType.IsGenericType() {
					if len(bindings) == 0 || partialBinding {
						expressionType = _type.OtTypeObject
					} else {
						typeArgument := bindings[expressionType.Name()]
						if typeArgument == nil {
							expressionType = _type.OtTypeObject
						} else {
							v.typeArgumentToTypeVisitor.Init()
							typeArgument.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
							expressionType = v.typeArgumentToTypeVisitor.Type()
						}
					}
				}
			}
		}

		v.typ = expressionType
		expr.Accept(v)

		v.bindParameters(parameterTypes, parameters)

		mie.SetBound(true)
	}
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	localVariable := expr.(intsrv.IClassFileLocalVariableReferenceExpression).LocalVariable().(intsrv.ILocalVariable)
	if localVariable.FromOffset() > 0 {
		// Do not update parameter
		localVariable.TypeOnLeft(v.contextualTypeBounds, v.checkTypeArguments(v.typ, localVariable))
	}
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitNewExpression(expr intmod.INewExpression) {
	ne := expr.(intsrv.IClassFileNewExpression)

	if !ne.IsBound() {
		parameterTypes := ne.ParameterTypes()
		parameters := ne.Parameters()
		neObjectType := ne.ObjectType()

		if v.staticMethod || !(neObjectType.InternalName() == v.internalTypeName) {
			typeTypes := v.typeMaker.MakeTypeTypes(neObjectType.InternalName())

			if typeTypes != nil {
				typeParameters := typeTypes.TypeParameters()
				typeArguments := neObjectType.TypeArguments()

				if (typeParameters != nil) && (typeArguments == nil) {
					if typeParameters.IsList() {
						tas := _type.NewTypeArgumentsWithCapacity(typeParameters.Size())
						for _, typeParameter := range typeParameters.ToSlice() {
							tas.Add(_type.NewGenericType(typeParameter.Identifier()))
						}
						neObjectType = neObjectType.CreateTypeWithArgs(tas)
					} else {
						neObjectType = neObjectType.CreateTypeWithArgs(_type.NewGenericType(typeParameters.First().Identifier()))
					}
				}

				t := neObjectType

				if v.typ.IsObjectType() {
					objectType := v.typ.(intmod.IObjectType)
					t = v.typeMaker.SearchSuperParameterizedType(objectType, neObjectType)
					if t == nil {
						t = neObjectType
					}
				}

				bindings := make(map[string]intmod.ITypeArgument)
				partialBinding := v.populateBindings(bindings, nil, typeParameters, typeArguments, nil, v.typ, t, parameterTypes, parameters)

				parameterTypes = v.bind(bindings, parameterTypes)
				ne.SetParameterTypes(parameterTypes)

				// Replace wildcards
				for key, value := range bindings {
					v.typeArgumentToTypeVisitor.Init()
					value.AcceptTypeArgumentVisitor(v.typeArgumentToTypeVisitor)
					bindings[key] = v.typeArgumentToTypeVisitor.Type()
				}

				if !partialBinding {
					ne.SetType(v.bind(bindings, neObjectType).(intmod.IObjectType))
				}
			}
		}

		v.bindParameters(parameterTypes, parameters)

		ne.SetBound(true)
	}
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitCastExpression(expr intmod.ICastExpression) {
	// assert TYPE_OBJECT.equals(type) || (type.Dimension() == expression.Type().Dimension()) : "TypeParametersToTypeArgumentsBinder.visit(CastExpression ce) : invalid array type";

	if v.typ.IsObjectType() {
		objectType := v.typ.(intmod.IObjectType)

		if objectType.TypeArguments() != nil && !(objectType.TypeArguments() == _type.WildcardTypeArgumentEmpty) {
			// assert expr.Type().IsObjectType() : "TypeParametersToTypeArgumentsBinder.visit(CastExpression ce) : invalid object type";

			expressionObjectType := expr.Type().(intmod.IObjectType)

			if objectType.InternalName() == expressionObjectType.InternalName() {
				expressionExpressionType := expr.Expression().Type()

				if expressionExpressionType.IsObjectType() {
					expressionExpressionObjectType := expressionExpressionType.(intmod.IObjectType)

					if expressionExpressionObjectType.TypeArguments() == nil {
						expr.SetType(objectType)
					} else if objectType.TypeArguments().IsTypeArgumentAssignableFrom(v.contextualTypeBounds, expressionExpressionObjectType.TypeArguments()) {
						expr.SetType(objectType)
					}
				} else if expressionExpressionType.IsGenericType() {
					expr.SetType(objectType)
				}
			}
		}
	}

	v.typ = expr.Type()
	expr.Expression().Accept(v)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	t := v.typ

	expr.SetType(t)
	expr.TrueExpression().Accept(v)
	v.typ = t
	expr.FalseExpression().Accept(v)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	if expr.Type() == _type.OtTypeString && "+" == expr.Operator() {
		v.typ = _type.OtTypeObject
	}

	t := v.typ

	expr.LeftExpression().Accept(v)
	v.typ = t
	expr.RightExpression().Accept(v)
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitArrayExpression(expr intmod.IArrayExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitBooleanExpression(expr intmod.IBooleanExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitCommentExpression(expr intmod.ICommentExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitExpressions(expr intmod.IExpressions) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitLengthExpression(expr intmod.ILengthExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitNewArray(expr intmod.INewArray) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitNoExpression(expr intmod.INoExpression) {

}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitNullExpression(expr intmod.INullExpression) {

}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitSuperExpression(expr intmod.ISuperExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitThisExpression(expr intmod.IThisExpression) {
}

func (v *Java5TypeParametersToTypeArgumentsBinder) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
}

func NewRemoveNonWildcardTypeArgumentsVisitor() *RemoveNonWildcardTypeArgumentsVisitor {
	return &RemoveNonWildcardTypeArgumentsVisitor{}
}

type RemoveNonWildcardTypeArgumentsVisitor struct {
	modexp.AbstractNopExpressionVisitor
}

func (v *RemoveNonWildcardTypeArgumentsVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	expr.SetNonWildcardTypeArguments(nil)
}

func ContainsValue(bindings map[string]intmod.ITypeArgument, source intmod.ITypeArgument) bool {
	for _, value := range bindings {
		if value == source {
			return true
		}
	}
	return false
}
