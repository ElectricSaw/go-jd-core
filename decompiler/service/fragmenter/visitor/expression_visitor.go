package visitor

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javafragment"
	_type "github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/type"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/token"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/fragmenter/visitor/fragutil"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func NewExpressionVisitor(loader api.Loader, mainInternalTypeName string, majorVersion int,
	importsFragment intmod.IImportsFragment) *ExpressionVisitor {
	v := &ExpressionVisitor{
		TypeVisitor:             *NewTypeVisitor(loader, mainInternalTypeName, majorVersion, importsFragment).(*TypeVisitor),
		contextStack:            util.NewDefaultList[IContext](),
		fragments:               NewFragments(),
		inExpressionFlag:        false,
		currentMethodParamNames: util.NewSet[string](),
	}
	v.diamondOperatorSupported = majorVersion > 49
	v.hexaExpressionVisitor = NewHexaExpressionVisitor(v)
	return v
}

type ExpressionVisitor struct {
	TypeVisitor

	contextStack             util.IList[IContext]
	fragments                IFragments
	diamondOperatorSupported bool
	inExpressionFlag         bool
	currentMethodParamNames  util.ISet[string]
	currentTypeName          string
	hexaExpressionVisitor    IHexaExpressionVisitor
}

func (v *ExpressionVisitor) Fragments() util.IList[intmod.IFragment] {
	return v.fragments
}

func (v *ExpressionVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	v.visit(expr, expr.Expression())
	v.tokens.Add(token.StartArrayBlock)
	expr.Index().Accept(v)
	v.tokens.Add(token.EndArrayBlock)
}

func (v *ExpressionVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	switch expr.Operator() {
	case "&":
	case "|":
	case "^":
	case "&=":
	case "|=":
	case "^=":
		v.visitHexa(expr, expr.LeftExpression())
		v.tokens.Add(token.Space)
		v.tokens.Add(v.newTextToken(expr.Operator()))
		v.tokens.Add(token.Space)
		v.visitHexa(expr, expr.RightExpression())
		break
	default:
		v.visit(expr, expr.LeftExpression())
		v.tokens.Add(token.Space)
		v.tokens.Add(v.newTextToken(expr.Operator()))
		v.tokens.Add(token.Space)
		v.visit(expr, expr.RightExpression())
		break
	}
}

func (v *ExpressionVisitor) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	if expr.IsTrue() {
		v.tokens.Add(token.True)
	} else {
		v.tokens.Add(token.False)
	}
}

func (v *ExpressionVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	if expr.IsExplicit() {
		v.tokens.AddLineNumberTokenAt(expr.LineNumber())
		v.tokens.Add(token.LeftRoundBracket)

		typ := expr.Type()
		typ.AcceptTypeVisitor(v)
		v.tokens.Add(token.RightRoundBracket)
	}

	v.visit(expr, expr.Expression())
}
func (v *ExpressionVisitor) VisitCommentExpression(expr intmod.ICommentExpression) {
	v.tokens.Add(token.StartComment)
	v.tokens.Add(v.newTextToken(expr.Text()))
	v.tokens.Add(token.EndComment)
}

func (v *ExpressionVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.This)
	v.tokens.Add(token.StartParametersBlock)

	parameters := expr.Parameters()

	if parameters != nil {
		parameters.Accept(v)
	}

	v.tokens.Add(token.EndParametersBlock)
}

func (v *ExpressionVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	ot := expr.ObjectType()

	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(v.newTypeReferenceToken(ot, v.currentInternalTypeName))
	v.tokens.Add(token.ColonColon)
	//v.tokens.Add(new ReferenceToken(ReferenceToken.CONSTRUCTOR, ot.InternalName(), "new", expr.Descriptor(), currentInternalTypeName));
	v.tokens.Add(token.New)
}

func (v *ExpressionVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.NewNumericConstantToken(fmt.Sprintf("%fD", expr.DoubleValue())))
}

