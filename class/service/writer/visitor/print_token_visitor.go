package visitor

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/class/api"
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/class/model/token"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

const UnknownLineNumber = api.UnknownLineNumber

func NewPrintTokenVisitor() *PrintTokenVisitor {
	return &PrintTokenVisitor{}
}

type PrintTokenVisitor struct {
	searchLineNumberVisitor *SearchLineNumberVisitor

	printer      api.Printer
	tokens       util.IList[intmod.IToken]
	index        int
	newLineCount int
}

func (v *PrintTokenVisitor) Start(printer api.Printer, tokens util.IList[intmod.IToken]) {
	v.printer = printer
	v.tokens = tokens
	v.index = 0
	v.newLineCount = 0
	v.printer.StartLine(v.searchLineNumber())
}

func (v *PrintTokenVisitor) End() {
	v.printer.EndLine()
}

func (v *PrintTokenVisitor) VisitBooleanConstantToken(tkn intmod.IBooleanConstantToken) {
	v.prepareNewLine()
	if tkn.Value() {
		v.printer.PrintKeyword("true")
	} else {
		v.printer.PrintKeyword("false")
	}
	v.index++
}

func (v *PrintTokenVisitor) VisitCharacterConstantToken(tkn intmod.ICharacterConstantToken) {
	v.prepareNewLine()
	v.printer.PrintStringConstant(fmt.Sprintf("'%s'", tkn.Character()), tkn.OwnerInternalName())
	v.index++
}

func (v *PrintTokenVisitor) VisitDeclarationToken(tkn intmod.IDeclarationToken) {
	v.prepareNewLine()
	v.printer.PrintDeclaration(tkn.Type(), tkn.InternalTypeName(), tkn.Name(), tkn.Descriptor())
	v.index++
}

func (v *PrintTokenVisitor) VisitStartBlockToken(tkn intmod.IStartBlockToken) {
	v.prepareNewLine()
	v.printer.PrintText(tkn.Text())
	v.printer.Indent()
	if tkn == token.StartResourcesBlock {
		v.printer.Indent()
	}
	v.index++
}

func (v *PrintTokenVisitor) VisitEndBlockToken(tkn intmod.IEndBlockToken) {
	v.printer.Unindent()
	if tkn == token.EndResourcesBlock {
		v.printer.Unindent()
	}
	v.prepareNewLine()
	v.printer.PrintText(tkn.Text())
	v.index++
}

func (v *PrintTokenVisitor) VisitStartMarkerToken(tkn intmod.IStartMarkerToken) {
	v.prepareNewLine()
	v.printer.StartMarker(tkn.Type())
	v.index++
}

func (v *PrintTokenVisitor) VisitEndMarkerToken(tkn intmod.IEndMarkerToken) {
	v.prepareNewLine()
	v.printer.EndMarker(tkn.Type())
	v.index++
}

func (v *PrintTokenVisitor) VisitNewLineToken(tkn intmod.INewLineToken) {
	v.newLineCount += tkn.Count()
	v.index++
}

func (v *PrintTokenVisitor) VisitKeywordToken(tkn intmod.IKeywordToken) {
	v.prepareNewLine()
	v.printer.PrintKeyword(tkn.Keyword())
	v.index++
}

func (v *PrintTokenVisitor) VisitLineNumberToken(_ intmod.ILineNumberToken) {
	v.index++
}

func (v *PrintTokenVisitor) VisitNumericConstantToken(tkn intmod.INumericConstantToken) {
	v.prepareNewLine()
	v.printer.PrintNumericConstant(tkn.Text())
	v.index++
}

func (v *PrintTokenVisitor) VisitReferenceToken(tkn intmod.IReferenceToken) {
	v.prepareNewLine()
	v.printer.PrintReference(tkn.Type(), tkn.InternalTypeName(), tkn.Name(), tkn.Descriptor(), tkn.OwnerInternalName())
	v.index++
}

func (v *PrintTokenVisitor) VisitStringConstantToken(tkn intmod.IStringConstantToken) {
	v.prepareNewLine()
	v.printer.PrintStringConstant(fmt.Sprintf("\"%s\"", tkn.Text()), tkn.OwnerInternalName())
	v.index++
}

func (v *PrintTokenVisitor) VisitTextToken(tkn intmod.ITextToken) {
	v.prepareNewLine()
	v.printer.PrintText(tkn.Text())
	v.index++
}

func (v *PrintTokenVisitor) prepareNewLine() {
	if v.newLineCount > 0 {
		v.printer.EndLine()

		if v.newLineCount > 2 {
			v.printer.StartLine(v.newLineCount - 2)
			v.newLineCount = 2
		}

		if v.newLineCount > 1 {
			v.printer.StartLine(UnknownLineNumber)
			v.printer.EndLine()
		}

		v.printer.StartLine(v.searchLineNumber())
		v.newLineCount = 0
	}
}

func (v *PrintTokenVisitor) searchLineNumber() int {
	// Backward search
	v.searchLineNumberVisitor.Reset()

	for i := v.index; i >= 0; i-- {
		v.tokens.Get(i).Accept(v.searchLineNumberVisitor)

		if v.searchLineNumberVisitor.lineNumber != UnknownLineNumber {
			return v.searchLineNumberVisitor.lineNumber
		}
		if v.searchLineNumberVisitor.newLineCounter > 0 {
			break
		}
	}

	// Forward search
	v.searchLineNumberVisitor.Reset()
	size := v.tokens.Size()

	for i := v.index; i < size; i++ {
		v.tokens.Get(i).Accept(v.searchLineNumberVisitor)

		if v.searchLineNumberVisitor.lineNumber != UnknownLineNumber {
			return v.searchLineNumberVisitor.lineNumber
		}
		if v.searchLineNumberVisitor.newLineCounter > 0 {
			break
		}
	}

	return UnknownLineNumber
}

type SearchLineNumberVisitor struct {
	token.AbstractNopTokenVisitor

	lineNumber     int
	newLineCounter int
}

func (v *SearchLineNumberVisitor) Reset() {
	v.lineNumber = UnknownLineNumber
	v.newLineCounter = 0
}

func (v *SearchLineNumberVisitor) VisitLineNumberToken(tkn intmod.ILineNumberToken) {
	v.lineNumber = tkn.LineNumber()
}

func (v *SearchLineNumberVisitor) VisitNewLineToken(_ intmod.INewLineToken) {
	v.newLineCounter++
}
