package testutils

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	"strings"
)

func NewCounterPrinter(escapeUnicodeCharacters bool) *CounterPrinter {
	return &CounterPrinter{
		PlainTextPrinter: PlainTextPrinter{
			EscapeUnicodeCharacters: escapeUnicodeCharacters,
		},
	}
}

func NewCounterPrinterEmpty() *CounterPrinter {
	return &CounterPrinter{
		PlainTextPrinter: PlainTextPrinter{
			EscapeUnicodeCharacters: false,
		},
	}
}

type CounterPrinter struct {
	PlainTextPrinter

	classCounter         int
	methodCounter        int
	errorInMethodCounter int
	accessCounter        int
}

func (p *CounterPrinter) PrintText(text string) {
	if text != "" {
		if text == "// Byte Code:" || strings.HasPrefix(text, "/* monitor enter ") || strings.HasPrefix(text, "/* monitor exit ") {
			p.errorInMethodCounter++
		}
	}
	p.PlainTextPrinter.PrintText(text)
}

func (p *CounterPrinter) PrintDeclaration(typ int, internalTypeName, name, descriptor string) {
	if typ == api.Type {
		p.classCounter++
	}
	if typ == api.Method || typ == api.Constructor {
		p.methodCounter++
	}

	p.PlainTextPrinter.PrintDeclaration(typ, internalTypeName, name, descriptor)
}

func (p *CounterPrinter) PrintReference(typ int, internalTypeName, name, descriptor, ownerInternalName string) {
	if name != "" && strings.HasPrefix(name, "access$") {
		p.accessCounter++
	}
	p.PlainTextPrinter.PrintReference(typ, internalTypeName, name, descriptor, ownerInternalName)
}