func (v *ExpressionVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	v.tokens.AddLineNumberToken(expr)

	typ := expr.ObjectType()

	v.tokens.Add(token.NewReferenceToken(intmod.FieldToken, typ.InternalName(),
		expr.Name(), typ.Descriptor(), v.currentInternalTypeName))
}

func (v *ExpressionVisitor) VisitExpressions(list intmod.IExpressions) {
	if list != nil {
		size := list.Size()

		if size > 0 {
			ief := v.inExpressionFlag
			iterator := list.Iterator()

			for ; size > 1; size-- {
				v.inExpressionFlag = true
				iterator.Next().Accept(v)

				if !v.tokens.IsEmpty() {
					v.tokens.Add(token.CommaSpace)
				}
			}

			v.inExpressionFlag = false
			iterator.Next().Accept(v)
			v.inExpressionFlag = ief
		}
	}
}

func (v *ExpressionVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	if expr.Expression() == nil {
		v.tokens.AddLineNumberToken(expr)
		v.tokens.Add(token.NewTextToken(expr.Name()))
	} else {
		v.tokens.AddLineNumberToken(expr.Expression())

		delta := v.tokens.Size()

		v.visit(expr, expr.Expression())
		delta -= v.tokens.Size()
		v.tokens.AddLineNumberToken(expr)

		if delta != 0 {
			v.tokens.Add(token.Dot)
		}

		v.tokens.Add(token.NewReferenceToken(intmod.FieldToken, expr.InternalTypeName(), expr.Name(), expr.Descriptor(), v.currentInternalTypeName))
	}
}

func (v *ExpressionVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.NewNumericConstantToken(fmt.Sprintf("%fF", expr.FloatValue())))
}

func (v *ExpressionVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	v.tokens.AddLineNumberToken(expr)

	pt := expr.Type().(intmod.IPrimitiveType)

	switch pt.JavaPrimitiveFlags() {
	case intmod.FlagChar:
		v.tokens.Add(token.NewCharacterConstantToken(fragutil.EscapeChar(expr.IntegerValue()), v.currentInternalTypeName))
		break
	case intmod.FlagBoolean:
		v.tokens.Add(token.NewBooleanConstantToken(expr.IntegerValue() != 0))
		break
	default:
		v.tokens.Add(token.NewNumericConstantToken(fmt.Sprintf("%d", expr.IntegerValue())))
		break
	}
}

func (v *ExpressionVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	expr.Expression().Accept(v)
	v.tokens.Add(token.Space)
	v.tokens.Add(token.InstanceOf)
	v.tokens.Add(token.Space)

	typ := expr.InstanceOfType()
	typ.AcceptTypeVisitor(v)
}

func (v *ExpressionVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	parameters := expr.FormalParameters()

	if parameters == nil {
		v.tokens.Add(token.LeftRightRoundBrackets)
	} else {
		size := parameters.Size()

		switch size {
		case 0:
			v.tokens.Add(token.LeftRightRoundBrackets)
			break
		case 1:
			parameters.First().AcceptDeclaration(v)
			break
		default:
			v.tokens.Add(token.LeftRoundBracket)
			iterator := parameters.Iterator()
			iterator.Next().AcceptDeclaration(v)
			for iterator.HasNext() {
				v.tokens.Add(token.CommaSpace)
				iterator.Next().AcceptDeclaration(v)
			}

			v.tokens.Add(token.RightRoundBracket)
			break
		}
	}

	v.visitLambdaBody(expr.Statements())
}

func (v *ExpressionVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	parameters := util.NewDefaultListWithSlice(expr.ParameterNames())

	if parameters == nil {
		v.tokens.Add(token.LeftRightRoundBrackets)
	} else {
		size := parameters.Size()

		switch size {
		case 0:
			v.tokens.Add(token.LeftRightRoundBrackets)
			break
		case 1:
			v.tokens.Add(v.newTextToken(parameters.Get(0)))
			break
		default:
			v.tokens.Add(token.LeftRoundBracket)
			v.tokens.Add(v.newTextToken(parameters.Get(0)))
			for i := 1; i < size; i++ {
				v.tokens.Add(token.CommaSpace)
				v.tokens.Add(v.newTextToken(parameters.Get(i)))
			}
			v.tokens.Add(token.RightRoundBracket)
			break
		}
	}

	v.visitLambdaBody(expr.Statements())
}

