package visitor

import (
	"math"

	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/expression"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor/utils"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
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

func NewUpdateIntegerConstantTypeVisitor(returnedType intmod.IType) intsrv.IUpdateIntegerConstantTypeVisitor {
	return &UpdateIntegerConstantTypeVisitor{
		returnedType: returnedType,
	}
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
	state.ElseStatements().AcceptStatement(v)
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
			t := utils.GetCommonPrimitiveType(leftType.(intmod.IPrimitiveType), rightType.(intmod.IPrimitiveType))
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
						t = utils.GetCommonPrimitiveType(leftType.(intmod.IPrimitiveType), rightType.(intmod.IPrimitiveType))
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
	for _, item := range decl.ToSlice() {
		item.AcceptDeclaration(v)
	}
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableDeclaration(decl intmod.ILocalVariableDeclaration) {
	t := v.arrayVariableInitializerType
	v.arrayVariableInitializerType = decl.Type()
	decl.LocalVariableDeclarators().AcceptDeclaration(v)
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFieldDeclaration(decl intmod.IFieldDeclaration) {
	t := v.arrayVariableInitializerType
	v.arrayVariableInitializerType = decl.Type()
	decl.FieldDeclarators().AcceptDeclaration(v)
	v.arrayVariableInitializerType = t
}

func (v *UpdateIntegerConstantTypeVisitor) VisitExpressionVariableInitializer(decl intmod.IExpressionVariableInitializer) {
	if decl != nil {
		decl.SetExpression(v.updateExpression(v.arrayVariableInitializerType, decl.Expression()))
	}
}

func (v *UpdateIntegerConstantTypeVisitor) updateExpressions(types intmod.IType, expressions intmod.IExpression) intmod.IExpression {
	if expressions.IsList() {
		typ := util.NewDefaultListWithSlice[intmod.IType](types.ToSlice())
		e := util.NewDefaultListWithSlice[intmod.IExpression](expressions.ToSlice())

		for i := e.Size() - 1; i >= 0; i-- {
			t := typ.Get(i)

			if t.Dimension() == 0 && t.IsPrimitiveType() {
				parameter := e.Get(i)
				updatedParameter := v.updateExpression(t, parameter)

				if updatedParameter.IsIntegerConstantExpression() {
					switch t.(intmod.IPrimitiveType).JavaPrimitiveFlags() {
					case intmod.FlagByte, intmod.FlagShort:
						updatedParameter = expression.NewCastExpression(t, updatedParameter)
					default:
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
				default:
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
			default:
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

func (v *UpdateIntegerConstantTypeVisitor) VisitFloatConstantExpression(_ intmod.IFloatConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitIntegerConstantExpression(_ intmod.IIntegerConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitConstructorReferenceExpression(_ intmod.IConstructorReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitDoubleConstantExpression(_ intmod.IDoubleConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitEnumConstantReferenceExpression(_ intmod.IEnumConstantReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitLongConstantExpression(_ intmod.ILongConstantExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitBreakStatement(_ intmod.IBreakStatement)       {}
func (v *UpdateIntegerConstantTypeVisitor) VisitByteCodeStatement(_ intmod.IByteCodeStatement) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitContinueStatement(_ intmod.IContinueStatement) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitNullExpression(_ intmod.INullExpression)       {}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectTypeReferenceExpression(_ intmod.IObjectTypeReferenceExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitSuperExpression(_ intmod.ISuperExpression) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitThisExpression(_ intmod.IThisExpression)   {}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeReferenceDotClassExpression(_ intmod.ITypeReferenceDotClassExpression) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectReference(_ intmod.IObjectReference) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitInnerObjectReference(_ intmod.IInnerObjectReference) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeArguments(_ intmod.ITypeArguments) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitWildcardExtendsTypeArgument(_ intmod.IWildcardExtendsTypeArgument) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitObjectType(_ intmod.IObjectType)           {}
func (v *UpdateIntegerConstantTypeVisitor) VisitInnerObjectType(_ intmod.IInnerObjectType) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitWildcardSuperTypeArgument(_ intmod.IWildcardSuperTypeArgument) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypes(_ intmod.ITypes) {}
func (v *UpdateIntegerConstantTypeVisitor) VisitTypeParameterWithTypeBounds(_ intmod.ITypeParameterWithTypeBounds) {
}
func (v *UpdateIntegerConstantTypeVisitor) VisitBodyDeclaration(_ intmod.IBodyDeclaration) {}

type DimensionTypes struct {
	_type.Types
}

func (t *DimensionTypes) First() intmod.IType    { return _type.PtTypeInt }
func (t *DimensionTypes) Last() intmod.IType     { return _type.PtTypeInt }
func (t *DimensionTypes) Get(_ int) intmod.IType { return _type.PtTypeInt }
func (t *DimensionTypes) Size() int              { return 0 }

// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
// $$$                           $$$
// $$$ AbstractJavaSyntaxVisitor $$$
// $$$                           $$$
// $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

func (v *UpdateIntegerConstantTypeVisitor) VisitCompilationUnit(compilationUnit intmod.ICompilationUnit) {
	compilationUnit.TypeDeclarations().AcceptDeclaration(v)
}

// --- DeclarationVisitor ---

func (v *UpdateIntegerConstantTypeVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.AnnotationDeclarators())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	superType := decl.SuperType()

	if superType != nil {
		superType.AcceptTypeVisitor(v)
	}

	v.SafeAcceptTypeParameter(decl.TypeParameters())
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.VisitTypeDeclaration(decl.(intmod.ITypeDeclaration))
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptListConstant(decl.Constants())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitEnumDeclarationConstant(decl intmod.IConstant) {
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptExpression(decl.Arguments())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFieldDeclarators(decl intmod.IFieldDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFormalParameter(decl intmod.IFormalParameter) {
	t := decl.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFormalParameters(decl intmod.IFormalParameters) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitInstanceInitializerDeclaration(decl intmod.IInstanceInitializerDeclaration) {
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptType(decl.Interfaces())
	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableDeclarator(decl intmod.ILocalVariableDeclarator) {
	v.SafeAcceptDeclaration(decl.VariableInitializer())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableDeclarators(decl intmod.ILocalVariableDeclarators) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitMethodDeclaration(decl intmod.IMethodDeclaration) {
	t := decl.ReturnedType()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptReference(decl.AnnotationReferences())
	v.SafeAcceptDeclaration(decl.FormalParameters())
	v.SafeAcceptType(decl.ExceptionTypes())
	v.SafeAcceptStatement(decl.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitMemberDeclarations(decl intmod.IMemberDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitModuleDeclaration(_ intmod.IModuleDeclaration) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTypeDeclarations(decl intmod.ITypeDeclarations) {
	list := make([]intmod.IDeclaration, 0, decl.Size())
	for _, element := range decl.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListDeclaration(list)
}

// --- IExpressionVisitor ---

func (v *UpdateIntegerConstantTypeVisitor) VisitBooleanExpression(_ intmod.IBooleanExpression) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitCommentExpression(_ intmod.ICommentExpression) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitExpressions(expr intmod.IExpressions) {
	list := make([]intmod.IExpression, 0, expr.Size())
	for _, element := range expr.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListExpression(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)

	v.SafeAcceptExpression(expr.Expression())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.SafeAcceptDeclaration(expr.FormalParameters())
	expr.Statements().AcceptStatement(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	t := expr.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptDeclaration(expr.ArrayInitializer())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitNoExpression(_ intmod.INoExpression) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	expr.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitStringConstantExpression(_ intmod.IStringConstantExpression) {
	// Empty
}

// --- IReferenceVisitor ---

func (v *UpdateIntegerConstantTypeVisitor) VisitAnnotationElementValue(ref intmod.IAnnotationElementValue) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitAnnotationReference(ref intmod.IAnnotationReference) {
	v.SafeAcceptReference(ref.ElementValue())
	v.SafeAcceptReference(ref.ElementValuePairs())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitAnnotationReferences(ref intmod.IAnnotationReferences) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitElementValueArrayInitializerElementValue(ref intmod.IElementValueArrayInitializerElementValue) {
	v.SafeAcceptReference(ref.ElementValueArrayInitializer())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitElementValues(ref intmod.IElementValues) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitElementValuePair(ref intmod.IElementValuePair) {
	ref.ElementValue().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitElementValuePairs(ref intmod.IElementValuePairs) {
	list := make([]intmod.IReference, 0, ref.Size())
	for _, element := range ref.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListReference(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitExpressionElementValue(ref intmod.IExpressionElementValue) {
	ref.Expression().Accept(v)
}

// --- IStatementVisitor ---

func (v *UpdateIntegerConstantTypeVisitor) VisitCommentStatement(_ intmod.ICommentStatement) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitExpressionStatement(stat intmod.IExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitForEachStatement(stat intmod.IForEachStatement) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLabelStatement(stat intmod.ILabelStatement) {
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLambdaExpressionStatement(stat intmod.ILambdaExpressionStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitLocalVariableDeclarationStatement(stat intmod.ILocalVariableDeclarationStatement) {
	v.VisitLocalVariableDeclaration(&stat.(*statement.LocalVariableDeclarationStatement).LocalVariableDeclaration)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitNoStatement(_ intmod.INoStatement) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitReturnStatement(_ intmod.IReturnStatement) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitStatements(stat intmod.IStatements) {
	list := make([]intmod.IStatement, 0, stat.Size())
	for _, element := range stat.ToSlice() {
		list = append(list, element)
	}
	v.AcceptListStatement(list)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSwitchStatement(stat intmod.ISwitchStatement) {
	stat.Condition().Accept(v)
	v.AcceptListStatement(stat.List())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSwitchStatementDefaultLabel(_ intmod.IDefaultLabel) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSwitchStatementExpressionLabel(stat intmod.IExpressionLabel) {
	stat.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSwitchStatementLabelBlock(stat intmod.ILabelBlock) {
	stat.Label().AcceptStatement(v)
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSwitchStatementMultiLabelsBlock(stat intmod.IMultiLabelsBlock) {
	v.SafeAcceptListStatement(stat.ToSlice())
	stat.Statements().AcceptStatement(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitSynchronizedStatement(stat intmod.ISynchronizedStatement) {
	stat.Monitor().Accept(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitThrowStatement(stat intmod.IThrowStatement) {
	stat.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTryStatement(stat intmod.ITryStatement) {
	v.SafeAcceptListStatement(stat.ResourceList())
	stat.TryStatements().AcceptStatement(v)
	v.SafeAcceptListStatement(stat.CatchClauseList())
	v.SafeAcceptStatement(stat.FinallyStatements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTryStatementResource(stat intmod.IResource) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	stat.Expression().Accept(v)
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTryStatementCatchClause(stat intmod.ICatchClause) {
	t := stat.Type()
	t.AcceptTypeVisitor(v)
	v.SafeAcceptStatement(stat.Statements())
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTypeDeclarationStatement(stat intmod.ITypeDeclarationStatement) {
	stat.TypeDeclaration().AcceptDeclaration(v)
}

// --- ITypeVisitor ---

// --- ITypeParameterVisitor --- //

func (v *UpdateIntegerConstantTypeVisitor) VisitTypeParameter(_ intmod.ITypeParameter) {
	// Empty
}

func (v *UpdateIntegerConstantTypeVisitor) VisitTypeParameters(parameters intmod.ITypeParameters) {
	for _, param := range parameters.ToSlice() {
		param.AcceptTypeParameterVisitor(v)
	}
}

// --- ITypeArgumentVisitor ---

func (v *UpdateIntegerConstantTypeVisitor) VisitTypeDeclaration(decl intmod.ITypeDeclaration) {
	v.SafeAcceptReference(decl.AnnotationReferences())
}

func (v *UpdateIntegerConstantTypeVisitor) AcceptListDeclaration(list []intmod.IDeclaration) {
	for _, value := range list {
		value.AcceptDeclaration(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) AcceptListExpression(list []intmod.IExpression) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) AcceptListReference(list []intmod.IReference) {
	for _, value := range list {
		value.Accept(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) AcceptListStatement(list []intmod.IStatement) {
	for _, value := range list {
		value.AcceptStatement(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptDeclaration(decl intmod.IDeclaration) {
	if decl != nil {
		decl.AcceptDeclaration(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptExpression(expr intmod.IExpression) {
	if expr != nil {
		expr.Accept(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptReference(ref intmod.IReference) {
	if ref != nil {
		ref.Accept(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptStatement(list intmod.IStatement) {
	if list != nil {
		list.AcceptStatement(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptType(list intmod.IType) {
	if list != nil {
		list.AcceptTypeVisitor(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptTypeParameter(list intmod.ITypeParameter) {
	if list != nil {
		list.AcceptTypeParameterVisitor(v)
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptListDeclaration(list []intmod.IDeclaration) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptListConstant(list []intmod.IConstant) {
	if list != nil {
		for _, value := range list {
			value.AcceptDeclaration(v)
		}
	}
}

func (v *UpdateIntegerConstantTypeVisitor) SafeAcceptListStatement(list []intmod.IStatement) {
	if list != nil {
		for _, value := range list {
			value.AcceptStatement(v)
		}
	}
}
