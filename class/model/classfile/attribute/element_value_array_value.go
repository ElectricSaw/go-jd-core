package attribute

import intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"

func NewElementValueArrayValue(values []intcls.IElementValue) intcls.IElementValueArrayValue {
	return &ElementValueArrayValue{
		values: values,
	}
}

type ElementValueArrayValue struct {
	values []intcls.IElementValue
}

func (e *ElementValueArrayValue) Values() []intcls.IElementValue {
	return e.values
}

func (e *ElementValueArrayValue) Accept(visitor intcls.IElementValueVisitor) {
	visitor.VisitArrayValue(e)
}