func (v *ExpressionVisitor) visitLambdaBody(statementList intmod.IStatement) {
	if statementList != nil {
		v.tokens.Add(token.SpaceArrowSpace)

		if statementList.IsLambdaExpressionStatement() {
			statementList.AcceptStatement(v)
		} else {
			v.fragments.AddTokensFragment(v.tokens)

			start := fragutil.AddStartStatementsInLambdaBlock(v.fragments)

			v.tokens = NewTokens(v)
			statementList.AcceptStatement(v)

			if v.inExpressionFlag {
				fragutil.AddEndStatementsInLambdaBlockInParameter(v.fragments, start)
			} else {
				fragutil.AddEndStatementsInLambdaBlock(v.fragments, start)
			}

			v.tokens = NewTokens(v)
		}
	}
}

func (v *ExpressionVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	v.visit(expr, expr.Expression())
	v.tokens.Add(token.Dot)
	v.tokens.Add(token.Length)
}

func (v *ExpressionVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(v.newTextToken(expr.Name()))
}

func (v *ExpressionVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.NewNumericConstantToken(fmt.Sprintf("%dL", expr.LongValue())))
}

func (v *ExpressionVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	exp := expr.Expression()
	nonWildcardTypeArguments := expr.NonWildcardTypeArguments()
	parameters := expr.Parameters()
	dot := false

	if exp.IsThisExpression() {
		// Nothing to do : do not print 'v.method(...)'
	} else if exp.IsObjectTypeReferenceExpression() {
		ot := exp.ObjectType()

		if ot.InternalName() != v.currentInternalTypeName {
			v.visit(expr, exp)
			v.tokens.AddLineNumberToken(expr)
			v.tokens.Add(token.Dot)
			dot = true
		}
	} else {
		if exp.IsFieldReferenceExpression() || exp.IsLocalVariableReferenceExpression() {
			v.tokens.AddLineNumberToken(expr)
			v.visit(expr, exp)
		} else {
			v.visit(expr, exp)
			v.tokens.AddLineNumberToken(expr)
		}

		v.tokens.Add(token.Dot)
		dot = true
	}

	v.tokens.AddLineNumberToken(expr)

	if (nonWildcardTypeArguments != nil) && dot {
		v.tokens.Add(token.LeftAngleBracket)
		nonWildcardTypeArguments.AcceptTypeArgumentVisitor(v)
		v.tokens.Add(token.RightAngleBracket)
	}

	v.tokens.Add(token.NewReferenceToken(intmod.MethodToken, expr.InternalTypeName(), expr.Name(), expr.Descriptor(), v.currentInternalTypeName))
	v.tokens.Add(token.StartParametersBlock)

	if parameters != nil {
		ief := v.inExpressionFlag
		v.inExpressionFlag = false
		parameters.Accept(v)
		v.inExpressionFlag = ief
	}

	v.tokens.Add(token.EndParametersBlock)
}

func (v *ExpressionVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	expr.Expression().Accept(v)
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.ColonColon)
	v.tokens.Add(token.NewReferenceToken(intmod.MethodToken, expr.InternalTypeName(), expr.Name(), expr.Descriptor(), v.currentInternalTypeName))
}

func (v *ExpressionVisitor) VisitNewArray(expr intmod.INewArray) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.New)
	v.tokens.Add(token.Space)

	typ := expr.Type()
	typ.AcceptTypeVisitor(v)

	dimensionExpressionList := expr.DimensionExpressionList()
	dimension := expr.Type().Dimension()

	if dimension > 0 {
		v.tokens.RemoveAt(v.tokens.Size() - 1)
	}

	if dimensionExpressionList != nil {
		if dimensionExpressionList.IsList() {
			iterator := dimensionExpressionList.Iterator()

			for iterator.HasNext() {
				v.tokens.Add(token.StartArrayBlock)
				iterator.Next().Accept(v)
				v.tokens.Add(token.EndArrayBlock)
				dimension--
			}
		} else {
			v.tokens.Add(token.StartArrayBlock)
			dimensionExpressionList.Accept(v)
			v.tokens.Add(token.EndArrayBlock)
			dimension--
		}
	}

	v.visitDimension(dimension)
}

