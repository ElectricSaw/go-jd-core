package classfile

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/constant"
)

func NewConstantPool(constants []constant.IConstant) intmod.IConstantPool {
	return &ConstantPool{
		constants: constants,
	}
}

type ConstantPool struct {
	constants []constant.IConstant
}

func (p *ConstantPool) Constant(index int) constant.IConstant {
	return p.constants[index]
}

func (p *ConstantPool) ConstantTypeName(index int) (string, bool) {
	if cc, ok := p.constants[index].(*constant.ConstantClass); ok {
		if utf8, ok := p.constants[cc.NameIndex()].(*constant.ConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantString(index int) (string, bool) {
	if cc, ok := p.constants[index].(*constant.ConstantString); ok {
		if utf8, ok := p.constants[cc.StringIndex()].(*constant.ConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantUtf8(index int) (string, bool) {
	if utf8, ok := p.constants[index].(*constant.ConstantUtf8); ok {
		return utf8.Value(), true
	}
	return "", false
}

func (p *ConstantPool) ConstantValue(index int) constant.ConstantValue {
	c := p.constants[index]

	if c != nil && c.Tag() == constant.ConstTagString {
		if cs, ok := c.(*constant.ConstantString); ok {
			c = p.constants[cs.StringIndex()]
		}
	}

	return c.(constant.ConstantValue)
}

func (p *ConstantPool) String() string {
	return "ConstantPool"
}
