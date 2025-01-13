package classfile

func NewConstantPool(constants []IConstant) ConstantPool {
	return ConstantPool{
		constants: constants,
	}
}

type IConstantPool interface {
	Constant(index int) IConstant
	ConstantTypeName(index int) (string, bool)
	ConstantString(index int) (string, bool)
	ConstantUtf8(index int) (string, bool)
	ConstantValue(index int) IConstantValue
	String() string
}

type ConstantPool struct {
	constants []IConstant
}

func (p *ConstantPool) Constant(index int) IConstant {
	return p.constants[index]
}

func (p *ConstantPool) ConstantTypeName(index int) (string, bool) {
	if cc, ok := p.constants[index].(*ConstantClass); ok {
		if utf8, ok := p.constants[cc.NameIndex].(*ConstantUtf8); ok {
			return utf8.Value, true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantString(index int) (string, bool) {
	if cc, ok := p.constants[index].(*ConstantString); ok {
		if utf8, ok := p.constants[cc.StringIndex].(*ConstantUtf8); ok {
			return utf8.Value, true
		}
	}
	return "", false
}

func (p *ConstantPool) ConstantUtf8(index int) (string, bool) {
	if utf8, ok := p.constants[index].(*ConstantUtf8); ok {
		return utf8.Value, true
	}
	return "", false
}

func (p *ConstantPool) ConstantValue(index int) IConstantValue {
	c := p.constants[index]

	if c != nil && c.GetTag() == ConstTagString {
		if cs, ok := c.(*ConstantString); ok {
			c = p.constants[cs.StringIndex]
		}
	}

	return c.(IConstantValue)
}

func (p *ConstantPool) String() string {
	return "ConstantPool"
}
