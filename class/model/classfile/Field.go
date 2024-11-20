package classfile

import (
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
)

func NewField(accessFlags int, name string, descriptor string, attributes map[string]attribute.Attribute) *Field {
	return &Field{accessFlags: accessFlags, name: name, descriptor: descriptor, attributes: attributes}
}

type Field struct {
	accessFlags int
	name        string
	descriptor  string
	attributes  map[string]attribute.Attribute
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

func (f Field) Attributes() map[string]attribute.Attribute {
	return f.attributes
}

func (f Field) String() string {
	return "Field{" + f.Name() + " " + f.Descriptor() + " }"
}
