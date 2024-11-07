package reference

import "fmt"

type ElementValuePairs struct {
	IElementValuePair

	ElementValuePair []IElementValuePair
}

func (e *ElementValuePairs) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValuePairs(e)
}

func (e *ElementValuePairs) String() string {
	return fmt.Sprintf("ElementValuePairs{%v}", *e)
}
