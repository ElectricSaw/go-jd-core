package attribute

import intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"

func NewAttributeCode(maxStack int, maxLocals int, code []byte,
	exceptionTable []intcls.ICodeException, attribute map[string]intcls.IAttribute) intcls.IAttributeCode {
	return &AttributeCode{maxStack, maxLocals, code, exceptionTable, attribute}
}

type AttributeCode struct {
	maxStack       int
	maxLocals      int
	code           []byte
	exceptionTable []intcls.ICodeException
	attribute      map[string]intcls.IAttribute
}

func (a AttributeCode) MaxStack() int {
	return a.maxStack
}

func (a AttributeCode) MaxLocals() int {
	return a.maxLocals
}

func (a AttributeCode) Code() []byte {
	return a.code
}

func (a AttributeCode) ExceptionTable() []intcls.ICodeException {
	return a.exceptionTable
}

func (a AttributeCode) Attribute(name string) intcls.IAttribute {
	if a.attribute == nil {
		return nil
	}

	return a.attribute[name]
}

func (a AttributeCode) IsAttribute() bool {
	return true
}
