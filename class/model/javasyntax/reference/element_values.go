package reference

import "fmt"

type ElementValues struct {
	IElementValuePair

	ElementValuePair []IElementValue
}

func (e *ElementValues) List() []IReference {
	ret := make([]IReference, 0, len(e.ElementValuePair))
	for _, v := range e.ElementValuePair {
		ret = append(ret, v)
	}
	return ret
}

func (e *ElementValues) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValues(e)
}

func (e *ElementValues) String() string {
	return fmt.Sprintf("ElementValues{%s}", *e)
}
