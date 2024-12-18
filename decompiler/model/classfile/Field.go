package classfile

import (
	intcls "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/classpath"
)

func NewField(accessFlags int, name string, descriptor string, attributes map[string]intcls.IAttribute) intcls.IField {
	return &Field{accessFlags: accessFlags, name: name, descriptor: descriptor, attributes: attributes}
}

type Field struct {
	accessFlags int
	name        string
	descriptor  string
	attributes  map[string]intcls.IAttribute
}

func (f Field) AccessFlags() int {
	return f.accessFlags
}

func (f Field) Name() string {
	return f.name
}

func (f Field) Descriptor() string {
	return f.descriptor
}

func (f Field) Attributes() map[string]intcls.IAttribute {
	return f.attributes
}

func (cf Field) Attribute(name string) intcls.IAttribute {
	return cf.attributes[name]
}

func (f Field) String() string {
	return "Field{" + f.Name() + " " + f.Descriptor() + " }"
}
