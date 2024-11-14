package classfile

import (
	"bitbucket.org/coontec/javaClass/class/model/classfile/attribute"
)

func NewClassFile(majorVersion int, minorVersion int, accessFlags int, internalTypeName string, superTypeName string, interfaceTypeNames []string, field []Field, method []Method, attributes map[string]attribute.Attribute) *ClassFile {
	return &ClassFile{
		majorVersion:       majorVersion,
		minorVersion:       minorVersion,
		accessFlags:        accessFlags,
		internalTypeName:   internalTypeName,
		superTypeName:      superTypeName,
		interfaceTypeNames: interfaceTypeNames,
		field:              field,
		method:             method,
		attributes:         attributes,
	}
}

type IClassFile interface {
	MajorVersion() int
	MinorVersion() int
	IsEnum() bool
	IsAnnotation() bool
	IsInterface() bool
	IsModule() bool
	IsStatic() bool
	AccessFlags() int
	InternalTypeName() string
	SuperTypeName() string
	InterfaceTypeNames() []string
	Field() []Field
	Method() []Method
	Attributes() map[string]attribute.Attribute
}

type ClassFile struct {
	majorVersion       int
	minorVersion       int
	accessFlags        int
	internalTypeName   string
	superTypeName      string
	interfaceTypeNames []string
	field              []Field
	method             []Method
	attributes         map[string]attribute.Attribute
	OuterClassFile     IClassFile
	InnerClassFiles    []IClassFile
}

func (cf ClassFile) MajorVersion() int {
	return cf.majorVersion
}

func (cf ClassFile) MinorVersion() int {
	return cf.minorVersion
}

func (cf ClassFile) IsEnum() bool {
	return (cf.accessFlags & AccEnum) != 0
}

func (cf ClassFile) IsAnnotation() bool {
	return (cf.accessFlags & AccAnnotation) != 0
}

func (cf ClassFile) IsInterface() bool {
	return (cf.accessFlags & AccInterface) != 0
}

func (cf ClassFile) IsModule() bool {
	return (cf.accessFlags & AccModule) != 0
}

func (cf ClassFile) IsStatic() bool {
	return (cf.accessFlags & AccStatic) != 0
}

func (cf ClassFile) AccessFlags() int {
	return cf.accessFlags
}

func (cf ClassFile) SetAccessFlags(accessFlags int) {
	cf.accessFlags = accessFlags
}

func (cf ClassFile) InternalTypeName() string {
	return cf.internalTypeName
}

func (cf ClassFile) SuperTypeName() string {
	return cf.superTypeName
}

func (cf ClassFile) InterfaceTypeNames() []string {
	return cf.interfaceTypeNames
}

func (cf ClassFile) Field() []Field {
	return cf.field
}

func (cf ClassFile) Method() []Method {
	return cf.method
}

func (cf ClassFile) Attributes() map[string]attribute.Attribute {
	return cf.attributes
}

func (cf ClassFile) String() string {
	return "ClassFile { Internal Type Name: " + cf.InternalTypeName() + " }"
}
