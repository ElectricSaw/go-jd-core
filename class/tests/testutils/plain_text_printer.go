package testutils

import "fmt"

const (
	Tab     = "  "
	NewLine = "\n"
)

func NewPlainTextPrinter(escapeUnicodeCharacters bool) *PlainTextPrinter {
	return &PlainTextPrinter{
		EscapeUnicodeCharacters: escapeUnicodeCharacters,
	}
}

func NewPlainTextPrinterEmpty() *PlainTextPrinter {
	return &PlainTextPrinter{
		EscapeUnicodeCharacters: false,
	}
}

type PlainTextPrinter struct {
	IndentationCount        int
	Source                  string
	RealLineNumber          int
	Format                  string
	EscapeUnicodeCharacters bool
}

func (p *PlainTextPrinter) Init() {
	p.Source = ""
	p.RealLineNumber = 0
	p.IndentationCount = 0
}

// --- Printer ---

func (p *PlainTextPrinter) Start(maxLineNumber, majorVersion, minorVersion int) {
	p.IndentationCount = 0

	if maxLineNumber == 0 {
		p.Format = "%4d"
	} else {
		width := 2

		for maxLineNumber >= 10 {
			width++
			maxLineNumber /= 10
		}

		p.Format = fmt.Sprintf("%%%dd", width)
	}
}

func (p *PlainTextPrinter) End() {
}

func (p *PlainTextPrinter) PrintText(text string) {
	if p.EscapeUnicodeCharacters {
		length := len(text)
		for i := 0; i < length; i++ {
			c := rune(text[i])

			if c < 127 {
				p.Source += string(c)
			} else {
				p.Source += "\\u"

				h := c >> 12
				if h <= 9 {
					p.Source += string(h + '0')
				} else {
					p.Source += string(h + ('A' - 10))
				}

				h = (c >> 8) & 15
				if h <= 9 {
					p.Source += string(h + '0')
				} else {
					p.Source += string(h + ('A' - 10))
				}

				h = (c >> 4) & 15
				if h <= 9 {
					p.Source += string(h + '0')
				} else {
					p.Source += string(h + ('A' - 10))
				}

				h = (c) & 15
				if h <= 9 {
					p.Source += string(h + '0')
				} else {
					p.Source += string(h + ('A' - 10))
				}
			}
		}
	} else {
		p.Source += text
	}
}

func (p *PlainTextPrinter) PrintNumericConstant(constant string) {
	p.Source += constant
}

func (p *PlainTextPrinter) PrintStringConstant(constant, _ string) {
	p.PrintText(constant)
}

func (p *PlainTextPrinter) PrintKeyword(keyword string) {
	p.Source += keyword
}

func (p *PlainTextPrinter) PrintDeclaration(_ int, _, name, _ string) {
	p.PrintText(name)
}

func (p *PlainTextPrinter) PrintReference(_ int, _, name, _, _ string) {
	p.PrintText(name)
}

func (p *PlainTextPrinter) Indent() {
	p.IndentationCount++
}

func (p *PlainTextPrinter) Unindent() {
	if p.IndentationCount > 0 {
		p.IndentationCount--
	}
}

func (p *PlainTextPrinter) StartLine(lineNumber int) {
	p.PrintLineNumber(lineNumber)

	for i := 0; i < p.IndentationCount; i++ {
		p.Source += Tab
	}
}

func (p *PlainTextPrinter) EndLine() {
	p.Source += NewLine
}

func (p *PlainTextPrinter) ExtraLine(count int) {
	for count > 0 {
		count--
		p.PrintLineNumber(0)
		p.Source += NewLine
	}
}

func (p *PlainTextPrinter) StartMarker(typ int) {
}

func (p *PlainTextPrinter) EndMarker(typ int) {
}

func (p *PlainTextPrinter) PrintLineNumber(lineNumber int) {
	p.Source += "/* "
	p.RealLineNumber++
	p.Source += fmt.Sprintf(p.Format, p.RealLineNumber)
	p.Source += ":"
	p.Source += fmt.Sprintf(p.Format, lineNumber)
	p.Source += " */ "
}

func (p *PlainTextPrinter) String() string {
	return p.Source
}