func (v *ExpressionVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.New)
	v.tokens.Add(token.Space)

	typ := expr.Type()
	typ.AcceptTypeVisitor(v)
	v.tokens.Add(token.Space)
	expr.ArrayInitializer().AcceptDeclaration(v)
}

func (v *ExpressionVisitor) VisitNewExpression(expr intmod.INewExpression) {
	bodyDeclaration := expr.BodyDeclaration()

	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.New)
	v.tokens.Add(token.Space)

	objectType := expr.ObjectType()

	if (objectType.TypeArguments() != nil) && (bodyDeclaration == nil) && v.diamondOperatorSupported {
		objectType = objectType.CreateTypeWithArgs(_type.Diamond)
	}

	typ := objectType

	typ.AcceptTypeVisitor(v)
	v.tokens.Add(token.StartParametersBlock)

	parameters := expr.Parameters()
	if parameters != nil {
		parameters.Accept(v)
	}

	v.tokens.Add(token.EndParametersBlock)

	if bodyDeclaration != nil {
		v.fragments.AddTokensFragment(v.tokens)

		start := fragutil.AddStartTypeBody(v.fragments)
		ot := expr.ObjectType()

		v.storeContext()
		v.currentInternalTypeName = bodyDeclaration.InternalTypeName()
		v.currentTypeName = ot.Name()
		bodyDeclaration.AcceptDeclaration(v)

		if !v.tokens.IsEmpty() {
			v.tokens = NewTokens(v)
		}

		v.restoreContext()

		if v.inExpressionFlag {
			fragutil.AddEndSubTypeBodyInParameter(v.fragments, start)
		} else {
			fragutil.AddEndSubTypeBody(v.fragments, start)
		}

		v.tokens = NewTokens(v)
	}
}

func (v *ExpressionVisitor) VisitNullExpression(expr intmod.INullExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(nil)
}

func (v *ExpressionVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	if expr.IsExplicit() {
		v.tokens.AddLineNumberToken(expr)

		typ := expr.Type()
		typ.AcceptTypeVisitor(v)
	}
}

func (v *ExpressionVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	v.tokens.Add(token.StartParametersBlock)
	expr.Expression().Accept(v)
	v.tokens.Add(token.EndParametersBlock)
}

func (v *ExpressionVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	v.visit(expr, expr.Expression())
	v.tokens.Add(v.newTextToken(expr.Operator()))
}

func (v *ExpressionVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	v.tokens.AddLineNumberToken(expr.Expression())
	v.tokens.Add(v.newTextToken(expr.Operator()))
	v.visit(expr, expr.Expression())
}

func (v *ExpressionVisitor) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.NewStringConstantToken(fragutil.EscapeString(expr.StringValue()), v.currentInternalTypeName))
}

func (v *ExpressionVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.Super)
	v.tokens.Add(token.StartParametersBlock)

	parameters := expr.Parameters()

	if parameters != nil {
		parameters.Accept(v)
	}

	v.tokens.Add(token.EndParametersBlock)
}

func (v *ExpressionVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	v.tokens.AddLineNumberToken(expr)
	v.tokens.Add(token.Super)
}

