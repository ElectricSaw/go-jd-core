package classfile

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
)

func NewMethod(accessFlags int, name string, descriptor string, attributes map[string]intcls.IAttribute, constants intcls.IConstantPool) intcls.IMethod {
	return &Method{accessFlags: accessFlags, name: name, descriptor: descriptor, attributes: attributes, constants: constants}
}

type Method struct {
	accessFlags int
	name        string
	descriptor  string
	attributes  map[string]intcls.IAttribute
	constants   intcls.IConstantPool
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

func (m Method) Attributes() map[string]intcls.IAttribute {
	return m.attributes
}

func (cf Method) Attribute(name string) intcls.IAttribute {
	return cf.attributes[name]
}

func (m Method) Constants() intcls.IConstantPool {
	return m.constants
}

func (m Method) String() string {
	return "Method { name: " + m.Name() + ", descriptor: " + m.Descriptor() + " }"
}
