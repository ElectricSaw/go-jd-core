package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
	"sort"
)

var CommaFrag = NewTokensFragment(&CommaTkn)
var SemicolonFrag = NewTokensFragment(&SemicolonTkn)
var StartDeclarationOrStatementBlockFrag = NewTokensFragment(&StartDeclarationOrStatementBlockTkn)
var EndDeclarationOrStatementBlockFrag = NewTokensFragment(&EndDeclarationOrStatementBlockTkn)
var EndDeclarationOrStatementBlockSemicolonFrag = NewTokensFragment(&EndDeclarationOrStatementBlockTkn, &SemicolonTkn)
var ReturnSemicolonFrag = NewTokensFragment(&ReturnTkn, &SemicolonTkn)

func NewImportsFragment(weight int) ImportsFragment {
	return ImportsFragment{
		FlexibleFragment: NewFlexibleFragment(0, -1,
			-1, weight, "Imports"),
		ImportMap: make(map[string]Import),
	}
}

func NewImport(internalName string, qualifiedName string) Import {
	return Import{
		InternalName:  internalName,
		QualifiedName: qualifiedName,
		Counter:       1,
	}
}

func NewLineNumberTokensFragment(tokens ...IToken) LineNumberTokensFragment {
	frag := LineNumberTokensFragment{
		Tokens: util.NewDefaultListWithElements[IToken](tokens...),
	}
	return frag
}

func NewStartStatementsBlockFragmentGroup() StartStatementsBlockFragmentGroup {
	return StartStatementsBlockFragmentGroup{minimalLineCount: math.MaxInt}
}

func SearchFirstLineNumber(tokens util.IList[IToken]) int {
	visitor := NewSearchLineNumberVisitor()

	for _, tkn := range tokens.ToSlice() {
		tkn.Accept(&visitor)

		if visitor.LineNumber != UnknownLineNumberToken {
			return visitor.LineNumber - visitor.NewLineCounter
		}
	}

	return UnknownLineNumberToken
}

func SearchLastLineNumber(tokens util.IList[IToken]) int {
	visitor := NewSearchLineNumberVisitor()
	index := tokens.Size()

	for index > 0 {
		index--
		tokens.Get(index).Accept(&visitor)

		if visitor.LineNumber != UnknownLineNumberToken {
			return visitor.LineNumber + visitor.NewLineCounter
		}
	}

	return UnknownLineNumberToken
}

type StartStatementsBlockFragmentGroup struct {
	util.DefaultList[IFlexibleFragment]
	minimalLineCount int
}

func (f *StartStatementsBlockFragmentGroup) MinimalLineCount() int {
	if f.minimalLineCount == math.MaxInt {
		for _, frag := range f.ToSlice() {
			if f.minimalLineCount > frag.GetLineCount() {
				f.minimalLineCount = frag.GetLineCount()
			}
		}
	}

	return f.minimalLineCount
}

func (f *StartStatementsBlockFragmentGroup) String() string {
	sb := "StartStatementsBlockFragmentGroup{"
	sb += fmt.Sprintf("minimalLineCount: %d", f.minimalLineCount)
	if !f.IsEmpty() {
		sb += ", group: "
		slice := f.ToSlice()
		for i := 0; i < f.Size(); i++ {
			sb += fmt.Sprintf("%s", slice[i])
			if i < f.Size()-1 {
				sb += ", "
			}
		}
	}
	sb += "}"

	return sb
}

func NewTokensFragment(tokens ...IToken) TokensFragment {
	return NewTokensFragmentWithSlice(tokens)
}

func NewTokensFragmentWithSlice(tokens []IToken) TokensFragment {
	return newTokensFragment(getLineCount(tokens), tokens)
}

func newTokensFragment(lineCount int, tokens []IToken) TokensFragment {
	return TokensFragment{
		FlexibleFragment: NewFlexibleFragment(lineCount, lineCount, lineCount,
			0, "Tokens"),
		Tokens: util.NewDefaultListWithElements[IToken](tokens...),
	}
}

func getLineCount(tokens []IToken) int {
	visitor := NewLineCountVisitor()

	for _, tkn := range tokens {
		tkn.Accept(&visitor)
	}

	return visitor.LineCount
}

type ImportsFragment struct {
	FlexibleFragment

	ImportMap map[string]Import
}

func (f *ImportsFragment) AddImport(internalName, qualifiedName string) {
	imp, ok := f.ImportMap[internalName]
	if ok {
		imp.IncCounter()
	} else {
		f.ImportMap[internalName] = NewImport(internalName, qualifiedName)
	}
}

func (f *ImportsFragment) IncCounter(internalName string) bool {
	imp, ok := f.ImportMap[internalName]
	if ok {
		imp.IncCounter()
		return true
	} else {
		return false
	}
}

func (f *ImportsFragment) IsEmpty() bool {
	return len(f.ImportMap) == 0
}

func (f *ImportsFragment) InitLineCounts() {
	f.MaximalLineCount = len(f.ImportMap)
	f.InitialLineCount = f.MaximalLineCount
	f.LineCount = f.MaximalLineCount
}

func (f *ImportsFragment) Contains(internalName string) bool {
	_, ok := f.ImportMap[internalName]
	return ok
}

func (f *ImportsFragment) Import(internalName string) (Import, bool) {
	if imp0rt, ok := f.ImportMap[internalName]; ok {
		return imp0rt, true
	}
	var zero Import
	return zero, false
}

func (f *ImportsFragment) Imports() []Import {
	lineCount := f.LineCount
	size := len(f.ImportMap)

	imports := make([]Import, 0, len(f.ImportMap))
	for _, v := range f.ImportMap {
		imports = append(imports, v)
	}

	if lineCount < size {
		sort.Slice(imports, func(i, j int) bool {
			return imports[i].Counter > imports[j].Counter
		})

		subList := imports[lineCount:size]

		for _, imp0rt := range subList {
			delete(f.ImportMap, imp0rt.InternalName)
		}
	}
	return imports
}

func (f *ImportsFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitImportsFragment(f)
}

func (f *ImportsFragment) String() string {
	return fmt.Sprintf("ImportsFragment { fragment: %s, importMap: %d }", f.FlexibleFragment.String(), len(f.ImportMap))
}

type Import struct {
	InternalName  string
	QualifiedName string
	Counter       int
}

func (i *Import) IncCounter() {
	i.Counter++
}

func (i *Import) String() string {
	return fmt.Sprintf("Import { internalName: %s, qualitifedName: %s, counter: %d }", i.InternalName, i.QualifiedName, i.Counter)
}

type LineNumberTokensFragment struct {
	FixedFragment

	Tokens util.IList[IToken]
}

func (f *LineNumberTokensFragment) TokenAt(index int) IToken {
	return f.Tokens.Get(index)
}

func (f *LineNumberTokensFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitLineNumberTokensFragment(f)
}

type TokensFragment struct {
	FlexibleFragment

	Tokens util.IList[IToken]
}

func (f *TokensFragment) TokenAt(index int) IToken {
	return f.Tokens.Get(index)
}

func (f *TokensFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitTokensFragment(f)
}
