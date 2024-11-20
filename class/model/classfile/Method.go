package classfile

import (
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
)

func NewMethod(accessFlags int, name string, descriptor string, attributes map[string]attribute.Attribute, constants ConstantPool) *Method {
	return &Method{accessFlags: accessFlags, name: name, descriptor: descriptor, attributes: attributes, constants: constants}
}

type Method struct {
	accessFlags int
	name        string
	descriptor  string
	attributes  map[string]attribute.Attribute
	constants   ConstantPool
}

func (m Method) AccessFlags() int {
	return m.accessFlags
}

func (m Method) Name() string {
	return m.name
}

func (m Method) Descriptor() string {
	return m.descriptor
}

func (m Method) Attributes() map[string]attribute.Attribute {
	return m.attributes
}

func (m Method) Constants() ConstantPool {
	return m.constants
}

func (m Method) String() string {
	return "Method { name: " + m.Name() + ", descriptor: " + m.Descriptor() + " }"
}