func (v *ExpressionVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	v.tokens.AddLineNumberToken(expr.Condition())

	if expr.TrueExpression().IsBooleanExpression() && expr.FalseExpression().IsBooleanExpression() {
		be1 := expr.TrueExpression().(intmod.IBooleanExpression)
		be2 := expr.FalseExpression().(intmod.IBooleanExpression)

		if be1.IsTrue() && be2.IsFalse() {
			v.printTernaryOperatorExpression(expr.Condition())
			return
		}

		if be1.IsFalse() && be2.IsTrue() {
			v.tokens.Add(token.Exclamation)
			v.printTernaryOperatorExpression(expr.Condition())
			return
		}
	}

	v.printTernaryOperatorExpression(expr.Condition())
	v.tokens.Add(token.SpaceQuestionSpace)
	v.printTernaryOperatorExpression(expr.TrueExpression())
	v.tokens.Add(token.SpaceColonSpace)
	v.printTernaryOperatorExpression(expr.FalseExpression())
}

func (v *ExpressionVisitor) printTernaryOperatorExpression(expr intmod.IExpression) {
	if expr.Priority() > 3 {
		v.tokens.Add(token.LeftRoundBracket)
		expr.Accept(v)
		v.tokens.Add(token.RightRoundBracket)
	} else {
		expr.Accept(v)
	}
}

func (v *ExpressionVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	if expr.IsExplicit() {
		v.tokens.AddLineNumberToken(expr)
		v.tokens.Add(token.This)
	}
}

func (v *ExpressionVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	v.tokens.AddLineNumberToken(expr)

	typ := expr.TypeDotClass()

	typ.AcceptTypeVisitor(v)
	v.tokens.Add(token.Dot)
	v.tokens.Add(token.Class)
}

func (v *ExpressionVisitor) storeContext() {
	v.contextStack.Add(NewContext(v.currentInternalTypeName, v.currentTypeName, v.currentMethodParamNames))
}

func (v *ExpressionVisitor) restoreContext() {
	currentContext := v.contextStack.RemoveLast()
	v.currentInternalTypeName = currentContext.CurrentInternalTypeName()
	v.currentTypeName = currentContext.CurrentTypeName()
	v.currentMethodParamNames = currentContext.CurrentMethodParamNames()
}

func (v *ExpressionVisitor) visit(parent, child intmod.IExpression) {
	if (parent.Priority() < child.Priority()) || ((parent.Priority() == 14) && (child.Priority() == 13)) {
		v.tokens.Add(token.LeftRoundBracket)
		child.Accept(v)
		v.tokens.Add(token.RightRoundBracket)
	} else {
		child.Accept(v)
	}
}

func (v *ExpressionVisitor) visitHexa(parent, child intmod.IExpression) {
	if (parent.Priority() < child.Priority()) || ((parent.Priority() == 14) && (child.Priority() == 13)) {
		v.tokens.Add(token.LeftRoundBracket)
		child.Accept(v.hexaExpressionVisitor)
		v.tokens.Add(token.RightRoundBracket)
	} else {
		child.Accept(v.hexaExpressionVisitor)
	}
}

func NewContext(currentInternalTypeName string, currentTypeName string,
	currentMethodParamNames util.ISet[string]) IContext {
	return &Context{
		currentInternalTypeName: currentInternalTypeName,
		currentTypeName:         currentTypeName,
		currentMethodParamNames: currentMethodParamNames,
	}
}

type IContext interface {
	CurrentInternalTypeName() string
	SetCurrentInternalTypeName(currentInternalTypeName string)
	CurrentTypeName() string
	SetCurrentTypeName(currentTypeName string)
	CurrentMethodParamNames() util.ISet[string]
	SetCurrentMethodParamNames(currentMethodParamNames util.ISet[string])
}

type Context struct {
	currentInternalTypeName string
	currentTypeName         string
	currentMethodParamNames util.ISet[string]
}

func (c *Context) CurrentInternalTypeName() string {
	return c.currentInternalTypeName
}

func (c *Context) SetCurrentInternalTypeName(currentInternalTypeName string) {
	c.currentInternalTypeName = currentInternalTypeName
}

func (c *Context) CurrentTypeName() string {
	return c.currentTypeName
}

func (c *Context) SetCurrentTypeName(currentTypeName string) {
	c.currentTypeName = currentTypeName
}

func (c *Context) CurrentMethodParamNames() util.ISet[string] {
	return c.currentMethodParamNames
}

