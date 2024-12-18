package attribute

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
)

func NewElementValuePrimitiveType(type_ int, constValue intcls.IConstantValue) intcls.IElementValuePrimitiveType {
	return &ElementValuePrimitiveType{
		type_:      type_,
		constValue: constValue,
	}
}

type ElementValuePrimitiveType struct {
	/*
	 * type = {'B', 'D', 'F', 'I', 'J', 'S', 'Z', 'C', 's'}
	 */
	type_      int
	constValue intcls.IConstantValue
}

func (e *ElementValuePrimitiveType) Type() int {
	return e.type_
}

func (e *ElementValuePrimitiveType) SetType(type_ int) {
	e.type_ = type_
}

func (e *ElementValuePrimitiveType) Value() intcls.IConstantValue {
	return e.constValue
}

func (e *ElementValuePrimitiveType) SetValue(constValue intcls.IConstantValue) {
	e.constValue = constValue
}

func (e *ElementValuePrimitiveType) Accept(visitor intcls.IElementValueVisitor) {
	visitor.VisitPrimitiveType(e)
}
