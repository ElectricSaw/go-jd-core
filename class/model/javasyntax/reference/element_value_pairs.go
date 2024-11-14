package reference

import "fmt"

type ElementValuePairs struct {
	IElementValuePair

	ElementValuePair []IElementValuePair
}

func (e *ElementValuePairs) List() []IReference {
	ret := make([]IReference, 0, len(e.ElementValuePair))
	for _, v := range e.ElementValuePair {
		ret = append(ret, v)
	}
	return ret
}

func (e *ElementValuePairs) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValuePairs(e)
}

func (e *ElementValuePairs) String() string {
	return fmt.Sprintf("ElementValuePairs{%v}", *e)
}