func (c *Context) SetCurrentMethodParamNames(currentMethodParamNames util.ISet[string]) {
	c.currentMethodParamNames = currentMethodParamNames
}

func NewFragments() IFragments {
	return &Fragments{
		DefaultList: *util.NewDefaultList[intmod.IFragment]().(*util.DefaultList[intmod.IFragment]),
	}
}

type IFragments interface {
	util.IList[intmod.IFragment]

	AddTokensFragment(tokens ITokens)
}

type Fragments struct {
	util.DefaultList[intmod.IFragment]
}

func (f *Fragments) AddTokensFragment(tokens ITokens) {
	if !tokens.IsEmpty() {
		if tokens.CurrentLineNumber() == UnknownLineNumber {
			f.Add(javafragment.NewTokensFragmentWithSlice(tokens.ToSlice()))
		} else {
			f.Add(javafragment.NewLineNumberTokensFragment(tokens.ToSlice()))
		}
	}
}

func NewHexaExpressionVisitor(parent *ExpressionVisitor) IHexaExpressionVisitor {
	return &HexaExpressionVisitor{
		parent: parent,
	}
}

type IHexaExpressionVisitor interface {
	intmod.IExpressionVisitor
}

type HexaExpressionVisitor struct {
	parent *ExpressionVisitor
}

func (v *HexaExpressionVisitor) VisitIntegerConstantExpression(expr intmod.IIntegerConstantExpression) {
	v.parent.tokens.AddLineNumberToken(expr)
	pt := expr.Type().(intmod.IPrimitiveType)

	switch pt.JavaPrimitiveFlags() {
	case intmod.FlagBoolean:
		v.parent.tokens.Add(token.NewBooleanConstantToken(expr.IntegerValue() == 1))
		break
	default:
		v.parent.tokens.Add(token.NewNumericConstantToken(fmt.Sprintf("0x%02X", expr.IntegerValue())))
		break
	}
}

func (v *HexaExpressionVisitor) VisitLongConstantExpression(expr intmod.ILongConstantExpression) {
	v.parent.tokens.AddLineNumberToken(expr)
	v.parent.tokens.Add(token.NewNumericConstantToken(
		fmt.Sprintf("0x%02XL", expr.LongValue())))
}

