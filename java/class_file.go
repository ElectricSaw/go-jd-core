package java

import "bitbucket.org/coontec/javaClass/java/attribute"

type ClassFile struct {
	MajorVersion       int
	MinorVersion       int
	AccessFlags        int
	InternalTypeName   string
	SuperTypeName      string
	InterfaceTypeNames []string
	Field              []Field
	Method             []Method
	Attributes         map[string]attribute.Attribute
}
