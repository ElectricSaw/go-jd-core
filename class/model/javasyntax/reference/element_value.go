package reference

const (
	EvUnknown = iota
	EvPrimitiveType
	EvEnumConstValue
	EvClassInfo
	EvAnnotationValue
	EvArrayValue
)

type IElementValue interface {
	IReference
}

type AbstractElementValue struct {
}

func (e *AbstractElementValue) Accept(visitor ReferenceVisitor) {

}

type IElementValuePair interface {
	IReference
}

type AbstractElementValuePair struct {
}

func (e *AbstractElementValuePair) Accept(visitor ReferenceVisitor) {

}