func (v *HexaExpressionVisitor) VisitArrayExpression(expr intmod.IArrayExpression) {
	v.parent.VisitArrayExpression(expr)
}
func (v *HexaExpressionVisitor) VisitBinaryOperatorExpression(expr intmod.IBinaryOperatorExpression) {
	v.parent.VisitBinaryOperatorExpression(expr)
}
func (v *HexaExpressionVisitor) VisitBooleanExpression(expr intmod.IBooleanExpression) {
	v.parent.VisitBooleanExpression(expr)
}
func (v *HexaExpressionVisitor) VisitCastExpression(expr intmod.ICastExpression) {
	v.parent.VisitCastExpression(expr)
}
func (v *HexaExpressionVisitor) VisitCommentExpression(expr intmod.ICommentExpression) {
	v.parent.VisitCommentExpression(expr)
}
func (v *HexaExpressionVisitor) VisitConstructorInvocationExpression(expr intmod.IConstructorInvocationExpression) {
	v.parent.VisitConstructorInvocationExpression(expr)
}
func (v *HexaExpressionVisitor) VisitConstructorReferenceExpression(expr intmod.IConstructorReferenceExpression) {
	v.parent.VisitConstructorReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitDoubleConstantExpression(expr intmod.IDoubleConstantExpression) {
	v.parent.VisitDoubleConstantExpression(expr)
}
func (v *HexaExpressionVisitor) VisitEnumConstantReferenceExpression(expr intmod.IEnumConstantReferenceExpression) {
	v.parent.VisitEnumConstantReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitExpressions(expr intmod.IExpressions) {
	v.parent.VisitExpressions(expr)
}
func (v *HexaExpressionVisitor) VisitFieldReferenceExpression(expr intmod.IFieldReferenceExpression) {
	v.parent.VisitFieldReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitFloatConstantExpression(expr intmod.IFloatConstantExpression) {
	v.parent.VisitFloatConstantExpression(expr)
}
func (v *HexaExpressionVisitor) VisitInstanceOfExpression(expr intmod.IInstanceOfExpression) {
	v.parent.VisitInstanceOfExpression(expr)
}
func (v *HexaExpressionVisitor) VisitLambdaFormalParametersExpression(expr intmod.ILambdaFormalParametersExpression) {
	v.parent.VisitLambdaFormalParametersExpression(expr)
}
func (v *HexaExpressionVisitor) VisitLambdaIdentifiersExpression(expr intmod.ILambdaIdentifiersExpression) {
	v.parent.VisitLambdaIdentifiersExpression(expr)
}
func (v *HexaExpressionVisitor) VisitLengthExpression(expr intmod.ILengthExpression) {
	v.parent.VisitLengthExpression(expr)
}
func (v *HexaExpressionVisitor) VisitLocalVariableReferenceExpression(expr intmod.ILocalVariableReferenceExpression) {
	v.parent.VisitLocalVariableReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitMethodInvocationExpression(expr intmod.IMethodInvocationExpression) {
	v.parent.VisitMethodInvocationExpression(expr)
}
func (v *HexaExpressionVisitor) VisitMethodReferenceExpression(expr intmod.IMethodReferenceExpression) {
	v.parent.VisitMethodReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitNewArray(expr intmod.INewArray) {
	v.parent.VisitNewArray(expr)
}
func (v *HexaExpressionVisitor) VisitNewExpression(expr intmod.INewExpression) {
	v.parent.VisitNewExpression(expr)
}
func (v *HexaExpressionVisitor) VisitNewInitializedArray(expr intmod.INewInitializedArray) {
	v.parent.VisitNewInitializedArray(expr)
}

func (v *HexaExpressionVisitor) VisitNoExpression(_ intmod.INoExpression) {
}

func (v *HexaExpressionVisitor) VisitNullExpression(expr intmod.INullExpression) {
	v.parent.VisitNullExpression(expr)
}
func (v *HexaExpressionVisitor) VisitObjectTypeReferenceExpression(expr intmod.IObjectTypeReferenceExpression) {
	v.parent.VisitObjectTypeReferenceExpression(expr)
}
func (v *HexaExpressionVisitor) VisitParenthesesExpression(expr intmod.IParenthesesExpression) {
	v.parent.VisitParenthesesExpression(expr)
}
func (v *HexaExpressionVisitor) VisitPostOperatorExpression(expr intmod.IPostOperatorExpression) {
	v.parent.VisitPostOperatorExpression(expr)
}
func (v *HexaExpressionVisitor) VisitPreOperatorExpression(expr intmod.IPreOperatorExpression) {
	v.parent.VisitPreOperatorExpression(expr)
}
func (v *HexaExpressionVisitor) VisitStringConstantExpression(expr intmod.IStringConstantExpression) {
	v.parent.VisitStringConstantExpression(expr)
}
func (v *HexaExpressionVisitor) VisitSuperConstructorInvocationExpression(expr intmod.ISuperConstructorInvocationExpression) {
	v.parent.VisitSuperConstructorInvocationExpression(expr)
}
func (v *HexaExpressionVisitor) VisitSuperExpression(expr intmod.ISuperExpression) {
	v.parent.VisitSuperExpression(expr)
}
func (v *HexaExpressionVisitor) VisitTernaryOperatorExpression(expr intmod.ITernaryOperatorExpression) {
	v.parent.VisitTernaryOperatorExpression(expr)
}
func (v *HexaExpressionVisitor) VisitThisExpression(expr intmod.IThisExpression) {
	v.parent.VisitThisExpression(expr)
}
func (v *HexaExpressionVisitor) VisitTypeReferenceDotClassExpression(expr intmod.ITypeReferenceDotClassExpression) {
	v.parent.VisitTypeReferenceDotClassExpression(expr)
}
