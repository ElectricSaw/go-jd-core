package attribute

func NewElementValueArrayValue(values []ElementValue) ElementValueArrayValue {
	return ElementValueArrayValue{
		values: values,
	}
}

type ElementValueArrayValue struct {
	values []ElementValue
}

func (e ElementValueArrayValue) Values() []ElementValue {
	return e.values
}

func (e *ElementValueArrayValue) Accept(visitor ElementValueVisitor) {
	visitor.VisitArrayValue(e)
}
