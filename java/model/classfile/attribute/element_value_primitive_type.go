package attribute

import (
	"bitbucket.org/coontec/javaClass/java/model/classfile/constant"
)

func NewElementValuePrimitiveType(type_ int, constValue constant.ConstantValue) ElementValuePrimitiveType {
	return ElementValuePrimitiveType{
		type_:      type_,
		constValue: constValue,
	}
}

type ElementValuePrimitiveType struct {
	/*
	 * type = {'B', 'D', 'F', 'I', 'J', 'S', 'Z', 'C', 's'}
	 */
	type_      int
	constValue constant.ConstantValue
}

func (e *ElementValuePrimitiveType) Type() int {
	return e.type_
}

func (e *ElementValuePrimitiveType) ConstValue() constant.ConstantValue {
	return e.constValue
}

func (e *ElementValuePrimitiveType) Accept(visitor ElementValueVisitor) {
	visitor.VisitPrimitiveType(e)
}
