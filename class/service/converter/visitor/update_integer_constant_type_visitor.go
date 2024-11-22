package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"math"
)

var GlobTypes = make(map[string]intmod.IType)
var GlobDimensionTypes = &DimensionTypes{}
var GlobTypeCharacterRef = expression.NewObjectTypeReferenceExpression(_type.OtTypeCharacter)
var GlobTypeByteRef = expression.NewObjectTypeReferenceExpression(_type.OtTypeByte)
var GlobTypeShortRef = expression.NewObjectTypeReferenceExpression(_type.OtTypeShort)
var GlobTypeIntegerRef = expression.NewObjectTypeReferenceExpression(_type.OtTypeInteger)

func init() {
	c := _type.PtTypeChar
	ci := _type.NewTypes()
	ci.Add(_type.PtTypeChar)
	ci.Add(_type.PtTypeInt)

	GlobTypes["java/lang/String:indexOf(I)I"] = c
	GlobTypes["java/lang/String:indexOf(II)I"] = ci
	GlobTypes["java/lang/String:lastIndexOf(I)I"] = c
	GlobTypes["java/lang/String:lastIndexOf(II)I"] = ci
}

type UpdateIntegerConstantTypeVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	returnedType                 intmod.IType
	arrayVariableInitializerType intmod.IType
}

func (v *UpdateIntegerConstantTypeVisitor) VisitAssertStatement(state intmod.IAssertStatement) {
	state.SetCondition(v.updateBooleanExpression(state.Condition()))
}

