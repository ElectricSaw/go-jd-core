package token

import intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"

// -------- EndBlockToken --------

var EndBlock = NewEndBlockToken("}")
var EndArrayBlock = NewEndBlockToken("]")
var EndArrayInitializerBlock = NewEndBlockToken("}")
var EndParametersBlock = NewEndBlockToken(")")
var EndResourcesBlock = NewEndBlockToken(")")
var EndDeclarationOrStatementBlock = NewEndBlockToken("")

// -------- EndMarkerToken --------

const (
	CommentToken          = 1
	JavaDocToken          = 2
	ErrorToken            = 3
	ImportStatementsToken = 4
)

var EndComment = NewEndMarkerToken(CommentToken)
var EndJavaDoc = NewEndMarkerToken(JavaDocToken)
var EndError = NewEndMarkerToken(ErrorToken)
var EndImportStatements = NewEndMarkerToken(ImportStatementsToken)

// -------- LineNumberToken --------

var UnknownLineNumber = NewLineNumberToken(intmod.UnknownLineNumberToken)

// -------- NewLineToken --------

var NewLine1 = NewNewLineToken(1)
var NewLine2 = NewNewLineToken(2)

// -------- StartBlockToken --------

var StartBlock = NewStartBlockToken("{")
var StartArrayBlock = NewStartBlockToken("[")
var StartArrayInitializerBlock = NewStartBlockToken("{")
var StartParametersBlock = NewStartBlockToken("(")
var StartResourcesBlock = NewStartBlockToken("(")
var StartDeclarationOrStatementBlock = NewStartBlockToken("")

// -------- StartMarkerToken --------

var StartComment = NewStartMarkerToken(CommentToken)
var StartJavaDocToken = NewStartMarkerToken(JavaDocToken)
var StartErrorToken = NewStartMarkerToken(ErrorToken)
var StartImportStatementsToken = NewStartMarkerToken(ImportStatementsToken)

// -------- TextToken --------

var At = NewTextToken("@")
var Comma = NewTextToken(",")
var Colon = NewTextToken(":")
var ColonColon = NewTextToken("::")
var CommaSpace = NewTextToken(", ")
var Diamond = NewTextToken("<>")
var Dot = NewTextToken(".")
var Dimension1 = NewTextToken("[]")
var Dimension2 = NewTextToken("[][]")
var InfiniteFor = NewTextToken("(;;)")
var LeftRightCurlyBrackets = NewTextToken("{}")
var LeftRoundBracket = NewTextToken("(")
var RightRoundBracket = NewTextToken(")")
var LeftRightRoundBrackets = NewTextToken("()")
var LeftAngleBracket = NewTextToken("<")
var RightAngleBracket = NewTextToken(">")
var QuestionMark = NewTextToken("?")
var QuestionMarkSpace = NewTextToken("? ")
var Space = NewTextToken(" ")
var SpaceAndSpace = NewTextToken(" & ")
var SpaceArrowSpace = NewTextToken(" -> ")
var SpaceColonSpace = NewTextToken(" : ")
var SpaceEqualSpace = NewTextToken(" = ")
var SpaceQuestionSpace = NewTextToken(" ? ")
var SpaceLeftRoundBracket = NewTextToken(" (")
var Semicolon = NewTextToken(";")
var SemicolonSpace = NewTextToken("; ")
var VarArgs = NewTextToken("... ")
var VerticalLine = NewTextToken("|")
var Exclamation = NewTextToken("!")

// -------- CompilationUnitVisitor --------

var Abstract = NewKeywordToken("abstract")
var Annotation = NewKeywordToken("@interface")
var Class2 = NewKeywordToken("class")
var Default2 = NewKeywordToken("default")
var Enum = NewKeywordToken("enum")
var Implements = NewKeywordToken("implements")
var Interface = NewKeywordToken("interface")
var Native = NewKeywordToken("native")
var Package = NewKeywordToken("package")
var Private = NewKeywordToken("private")
var Protected = NewKeywordToken("protected")
var Public = NewKeywordToken("public")
var Static = NewKeywordToken("static")
var Throws = NewKeywordToken("Throws")

var CommentBridge = NewTextToken("/* bridge */")
var CommentSynthetic = NewTextToken("/* synthetic */")

// -------- ExpressionVisitor --------

var Class = NewKeywordToken("class")
var False = NewKeywordToken("false")
var InstanceOf = NewKeywordToken("instanceof")
var Length = NewKeywordToken("length")
var New = NewKeywordToken("new")
var Null = NewKeywordToken("null")
var This = NewKeywordToken("this")
var True = NewKeywordToken("true")

// -------- StatementVisitor --------

var Assert = NewKeywordToken("assert")
var Break = NewKeywordToken("break")
var Case = NewKeywordToken("case")
var Catch = NewKeywordToken("catch")
var Continue = NewKeywordToken("continue")
var Default = NewKeywordToken("default")
var Do = NewKeywordToken("do")
var Else = NewKeywordToken("else")
var Final = NewKeywordToken("final")
var Finally = NewKeywordToken("finally")
var For = NewKeywordToken("for")
var If = NewKeywordToken("if")
var Return = NewKeywordToken("return")
var Strict = NewKeywordToken("strictfp")
var Synchronized = NewKeywordToken("synchronized")
var Switch = NewKeywordToken("switch")
var Throw = NewKeywordToken("throw")
var Transient = NewKeywordToken("transient")
var Try = NewKeywordToken("try")
var Volatile = NewKeywordToken("volatile")
var While = NewKeywordToken("while")

// -------- ITypeVisitor --------

var Boolean = NewKeywordToken("boolean")
var Byte = NewKeywordToken("byte")
var Char = NewKeywordToken("char")
var Double = NewKeywordToken("double")
var Exports = NewKeywordToken("exports")
var Extends = NewKeywordToken("extends")
var Float = NewKeywordToken("float")
var Int = NewKeywordToken("int")
var Long = NewKeywordToken("long")
var Module = NewKeywordToken("module")
var Open = NewKeywordToken("open")
var Opens = NewKeywordToken("opens")
var Provides = NewKeywordToken("provides")
var Requires = NewKeywordToken("requires")
var Short = NewKeywordToken("short")
var Super = NewKeywordToken("super")
var To = NewKeywordToken("to")
var Transitive = NewKeywordToken("transitive")
var Uses = NewKeywordToken("uses")
var Void = NewKeywordToken("void")
var With = NewKeywordToken("with")

// -------- EndBlockToken --------

// -------- EndBlockToken --------

// -------- EndBlockToken --------
