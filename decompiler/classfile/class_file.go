package classfile

// Access flags for Class, Field, Method, Nested class, Module, Module Requires, Module Exports, Module Opens
const (
	AccPublic       = 0x0001 // C  F  M  N  .  .  .  .
	AccPrivate      = 0x0002 // .  F  M  N  .  .  .  .
	AccProtected    = 0x0004 // .  F  M  N  .  .  .  .
	AccStatic       = 0x0008 // C  F  M  N  .  .  .  .
	AccFinal        = 0x0010 // C  F  M  N  .  .  .  .
	AccSynchronized = 0x0020 // .  .  M  .  .  .  .  .
	AccSuper        = 0x0020 // C  .  .  .  .  .  .  .
	AccOpen         = 0x0020 // .  .  .  .  Mo .  .  .
	AccTransitive   = 0x0020 // .  .  .  .  .  MR .  .
	AccVolatile     = 0x0040 // .  F  .  .  .  .  .  .
	AccBridge       = 0x0040 // .  .  M  .  .  .  .  .
	AccStaticPhase  = 0x0040 // .  .  .  .  .  MR .  .
	AccTransient    = 0x0080 // .  F  .  .  .  .  .  .
	AccVarArgs      = 0x0080 // .  .  M  .  .  .  .  .
	AccNative       = 0x0100 // .  .  M  .  .  .  .  .
	AccInterface    = 0x0200 // C  .  .  N  .  .  .  .
	AccAbstract     = 0x0400 // C  .  M  N  .  .  .  .
	AccStrict       = 0x0800 // .  .  M  .  .  .  .  .
	AccSynthetic    = 0x1000 // C  F  M  N  Mo MR ME MO
	AccAnnotation   = 0x2000 // C  .  .  N  .  .  .  .
	AccEnum         = 0x4000 // C  F  .  N  .  .  .  .
	AccModule       = 0x8000 // C  .  .  .  .  .  .  .
	AccMandated     = 0x8000 // .  .  .  .  Mo MR ME MO
)

func NewClassFile(majorVersion int, minorVersion int, accessFlags int,
	internalTypeName string, superTypeName string, interfaceTypeNames []string,
	field []Field, method []Method, attributes map[string]IAttribute) ClassFile {
	return ClassFile{
		MajorVersion:       majorVersion,
		MinorVersion:       minorVersion,
		AccessFlags:        accessFlags,
		InternalTypeName:   internalTypeName,
		SuperTypeName:      superTypeName,
		InterfaceTypeNames: interfaceTypeNames,
		Fields:             field,
		Methods:            method,
		Attributes:         attributes,
	}
}

type IClassFile interface {
	IsEnum() bool
	IsAnnotation() bool
	IsInterface() bool
	IsModule() bool
	IsStatic() bool
	Attribute(name string) IAttribute
	String() string
}

type ClassFile struct {
	MajorVersion       int
	MinorVersion       int
	AccessFlags        int
	InternalTypeName   string
	SuperTypeName      string
	InterfaceTypeNames []string
	Fields             []Field
	Methods            []Method
	Attributes         map[string]IAttribute
	OuterClassFile     IClassFile
	InnerClassFiles    []IClassFile
}

func (cf ClassFile) IsEnum() bool {
	return (cf.AccessFlags & AccEnum) != 0
}

func (cf ClassFile) IsAnnotation() bool {
	return (cf.AccessFlags & AccAnnotation) != 0
}

func (cf ClassFile) IsInterface() bool {
	return (cf.AccessFlags & AccInterface) != 0
}

func (cf ClassFile) IsModule() bool {
	return (cf.AccessFlags & AccModule) != 0
}

func (cf ClassFile) IsStatic() bool {
	return (cf.AccessFlags & AccStatic) != 0
}

func (cf ClassFile) Attribute(name string) IAttribute {
	return cf.Attributes[name]
}

func (cf ClassFile) String() string {
	return "ClassFile{ Internal Type Name: " + cf.InternalTypeName + " }"
}
