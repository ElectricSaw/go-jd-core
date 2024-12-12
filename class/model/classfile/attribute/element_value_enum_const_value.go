package attribute

import intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"

func NewElementValueEnumConstValue(descriptor string, constName string) intcls.IElementValueEnumConstValue {
	return &ElementValueEnumConstValue{
		descriptor: descriptor,
		constName:  constName,
	}
}

type ElementValueEnumConstValue struct {
	descriptor string
	constName  string
}

func (e *ElementValueEnumConstValue) Descriptor() string {
	return e.descriptor
}

func (e *ElementValueEnumConstValue) ConstName() string {
	return e.constName
}

func (e *ElementValueEnumConstValue) Accept(visitor intcls.IElementValueVisitor) {
	visitor.VisitEnumConstValue(e)
}
