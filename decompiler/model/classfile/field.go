package classfile

func NewField(accessFlags int, name string, descriptor string, attributes map[string]IAttribute) Field {
	return Field{
		AccessFlags: accessFlags,
		Name:        name,
		Descriptor:  descriptor,
		Attributes:  attributes,
	}
}

type IField interface {
	Attribute(name string) IAttribute
	String() string
}

type Field struct {
	AccessFlags int
	Name        string
	Descriptor  string
	Attributes  map[string]IAttribute
}

func (f Field) Attribute(name string) IAttribute {
	return f.Attributes[name]
}

func (f Field) String() string {
	return "Field{" + f.Name + " " + f.Descriptor + " }"
}
