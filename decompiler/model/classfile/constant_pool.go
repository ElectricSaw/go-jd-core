package classfile

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/classfile/constant"
)

func NewConstantPool(constants []intcls.IConstant) intcls.IConstantPool {
	return &ConstantPool{
		constants: constants,
	}
}

type ConstantPool struct {
	constants []intcls.IConstant
}

func (p *ConstantPool) Constant(index int) intcls.IConstant {
	return p.constants[index]
}

func (p *ConstantPool) ConstantTypeName(index int) (string, bool) {
	if cc, ok := p.constants[index].(intcls.IConstantClass); ok {
		if utf8, ok := p.constants[cc.NameIndex()].(intcls.IConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantString(index int) (string, bool) {
	if cc, ok := p.constants[index].(intcls.IConstantString); ok {
		if utf8, ok := p.constants[cc.StringIndex()].(intcls.IConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantUtf8(index int) (string, bool) {
	if utf8, ok := p.constants[index].(intcls.IConstantUtf8); ok {
		return utf8.Value(), true
	}
	return "", false
}

func (p *ConstantPool) ConstantValue(index int) intcls.IConstantValue {
	c := p.constants[index]

	if c != nil && c.Tag() == intcls.ConstTagString {
		if cs, ok := c.(*constant.ConstantString); ok {
			c = p.constants[cs.StringIndex()]
		}
	}

	return c.(intcls.IConstantValue)
}

func (p *ConstantPool) String() string {
	return "ConstantPool"
}
