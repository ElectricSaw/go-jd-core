package model

// -------- EndBlockToken --------

var EndBlockTkn = NewEndBlockToken("}")
var EndArrayBlockTkn = NewEndBlockToken("]")
var EndArrayInitializerBlockTkn = NewEndBlockToken("}")
var EndParametersBlockTkn = NewEndBlockToken(")")
var EndResourcesBlockTkn = NewEndBlockToken(")")
var EndDeclarationOrStatementBlockTkn = NewEndBlockToken("")

// -------- EndMarkerToken --------

const (
	CommentToken          = 1
	JavaDocToken          = 2
	ErrorToken            = 3
	ImportStatementsToken = 4
)

var EndCommentTkn = NewEndMarkerToken(CommentToken)
var EndJavaDocTkn = NewEndMarkerToken(JavaDocToken)
var EndErrorTkn = NewEndMarkerToken(ErrorToken)
var EndImportStatementsTkn = NewEndMarkerToken(ImportStatementsToken)

// -------- LineNumberToken --------

var UnknownLineNumberTkn = NewLineNumberToken(UnknownLineNumber)

// -------- NewLineToken --------

var NewLineTkn1 = NewNewLineToken(1)
var NewLineTkn2 = NewNewLineToken(2)

// -------- StartBlockToken --------

var StartBlockTkn = NewStartBlockToken("{")
var StartArrayBlockTkn = NewStartBlockToken("[")
var StartArrayInitializerBlockTkn = NewStartBlockToken("{")
var StartParametersBlockTkn = NewStartBlockToken("(")
var StartResourcesBlockTkn = NewStartBlockToken("(")
var StartDeclarationOrStatementBlockTkn = NewStartBlockToken("")

// -------- StartMarkerToken --------

var StartCommentTkn = NewStartMarkerToken(CommentToken)
var StartJavaDocTkn = NewStartMarkerToken(JavaDocToken)
var StartErrorTkn = NewStartMarkerToken(ErrorToken)
var StartImportStatementsTkn = NewStartMarkerToken(ImportStatementsToken)

// -------- TextToken --------

var AtTkn = NewTextToken("@")
var CommaTkn = NewTextToken(",")
var ColonTkn = NewTextToken(":")
var ColonColonTkn = NewTextToken("::")
var CommaSpaceTkn = NewTextToken(", ")
var DiamondTkn = NewTextToken("<>")
var DotTkn = NewTextToken(".")
var DimensionTkn1 = NewTextToken("[]")
var DimensionTkn2 = NewTextToken("[][]")
var InfiniteForTkn = NewTextToken("(;;)")
var LeftRightCurlyBracketsTkn = NewTextToken("{}")
var LeftRoundBracketTkn = NewTextToken("(")
var RightRoundBracketTkn = NewTextToken(")")
var LeftRightRoundBracketsTkn = NewTextToken("()")
var LeftAngleBracketTkn = NewTextToken("<")
var RightAngleBracketTkn = NewTextToken(">")
var QuestionMarkTkn = NewTextToken("?")
var QuestionMarkSpaceTkn = NewTextToken("? ")
var SpaceTkn = NewTextToken(" ")
var SpaceAndSpaceTkn = NewTextToken(" & ")
var SpaceArrowSpaceTkn = NewTextToken(" -> ")
var SpaceColonSpaceTkn = NewTextToken(" : ")
var SpaceEqualSpaceTkn = NewTextToken(" = ")
var SpaceQuestionSpaceTkn = NewTextToken(" ? ")
var SpaceLeftRoundBracketTkn = NewTextToken(" (")
var SemicolonTkn = NewTextToken(";")
var SemicolonSpaceTkn = NewTextToken("; ")
var VarArgsTkn = NewTextToken("... ")
var VerticalLineTkn = NewTextToken("|")
var ExclamationTkn = NewTextToken("!")

// -------- CompilationUnitVisitor --------

var AbstractTkn = NewKeywordToken("abstract")
var AnnotationTkn = NewKeywordToken("@interface")
var Class2Tkn = NewKeywordToken("class")
var Default2Tkn = NewKeywordToken("default")
var EnumTkn = NewKeywordToken("enum")
var ImplementsTkn = NewKeywordToken("implements")
var InterfaceTkn = NewKeywordToken("interface")
var NativeTkn = NewKeywordToken("native")
var PackageTkn = NewKeywordToken("package")
var PrivateTkn = NewKeywordToken("private")
var ProtectedTkn = NewKeywordToken("protected")
var PublicTkn = NewKeywordToken("public")
var StaticTkn = NewKeywordToken("static")
var ThrowsTkn = NewKeywordToken("Throws")

var CommentBridgeTkn = NewTextToken("/* bridge */")
var CommentSyntheticTkn = NewTextToken("/* synthetic */")

// -------- ExpressionVisitor --------

var ClassTkn = NewKeywordToken("class")
var FalseTkn = NewKeywordToken("false")
var InstanceOfTkn = NewKeywordToken("instanceof")
var LengthTkn = NewKeywordToken("length")
var NewTkn = NewKeywordToken("new")
var NullTkn = NewKeywordToken("null")
var ThisTkn = NewKeywordToken("this")
var TrueTkn = NewKeywordToken("true")

// -------- StatementVisitor --------

var AssertTkn = NewKeywordToken("assert")
var BreakTkn = NewKeywordToken("break")
var CaseTkn = NewKeywordToken("case")
var CatchTkn = NewKeywordToken("catch")
var ContinueTkn = NewKeywordToken("continue")
var DefaultTkn = NewKeywordToken("default")
var DoTkn = NewKeywordToken("do")
var ElseTkn = NewKeywordToken("else")
var FinalTkn = NewKeywordToken("final")
var FinallyTkn = NewKeywordToken("finally")
var ForTkn = NewKeywordToken("for")
var IfTkn = NewKeywordToken("if")
var ReturnTkn = NewKeywordToken("return")
var StrictTkn = NewKeywordToken("strictfp")
var SynchronizedTkn = NewKeywordToken("synchronized")
var SwitchTkn = NewKeywordToken("switch")
var ThrowTkn = NewKeywordToken("throw")
var TransientTkn = NewKeywordToken("transient")
var TryTkn = NewKeywordToken("try")
var VolatileTkn = NewKeywordToken("volatile")
var WhileTkn = NewKeywordToken("while")

// -------- ITypeVisitor --------

var BooleanTkn = NewKeywordToken("boolean")
var ByteTkn = NewKeywordToken("byte")
var CharTkn = NewKeywordToken("char")
var DoubleTkn = NewKeywordToken("double")
var ExportsTkn = NewKeywordToken("exports")
var ExtendsTkn = NewKeywordToken("extends")
var FloatTkn = NewKeywordToken("float")
var IntTkn = NewKeywordToken("int")
var LongTkn = NewKeywordToken("long")
var ModuleTkn = NewKeywordToken("module")
var OpenTkn = NewKeywordToken("open")
var OpensTkn = NewKeywordToken("opens")
var ProvidesTkn = NewKeywordToken("provides")
var RequiresTkn = NewKeywordToken("requires")
var ShortTkn = NewKeywordToken("short")
var SuperTkn = NewKeywordToken("super")
var ToTkn = NewKeywordToken("to")
var TransitiveTkn = NewKeywordToken("transitive")
var UsesTkn = NewKeywordToken("uses")
var VoidTkn = NewKeywordToken("void")
var WithTkn = NewKeywordToken("with")

// -------- EndBlockToken --------

// -------- EndBlockToken --------

// -------- EndBlockToken --------
