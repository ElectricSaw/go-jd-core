package reference

import "fmt"

type ElementValues struct {
	IElementValue

	ElementValuePair []IElementValue
}

func (e *ElementValues) Accept(visitor ReferenceVisitor) {
	visitor.VisitElementValues(e)
}

func (e *ElementValues) String() string {
	return fmt.Sprintf("ElementValues{%s}", *e)
}
