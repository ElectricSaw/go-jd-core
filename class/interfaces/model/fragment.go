package model

type IFragment interface {
	AcceptFragmentVisitor(visitor IFragmentVisitor)
}

type IFragmentVisitor interface {
	VisitFlexibleFragment(fragment IFlexibleFragment)
	VisitEndFlexibleBlockFragment(fragment IEndFlexibleBlockFragment)
	VisitEndMovableBlockFragment(fragment IEndMovableBlockFragment)
	VisitSpacerBetweenMovableBlocksFragment(fragment ISpacerBetweenMovableBlocksFragment)
	VisitStartFlexibleBlockFragment(fragment IStartFlexibleBlockFragment)
	VisitStartMovableBlockFragment(fragment IStartMovableBlockFragment)
	VisitFixedFragment(fragment IFixedFragment)
}

type IEndFlexibleBlockFragment interface {
	IFlexibleFragment
}

type IEndMovableBlockFragment interface {
	IFlexibleFragment

	String() string
}

type IFixedFragment interface {
	IFragment

	FirstLineNumber() int
	SetFirstLineNumber(firstLineNumber int)
	LastLineNumber() int
	SetLastLineNumber(lastLineNumber int)

	String() string
}

type IFlexibleFragment interface {
	IFragment

	MinimalLineCount() int
	SetMinimalLineCount(minimalLineCount int)
	MaximalLineCount() int
	SetMaximalLineCount(maximalLineCount int)
	InitialLineCount() int
	SetInitialLineCount(initialLineCount int)
	LineCount() int
	SetLineCount(lineCount int)
	Weight() int
	SetWeight(weight int)
	Label() string
	SetLabel(label string)

	IncLineCount(force bool) bool
	DecLineCount(force bool) bool
	String() string
}

type ISpacerBetweenMovableBlocksFragment interface {
	IFlexibleFragment
}

type IStartFlexibleBlockFragment interface {
	IFlexibleFragment
}

type IStartMovableBlockFragment interface {
	IFlexibleFragment

	Type() int
	SetType(typ int)

	String() string
}
