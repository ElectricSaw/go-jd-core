package classfile

import (
	"bitbucket.org/coontec/javaClass/java/model/classfile/constant"
)

type ConstantPool struct {
	constants []constant.Constant
}

func (p *ConstantPool) GetConstant(index int) constant.Constant {
	return p.constants[index]
}

func (p *ConstantPool) GetConstantTypeName(index int) (string, bool) {
	if cc, ok := p.constants[index].(constant.ConstantClass); ok {
		if utf8, ok := p.constants[cc.NameIndex()].(constant.ConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) getConstantString(index int) (string, bool) {
	if cc, ok := p.constants[index].(constant.ConstantString); ok {
		if utf8, ok := p.constants[cc.StringIndex()].(constant.ConstantUtf8); ok {
			return utf8.Value(), true
		}
	}
	return "", false
}

func (p *ConstantPool) getConstantUtf8(index int) (string, bool) {
	if utf8, ok := p.constants[index].(constant.ConstantUtf8); ok {
		return utf8.Value(), true
	}
	return "", false
}

func (p *ConstantPool) getConstantValue(index int) constant.ConstantValue {
	c := p.constants[index]

	if c != nil && c.Tag() == constant.ConstTagString {
		if cs, ok := c.(constant.ConstantString); ok {
			c = p.constants[cs.StringIndex()]
		}
	}

	return c.(constant.ConstantValue)
}

func (p *ConstantPool) String() string {
	return "ConstantPool"
}
