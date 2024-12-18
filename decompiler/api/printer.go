package api

const (
	UnknownLineNumber = iota

	// Declaration & reference types

	Type
	Field
	Method
	Constructor
	Package
	Module

	// Marker Types

	Comment          = 1
	JavaDoc          = 2
	Error            = 3
	ImportStatements = 4
)

type Printer interface {
	Start(maxLineNumber, majorVersion, minorVersion int)
	End()

	PrintText(text string)
	PrintNumericConstant(constant string)
	PrintStringConstant(constant, ownerInternalName string)
	PrintKeyword(keyword string)

	PrintDeclaration(typ int, internalName, name, descriptor string)
	PrintReference(typ int, internalName, name, descriptor, ownerInternalName string)

	Indent()
	Unindent()

	StartLine(lineNumber int)
	EndLine()
	ExtraLine(count int)

	StartMarker(typ int)
	EndMarker(typ int)
}
