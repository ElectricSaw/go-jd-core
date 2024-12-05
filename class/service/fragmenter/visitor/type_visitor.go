package visitor

import (
	"bitbucket.org/coontec/go-jd-core/class/api"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/token"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"strings"
)

var Boolean = token.NewKeywordToken("boolean")
var Byte = token.NewKeywordToken("byte")
var Char = token.NewKeywordToken("char")
var Double = token.NewKeywordToken("double")
var Exports = token.NewKeywordToken("exports")
var Extends = token.NewKeywordToken("extends")
var Float = token.NewKeywordToken("float")
var Int = token.NewKeywordToken("int")
var Long = token.NewKeywordToken("long")
var Module = token.NewKeywordToken("module")
var Open = token.NewKeywordToken("open")
var Opens = token.NewKeywordToken("opens")
var Provides = token.NewKeywordToken("provides")
var Requires = token.NewKeywordToken("requires")
var Short = token.NewKeywordToken("short")
var Super = token.NewKeywordToken("super")
var To = token.NewKeywordToken("to")
var Transitive = token.NewKeywordToken("transitive")
var Uses = token.NewKeywordToken("uses")
var Void = token.NewKeywordToken("void")
var With = token.NewKeywordToken("with")

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
	v.tokens.Add(Extends)
	v.tokens.Add(token.Space)

	typ := argument.Type()
	typ.AcceptTypeVisitor(v)
}

func (v *TypeVisitor) VisitPrimitiveType(typ intmod.IPrimitiveType) {
	switch typ.JavaPrimitiveFlags() {
	case intmod.FlagBoolean:
		v.tokens.Add(Boolean)
		break
	case intmod.FlagChar:
		v.tokens.Add(Char)
		break
	case intmod.FlagFloat:
		v.tokens.Add(Float)
		break
	case intmod.FlagDouble:
		v.tokens.Add(Double)
		break
	case intmod.FlagByte:
		v.tokens.Add(Byte)
		break
	case intmod.FlagShort:
		v.tokens.Add(Short)
		break
	case intmod.FlagInt:
		v.tokens.Add(Int)
		break
	case intmod.FlagLong:
		v.tokens.Add(Long)
		break
	case intmod.FlagVoid:
		v.tokens.Add(Void)
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
	v.tokens.Add(Super)
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
	v.tokens.Add(Extends)
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
