package model

const (
	UnknownLineNumberToken = 0
)

const (
	TypeToken        = 1
	FieldToken       = 2
	MethodToken      = 3
	ConstructorToken = 4
	PackageToken     = 5
	ModuleToken      = 6
)

type IToken interface {
	Accept(visitor ITokenVisitor)
}

type ITokenVisitor interface {
	VisitBooleanConstantToken(token IBooleanConstantToken)
	VisitCharacterConstantToken(token ICharacterConstantToken)
	VisitDeclarationToken(token IDeclarationToken)
	VisitEndBlockToken(token IEndBlockToken)
	VisitEndMarkerToken(token IEndMarkerToken)
	VisitKeywordToken(token IKeywordToken)
	VisitLineNumberToken(token ILineNumberToken)
	VisitNewLineToken(token INewLineToken)
	VisitNumericConstantToken(token INumericConstantToken)
	VisitReferenceToken(token IReferenceToken)
	VisitStartBlockToken(token IStartBlockToken)
	VisitStartMarkerToken(token IStartMarkerToken)
	VisitStringConstantToken(token IStringConstantToken)
	VisitTextToken(token ITextToken)
}

type IBooleanConstantToken interface {
	Value() bool
	SetValue(value bool)
	Accept(visitor ITokenVisitor)
	String() string
}

type ICharacterConstantToken interface {
	Character() string
	SetCharacter(character string)
	OwnerInternalName() string
	SetOwnerInternalName(ownerInternalName string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IDeclarationToken interface {
	Type() int
	SetType(typ int)
	InternalTypeName() string
	SetInternalTypeName(internalTypeName string)
	Name() string
	SetName(name string)
	Descriptor() string
	SetDescriptor(descriptor string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IEndBlockToken interface {
	Text() string
	SetText(text string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IEndMarkerToken interface {
	Type() int
	SetType(typ int)
	Accept(visitor ITokenVisitor)
	String() string
}

type IKeywordToken interface {
	Keyword() string
	SetKeyword(keyword string)
	Accept(visitor ITokenVisitor)
	String() string
}

type ILineNumberToken interface {
	LineNumber() int
	SetLineNumber(lineNumber int)
	Accept(visitor ITokenVisitor)
	String() string
}

type INewLineToken interface {
	Count() int
	SetCount(count int)
	Accept(visitor ITokenVisitor)
	String() string
}

type INumericConstantToken interface {
	Text() string
	SetText(text string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IReferenceToken interface {
	IDeclarationToken

	OwnerInternalName() string
	SetOwnerInternalName(ownerInternalName string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IStartBlockToken interface {
	Text() string
	SetText(text string)
	Accept(visitor ITokenVisitor)
	String() string
}

type IStartMarkerToken interface {
	Type() int
	SetType(typ int)
	Accept(visitor ITokenVisitor)
	String() string
}

type IStringConstantToken interface {
	Text() string
	SetText(text string)
	OwnerInternalName() string
	SetOwnerInternalName(ownerInternalName string)
	Accept(visitor ITokenVisitor)
	String() string
}

type ITextToken interface {
	Text() string
	SetText(text string)
	Accept(visitor ITokenVisitor)
	String() string
}
