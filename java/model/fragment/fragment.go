package fragment

type Fragment interface {
	Accept(visitor FragmentVisitor)
}

type FragmentVisitor interface {
	VisitFlexibleFragment(fragment *FlexibleFragment)
	VisitEndFlexibleBlockFragment(fragment *EndFlexibleBlockFragment)
	VisitEndMovableBlockFragment(fragment *EndMovableBlockFragment)
	VisitSpacerBetweenMovableBlocksFragment(fragment *SpacerBetweenMovableBlocksFragment)
	VisitStartFlexibleBlockFragment(fragment *StartFlexibleBlockFragment)
	VisitStartMovableBlockFragment(fragment *StartMovableBlockFragment)
	VisitFixedFragment(fragment *FixedFragment)
}
