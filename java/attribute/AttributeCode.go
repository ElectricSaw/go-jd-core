package attribute

func NewAttributeCode(maxStack int, maxLocals int, code []byte, exceptionTable []CodeException, attribute map[string]Attribute) AttributeCode {
	return AttributeCode{maxStack, maxLocals, code, exceptionTable, attribute}
}

type AttributeCode struct {
	maxStack       int
	maxLocals      int
	code           []byte
	exceptionTable []CodeException
	attribute      map[string]Attribute
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

func (a AttributeCode) ExceptionTable() []CodeException {
	return a.exceptionTable
}

func (a AttributeCode) Attribute(name string) Attribute {
	if a.attribute == nil {
		return nil
	}

	return a.attribute[name]
}

func (a AttributeCode) attributeIgnoreFunc() {}
