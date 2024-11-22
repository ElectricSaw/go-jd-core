package token

import "fmt"

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
var LeftRightRoundBracket = NewTextToken("()")
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

func NewTextToken(text string) *TextToken {
	return &TextToken{text}
}

type TextToken struct {
	text string
}

func (t *TextToken) Text() string {
	return t.text
}

func (t *TextToken) Accept(visitor TokenVisitor) {
	visitor.VisitTextToken(t)
}

func (t *TextToken) String() string {
	return fmt.Sprintf("TextToken { '%s' }", t.text)
}