func (v *UpdateIntegerConstantTypeVisitor) VisitDoWhileStatement(state intmod.IDoWhileStatement) {
	state.SetCondition(v.safeUpdateBooleanExpression(state.Condition()))
	v.SafeAcceptStatement(state.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitForStatement(state intmod.IForStatement) {
	v.SafeAcceptDeclaration(state.Declaration())
	v.SafeAcceptExpression(state.Init())
	state.SetCondition(v.safeUpdateBooleanExpression(state.Condition()))
	v.SafeAcceptExpression(state.Update())
	v.SafeAcceptStatement(state.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitIfStatement(state intmod.IIfStatement) {
	state.SetCondition(v.updateBooleanExpression(state.Condition()))
	v.SafeAcceptStatement(state.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitIfElseStatement(state intmod.IIfElseStatement) {
	state.SetCondition(v.updateBooleanExpression(state.Condition()))
	v.SafeAcceptStatement(state.Statements())
	state.ElseStatements().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitWhileStatement(state intmod.IWhileStatement) {
	state.SetCondition(v.updateBooleanExpression(state.Condition()))
	v.SafeAcceptStatement(state.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitReturnExpressionStatement(state intmod.IReturnExpressionStatement) {
	state.SetExpression(v.updateExpression(v.returnedType, state.Expression()))
}

func (v *UpdateIntegerConstantTypeVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	left := expr.LeftExpression()
	right := expr.RightExpression()

	leftType := left.Type()
	rightType := right.Type()

	switch expr.Operator() {
	case "&", "|", "^":
		if leftType.IsPrimitiveType() && rightType.IsPrimitiveType() {
			t := GetCommonPrimitiveType(leftType.(intmod.IPrimitiveType), rightType.(intmod.IPrimitiveType))
			if t == nil {
				t = _type.PtTypeInt
			}
			expr.SetLeftExpression(v.updateExpression(t, left))
			expr.SetRightExpression(v.updateExpression(t, right))
		}
		break
	case "=":
		left.Accept(v)
		expr.SetRightExpression(v.updateExpression(leftType, right))
		break
	case ">", ">=", "<", "<=", "==", "!=":
		if (leftType.Dimension() == 0) && (rightType.Dimension() == 0) {
			if leftType.IsPrimitiveType() {
				if rightType.IsPrimitiveType() {
					var t intmod.IType
					if leftType == rightType {
						t = leftType
					} else {
						t = GetCommonPrimitiveType(leftType.(intmod.IPrimitiveType), rightType.(intmod.IPrimitiveType))
						if t == nil {
							t = _type.PtTypeInt
						}
					}
					expr.SetLeftExpression(v.updateExpression(t, left))
					expr.SetRightExpression(v.updateExpression(t, right))
				} else {
					expr.SetLeftExpression(v.updateExpression(rightType, left))
					right.Accept(v)
				}
				break
			} else if rightType.IsPrimitiveType() {
				left.Accept(v)
				expr.SetRightExpression(v.updateExpression(leftType, right))
				break
			}
		}

		left.Accept(v)
		right.Accept(v)
		break
	default:
		expr.SetRightExpression(v.updateExpression(expr.Type(), right))
		expr.SetLeftExpression(v.updateExpression(expr.Type(), left))
		break
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	rt := v.returnedType
	v.returnedType = expr.ReturnedType()
	v.SafeAcceptStatement(expr.Statements())
	v.returnedType = rt
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	parameters := expr.Parameters()

	if parameters != nil {
		expr.SetParameters(v.updateExpressions(
			expr.(intsrv.IClassFileSuperConstructorInvocationExpression).ParameterTypes(), parameters))
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	parameters := expr.Parameters()

	if parameters != nil {
		expr.SetParameters(v.updateExpressions(
			expr.(intsrv.IClassFileConstructorInvocationExpression).ParameterTypes(), parameters))
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	parameters := expr.Parameters()

	if parameters != nil {
		internalTypeName := expr.InternalTypeName()
		name := expr.Name()
		descriptor := expr.Descriptor()
		types := GlobTypes[internalTypeName+":"+name+descriptor]

		if types == nil {
			types = expr.(intsrv.IClassFileMethodInvocationExpression).ParameterTypes()
		}

		expr.SetParameters(v.updateExpressions(types, parameters))
	}

	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitNewExpression(expr intmod.INewExpression) {
	parameters := expr.Parameters()

	if parameters != nil {
		internalTypeName := expr.ObjectType().InternalName()
		descriptor := expr.Descriptor()
		types := GlobTypes[internalTypeName+":<init>"+descriptor]

		if types == nil {
			types = expr.(intsrv.IClassFileNewExpression).ParameterTypes()
		}

		expr.SetParameters(v.updateExpressions(types, parameters))
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitNewArray(expr intmod.INewArray) {
	dimensions := expr.DimensionExpressionList()

	if dimensions != nil {
		v.updateExpressions(GlobDimensionTypes, dimensions)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	expr.Expression().Accept(v)
	expr.SetIndex(v.updateExpression(_type.PtTypeInt, expr.Index()))
}

func (v *UpdateIntegerConstantTypeVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	expr.SetExpression(v.updateExpression(expr.Type(), expr.Expression()))
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	trueType := expr.TrueExpression().Type()
	falseType := expr.FalseExpression().Type()

	expr.SetCondition(v.updateBooleanExpression(expr.Condition()))

	if trueType.IsPrimitiveType() {
		if falseType.IsPrimitiveType() {
			expr.SetTrueExpression(v.updateExpression(_type.PtTypeInt, expr.TrueExpression()))
			expr.SetFalseExpression(v.updateExpression(_type.PtTypeInt, expr.FalseExpression()))
		} else {
			expr.TrueExpression().Accept(v)
			expr.SetTrueExpression(v.updateExpression(falseType, expr.TrueExpression()))
		}
	} else {
		if falseType.IsPrimitiveType() {
			expr.SetFalseExpression(v.updateExpression(trueType, expr.FalseExpression()))
			expr.FalseExpression().Accept(v)
		} else {
			expr.TrueExpression().Accept(v)
			expr.FalseExpression().Accept(v)
		}
	}
}

func (v *UpdateIntegerConstantTypeVisitor) VisitArrayVariableInitializer(decl intmod.IArrayVariableInitializer) {
	t := v.arrayVariableInitializerType
	v.arrayVariableInitializerType = decl.Type()
	// v.acceptListDeclaration(decl);
	for _, item := range decl.Elements() {
		item.Accept(v)
	}
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	t := v.arrayVariableInitializerType
	v.arrayVariableInitializerType = decl.Type()
	decl.LocalVariableDeclarators().Accept(v)
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := v.arrayVariableInitializerType
	v.arrayVariableInitializerType = decl.Type()
	decl.FieldDeclarators().Accept(v)
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	if decl != nil {
		decl.SetExpression(v.updateExpression(v.arrayVariableInitializerType, decl.Expression()))
	}
}

func (v *UpdateIntegerConstantTypeVisitor) updateExpressions(types intmod.IType, expressions intmod.IExpression) intmod.IExpression {
	if expressions.IsList() {
		typ := sliceToDefaultList[intmod.IType](types.List())
		e := sliceToDefaultList[intmod.IExpression](expressions.List())

		for i := e.Size() - 1; i >= 0; i-- {
			t := typ.Get(i)

			if t.Dimension() == 0 && t.IsPrimitiveType() {
				parameter := e.Get(i)
				updatedParameter := v.updateExpression(t, parameter)

				if updatedParameter.IsIntegerConstantExpression() {
					switch t.(intmod.IPrimitiveType).JavaPrimitiveFlags() {
					case intmod.FlagByte, intmod.FlagShort:
						updatedParameter = expression.NewCastExpression(t, updatedParameter)
					}
				}

				e.Set(i, updatedParameter)
			}
		}
	} else {
		t := types.First()

		if t.Dimension() == 0 && t.IsPrimitiveType() {
			updatedParameter := v.updateExpression(t, expressions)

			if updatedParameter.IsIntegerConstantExpression() {
				switch t.(intmod.IPrimitiveType).JavaPrimitiveFlags() {
				case intmod.FlagByte, intmod.FlagShort:
					updatedParameter = expression.NewCastExpression(t, updatedParameter)
				}
			}

			expressions = updatedParameter
		}
	}

	expressions.Accept(v)
	return expressions
}

func (v *UpdateIntegerConstantTypeVisitor) updateExpression(t intmod.IType, expr intmod.IExpression) intmod.IExpression {
	// assert type != TYPE_VOID : "UpdateIntegerConstantTypeVisitorupdateexpr.(type, intmod.IupdateExpression expr) : try to set 'void' to a numeric expression";

	if (t != expr.Type()) && expr.IsIntegerConstantExpression() {
		if _type.OtTypeString.(intmod.IType) == t {
			t = _type.PtTypeChar
		}

		if t.IsPrimitiveType() {
			primitiveType := t.(intmod.IPrimitiveType)
			ice := expr.(intmod.IIntegerConstantExpression)
			icePrimitiveType := ice.Type().(intmod.IPrimitiveType)
			value := ice.IntegerValue()
			lineNumber := ice.LineNumber()

			switch primitiveType.JavaPrimitiveFlags() {
			case intmod.FlagBoolean:
				return expression.NewBooleanExpressionWithLineNumber(lineNumber, value != 0)
			case intmod.FlagChar:
				switch value {
				case math.MinInt16:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeChar, GlobTypeCharacterRef, "java/lang/Character", "MIN_VALUE", "C")
				case math.MaxInt16:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeChar, GlobTypeCharacterRef, "java/lang/Character", "MAX_VALUE", "C")
				default:
					if (icePrimitiveType.Flags() & primitiveType.Flags()) != 0 {
						ice.SetType(t)
					} else {
						ice.SetType(_type.PtTypeInt)
					}
					break
				}
				break
			case intmod.FlagByte:
				switch value {
				case math.MinInt8:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeByte, GlobTypeByteRef, "java/lang/Byte", "MIN_VALUE", "B")
				case math.MaxInt8:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeByte, GlobTypeByteRef, "java/lang/Byte", "MAX_VALUE", "B")
				default:
					if (icePrimitiveType.Flags() & primitiveType.Flags()) != 0 {
						ice.SetType(t)
					} else {
						ice.SetType(_type.PtTypeInt)
					}
					break
				}
				break
			case intmod.FlagShort:
				switch value {
				case math.MinInt16:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeShort, GlobTypeShortRef, "java/lang/Short", "MIN_VALUE", "S")
				case math.MaxInt16:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeShort, GlobTypeShortRef, "java/lang/Short", "MAX_VALUE", "S")
				default:
					if (icePrimitiveType.Flags() & primitiveType.Flags()) != 0 {
						ice.SetType(t)
					} else {
						ice.SetType(_type.PtTypeInt)
					}
					break
				}
				break
			case intmod.FlagInt:
				switch value {
				case math.MinInt32:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeInt, GlobTypeIntegerRef, "java/lang/Integer", "MIN_VALUE", "I")
				case math.MaxInt32:
					return expression.NewFieldReferenceExpressionWithAll(lineNumber, _type.PtTypeInt, GlobTypeIntegerRef, "java/lang/Integer", "MAX_VALUE", "I")
				default:
					if (icePrimitiveType.Flags() & primitiveType.Flags()) != 0 {
						ice.SetType(t)
					} else {
						ice.SetType(_type.PtTypeInt)
					}
					break
				}
				break
			case intmod.FlagLong:
				return expression.NewLongConstantExpressionWithAll(ice.LineNumber(), int64(ice.IntegerValue()))
			}

			return expr
		}
	}

	if t.IsPrimitiveType() && expr.IsTernaryOperatorExpression() {
		toe := expr.(intmod.ITernaryOperatorExpression)

		toe.SetType(t)
		toe.SetCondition(v.updateBooleanExpression(toe.Condition()))
		toe.SetTrueExpression(v.updateExpression(t, toe.TrueExpression()))
		toe.SetFalseExpression(v.updateExpression(t, toe.FalseExpression()))

		return expr
	}

	expr.Accept(v)
	return expr
}

func (v *UpdateIntegerConstantTypeVisitor) safeUpdateBooleanExpression(expr intmod.IExpression) intmod.IExpression {
	if expr == nil {
		return nil
	}
	return v.updateBooleanExpression(expr)
}

func (v *UpdateIntegerConstantTypeVisitor) updateBooleanExpression(expr intmod.IExpression) intmod.IExpression {
	if _type.PtTypeBoolean != expr.Type() {
		if expr.IsIntegerConstantExpression() {
			return expression.NewBooleanExpressionWithLineNumber(expr.LineNumber(), expr.IntegerValue() != 0)
		} else if expr.IsTernaryOperatorExpression() {
			toe := expr.(intmod.ITernaryOperatorExpression)

			toe.SetType(_type.PtTypeBoolean)
			toe.SetCondition(v.updateBooleanExpression(toe.Condition()))
			toe.SetTrueExpression(v.updateBooleanExpression(toe.TrueExpression()))
			toe.SetFalseExpression(v.updateBooleanExpression(toe.FalseExpression()))

			return expr
		}
	}

	expr.Accept(v)
	return expr
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitBreakStatement(state intmod.IBreakStatement)       {}
func (v *UpdateIntegerConstantTypeVisitor) VisitByteCodeStatement(state intmod.IByteCodeStatement) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitContinueStatement(state intmod.IContinueStatement) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitNullExpression(expr intmod.INullExpression)        {}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitThisExpression(expr intmod.IThisExpression)   {}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectReference(reference intmod.IObjectReference) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitInnerObjectReference(reference intmod.IInnerObjectReference) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeArguments(t intmod.ITypeArguments) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitWildcardExtendsTypeArgument(t intmod.IWildcardExtendsTypeArgument) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectType(t intmod.IObjectType)           {}
func (v *UpdateIntegerConstantTypeVisitor) VisitInnerObjectType(t intmod.IInnerObjectType) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitWildcardSuperTypeArgument(t intmod.IWildcardSuperTypeArgument) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypes(list intmod.ITypes) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeParameterWithTypeBounds(t intmod.ITypeParameterWithTypeBounds) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {}

type DimensionTypes struct {
	_type.Types
}

func (t *DimensionTypes) First() intmod.IType    { return _type.PtTypeInt }
func (t *DimensionTypes) Last() intmod.IType     { return _type.PtTypeInt }
func (t *DimensionTypes) Get(i int) intmod.IType { return _type.PtTypeInt }
func (t *DimensionTypes) Size() int              { return 0 }

func GetPrimitiveTypeFromValue(value int) intmod.IPrimitiveType {
	if value >= 0 {
		if value <= 1 {
			return _type.PtMaybeBooleanType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt8 {
			return _type.PtMaybeByteType.(intmod.IPrimitiveType)
		}
		if value <= math.MaxInt16 {
			return _type.PtMaybeShortType.(intmod.IPrimitiveType)
		}
		if value <= '\uFFFF' {
			return _type.PtMaybeCharType.(intmod.IPrimitiveType)
		}
	} else {
		if value >= math.MinInt8 {
			return _type.PtMaybeNegativeByteType.(intmod.IPrimitiveType)
		}
		if value >= math.MinInt16 {
			return _type.PtMaybeNegativeShortType.(intmod.IPrimitiveType)
		}
	}
	return _type.PtMaybeIntType
}

func GetCommonPrimitiveType(pt1, pt2 intmod.IPrimitiveType) intmod.IPrimitiveType {
	return GetPrimitiveTypeFromFlags(pt1.Flags() & pt2.Flags())
}

func GetPrimitiveTypeFromFlags(flags int) intmod.IPrimitiveType {
	switch flags {
	case intmod.FlagBoolean:
		return _type.PtTypeBoolean
	case intmod.FlagChar:
		return _type.PtTypeChar
	case intmod.FlagFloat:
		return _type.PtTypeFloat
	case intmod.FlagDouble:
		return _type.PtTypeDouble
	case intmod.FlagByte:
		return _type.PtTypeByte
	case intmod.FlagShort:
		return _type.PtTypeShort
	case intmod.FlagInt:
		return _type.PtTypeInt
	case intmod.FlagLong:
		return _type.PtTypeLong
	case intmod.FlagVoid:
		return _type.PtTypeVoid
	default:
		if flags == intmod.FlagChar|intmod.FlagInt {
			return _type.PtMaybeCharType
		}
		if flags == intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeShortType
		}
		if flags == intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeByteType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagChar|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeBooleanType
		}
		if flags == intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeByteType
		}
		if flags == intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeShortType
		}
		if flags == intmod.FlagBoolean|intmod.FlagByte|intmod.FlagShort|intmod.FlagInt {
			return _type.PtMaybeNegativeBooleanType
		}
	}

	return nil
}

func GetPrimitiveTypeFromTag(tag int) intmod.IType {
	switch tag {
	case 4:
		return _type.PtTypeBoolean.(intmod.IType)
	case 5:
		return _type.PtTypeChar.(intmod.IType)
	case 6:
		return _type.PtTypeFloat.(intmod.IType)
	case 7:
		return _type.PtTypeDouble.(intmod.IType)
	case 8:
		return _type.PtTypeByte.(intmod.IType)
	case 9:
		return _type.PtTypeShort.(intmod.IType)
	case 10:
		return _type.PtTypeInt.(intmod.IType)
	case 11:
		return _type.PtTypeLong.(intmod.IType)
	default:
		return nil
	}
}

func sliceToDefaultList[T any](slice []T) util.DefaultList[T] {
	ret := util.DefaultList[T]{}
	for _, item := range slice {
		ret.Add(item)
	}
	return ret
}