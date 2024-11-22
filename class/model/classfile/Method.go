package classfile

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
)

func NewMethod(accessFlags int, name string, descriptor string, attributes map[string]attribute.Attribute, constants intmod.IConstantPool) intmod.IMethod {
	return &Method{accessFlags: accessFlags, name: name, descriptor: descriptor, attributes: attributes, constants: constants}
}

type Method struct {
	accessFlags int
	name        string
	descriptor  string
	attributes  map[string]attribute.Attribute
	constants   intmod.IConstantPool
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

func (cf Method) Attribute(name string) attribute.Attribute {
	return cf.attributes[name]
}

func (m Method) Constants() intmod.IConstantPool {
	return m.constants
}

func (m Method) String() string {
	return "Method { name: " + m.Name() + ", descriptor: " + m.Descriptor() + " }"
}
