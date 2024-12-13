package classfile

import (
	intcls "github.com/ElectricSaw/go-jd-core/class/interfaces/classpath"
)

func NewClassFile(majorVersion int, minorVersion int, accessFlags int,
	internalTypeName string, superTypeName string, interfaceTypeNames []string,
	field []intcls.IField, method []intcls.IMethod, attributes map[string]intcls.IAttribute) intcls.IClassFile {
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
	field              []intcls.IField
	method             []intcls.IMethod
	attributes         map[string]intcls.IAttribute
	outerClassFile     intcls.IClassFile
	innerClassFiles    []intcls.IClassFile
}

func (cf ClassFile) MajorVersion() int {
	return cf.majorVersion
}

func (cf ClassFile) MinorVersion() int {
	return cf.minorVersion
}

func (cf ClassFile) IsEnum() bool {
	return (cf.accessFlags & intcls.AccEnum) != 0
}

func (cf ClassFile) IsAnnotation() bool {
	return (cf.accessFlags & intcls.AccAnnotation) != 0
}

func (cf ClassFile) IsInterface() bool {
	return (cf.accessFlags & intcls.AccInterface) != 0
}

func (cf ClassFile) IsModule() bool {
	return (cf.accessFlags & intcls.AccModule) != 0
}

func (cf ClassFile) IsStatic() bool {
	return (cf.accessFlags & intcls.AccStatic) != 0
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

func (cf ClassFile) Fields() []intcls.IField {
	return cf.field
}

func (cf ClassFile) Methods() []intcls.IMethod {
	return cf.method
}

func (cf ClassFile) Attributes() map[string]intcls.IAttribute {
	return cf.attributes
}

func (cf ClassFile) Attribute(name string) intcls.IAttribute {
	return cf.attributes[name]
}

func (cf ClassFile) OuterClassFile() intcls.IClassFile {
	return cf.outerClassFile
}

func (cf ClassFile) SetOuterClassFile(outerClassFile intcls.IClassFile) {
	cf.outerClassFile = outerClassFile
}

func (cf ClassFile) InnerClassFiles() []intcls.IClassFile {
	return cf.innerClassFiles
}

func (cf ClassFile) SetInnerClassFiles(innerClassFiles []intcls.IClassFile) {
	cf.innerClassFiles = innerClassFiles
}

func (cf ClassFile) String() string {
	return "ClassFile { Internal Type Name: " + cf.InternalTypeName() + " }"
}
