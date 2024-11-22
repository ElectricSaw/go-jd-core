package classfile

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
)

func NewClassFile(majorVersion int, minorVersion int, accessFlags int,
	internalTypeName string, superTypeName string, interfaceTypeNames []string,
	field []intmod.IField, method []intmod.IMethod, attributes map[string]attribute.Attribute) intmod.IClassFile {
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

type ClassFile struct {
	majorVersion       int
	minorVersion       int
	accessFlags        int
	internalTypeName   string
	superTypeName      string
	interfaceTypeNames []string
	field              []intmod.IField
	method             []intmod.IMethod
	attributes         map[string]attribute.Attribute
	outerClassFile     intmod.IClassFile
	innerClassFiles    []intmod.IClassFile
}

func (cf ClassFile) MajorVersion() int {
	return cf.majorVersion
}

func (cf ClassFile) MinorVersion() int {
	return cf.minorVersion
}

func (cf ClassFile) IsEnum() bool {
	return (cf.accessFlags & intmod.AccEnum) != 0
}

func (cf ClassFile) IsAnnotation() bool {
	return (cf.accessFlags & intmod.AccAnnotation) != 0
}

func (cf ClassFile) IsInterface() bool {
	return (cf.accessFlags & intmod.AccInterface) != 0
}

func (cf ClassFile) IsModule() bool {
	return (cf.accessFlags & intmod.AccModule) != 0
}

func (cf ClassFile) IsStatic() bool {
	return (cf.accessFlags & intmod.AccStatic) != 0
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

func (cf ClassFile) Fields() []intmod.IField {
	return cf.field
}

func (cf ClassFile) Methods() []intmod.IMethod {
	return cf.method
}

func (cf ClassFile) Attributes() map[string]attribute.Attribute {
	return cf.attributes
}

func (cf ClassFile) Attribute(name string) attribute.Attribute {
	return cf.attributes[name]
}

func (cf ClassFile) OuterClassFile() intmod.IClassFile {
	return cf.outerClassFile
}

func (cf ClassFile) InnerClassFiles() []intmod.IClassFile {
	return cf.innerClassFiles
}

func (cf ClassFile) String() string {
	return "ClassFile { Internal Type Name: " + cf.InternalTypeName() + " }"
}
