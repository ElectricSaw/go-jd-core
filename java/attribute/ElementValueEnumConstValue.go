package attribute

func NewElementValueEnumConstValue(descriptor string, constName string) ElementValueEnumConstValue {
	return ElementValueEnumConstValue{
		descriptor: descriptor,
		constName:  constName,
	}
}

type ElementValueEnumConstValue struct {
	descriptor string
	constName  string
}

func (e ElementValueEnumConstValue) Descriptor() string {
	return e.descriptor
}

func (e ElementValueEnumConstValue) ConstName() string {
	return e.constName
}

func (e *ElementValueEnumConstValue) Accept(visitor ElementValueVisitor) {
	visitor.VisitEnumConstValue(e)
}
