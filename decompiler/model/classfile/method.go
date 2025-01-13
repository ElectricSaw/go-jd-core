package classfile

func NewMethod(accessFlags int, name string, descriptor string, attributes map[string]IAttribute, constants ConstantPool) Method {
	return Method{
		AccessFlags: accessFlags,
		Name:        name,
		Descriptor:  descriptor,
		Attributes:  attributes,
		Constants:   constants,
	}
}

type IMethod interface {
	Attribute(name string) IAttribute
	String() string
}

type Method struct {
	AccessFlags int
	Name        string
	Descriptor  string
	Attributes  map[string]IAttribute
	Constants   ConstantPool
}

func (m Method) Attribute(name string) IAttribute {
	return m.Attributes[name]
}

func (m Method) String() string {
	return "Method { name: " + m.Name + ", descriptor: " + m.Descriptor + " }"
}
